/**
 *
 * @author: swayamraina@gmail.com
 * @dated : 14.03.2020
 *
**/


package main


/**
 *
 * This is the main function which gets executed
 * as the daemon process starts
 *
**/
func main () {
	config := get_config()
	success := true //do_setup(&config)
	if success {
		daemon(&config)
	}
}