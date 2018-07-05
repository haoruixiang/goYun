package main

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
)

func delFunc(ev *MainServer, w http.ResponseWriter, r *http.Request) {
	project := r.FormValue("project")
	funcName := r.FormValue("func_name")
	currentPath := filepath.Join(ev.BasePath, project, funcName)
	var res OperationResult
	if err := os.Remove(currentPath); err != nil {
		res.Status = 1003
		res.Message = err.Error()
	} else {
		res.Message = "OK"
	}
	msg, _ := json.Marshal(&res)
	w.Write(msg)
}
