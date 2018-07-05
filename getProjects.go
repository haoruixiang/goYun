package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type OperationResult struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message"`
}

type ProjectData struct {
	Name     string   `json:"name"`
	FuncName []string `json:"func_name"`
}

func getProjects(ev *MainServer, w http.ResponseWriter, r *http.Request) {
	log.Println("getProjects", ev.BasePath)
	res := OperationResult{
		Status:  0,
		Message: getPathPaths(ev.BasePath),
	}
	msg, _ := json.Marshal(&res)
	log.Println(string(msg))
	w.Write(msg)
}
