// Netcat1 is a read-only TCP client.
package main

import (
	"sort"
	"fmt"
	"bufio"
	"log"
	"net"
	"os"
	"strings"
)

type Clock struct {
	Location string
	Time string
}

func main() {
	waitings := make(chan Clock)
	for _, arg := range(os.Args[1:]) {
		go func(arg string) {
			// fmt.Fprintf(os.Stderr, "%v\n", arg)
			keyValue := strings.Split(arg, "=")
			location := keyValue[0]
			url := keyValue[1]
			conn, err := net.Dial("tcp", url)
			if err != nil {
				log.Fatal(err)
				return
			}
			defer conn.Close()

			scanner := bufio.NewScanner(conn)
			scanner.Split(bufio.ScanLines)
			for scanner.Scan() {
				data := scanner.Text()
				// fmt.Fprintf(os.Stderr, "%v\n", string(data))
				waitings<-Clock{Location:location, Time:string(data)}
			}
		}(arg)
	}

	clocks := map[string]Clock{}
	for clock := range(waitings) {
		clocks[clock.Location] = clock

		locations := []string{}
		for location, _ := range(clocks) {
			locations = append(locations, location)
		}
		sort.Slice(locations, func(i, j int) bool { return locations[i] < locations[j] })

		fmt.Fprintf(os.Stdout, "\r")
		for _, location := range(locations) {
			fmt.Fprintf(os.Stdout, "%v:%v ", location, clocks[location].Time)
		}
	}
}

