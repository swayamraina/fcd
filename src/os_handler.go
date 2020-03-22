package main

import (
	"os"
	"path/filepath"
	"strings"
)

const sync_indicator = "fcd-"


func check_on_os (path string) bool {
	_, err := os.Stat(path)
	if err == nil { return true }
	if os.IsNotExist(err) { return false }
	return false
}



func create_on_os (path string) bool {
	err := os.Mkdir(path, os.ModePerm)
	return err == nil
}



func check_sync (files *[]string) filepath.WalkFunc {
	return func (path string, info os.FileInfo, err error) error {
		if !strings.HasPrefix(info.Name(), sync_indicator) {
			*files = append(*files, path)
		}
		return nil
	}
}