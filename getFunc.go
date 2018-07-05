package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func getFunc(ev *MainServer, w http.ResponseWriter, r *http.Request) {
	project := r.FormValue("project")
	funcName := r.FormValue("name")
	currentPath := filepath.Join(ev.BasePath, project, funcName)
	var res OperationResult
	if file, err := os.Open(currentPath); err != nil {
		res.Status = 1001
		res.Message = err.Error()
	} else {
		if data, err := ioutil.ReadAll(file); err != nil {
			res.Status = 1002
			res.Message = err.Error()
		} else {
			res.Message = string(data)
		}
		file.Close()
	}
	msg, _ := json.Marshal(&res)
	w.Write(msg)
}
