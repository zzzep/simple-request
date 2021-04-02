package simplerequest

import (
	"io/ioutil"
	"net/http"
)

const InternalServerError = 500
const OK = 200

var lastErr error

func Get(url string) (int, string) {
	resp, e := http.Get(url)
	if e != nil {
		lastErr = e
		return InternalServerError, "Fail to do the requisition"
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		lastErr = e
	}
	return resp.StatusCode, string(body)
}
