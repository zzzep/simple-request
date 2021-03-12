package simplerequest

import (
	"io/ioutil"
	"log"
	"net/http"
)

func get(url string) (code int, body interface{}, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, nil, err
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return resp.StatusCode, body, nil
}
