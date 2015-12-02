package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	d := 8000
	var port int
	flag.IntVar(&port, "port", d, "listen port")
	flag.IntVar(&port, "p", d, "listen port")
	flag.Parse()

	panic(http.ListenAndServe(fmt.Sprintf(":%d", port), http.FileServer(http.Dir("."))))
}
