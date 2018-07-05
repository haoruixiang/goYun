package main

import (
	"log"
	"net/http"
)

type funcHandler func(ev *MainServer, w http.ResponseWriter, r *http.Request)

var handlerMap = map[string]funcHandler{
	"/get":            getProjects,
	"/procject/get":   getProjectFuncs,
	"/procject/add":   addProject,
	"/procject/del":   delProject,
	"/func/get":       getFunc,
	"/func/add":       addFunc,
	"/func/update":    updateFunc,
	"/func/del":       delFunc,
	"/procject/build": buildProject,
}

func (ev *MainServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Host, r.URL.Path)
	if _, ok := handlerMap[r.URL.Path]; ok {
		handlerMap[r.URL.Path](ev, w, r)
	} else {
		w.Write([]byte(`{"status":0,"error":2990,"msg":"bad request"}`))
	}
}
