package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var addr = flag.String("addr", ":8080", "http service address")
	flag.Parse()

	r := newRoom()

	http.Handle("/", r)

	go r.run()

	log.Println("Web server started on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
