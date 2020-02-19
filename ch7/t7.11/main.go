package main

import (
	"fmt"
	"log"
	"sync"
	"strconv"
	"net/http"
)

//!+main
func main() {
	database := map[string]dollars{"shoes": 50, "socks": 5}
	db := Database{database:database}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!-main

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type Database struct {
	database map[string]dollars
	mux sync.Mutex
}

func (db Database) list(w http.ResponseWriter, req *http.Request) {
	db.mux.Lock()
	defer db.mux.Unlock()
	for item, price := range db.database {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db Database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	db.mux.Lock()
	defer db.mux.Unlock()
	if price, ok := db.database[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db Database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	priceStr := req.URL.Query().Get("price")
	db.mux.Lock()
	defer db.mux.Unlock()
	if _, ok := db.database[item]; ok {
		price, err := strconv.Atoi(priceStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest) // 400
			fmt.Fprintf(w, "no such price: %q\n", priceStr)
			return
		}
		db.database[item] = dollars(price)
		fmt.Fprintf(w, "%s\n", db.database[item])
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}
