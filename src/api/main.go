package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
)

// Test db
//const connStr = "postgres://admin:12345@140.96.0.34:50003/Test_db?sslmode=disable"
//const connStr = "postgres://admin:12345@Test_Postgres:5432/Test_db?sslmode=disable"

const connStr = "postgres://admin:12345@task5-4-0-TH:5432/Test_db?sslmode=disable"

// production
//const connStr = "postgres://admin:12345@MyPostgres:5432/database_project?sslmode=disable"

var log = logging.MustGetLogger("main")

var whiteip1 = "140.96.29.153"
var whiteip2 = "140.96.29.211"
var whiteip3 = "140.96.29.202"

var whiteip4 = "210.61.209.193"
var whiteip5 = "210.61.209.194"
var whiteip6 = "210.61.209.195"
var whiteip7 = "210.61.209.196"
var whiteip8 = "210.61.209.197"

// Div-N Vincent
var whiteip9 = "140.96.116.61"
var whiteip10 = "140.96.158.25"

var cntConnection = 0

type HttpResp struct {
	StatusCode int
	Context    string
}
// @title Dataset Collection API
// @version 1.0
// @description This api call data from youtube anf table

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 10.201.252.7:30014
// 10.174.61.1:30017 for Div-X
func main() {
	// initLogSetting(logging.DEBUG)
	router := gin.Default()

	// Check client ip is accept to connect
	whitelist := make(map[string]bool)
	whitelist[whiteip1] = true
	whitelist[whiteip2] = true
	whitelist[whiteip3] = true
	whitelist[whiteip4] = true
	whitelist[whiteip5] = true
	whitelist[whiteip6] = true
	whitelist[whiteip7] = true
	whitelist[whiteip8] = true
	whitelist[whiteip9] = true
	whitelist[whiteip10] = true
	//router.Use(IPWhiteList(whitelist))

	// GET Default version
	router.GET("/", check)

	// dataset list
	router.GET("/dataset/list",GetDatasetList) 							//summary.go
	
	// get data list by filter parameter
	router.POST("/filterfun/detectImg", PostDetectImgHandler) 			//detectImage.go
	router.POST("/filterfun/youtubeUrl", Url2file) 						//filterfun.go
	// router.GET("/filterfun/youtubeInfo/:youtubeId", getYoutubeInfoById) //carType.go

	//========================= dataset list =========================
	// dataset subtitle 
	// TO-DO add srt mapping table
	router.GET("/dataset/subtitle", QuerySubtitleTagHandler) 							//subTitle.go
	router.GET("/dataset/subtitle/:subtitleTagId", QuerySubtitleBySubtitleTagIdHandler) //subTitle.go
	router.GET("/dataset/youtubeUrl/subtitle/:subtitleTagId", Url2DownloadSubtitleTag) 	//subTitle.go
	router.GET("/dataset/youtubeUrl/subtitleById/:subtitleId", Url2DownloadSubtitleId) 	//subTitle.go
	router.GET("/dataset/subTitleThumbnail", GetSubTitleThumbnail) 					//subTitle.go

	// dataset car type
	router.GET("/dataset/queryTrainTwOrg",QueryTrainTwOrgHandler) 				//carType.go
	router.GET("/dataset/url2DownloadTrainTwOrg",Url2DownloadTrainTwOrg) 		//carType.go
	router.GET("/dataset/youtubeUrl/cartype/:youtubeId",Url2DownloadCarType)	//carType.go
	router.GET("/dataset/queryTrainTwOrg/getThumbnail", GetTrainTwOrgThumbnail) //carType.go
	router.GET("/filterfun/youtubeUrl/getSearchByKeyWord", GetSearchByKeyWord)  //carType.go

	// dataset car accident
	// TO-DO key word
	router.GET("/dataset/caracdnt", QueryCarAccidentTagHandler) 				//carAccident.go
	router.GET("/dataset/caracdnt/:carAccidentId", QueryCarAccidentByIdHandler) //carAccident.go
	
	//========================= after training =========================
	// get yolo resualt from car type
	router.GET("/dataset/queryTrainYoloTag/:youtubeId",QueryTrainYoloTagByYoutubeIdHandler) //carType.go
	router.GET("/filterfun/trainTwOrg2TrainYolo/:youtubeId",TrainTwOrg2TrainYolo) 			//filterfun.go
	router.GET("/filterfun/parsingTrainYoloResult",ParsingTrainYoloResult) 					//filterfun.go
	router.POST("/filterfun/getYoloImg",GetYoloImg)											//detectImage.go

	// get lpr resualt from car type
	router.GET("/dataset/queryTrainLprTag/:youtubeId", QueryTrainLprTagByYoutubeIdHandler) 	//carType.go
	router.GET("/filterfun/trainTwOrg2TrainLpr/:youtubeId", TrainTwOrg2TrainLpr) 			//filterfun.go
	router.GET("/filterfun/parsingTrainLprResult", ParsingTrainLprResult) 					//filterfun.go
	router.POST("/filterfun/getLprImg", GetLprImg)											//detectImage.go

	router.Run(":80")
}

func check(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "apiserver ready and Summary Connection ",
	})
}

func IPWhiteList(whitelist map[string]bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.ClientIP())
		if !whitelist[c.ClientIP()] {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"status":  http.StatusForbidden,
				"message": "Permission denied",
			})
			return
		}
		cntConnection++
	}
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

// func LoadConfiguration() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		configFile, err := os.Open(configPath)
// 		defer configFile.Close()
// 		if err != nil {
// 			log.Errorf("Error load config file :", configPath)
// 		}
// 		jsonParser := json.NewDecoder(configFile)
// 		decodeErr := jsonParser.Decode(&configPara)
// 		if decodeErr != nil {
// 			log.Errorf("Decode fail : ", decodeErr)
// 			log.Errorf("JsonParser fail : ", jsonParser)
// 		}
// 		log.Info("init config :" + configPara.YamlPath)
// 		c.Next()
// 	}

//init log
// func initLogSetting(level logging.Level) {

// 	//Setup Console format
// 	var consoleFormat = logging.MustStringFormatter(
// 		`%{color}%{time:2006-01-02 15:04:05.000} %{level:.4s} > %{shortfile} [%{shortfunc}] %{message}%{color:reset}`,
// 	)

// 	consoleBackend := logging.NewLogBackend(os.Stderr, "", 0)
// 	consoleBackendFormatter := logging.NewBackendFormatter(consoleBackend, consoleFormat)
// 	consoleBackendLeveled := logging.AddModuleLevel(consoleBackendFormatter)
// 	consoleBackendLeveled.SetLevel(level, "")

// 	//Setup file format
// 	var fileFormat = logging.MustStringFormatter(
// 		`%{time:2006-01-02 15:04:05.000} %{level:.4s} > %{shortfile} [%{shortfunc}] %{message}`,
// 	)

// 	// Write log file
// 	//createFolder(logPath)
// 	now := time.Now()
// 	nowStr := now.Format(configPara.FormatStr2)
// 	fileName := logPath + "/" + nowStr + ".log"
// 	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
// 	if err != nil {
// 		log.Errorf("error opening file: %v", err)
// 	}
// 	defer file.Close()

// 	//create log rolling
// 	f := &lumberjack.Logger{
// 		Filename:   fileName,
// 		MaxSize:    500, // megabytes
// 		MaxBackups: 3,
// 		MaxAge:     28, //days
// 	}
// 	fileBackend := logging.NewLogBackend(f, "", 0)
// 	fileBackendFormatter := logging.NewBackendFormatter(fileBackend, fileFormat)
// 	fileBackendLeveled := logging.AddModuleLevel(fileBackendFormatter)
// 	fileBackendLeveled.SetLevel(level, "")

// 	// Set the backends to be used.
// 	logging.SetBackend(consoleBackendLeveled, fileBackendLeveled)
// }



