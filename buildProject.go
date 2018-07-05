package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func buildProject(ev *MainServer, w http.ResponseWriter, r *http.Request) {
	project := r.FormValue("project")
	//gofmt -w *.go
	//go build
	//return build status
	w.Write(msg)
}
