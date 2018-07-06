package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"path/filepath"
)

func buildProject(ev *MainServer, w http.ResponseWriter, r *http.Request) {
	project := r.FormValue("project")
	cmdstr := fmt.Sprintf("%s/build.sh", filepath.Join(ev.BasePath, project))
	cmd := exec.Command("/bin/bash", cmdstr)
	log.Println(cmd)
	var res OperationResult
	if bytes, err := cmd.CombinedOutput(); err != nil {
		res.Status = 1005
		res.Message = string(bytes)
	} else {
        if len(bytes) == 0{
		    res.Message = "ok"
        }else{
            res.Message = string(bytes)
        }
	}
	msg, _ := json.Marshal(&res)
	w.Write(msg)
}
