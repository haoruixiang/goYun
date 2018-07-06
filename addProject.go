package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

var build_string = ``

func addProject(ev *MainServer, w http.ResponseWriter, r *http.Request) {
	project := r.FormValue("project")
	currentPath := filepath.Join(ev.BasePath, project)
	var res OperationResult
	if err := os.Mkdir(currentPath, os.ModePerm); err != nil {
		res.Status = 1003
		res.Message = err.Error()
	} else {
		build_file := filepath.Join(currentPath, "build.sh")
		if file, err := os.OpenFile(build_file, os.O_CREATE|os.O_WRONLY, 0644); err != nil {
			res.Status = 1003
			res.Message = err.Error()
		} else {
			build_string := fmt.Sprintf("#!/bin/bash\n cd %s; gofmt -w *.go; go build;", currentPath)
			if _, err := file.Write([]byte(build_string)); err != nil {
				res.Status = 1003
				res.Message = err.Error()
			}
			if err := file.Close(); err != nil {
				res.Status = 1003
				res.Message = err.Error()
			} else {
				res.Message = "ok"
			}
		}
	}
	msg, _ := json.Marshal(&res)
	w.Write(msg)
}
