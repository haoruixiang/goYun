package main

import (
	"encoding/json"
	"net/http"
)

func getProjectFuncs(ev *MainServer, w http.ResponseWriter, r *http.Request) {
	project := r.FormValue("project")
	res := OperationResult{
		Status: 0,
		Message: ProjectData{
			Name:     project,
			FuncName: getPathFiles(ev.BasePath, project),
		},
	}
	msg, _ := json.Marshal(&res)
	w.Write(msg)
}
