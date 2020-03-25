/**
 * @author: swayamraina@gmail.com
 * @dated : 14.03.2020
**/

package main



func do_local_setup (paths *[]string) {
	for i:=0; i<len(*paths); i++ {
		exists := check_on_os((*paths)[i])
		if !exists {
			create_on_os((*paths)[i])
		}
	}
}



func  do_git_setup (config *git_config) {
	exists := check_repo_exists(&config.Username, &config.Repo)
	if !exists {
		create_repo(&config.Username, &config.Access_token, &config.Repo)
	}
}



func do_setup (config *fcd_config) {
	do_local_setup(&config.Search_locations)
	do_git_setup(&config.Git_config)
}
