/**
 *
 * @author: swayamraina@gmail.com
 * @dated : 22.03.2020
 *
**/


package main


import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)


/**
 *
 * This utility function is used to generate the complete
 * API url given host, endpoint and url variables
 *
**/
func get_url (host string, endpoint string, vars ...string) string {
	for i,_ := range vars {
		endpoint = strings.Replace(endpoint, "%s", vars[i], 1)
	}
	return host + endpoint
}


/**
 *
 * This utility function returns the max of two int64
 * values.
 * This is required since Go doesn't have an inbuilt
 * max() function for integers.
 *
**/
func max (x, y int64) int64 {
	if x > y { return x }
	return y
}


/**
 *
 * This utility function reads the file content from
 * file-system and returns the base-64 encoded value
 * which is required to be sent over wire for storage
 *
**/
func generate_b64 (path string) string {
	file, err := os.Open(path)
	if nil !=  err { panic(err) }
	reader := bufio.NewReader(file)
	content, err := ioutil.ReadAll(reader)
	if nil !=  err { panic(err) }
	return base64.StdEncoding.EncodeToString(content)
}


/**
 *
 * This utility function extracts out the file name from
 * the complete file-system path
 *
**/
func extract_file_name (path string) string {
	index := strings.LastIndex(path, "/")
	return path[index+1:]
}


func extract_storage_path (path string) string {
	var index int
	index = strings.LastIndex(path, "/")
	index = strings.LastIndex(path[:index], "/")
	storage_path := path[index+1:]
	return storage_path
}


/**
 *
 * This utility function extracts out the location of file
 * from the file-system path
 *
**/
func extract_base_name (path string) string {
	index := strings.LastIndex(path, "/")
	return path[:index+1]
}


/**
 *
 * This utility function is a wrapper function over
 * the actual HTTP POST call.
 *
**/
func make_post_request (url string, body []byte, token *string) (*http.Response, error) {
	return make_request(post_request, url, body, token)
}


/**
 *
 * This utility function is a wrapper function over
 * the actual HTTP GET call.
 *
**/
func make_get_request (url string, token *string) (*http.Response, error) {
	return make_request(get_request, url, nil, token)
}


/**
 *
 * This utility function is a wrapper function over
 * the actual HTTP PUT call.
 *
**/
func make_put_request (url string, body []byte, token *string) (*http.Response, error) {
	return make_request(put_request, url, body, token)
}


/**
 *
 * This utility function does actual API call to the external
 * world. This returns the response object for clients to use.
 *
**/
func make_request (method string, url string, body []byte, token *string) (*http.Response, error) {
	request, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if nil !=  err { panic(err) }
	request.Header.Add(content_type_key, content_type_value)
	request.Header.Add(auth_key, get_token_header(token))
	client := &http.Client{}
	return client.Do(request)
}


/**
 *
 * This utility function generates the auth header as per the
 * github documentation.
 *
**/
func get_token_header (token *string) string {
	return fmt.Sprintf(auth_value, *token)
}