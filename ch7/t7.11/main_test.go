package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestDatabase(tt *testing.T) {
	go func() {
		database := map[string]dollars{"shoes": 50, "socks": 5}
		db := Database{database:database}
		http.HandleFunc("/list", db.list)
		http.HandleFunc("/price", db.price)
		http.HandleFunc("/update", db.update)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
	}()
	resp, err := http.Post("http://localhost:8000/update?item=socks&price=6",
		"", nil)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	tt.Logf("%v\n", string(body))
	// ...
}
