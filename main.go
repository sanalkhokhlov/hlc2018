package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"

	"github.com/sanalkhokhlov/hlc2018/httpserver"
	"github.com/sanalkhokhlov/hlc2018/store/engine"
	"github.com/sanalkhokhlov/hlc2018/store/service"
	"github.com/sanalkhokhlov/hlc2018/uploader"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered: ", r)
		}
	}()

	var err error
	memoryEngine := engine.NewMemoryEngine()
	dataStore := &service.DataStore{memoryEngine}
	err = uploader.Upload("/tmp/data", dataStore)
	// err = uploader.Upload("./data", dataStore)
	if err != nil {
		panic(err)
	}

	err = dataStore.MakeIndexes()
	if err != nil {
		panic(err)
	}

	httpServer := httpserver.Server{DataStore: dataStore}
	go httpServer.Run(80)

	runtime.GC()
	PrintMemUsage()

	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan struct{})
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		<-signalChan
		close(cleanupDone)
	}()
	<-cleanupDone
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB\n", bToMb(m.Alloc))
	fmt.Printf("TotalAlloc = %v MiB\n", bToMb(m.TotalAlloc))
	fmt.Printf("Sys = %v MiB\n", bToMb(m.Sys))
	fmt.Printf("NumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
