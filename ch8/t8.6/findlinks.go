package main

import (
  "fmt"
  "log"
  "flag"
  "sync"

  "gopl.io/ch5/links"
)

func crawl(url string) []string {
  fmt.Println(url)
  list, err := links.Extract(url)
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

type SyncCount struct {
  sync.Mutex
  count int
}

//!+
func main() {
  depth := flag.Int("depth", 0, "depth limit")
  flag.Parse()
  args := flag.Args()
  fmt.Printf("%v %v\n", args, *depth)
  if len(args) <= 0 {
    return
  }

  worklist := make(chan work)  // lists of URLs, may have duplicates
  unseenLinks := make(chan unseenLink) // de-duplicated URLs
  var thrownTasks sync.WaitGroup
  var count SyncCount

  // Create 20 crawler goroutines to fetch each unseen link.
  for i := 0; i < 20; i++ {
    go func() {
      for unseenLink := range unseenLinks {
        func() {
          foundLinks := crawl(unseenLink.link)
          defer func() {
          // fmt.Printf("finish work\n")
          count.Lock()
          thrownTasks.Done()
          count.count--
          fmt.Printf("%v tasks\n", count.count)
          count.Unlock()
          }()
          if *depth <= 0 || *depth > unseenLink.depth {
            // fmt.Printf("gen %v work\n", len(foundLinks))
            count.Lock()
            thrownTasks.Add(len(foundLinks))
            count.count += len(foundLinks)
            fmt.Printf("%v tasks\n", count.count)
            count.Unlock()
            worklist <- work{list:foundLinks,depth:unseenLink.depth+1}
          }
        }()
      }
    }()
  }

  // fmt.Printf("gen %v work\n", len(args))
  count.Lock()
  thrownTasks.Add(len(args))
  count.count += len(args)
  fmt.Printf("%v tasks\n", count.count)
  count.Unlock()
  // Add command-line arguments to worklist.
  go func() {
    worklist <- work{list:args,depth:1}
  }()

  go func() {
    thrownTasks.Wait()
    close(worklist)
    fmt.Printf("worklist released\n")
  }()

  // The main goroutine de-duplicates worklist items
  // and sends the unseen ones to the crawlers.
  seen := make(map[string]bool)
  for work := range worklist {
    for _, link := range work.list {
      func() {
        if !seen[link] {
          seen[link] = true
          unseenLinks <- unseenLink{link:link,depth:work.depth}
        } else {
          defer func() {
          // fmt.Printf("finish work\n")
          count.Lock()
          thrownTasks.Done()
          count.count--
          fmt.Printf("%v tasks\n", count.count)
          count.Unlock()
          }()
        }
      }()
    }
  }

}

//!-
