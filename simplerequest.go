package simplerequest

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Get(url string) (c int, b string, e error) {
	resp, e := http.Get(url)
	if e != nil {
		return 0, "", e
	}

	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		log.Fatalln(e)
	}

	b = string(body)

	return resp.StatusCode, b, nil
}
