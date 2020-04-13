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
type git struct {
	Username string
	Email string
	Access_token string
	Private bool
}


/**
 *
 * DS for daemon ping settings
 *
**/
type refresh struct {
	Interval int64
	Unit string
}


/**
 *
 * DS for syncing data between the repo and
 * the location on the hardware.
 * Note that the 'dirs' field tells fcd to treat
 * these locations as separate directories or not
 *
**/
type sync struct {
	Repo string
	Locations []string
	Merge bool
}


/**
 *
 * DS for overall fcd settings
 *
**/
type fcd_config struct {
	Git git
	Refresh refresh
	Sync []sync
}