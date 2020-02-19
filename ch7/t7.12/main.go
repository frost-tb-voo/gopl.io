package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

//!+main
func main() {
	database := map[string]dollars{"shoes": 50, "socks": 5}
	db := Database{database: database}
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
	mux      sync.Mutex
}

var list = template.Must(template.New("list").Parse(`
<h1>prices</h1>
<table>
<tr style='text-align: left'>
  <th>Item</th>
  <th>Price</th>
</tr>
{{range $key, $value := .}}
<tr>
<td>{{$key}}</td>
<td>{{$value}}</td>
</tr>
{{end}}
</table>
`))

func (db Database) list(w http.ResponseWriter, req *http.Request) {
	db.mux.Lock()
	defer db.mux.Unlock()
	if err := list.Execute(w, db.database); err != nil {
		log.Fatal(err)
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
		}
		db.database[item] = dollars(price)
		fmt.Fprintf(w, "%s\n", db.database[item])
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}
