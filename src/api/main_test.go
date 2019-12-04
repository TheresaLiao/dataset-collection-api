package main

import (
	"os"
	//"io"
	"bytes"
	"testing"
	"net/http"
	"io/ioutil"
	//"net/http/httptest"
	//"github.com/gin-gonic/gin"
	//"mime/multipart"
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
	filename := "./picture/dog.jpg"

	// get file
    data, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
	
	// Request
	req, err := http.NewRequest(http.MethodPost, testUrl+ "/filterfun/detectImg", data)
	if err != nil {
        log.Fatal(err)
	}
	req.Header.Add("Content-Type", "image/jpeg")
    client := &http.Client{}

	// response
	resp, err := client.Do(req)
    if err != nil {
        log.Fatal(err)
    }
	defer resp.Body.Close()
	
	// convert into sting
    content, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
	}
	t.Log(string(content))

	// 200 = 200
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestUrl2file(t *testing.T) {
	// Request
	var jsonStr = []byte(`{"filename":"test","url":"https://www.youtube.com/watch?v=-EWwmIZFBQ8"}`)
	req, err := http.NewRequest(http.MethodPost, testUrl+ "/filterfun/youtubeUrl", bytes.NewBuffer(jsonStr))
	if err != nil {
        log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	// response
	resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	t.Log(len(body))

	// 200 = 200
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGetDatasetList(t *testing.T) {
	// Request
	req, _ := http.NewRequest(http.MethodGet, testUrl+ "/dataset/list", nil)
	
	// response
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Log(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	t.Log(string(body))

	// 200 = 200
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}


func TestQuerySubtitleTagHandler(t *testing.T) {
	// Request
	req, _ := http.NewRequest(http.MethodGet, testUrl+ "/dataset/subtitle", nil)

	// response
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Log(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	t.Log(string(body))
	
	// 200 = 200
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestQuerySubtitleBySubtitleTagIdHandler(t *testing.T) {
	// Request
	req, _ := http.NewRequest(http.MethodGet, testUrl+ "/dataset/subtitle/3", nil)
	
	// response
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Log(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	t.Log(string(body))
	
	// 200 = 200
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}


func TestUrl2DownloadSubtitleTag(t *testing.T) {
	// Request
	req, _ := http.NewRequest(http.MethodGet, testUrl+ "/dataset/youtubeUrl/subtitle/3", nil)
	
	// response
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Log(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	t.Log(len(body))

	// 200 = 200
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
func TestUrl2DownloadSubtitleId(t *testing.T) {
	// Request
	req, _ := http.NewRequest(http.MethodGet, testUrl+ "/dataset/youtubeUrl/subtitleById/8", nil)	
	
	// response
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Log(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	t.Log(len(body))

	// 200 = 200
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
func TestGetSubTitleThumbnail(t *testing.T) {
	// Request
	req, _ := http.NewRequest(http.MethodGet, testUrl+ "/dataset/getSubTitleThumbnail", nil)
	
	// response
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Log(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	t.Log(string(body))

	// 200 = 200
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestQueryTrainTwOrgHandler(t *testing.T) {
	// Request
	req, _ := http.NewRequest(http.MethodGet, testUrl+ "/dataset/queryTrainTwOrg", nil)

	// response
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Log(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	t.Log(string(body))

	// 200 = 200
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

// func TestUrl2DownloadTrainTwOrg(t *testing.T) {
// 	// Request
// 	req, _ := http.NewRequest(http.MethodGet, testUrl+ "/dataset/url2DownloadTrainTwOrg", nil)
// 	// response
// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		t.Log(err)
// 	}
// 	defer resp.Body.Close()
// 	body, _ := ioutil.ReadAll(resp.Body)
// 	t.Log(string(body))

// 	// 200 = 200
// 	assert.Equal(t, http.StatusOK, resp.StatusCode)
// }

func TestUrl2DownloadCarType(t *testing.T) {
	// Request
	req, _ := http.NewRequest(http.MethodGet, testUrl+ "/dataset/youtubeUrl/cartype/0-7_nvNNdcM ", nil)
	// response
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Log(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	t.Log(len(body))
	// 200 = 200
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func Test(t *testing.T) {
	// Request
	req, _ := http.NewRequest(http.MethodGet, testUrl+ "/", nil)
	// response
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Log(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	t.Log(string(body))
	// 200 = 200
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}