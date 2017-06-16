package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/", nil)

	handler(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Error("status wasn't OK")
	}

	if recorder.Header().Get("Content-Type") != "text/html" {
		t.Error("content-type wasn't text/html")
	}

	data, err := ioutil.ReadAll(recorder.Body)
	if err != nil {
		t.Error(err)
	}

	if string(data) != html {
		t.Error("content wasn't", html)
	}
}
