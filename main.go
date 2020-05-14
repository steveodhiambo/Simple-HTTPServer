package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Hello World")
}

func main() {

	port := flag.String("port",":8080","Insert the port you want to use for the application eg :8000, :4000 etc")
	flag .Parse()
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(*port,nil))
}
