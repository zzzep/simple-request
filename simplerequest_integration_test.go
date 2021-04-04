package simplerequest

import (
	"testing"
)

type responseGetJson struct {
	Success string `json:"success"`
}

func TestExample(t *testing.T) {
	code, resp := Get("http://example.com")
	if code != OK {
		t.Error("wrong status code")
		t.Error(code)
	}
	if resp == "" {
		t.Error("nil response")
	}
	if lastErr != nil {
		t.Error(lastErr)
	}
}

func TestExampleWithHeader(t *testing.T) {
	code, resp, h := GetH("http://example.com", map[string]string{})
	if code != OK {
		t.Error("wrong status code")
		t.Error(code)
	}
	if _, ok := h["Date"]; !ok {
		t.Error("fail to get header")
	}
	if resp == "" {
		t.Error("nil response")
	}
	if lastErr != nil {
		t.Error(lastErr)
	}
}

func TestPostExample(t *testing.T) {
	body := map[string][]string{"teste": {"teste123"}}
	code, resp := Post("https://reqbin.com/echo/post/json", body)
	if code != OK {
		t.Error("wrong status code")
		t.Error(code)
	}
	if resp == "" {
		t.Error("nil response")
	}
	if lastErr != nil {
		t.Error(lastErr)
	}
}

func TestPostWithHeaderExample(t *testing.T) {
	header := map[string][]string{"teste": {"teste123"}}
	code, resp, h := PostH("https://reqbin.com/echo/post/json", header, "")
	if code != OK {
		t.Error("wrong status code")
		t.Error(code)
	}
	if _, ok := h["Date"]; !ok {
		t.Error("fail to get header")
	}
	if resp == "" {
		t.Error("nil response")
	}
	if lastErr != nil {
		t.Error(lastErr)
	}
}

func TestPutExample(t *testing.T) {
	code, resp := Put("https://reqbin.com/echo/put/json", "body")
	if code != OK {
		t.Error("wrong status code")
		t.Error(code)
	}
	if resp == "" {
		t.Error("nil response")
	}
	if lastErr != nil {
		t.Error(lastErr)
	}
}

func TestPutWithHeaderExample(t *testing.T) {
	header := map[string][]string{"teste": {"teste123"}}
	code, resp, h := PutH("https://reqbin.com/echo/put/json", header, "")
	if code != OK {
		t.Error("wrong status code")
		t.Error(code)
	}
	if _, ok := h["Date"]; !ok {
		t.Error("fail to get header")
	}
	if resp == "" {
		t.Error("nil response")
	}
	if lastErr != nil {
		t.Error(lastErr)
	}
}

func TestGetToJson(t *testing.T) {
	var r responseGetJson
	c := GetToJson("https://reqbin.com/echo/get/json", &r)
	if r.Success != "true" {
		t.Error("r.success", r.Success, "is not true")
	}
	if c != 200 {
		t.Error("Status code invalid")
	}
}
