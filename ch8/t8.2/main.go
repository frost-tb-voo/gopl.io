// Clock is a TCP server that periodically writes the time.
package main

import (
  "os/exec"
  "strconv"
  "bufio"
	"flag"
	"io"
	"log"
  "net"
  "fmt"
  "os"
  "strings"
)

type FtpSession struct {
  c net.Conn
  dataconn net.Conn
}

func handleConn(c net.Conn) {
  session := FtpSession{c:c}
  defer func(session FtpSession) {
    if session.c != nil {
      session.c.Close()
    }
    if session.dataconn != nil {
      session.dataconn.Close()
    }
  }(session)

	for {
    fmt.Fprintf(os.Stdout, "remote %v\n", c.RemoteAddr())
    fmt.Fprintf(os.Stdout, "local  %v\n", c.LocalAddr())
		_, err := io.WriteString(c, "220 FTP session started\n")
		if err != nil {
      log.Fatal(err)
      return
    }
    scanner := bufio.NewScanner(c)
    scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
      command := scanner.Text()
      fmt.Fprintf(os.Stdout, "%v\n", command)
      if strings.HasPrefix(command, "USER") {
        err = user(c)
        if err != nil {
          log.Fatal(err)
          return
        }
      }
      if strings.HasPrefix(command, "SYST") {
        err = syst(c)
        if err != nil {
          log.Fatal(err)
          return
        }
      }
      if strings.HasPrefix(command, "CWD") {
        err = cd(c)
        if err != nil {
          log.Fatal(err)
          return
        }
      }
      if strings.HasPrefix(command, "QUIT") {
        err = close(c)
        if err != nil {
          log.Fatal(err)
          return
        }
        return
      }
      if strings.HasPrefix(command, "PORT") {
        if session.dataconn != nil {
          session.dataconn.Close()
        }
        dataconn, err := port(c, command)
        if dataconn != nil {
          session.dataconn = dataconn
        }
        if err != nil {
          log.Fatal(err)
          return // e.g., client disconnected
        }
        go func(dataconn net.Conn) {
          scanner := bufio.NewScanner(dataconn)
          scanner.Split(bufio.ScanLines)
          for scanner.Scan() {
            command := scanner.Text()
            fmt.Fprintf(os.Stdout, "%v\n", command)
          }
        }(session.dataconn)
      }
      if strings.HasPrefix(command, "LIST") {
        err = ls(c, session.dataconn)
        if err != nil {
          log.Fatal(err)
          return // e.g., client disconnected
        }
      }
      if strings.HasPrefix(command, "RETR") {
        err = get(c, session.dataconn)
        if err != nil {
          log.Fatal(err)
          return // e.g., client disconnected
        }
      }
    }
	}
}

func main() {
	port := flag.String("port", "10021", "port num")
	flag.Parse()
	
	listener, err := net.Listen("tcp", "localhost:" + *port)
	if err != nil {
		log.Fatal(err)
	}
	//!+
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
	//!-
}

func user(c net.Conn) error {
  _, err := io.WriteString(c, "230 Anonymous user logged in.\n")
  return err
}

func syst(c net.Conn) error {
  _, err := io.WriteString(c, "215 ftp syst.\n")
  return err
}

func cd(c net.Conn) error {
  _, err := io.WriteString(c, "200 ftp cwd.\n")
  return err
}

func close(c net.Conn) error {
  _, err := io.WriteString(c, "221 ftp quit.\n")
  return err
}

func port(c net.Conn, command string) (net.Conn, error) {
  addresses := strings.Split(strings.Split(command, " ")[1], ",")
  var head, foot int
  var err error
  if head, err = strconv.Atoi(addresses[4]) ; err != nil {
    return nil, err
  }
  if foot, err = strconv.Atoi(addresses[5]) ; err != nil {
    return nil, err
  }
  address := strings.Join(addresses[:4], ".") + ":" + strconv.Itoa(head * 256 + foot)
  var dataconn net.Conn
  dataconn, err = net.Dial("tcp", address)
  if err != nil {
    return nil, err
  }
  _, err = io.WriteString(c, "225 ftp port.\n")
  if err != nil {
    return nil, err
  }
  return dataconn, nil
}

func ls(c net.Conn, dataconn net.Conn) error {
  var err error
  _, err = io.WriteString(c, "125 ftp list.\n")
  if err != nil {
    io.WriteString(c, "426 ftp list.\n")
    return err
  }
  // _, err = io.WriteString(dataconn, "data list.\n")
  output, err := exec.Command("ls", "-al").Output()
  if err != nil {
    io.WriteString(c, "426 ftp list.\n")
    return err
  }
  _, err = io.WriteString(dataconn, string(output))
  if err != nil {
    io.WriteString(c, "426 ftp list.\n")
    return err
  }
  fmt.Fprintf(os.Stderr, "sent data list\n")
  _, err = io.WriteString(c, "226 ftp list.\n")
  dataconn.Close()
  return err
}

func get(c net.Conn, dataconn net.Conn) error {
  var err error
  _, err = io.WriteString(c, "125 ftp retr.\n")
  if err != nil {
    io.WriteString(c, "426 ftp retr.\n")
    return err
  }
  _, err = io.WriteString(dataconn, "data retr.\n")
  if err != nil {
    io.WriteString(c, "426 ftp retr.\n")
    return err
  }
  fmt.Fprintf(os.Stderr, "sent data retr\n")
  _, err = io.WriteString(c, "226 ftp retr.\n")
  dataconn.Close()
  return err
}
