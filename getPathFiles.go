package main

import (
	"log"
	"os"
	"path/filepath"
)

func getPathFiles(path ...string) []string {
	currentPath := filepath.Join(path...)
	var ls []string
	err := filepath.Walk(currentPath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		ls = append(ls, path)
		return nil
	})
	if err != nil {
		log.Println("ERROR", err)
	}
	return ls
}
