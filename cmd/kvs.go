package main

import (
	"os"

	"github.com/fofcorp/kvstorage/src/httpserv"
	"github.com/fofcorp/kvstorage/src/storage"
	log "github.com/sirupsen/logrus"
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

func main() {
	customFormatter := &log.TextFormatter{}
	customFormatter.TimestampFormat = timeFormat
	customFormatter.FullTimestamp = true
	log.SetFormatter(customFormatter)
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)

	log.Info("kvs init")
	store := storage.InitInMemory()
	serv, err := httpserv.Init(
		&httpserv.Options{
			Port:  "8888",
			Store: store,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	if err = serv.Run(); err != nil {
		log.Fatal(err)
	}
}
