/**
 *
 * @author: swayamraina@gmail.com
 * @dated : 22.03.2020
 *
**/


package main


import (
	"os"
	"path/filepath"
	"strings"
)


/**
 *
 * This connstant is used to mark the scanned
 * files as synced.
 * This is also used to check for new files in
 * the search locations.
 *
**/
const sync_indicator = "fcd-"


/**
 *
 * This utility function checks if the incoming
 * file/dir exists in file system or not
 *
**/
func check_on_os (path string) bool {
	_, err := os.Stat(path)
	if err == nil { return true }
	if os.IsNotExist(err) { return false }
	return false
}


/**
 *
 * This utility function creates a dir on the
 * file system.
 *
**/
func create_on_os (path string) bool {
	err := os.Mkdir(path, os.ModePerm)
	return err == nil
}


/**
 *
 * This function checks if a particular file has
 * been alreday synced to github or is a new file
 *
**/
func check_sync (files *[]string) filepath.WalkFunc {
	return func (path string, info os.FileInfo, err error) error {
		if !info.IsDir() && !strings.HasPrefix(info.Name(), sync_indicator) {
			*files = append(*files, path)
		}
		return nil
	}
}


/**
 *
 * This function marks a specified file as synced.
 * fcd adds a marker in the beginning of the filename
 * to maark the file as synced.
 *
**/
func mark_sync (path string) bool {
	base := extract_base_name(path)
	new_filename := base + sync_indicator + extract_file_name(path)
	err := os.Rename(path, new_filename)
	return err == nil
}