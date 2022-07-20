package main

import (
	"log"
	"net/http"
	"os"

	"github.com/anksus/http-server/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api ", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodBye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)
	http.ListenAndServe(":9090", sm)
}
