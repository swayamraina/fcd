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
const environment_file = "/.fcd/fcd.env"


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
	path_to_env := home + environment_file
	data, err := ioutil.ReadFile(path_to_env)
	if err != nil { panic(err) }
	return get_path(string(data))
}


/**
 *
 * This function extracts the path to fcd_config yaml
 * file from the environment file
 *
**/
func get_path (data string) string {
	var path string
	envs := strings.Split(data, "\n")
	for _, env := range envs {
		kv := strings.Split(env, "=")
		if "path" == kv[0] {
			path = kv[1]
		}
	}
	return path
}