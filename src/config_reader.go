/**
 *
 * @author: swayamraina@gmail.com
 * @dated : 14.03.2020
 *
**/


package main


import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)


/**
 *
 * This function reads the config file to create the
 * daemon config object at startup.
 *
**/
func get_config (path string) fcd_config {
	data, err := ioutil.ReadFile(path)
	if err != nil { panic(err) }
	var config fcd_config
	yaml.Unmarshal(data, &config)
	return config
}