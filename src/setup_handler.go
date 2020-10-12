/**
 *
 * @author: swayamraina@gmail.com
 * @dated : 14.03.2020
 *
**/


package main

import "errors"

/**
 *
 * This function is the wrapper function which validates
 * the config by checking the path does exists on OS and
 * if it does not, creates the dir on file-system
 *
**/
func do_local_setup (paths *[]string) error {
	for i:=0; i<len(*paths); i++ {
		exists := check_on_os((*paths)[i])
		if !exists {
			success := create_on_os((*paths)[i])
			if !success {
				return errors.New("file setup failed")
			}
		}
	}
	return nil
}


/**
 *
 * This function is the wrapper function which validates
 * the config by checking the git path does exists on the
 * external servers and if it does, creates the repo for
 * the user
 *
**/
func  do_git_setup (config *git, repo *string) {
	exists := check_repo_exists(&config.Username, &config.Access_token, repo)
	if !exists {
		create_repo(config, repo)
	}
}


/**
 *
 * This function is the wrapper function over the
 * actual setup wrapper functions.
 *
**/
func do_setup (config *fcd_config) bool {
	for i:=0; i<len(config.Sync); i++ {
		err := do_local_setup(&config.Sync[i].Locations)
		if err != nil {
			return false
		}
		do_git_setup(&config.Git, &config.Sync[i].Repo)
	}
	return true
}
