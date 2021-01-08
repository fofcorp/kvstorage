package main

import (
	"log"

	"github.com/fofcorp/kvstorage/src/httpserv"
	"github.com/fofcorp/kvstorage/src/storage"
)

func main() {
	store := storage.InitInMemory()
	log.Println(store)
	serv, err := httpserv.Init(&httpserv.Options{Port: "8888"})
	if err != nil {
		log.Fatalln(err)
	}
	if err = serv.Run(); err != nil {
		log.Fatalln(err)
	}
}
