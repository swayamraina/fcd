/**
 *
 * @author: swayamraina@gmail.com
 * @dated : 14.03.2020
 *
**/


package main


/**
 *
 * DS for github connection
 *
**/
type git_config struct {

	// github username
	Username string

	// email
	Email string

	// github access token
	Access_token string

	// github repo
	Repo string

}


/**
 *
 * DS for daemon ping settings
 *
**/
type refresh_config struct {

	// refresh interval
	Refresh_interval int64

	// refresh unit
	Refresh_unit string

}


/**
 *
 * DS for overall fcd settings
 *
**/
type fcd_config struct {

	// git config
	Git_config git_config

	//  refresh config
	Refresh_config refresh_config

	// daemon search paths
	Search_locations []string

}