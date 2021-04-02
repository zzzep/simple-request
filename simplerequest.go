package simplerequest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const InternalServerError = 500
const OK = 200

var lastErr error

type BaseResponse interface{}

func Get(u string) (int, string) {
	resp, e := http.Get(u)
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

func Post(u string, body map[string][]string) (int, string) {
	data := url.Values(body)
	resp, err := http.PostForm(u, data)
	if err != nil {
		lastErr = err
	}
	if resp == nil {
		return 500, ""
	}
	res, e := ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if e != nil {
		lastErr = e
	}
	return resp.StatusCode, string(res)
}

//url and the struct with return pattern
func GetToJson(u string, response BaseResponse) int {
	c, r := Get(u)
	b := []byte(r)
	e := json.Unmarshal(b, &response)
	if e != nil {
		lastErr = e
	}
	return c
}
