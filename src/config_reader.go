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
	"os"
	"strings"
)


/**
 *
 * This constant holds the environment file name
 * where all fcd related environment variables
 * are stored
 *
**/
const env_file = "/.fcd_env"


/**
 *
 * This function reads the config file to create the
 * daemon config object at startup.
 *
**/
func get_config () fcd_config {
	path := get_env_data()
	data, err := ioutil.ReadFile(path)
	if err != nil { panic(err) }
	var config fcd_config
	yaml.Unmarshal(data, &config)
	return config
}


/**
 *
 * This function extracts the path to fcd_config yaml
 * file
 *
**/
func get_env_data () string {
	home := os.Getenv("HOME")
	path_to_env := home + env_file
	path, err := ioutil.ReadFile(path_to_env)
	if err != nil { panic(err) }
	return strings.TrimSuffix(string(path), "\n")
}