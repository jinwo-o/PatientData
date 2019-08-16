package tests

import (
	"os"
	"testing"
)

var host = "http://127.0.0.1:8080"

func TestMain(m *testing.M) {
	testResults := m.Run()
	os.Exit(testResults)
}

func TestUrl(t *testing.T) {
	result := true
	result = result && CheckPageResponse(host+"/patients")
	if result != true {
		t.Error("Web Pages Not Successful")
	}
}
