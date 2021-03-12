package main

import "testing"

func TestExample(t *testing.T) {
	code, resp, err := Get("http://example.com")
	if err != nil {
		t.Error(err)
	}
	if code != 200 {
		t.Error("wrong status code")
		t.Error(code)
	}
	if resp == nil {
		t.Error("nil response")
	}
}
