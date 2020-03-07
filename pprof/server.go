package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/jiongdu/golang-snippets/pprof/handler"
)

const hostPort = ":9090"

func main() {
	flag.Parse()
	http.HandleFunc("/advance", handler.WithAdvanced(handler.Simple))

	http.HandleFunc("/simple", handler.Simple)
	http.HandleFunc("/", index)

	fmt.Println("Starting Server on ", hostPort)
	if err := http.ListenAndServe(hostPort, nil); err != nil {
		log.Fatalf("HTTP Server Failed: %v", err)
	}
}

func index(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-type", "text/html")
	io.WriteString(w, "<h2>Links</h2>\n<ul>")
	for _, link := range []string{"/advance", "/simple"} {
		fmt.Fprintf(w, `<li><a href="%v">%v</a>`, link, link)
	}
	io.WriteString(w, "</ul>")
}
