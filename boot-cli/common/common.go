package common

import (
	"io/ioutil"
	"net/http"
)

func HttpGet(serverUrl, path string) ([]byte, error) {
	resp, err := http.Get(serverUrl + path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}
