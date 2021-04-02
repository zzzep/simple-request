package simplerequest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const InternalServerError = 500
const OK = 200

var lastErr error

type BaseResponse interface{}

func Get(url string) (int, string) {
	resp, e := http.Get(url)
	if e != nil {
		lastErr = e
		return InternalServerError, "Fail to do the requisition"
	}
	body, err := ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		lastErr = e
	}
	return resp.StatusCode, string(body)
}

//url and the struct with return pattern
func GetToJson(url string, response BaseResponse) int {
	c, r := Get(url)
	b := []byte(r)
	e := json.Unmarshal(b, &response)
	if e != nil {
		lastErr = e
	}
	return c
}
