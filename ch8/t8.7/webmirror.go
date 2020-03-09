package main

import (
  "fmt"
  "log"
  "flag"
  "sync"
  "os"
  "path"
  "strings"
  "net/url"
  "net/http"
  "path/filepath"
  "golang.org/x/net/html"
)

func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
  forEachNode(doc, visitNode, nil)
  
  hostname, nextPath, err := localPath(url)
  if err != nil {
    return nil, err
  }
  modifyNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
      newAttributes := []html.Attribute{}
			for _, a := range n.Attr {
				if a.Key != "href" {
          newAttributes = append(newAttributes, a)
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
          fmt.Fprintf(os.Stderr, "%v\n", err)
          newAttributes = append(newAttributes, a)
					continue // ignore bad URLs
        }
        hostname2, nextPath2, err := localPath(link.String())
        if err != nil {
          fmt.Fprintf(os.Stderr, "%v\n", err)
          newAttributes = append(newAttributes, a)
					continue // ignore bad URLs
        }
        if hostname != hostname2 {
          newAttributes = append(newAttributes, a)
          continue
        }
        rel2, err := filepath.Rel(nextPath, nextPath2)
        if err != nil {
          fmt.Fprintf(os.Stderr, "%v\n", err)
          newAttributes = append(newAttributes, a)
					continue // ignore bad URLs
        }
        rel2 = "./" + rel2
        // rel2 = link.String()
        newAttr := html.Attribute{Namespace:a.Namespace,Key:a.Key,Val:rel2}
        newAttributes = append(newAttributes, newAttr)
        // fmt.Fprintf(os.Stderr, "%v --> %v\n", nextPath, nextPath2)
        // fmt.Fprintf(os.Stderr, "%v %v\n", rel2, a)
      }
      n.Attr = newAttributes
		}
  }
  forEachNode(doc, modifyNode, nil)

  if strings.HasSuffix(url, "/") {
    nextPath = path.Join(nextPath, "index.html")
  }
  info, err := os.Stat(path.Dir(nextPath))
  if err != nil || !info.IsDir() {
    if err = os.MkdirAll(path.Dir(nextPath), os.ModePerm); err != nil {
      return nil, err
    }
  }
  info, err = os.Stat(nextPath)
  if err == nil && info.IsDir() {
    nextPath = path.Join(nextPath, "index.html")
  }
  htmlFile, err := os.Create(nextPath)
  if err != nil {
    htmlFile, err = os.OpenFile(nextPath, os.O_RDWR | os.O_TRUNC, 0644)
    if err != nil {
      return nil, err
    }
  }
  err = html.Render(htmlFile, doc)
  if err != nil {
    return nil, err
  }

	return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func localPath(remoteURL string) (string, string, error) {
  var hostname string
  var localpath string
  url, err := url.Parse(remoteURL)
  if err != nil {
    return "", "", err
  }
  hostname = url.Hostname()
  localpath = path.Join(".", "temp", hostname, url.EscapedPath())
  return hostname, localpath, nil
}

func crawl(url string) []string {
  fmt.Println(url)
  list, err := Extract(url)
  if err != nil {
    log.Print(err)
  }
  return list
}

type work struct {
  list []string
  depth int
}

type unseenLink struct {
  link string
  depth int
}

//!+
func main() {
  depth := flag.Int("depth", 0, "depth limit")
  flag.Parse()
  args := flag.Args()
  fmt.Printf("%v %v\n", args, *depth)
  if len(args) <= 0 {
    fmt.Printf("No args.\n")
    return
  }
  var hostname string
  if url, err := url.Parse(args[0]); err == nil {
    hostname = url.Hostname()
  } else {
    fmt.Printf("Invalid url %v\n", args[0])
    return
  }

  worklist := make(chan work)  // lists of URLs, may have duplicates
  unseenLinks := make(chan unseenLink) // de-duplicated URLs
  var thrownTasks sync.WaitGroup

  // Create 20 crawler goroutines to fetch each unseen link.
  for i := 0; i < 20; i++ {
    go func() {
      for unseenLink := range unseenLinks {
        func() {
          foundLinks := crawl(unseenLink.link)
          defer func() {
          // fmt.Printf("finish work\n")
          thrownTasks.Done()
          }()
          if *depth <= 0 || *depth > unseenLink.depth {
            // fmt.Printf("gen %v work\n", len(foundLinks))
            thrownTasks.Add(len(foundLinks))
            worklist <- work{list:foundLinks,depth:unseenLink.depth+1}
          }
        }()
      }
    }()
  }

  // fmt.Printf("gen %v work\n", len(args))
  thrownTasks.Add(len(args))
  // Add command-line arguments to worklist.
  go func() {
    worklist <- work{list:args,depth:1}
  }()

  go func() {
    thrownTasks.Wait()
    close(worklist)
    // fmt.Printf("worklist released\n")
  }()

  // The main goroutine de-duplicates worklist items
  // and sends the unseen ones to the crawlers.
  seen := make(map[string]bool)
  for work := range worklist {
    for _, link := range work.list {
      func() {
        if url, err := url.Parse(link); err == nil {
          if url.Hostname() != hostname {
            defer func() {
              // fmt.Printf("finish work\n")
              thrownTasks.Done()
            }()
            return
          }
        } else {
          defer func() {
            // fmt.Printf("finish work\n")
            thrownTasks.Done()
          }()
          return
        }
        if !seen[link] {
          seen[link] = true
          unseenLinks <- unseenLink{link:link,depth:work.depth}
        } else {
          defer func() {
            // fmt.Printf("finish work\n")
            thrownTasks.Done()
          }()
        }
      }()
    }
  }

}

//!-
