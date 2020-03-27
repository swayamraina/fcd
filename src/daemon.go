/**
 *
 * @author: swayamraina@gmail.com
 * @dated : 14.03.2020
 *
**/


package main


import (
	"path/filepath"
	"time"
)


/**
 *
 * This constant defines the default value of
 * the ping setting of the fcd daemon.
 * The daemon will wake up again in 10 mins
 *
**/
const default_sleep = 10 * 60 * 1000


/**
 *
 * This function is the daemon executor function
 * Daemon runs in a never ending loop to scan and
 * sync all files specified in the fcd config search
 * locations
 *
**/
func daemon (config *fcd_config)  {
	for true {
		new_files := scan_all(&config.Search_locations)
		if new_files != nil {
			sync_all(&new_files, &config.Git_config)
		}
		sleep(&config.Refresh_config)
	}
}


/**
 *
 * This function scans all the search locations
 * passed on to it. It returns all the new files
 * that have been added too the search locations.
 *
**/
func scan_all (paths *[]string) []string {
	var new_files []string
	for _, path := range *paths {
		new_files = append(new_files, scan(&path)...)
	}
	return new_files
}


/**
 *
 * This is the utility function which actually
 * scans a dir and returns new files added
 *
**/
func scan (path *string) []string {
	var new_files []string
	err := filepath.Walk(*path, check_sync(&new_files))
	if err != nil { panic(err) }
	return new_files
}


/**
 *
 * This function receives the list of files to be synced
 * to github. It syncs all files to github one by one
 *
**/
func sync_all (files *[]string, config *git_config)  {
	for _, file := range *files {
		sync(file, *config)
	}
}


/**
 *
 * This is the utility function which syncs a file
 * to github.
 * This function also marks the file as synced locally
 * so that daemoon does not  pick this up again for
 * sync.
 *
**/
func sync (file string, config git_config) bool {
	sync := add_file(file, &config.Username, &config.Repo, &config.Email, &config.Access_token)
	mark := mark_sync(file)
	return sync && mark
}


/**
 *
 * This utility function sleeps the the daemon process.
 * This function reads the refresh config and wakes
 * the daemon after specified time.
 * Note that, the refresh interval is 10 min by default.
 * Any value less than 10 minutes will be neglected
 *
**/
func sleep (config *refresh_config) {
	freq := config.Refresh_interval
	unit := config.Refresh_unit
	var multiplier int64
	switch unit {
		case "s": multiplier = 1000
		case "m": multiplier = 1000 * 60
		case "h": multiplier = 1000 * 60 * 60
		case "d": multiplier = 1000 * 60 * 60 * 24
	}
	sleep_time := max(default_sleep, freq*multiplier)
	time.Sleep(time.Duration(sleep_time) * time.Millisecond)
}
