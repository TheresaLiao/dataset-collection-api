package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
	"path/filepath"
)

var detectImgUrl = "http://task5-4-4-TH:8080/yolo_coco_image"

// PostDetectImgHandler example
// @Summary detect Imgage object
// @Tags get data list by filter parameter
// @Description post by binary image output json string
// @Accept  multipart/form-data
// @Produce  json
// @Success 200	{object}	main.YoloItem	"ok"
// @Router /filterfun/detectImg [post]
/*
 * input binary 
 * output json file
 */
// curl -X POST \
// --data-binary "@/file_path" \
// http://localhost:port/filterfun/detectImg
func PostDetectImgHandler(c *gin.Context) {
	log.Info("postDetectImgHandler")
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

// GetYoloImg example
// @Summary download yolo detect image file
// @Description 
// @Tags yolo resualt
// @ID get-yolo-img
// @Accept  json
// @Produce  json
// @Param   filename	path	string	true  "File Name"
// @Param   youtubeId	body	string	true  "Youtube Id"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} string "We need File Name!!"
// @Failure 404 {object} string "We need Youtube Id!!"
// @Router /filterfun/getYoloImg [post]
// curl --header "Content-Type: application/json" \
// --request POST \
// --data '{"filename":"res_00011441.jpg","youtubeId":"zoqVFEuVPJY"}' \
// http://localhost:port/filterfun/getYoloImg \
// --output res_00011441.mp4
func GetYoloImg(c *gin.Context){
	log.Info("getYoloImg")	
	var getImageVo GetImageVo
	c.BindJSON(&getImageVo)
	log.Info("Parameter FileName :" + getImageVo.FileName)
	log.Info("Parameter YoutubeId :" + getImageVo.YoutubeId)
	
	// srcDirPath : /tmp/traintworg
	srcDirPath := filepath.Join(DOWNLOADS_PATH, TRAINTWORG_PATH)
	// srcDirPathViedo : /tmp/traintworg/viedo
	srcDirPathViedo := filepath.Join(srcDirPath, VIEDO_PATH)
	srcFilePathViedo := filepath.Join(srcDirPathViedo, getImageVo.YoutubeId + YOLO_FOLDER, getImageVo.FileName)
	
	log.Info("srcFilePathViedo :" + srcFilePathViedo)
	content, err := ioutil.ReadFile(srcFilePathViedo)
	if err != nil{
		log.Info(err)
	}

	c.Header("Access-Control-Allow-Origin", "*") 
	c.Data(http.StatusOK, "text/html; charset=utf-8", content)
}

// GetLprImg example
// @Summary download lpr detect image file
// @Description 
// @Tags lpr resualt
// @ID get-lpr-img
// @Accept  json
// @Produce  json
// @Param   filename	path	string	true  "File Name"
// @Param   youtubeId	body	string	true  "Youtube Id"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} string "We need File Name!!"
// @Failure 404 {object} string "We need Youtube Id!!"
// @Router /filterfun/getLprImg [post]
// curl --header "Content-Type: application/json" \
// --request POST \
// --data '{"filename":"res_00011441.jpg","youtubeId":"zoqVFEuVPJY"}' \
// http://localhost:port/filterfun/getLprImg \
// --output res_00011441.mp4
func GetLprImg(c *gin.Context){
	log.Info("getYoloImg")	
	var getImageVo GetImageVo
	c.BindJSON(&getImageVo)
	log.Info("Parameter FileName :" + getImageVo.FileName)
	log.Info("Parameter YoutubeId :" + getImageVo.YoutubeId)
	
	// srcDirPath : /tmp/traintworg
	srcDirPath := filepath.Join(DOWNLOADS_PATH, TRAINTWORG_PATH)
	// srcDirPathViedo : /tmp/traintworg/viedo
	srcDirPathViedo := filepath.Join(srcDirPath, VIEDO_PATH)
	srcFilePathViedo := filepath.Join(srcDirPathViedo, getImageVo.YoutubeId + LPR_FOLDER,  getImageVo.FileName)
	
	log.Info(srcFilePathViedo)
	content, err := ioutil.ReadFile(srcFilePathViedo)
	if err != nil{
		log.Info(err)
	}

	c.Header("Access-Control-Allow-Origin", "*") 
	c.Data(http.StatusOK, "text/html; charset=utf-8", content)
}

func convertBody2Str(resp *http.Response) (context string) {
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Info("Error ioutil.ReadAll")
		log.Info(string(data))
		return
	}
	return string(data)
}
