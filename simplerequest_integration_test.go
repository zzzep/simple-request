package simplerequest

import "testing"

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
