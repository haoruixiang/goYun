package main

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
)

func addProject(ev *MainServer, w http.ResponseWriter, r *http.Request) {
	project := r.FormValue("project")
	currentPath := filepath.Join(ev.BasePath, project)
	var res OperationResult
	if err := os.Mkdir(currentPath, os.ModePerm); err != nil {
		res.Status = 1003
		res.Message = err.Error()
	} else {
		res.Message = "OK"
	}
	msg, _ := json.Marshal(&res)
	w.Write(msg)
}
