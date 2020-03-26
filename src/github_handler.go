/**
 *
 * @author: swayamraina@gmail.com
 * @dated : 24.03.2020
 *
**/


package main


import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)


/**
 *
 * Base URI for github APIs
 *
**/
const github_host = "https://api.github.com"


/**
 *
 * API endpoints for github APIs
 *
**/
const (
	get_repo_endpoint = "/repos/%s/%s"
	create_repo_endpoint = "/user/repos"
	add_file_endpoint = "/repos/%s/%s/contents/%s"
)


/**
 *
 * header constants for github APIs
 *
**/
const (
	content_type_key = "Content-Type"
	content_type_value = "application/json"
	auth_key = "Authorization"
	auth_value = "token %s"
)


/**
 *
 * HTTP method constants for github APIs
 *
**/
const (
	post_request = "POST"
	put_request = "PUT"
)


/**
 *
 * DS for create repository request
 *
**/
type create_repo_request struct {
	Name string `json:"name"`
}


/**
 *
 * DS for committer info for create file request
 *
**/
type committer struct {
	Name string `json:"name"`
	Email string `json:"email"`
}


/**
 *
 * DS for create file request
 *
**/
type create_file_request struct {
	Content string `json:"content"`
	Message string `json:"message"`
	Committer committer `json:"committer"`
}


/**
 *
 * This function makes a GET API request to github
 * to check the status of the repository mentioned
 * by the user.
 *
**/
func check_repo_exists (username *string, repo *string) bool {
	url := get_url(github_host, get_repo_endpoint, *username, *repo)
	response, err := http.Get(url)
	if nil !=  err { panic(err) }
	return response.StatusCode == 200
}


/**
 *
 * This function creates a new repository with the specified
 * name for the user on github.
 *
**/
func create_repo (username *string, token *string, repo *string) bool {
	body, err := json.Marshal(create_repo_request{*repo})
	if nil !=  err { panic(err) }
	url := get_url(github_host, create_repo_endpoint)
	response, err := make_post_request(url, body, token)
	return response.StatusCode == 200
}


/**
 *
 * This function creates a new file with the specified contents
 * in the repository specified by the user.
 * Note that, if filename is same or in case of a duplicate
 * this function does not save the file.
 *
**/
func add_file (path string, username *string, repo *string, email *string, token *string) bool {
	b64_content := generate_b64(path)
	filename := extract_file_name(path)
	body, err := json.Marshal(create_file_request {
		  b64_content,
		 create_friendly_commit_message(filename),
		committer {
				*username,
				*email,
				},
		},
	)
	if nil !=  err { panic(err) }
	url := get_url(github_host, add_file_endpoint, *username, *repo, filename)
	response, err := make_put_request(url, body, token)
	return response.StatusCode == 200 || response.StatusCode == 201
}


/**
 *
 * This function generates the commit message for a sync
 * request by fcd to github
 *
**/
func create_friendly_commit_message (filename string) string {
	t := time.Now()
	return fmt.Sprintf("[ synced : %s ] - %s", t.Format(time.ANSIC), filename)
}