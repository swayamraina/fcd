package main

import "fmt"

func getUrl (host string, endpoint string, vars ...string) string {
	endpoint = fmt.Sprintf("%s%s", endpoint, vars)
	return fmt.Sprintf("%s%s", host, endpoint)
}

func max (x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}