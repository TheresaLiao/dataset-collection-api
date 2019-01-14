package main

import (
	
	"bytes"
	"io/ioutil"
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
)

var detectImgUrl = "http://140.96.0.34:50008/detection"

/*
 * input binary 
 * output json file
 */
func postDetectImgHandler(c *gin.Context) {
	log.Info("===================")
	log.Info("POST")
	log.Info("postDetectImgHandler")

	log.Info("call detection post")
	log.Info("detectImgUrl : " + detectImgUrl)
	httpDetectionApiPost := detectionApiPost(detectImgUrl, c)
	log.Info(strconv.Itoa(httpDetectionApiPost.StatusCode))
	
	// show return status
	c.String(httpDetectionApiPost.StatusCode, httpDetectionApiPost.Context)
	if httpDetectionApiPost.StatusCode != 200 {
		log.Errorf(httpDetectionApiPost.Context)
	}
	return
}

func detectionApiPost(apiUrl string,c *gin.Context)(httpResp HttpResp) {
	
	log.Info("start detectionApiPost")
	
	// convert body into buffer
	var bodyBytes []byte
	if c.Request.Body != nil {
    	bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}
	buffer := bytes.NewBuffer(bodyBytes)
	//log.Info(buffer)

	// set request setting 
	log.Info("apiUrl : " + apiUrl)
	req, err := http.NewRequest("POST", apiUrl, buffer)
	if err != nil {
		log.Errorf("Error NewRequest", err)
		return HttpResp{http.StatusUnauthorized, "call storage fail"}
	}
	req.Header.Set("Content-Type", "image/jpeg")
	
	log.Info("start Do")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Errorf("DefaultClient fail", err)
		return HttpResp{http.StatusUnauthorized, "call detectionApiPost fail"}
	}
	log.Info("call detectionApiPost work")
	log.Info(resp.StatusCode)
	defer resp.Body.Close()

	// convert return into string
	context := convertBody2Str(resp)
	log.Info("context: "+context)
	return HttpResp{resp.StatusCode, context}
}
