package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"syscall"
	"time"
)

type MainServer struct {
	BasePath string
}

func main() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Ldate)
	port := flag.Uint("port", 8085, "")
	limit := flag.Uint("limit", 20480, "")
	path := flag.String("path", "./", "")
	flag.Parse()
	ev := &MainServer{
		BasePath: *path,
	}
	var rLimit syscall.Rlimit
	rLimit.Max = uint64(*limit)
	rLimit.Cur = uint64(*limit)
	err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		log.Println("Error Setting Rlimit ", err)
	}
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", *port),
		Handler:        ev,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 48 * 1024,
	}
	server.ListenAndServe()
}
