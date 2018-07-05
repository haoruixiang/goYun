package main

import (
	"log"
	"os"
	"path/filepath"
)

func getPathPaths(path ...string) []string {
	currentPath := filepath.Join(path...)
	log.Println(currentPath)
	var ls []string
	err := filepath.Walk(currentPath, func(path string, f os.FileInfo, err error) error {
		log.Println(path)
		if f == nil {
			return err
		}
		if f.IsDir() {
			ls = append(ls, path)
			return nil
		}
		return nil
	})
	if err != nil {
		log.Println("ERROR", err)
	}
	return ls
}
