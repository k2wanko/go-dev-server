package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "listen port")
	flag.IntVar(&port, "p", 8080, "listen port")
	flag.Parse()

	panic(http.ListenAndServe(fmt.Sprintf(":%d", port), http.FileServer(http.Dir("."))))
}
