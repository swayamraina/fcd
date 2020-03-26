/**
 *
 * @author: swayamraina@gmail.com
 * @dated : 14.03.2020
 *
**/


package main


/**
 *
 * This function is the wrapper function which validates
 * the config by checking the path does exists on OS and
 * if it does not, creates the dir on file-system
 *
**/
func do_local_setup (paths *[]string) {
	for i:=0; i<len(*paths); i++ {
		exists := check_on_os((*paths)[i])
		if !exists {
			create_on_os((*paths)[i])
		}
	}
}


/**
 *
 * This function is the wrapper function which validates
 * the config by checking the git path does exists on the
 * external servers and if it does, creates the repo for
 * the user
 *
**/
func  do_git_setup (config *git_config) {
	exists := check_repo_exists(&config.Username, &config.Repo)
	if !exists {
		create_repo(&config.Username, &config.Access_token, &config.Repo)
	}
}


/**
 *
 * This function is the wrapper function over the
 * actual setup wrapper functions.
 *
**/
func do_setup (config *fcd_config) {
	do_local_setup(&config.Search_locations)
	do_git_setup(&config.Git_config)
}
