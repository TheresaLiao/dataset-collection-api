package main

import (
	"testing"
	"net/http"
	//"net/http/httptest"
	//"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)
var testUrl = "http://localhost:80"

func TestMain(t *testing.T) {
	t.Log("TestMain PASS")
}

func TestCheck(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, testUrl+ "/", nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestPostDetectImgHandler(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, testUrl+ "/", nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}