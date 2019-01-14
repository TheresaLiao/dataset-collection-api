package main

import (
	"io/ioutil"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	// "github.com/op/go-logging"
	"log"
)

var log = logging.MustGetLogger("main")
type HttpResp struct {
	StatusCode int
	Context    string
}

func main() {
	//init log
	// initLogSetting(logging.DEBUG)
	fmt.Println("start api")
	router := gin.Default()

	//GET Default version
	router.GET("/", check)

	// filter fun.
	router.POST("/filerfun/detectImg", postDetectImgHandler)

	// dataset
	//router.POST("/dataSet/crawingCarAcdnt", crawingCarAcdntHandler)

	router.Run(":80")
}


func check(c *gin.Context) {
	c.String(http.StatusOK, "apiserver ready!")
}



func convertBody2Str(resp *http.Response) (context string) {
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Info("Error ioutil.ReadAll")
		log.Info(string(data))
		return
	}
	log.Info(string(data))
	return string(data)
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
