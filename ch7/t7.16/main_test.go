package main

import (
	"io/ioutil"
	"net/http"
	"testing"
	"fmt"
	"net/http/httptest"
	"net/url"
)

func TestDatabase(tt *testing.T) {
	http.HandleFunc("/calc", calc)
	{
		uu, err := url.Parse("http://example.com/calc?")
		if err != nil {
			tt.Fatal(err)
		}		
		qq := uu.Query()
		qq.Set("expr", "1*sqrt(3)+6")
		uu.RawQuery = qq.Encode()

		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", uu.String(), nil)
		calc(w, req)
		resp := w.Result()
	
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			tt.Fatal(err)
		}
		fmt.Printf("%v: %v\n", resp.Status, string(body))
	}
	{
		uu, err := url.Parse("http://example.com/calc?")
		if err != nil {
			tt.Fatal(err)
		}		
		qq := uu.Query()
		qq.Set("expr", "hoge*pow(piyo,3)+fuga")
		uu.RawQuery = qq.Encode()

		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", uu.String(), nil)
		calc(w, req)
		resp := w.Result()
	
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			tt.Fatal(err)
		}
		fmt.Printf("%v: %v\n", resp.Status, string(body))
	}

	// Output:
}
