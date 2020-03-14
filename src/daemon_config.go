/**
 * @author: swayamraina@gmail.com
 * @dated : 14.03.2020
**/

package main


type git_config struct {

	// github username
	Username string

	// github password
	Password string

	// github repo
	Repo string

}

type refresh_config struct {

	// refresh interval
	Refresh_interval int64

	// refresh unit
	Refresh_unit string

}

type fcd_config struct {

	// git config
	Git_config git_config

	//  refresh config
	Refresh_config refresh_config

	// daemon search paths
	Search_locations []string

}