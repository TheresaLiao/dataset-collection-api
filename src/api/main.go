package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
)

// Test db
//const connStr = "postgres://admin:12345@140.96.0.34:50003/Test_db?sslmode=disable"
const connStr = "postgres://admin:12345@Test_MyPostgres:5432/Test_db?sslmode=disable"

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

var cntConnection = 0

type HttpResp struct {
	StatusCode int
	Context    string
}

func main() {
	//init log
	// initLogSetting(logging.DEBUG)
	fmt.Println("start api")
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

	//router.Use(IPWhiteList(whitelist))

	//GET Default version
	router.GET("/", check)
	
	// get data list by filter parameter
	router.POST("/filterfun/detectImg", postDetectImgHandler)
	router.POST("/filterfun/youtubeUrl", url2file)

	router.GET("/filterfun/youtubeUrl/subtitle/:subtitleTagId", url2DownloadSubtitle)
	router.GET("/filterfun/youtubeUrl/caracdnt/:carAccidentTagId", url2DownloadCaracdnt)

	// dataset for download
	// dataset subtitle
	router.GET("/dataset/subtitle", querySubtitleTagHandler)
	router.GET("/dataset/subtitle/:subtitleTagId", querySubtitleBySubtitleTagIdHandler)

	// dataset car accident
	router.GET("/dataset/caracdnt", queryCarAccidentTagHandler)
	router.GET("/dataset/caracdnt/:carAccidentTagId", queryCarAccidentByCarAccidentTagIdHandler)

	// dataset car type
	router.GET("/filterfun/url2DownloadTrainTwOrg",url2DownloadTrainTwOrg)
	router.GET("/dataset/queryTrainTwOrg",queryTrainTwOrgHandler)

	// get yolo resualt
	router.GET("/filterfun/parsingTrainYoloResult",parsingTrainYoloResult)
	router.GET("/dataset/queryTrainYoloTag/:youtubeId",queryTrainYoloTagByYoutubeIdHandler)
	router.GET("/filterfun/getYoloImg/:youtubeId/:filename",getYoloImg)

	// get lpr resualt
	router.GET("/filterfun/parsingTrainLprResult",parsingTrainLprResult)
	router.GET("/dataset/queryTrainLprTag/:youtubeId",queryTrainLprTagByYoutubeIdHandler)
	router.GET("/filterfun/getLprImg/:youtubeId/:filename",getLprImg)

	//router.GET("/dataset/trans",queryTransHandler)
	//router.GET("/dataset/trans/:transType",queryTransByTranstypeHandler)


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

func convertBody2Str(resp *http.Response) (context string) {
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Info("Error ioutil.ReadAll")
		log.Info(string(data))
		return
	}
	//log.Info(string(data))
	return string(data)
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



