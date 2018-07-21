package main

import (
	//"encoding/json"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"net/http"
	"os"
	"log"
	"path/filepath"
)

func getOneDoc(ev *MainServer, w http.ResponseWriter, r *http.Request) {
	docName := r.FormValue("doc_name")
	currentPath := filepath.Join(ev.BasePath, "doc", docName)
	var res OperationResult
	if file, err := os.Open(currentPath); err != nil {
		res.Status = 1001
		res.Message = err.Error()
	} else {
		if data, err := ioutil.ReadAll(file); err != nil {
			res.Status = 1002
			res.Message = err.Error()
		} else {
			html := blackfriday.MarkdownBasic(data)
			log.Println(string(html))
			w.Write(html)
			res.Message = string(html)
		}
		file.Close()
	}
	//msg, _ := json.Marshal(&res)
	//w.Write(msg)
}
