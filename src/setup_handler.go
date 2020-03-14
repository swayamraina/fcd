/**
 * @author: swayamraina@gmail.com
 * @dated : 14.03.2020
**/

package main

import "os"


func do_local_setup (paths *[]string) {
	for i:=0; i<len(*paths); i++ {
		_, err := os.Stat((*paths)[i])
		if err == nil { continue }
		if os.IsNotExist(err) { os.Mkdir((*paths)[i], os.ModePerm) }
		panic(err)
	}
}



func  do_git_setup (config *git_config) {

}



func do_setup (config *fcd_config) {
	do_local_setup(&config.Search_locations)
	do_git_setup(&config.Git_config)
}
