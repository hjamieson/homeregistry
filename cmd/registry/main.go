package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/hjamieson/homeregistry/internal/model"
	"github.com/hjamieson/homeregistry/internal/stores"
)

type app struct {
	devices Store
}

type Store interface {
	Put(d model.Device) error
	List() ([]model.Device, error)
	Get(id string) (*model.Device, error)
}

func main() {
	listen := flag.String("address", ":8080", "address for listen")
	flag.Parse()

	log.SetPrefix("|registrymain| ")
	log.SetFlags(log.Flags() | log.Lmsgprefix)
	log.Println("the registry!!")

	server := &app{
		devices: stores.NewMemStore(),
	}
	mux := http.NewServeMux()
	log.Printf("listening on %v", *listen)
	mux.HandleFunc("POST /devices", server.handlePost)
	mux.HandleFunc("GET /devices", server.handleList)
	mux.HandleFunc("GET /devices/{id}", server.handleGet)

	err := http.ListenAndServe(*listen, mux)
	if err != nil {
		log.Printf("error: %v", err)
	}
}
