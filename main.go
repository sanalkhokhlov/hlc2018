package main

import (
	"fmt"
	"os"
	"os/signal"

	"bitbucket.org/sLn/hlc2018/httpserver"
	"bitbucket.org/sLn/hlc2018/store/engine"
	"bitbucket.org/sLn/hlc2018/store/service"
	"bitbucket.org/sLn/hlc2018/uploader"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered: ", r)
		}
	}()

	memoryEngine := engine.NewMemoryEngine()
	dataStore := &service.DataStore{memoryEngine}
	// err := uploader.Upload("/tmp/data/data.zip", dataStore)
	err := uploader.Upload("./data", dataStore)
	if err != nil {
		panic(err)
	}

	err = dataStore.MakeIndexes()
	if err != nil {
		panic(err)
	}

	httpServer := httpserver.Server{DataStore: dataStore}
	go httpServer.Run(4000)

	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan struct{})
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		<-signalChan
		close(cleanupDone)
	}()
	<-cleanupDone
}
