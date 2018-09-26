package utils 

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func HttpPostJson(postBody []byte, url string) ([]byte, error) {
	Debug(string(postBody))
	Debug(url)
	var resp *http.Response
	resp, err := http.Post(url, "application/json", bytes.NewReader(postBody))
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	Debug(string(data))
	return data, err
}

