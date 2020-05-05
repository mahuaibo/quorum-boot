package common

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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

func HandlePrivateKey(privateKey string) string {
	privateKey = strings.ToLower(privateKey)
	privateKey = strings.ReplaceAll(privateKey, "0x", "")
	return privateKey
}

func CreateFile(filename string, content []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(content)
	return err
}
