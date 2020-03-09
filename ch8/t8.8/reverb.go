// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 224.

// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

//!+
func handleConn(c net.Conn) {
	done := make(chan struct{})
	go func() {
		input := bufio.NewScanner(c)
		for input.Scan() {
			done<-struct{}{}
			go echo(c, input.Text(), 1*time.Second)
		}
		defer c.Close()
	}()
	ticker := time.NewTicker(10 * time.Second)
	for {
		select {
		case <- ticker.C:
			fmt.Printf("Over 10 sec\n")
			ticker.Stop()
			defer c.Close()
			return
		case <- done:
			ticker.Stop()
		}
	}
  // NOTE: ignoring potential errors from input.Err()
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
