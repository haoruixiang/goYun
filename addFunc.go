package main

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
)

func addFunc(ev *MainServer, w http.ResponseWriter, r *http.Request) {
	project := r.FormValue("project")
	funcName := r.FormValue("func_name")
	funcData := r.FormValue("func_data")
	currentPath := filepath.Join(ev.BasePath, project, funcName)
	var res OperationResult
	if file, err := os.OpenFile(currentPath, os.O_CREATE|os.O_WRONLY, 0644); err != nil {
		res.Status = 1003
		res.Message = err.Error()
	} else {
		if _, err := file.Write([]byte(funcData)); err != nil {
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
	msg, _ := json.Marshal(&res)
	w.Write(msg)
}
