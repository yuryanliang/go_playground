package main

import (
	"io/ioutil"
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
