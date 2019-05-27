package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReadFile(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/test.data")
	if err != nil {
		t.Fatal(err)
	}
	//fmt.Println(data)

	if string(data) != "this is test data" {
		t.Error("not match")
	}
}

func TestHttp(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		_, err := io.WriteString(w, "{\"status\":\"expected service response\"}")
		if err != nil {
			log.Fatal(err)
		}
	}

	req := httptest.NewRequest("GET", "https://tutorialedge.net", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	resp := w.Result()
	fmt.Println(resp.StatusCode)
	body, _ := ioutil.ReadAll(resp.Body)
	body1, _ := ioutil.ReadAll(w.Body)

	if string(body) != string(body1) {
		t.Error("body not the same")
	}
}
