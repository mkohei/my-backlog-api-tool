package backlog

import (
	"io/ioutil"
	"net/http"
)

func Get(url string) (body []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return body, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}
	return body, nil
}
