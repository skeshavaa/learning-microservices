package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		d, _ := ioutil.ReadAll(r.Body)

		log.Printf("Data $s\n", d)
		fmt.Fprintf(rw, "Hello %s", d)
	})

	http.HandleFunc("/b", func(http.ResponseWriter, *http.Request) {
		log.Println("Bye World")
	})
	http.ListenAndServe(":9090", nil)
}
