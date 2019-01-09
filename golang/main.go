package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
	"gopkg.in/natefinch/lumberjack.v2"
)

const logPath = "./utils"
const configPath = "./config.json"
const urlSslHeader = "https://"
const urlHeader = "http://"

var configPara ConfigPara
var log = logging.MustGetLogger("main")

type HttpResp struct {
	StatusCode int
	Context    string
}

type VersionObj struct {
	Version string `form:"version" json:"version" binding:"required"`
}

type ConfigPara struct {
	Version              string   `json:"version"`
	TempServiceYaml      string   `json:"tempServiceYaml"`
	TempPodYaml          string   `json:"tempPodYaml"`
	YamlPath             string   `json:"yamlPath"`
	FileNameServiceYaml  string   `json:"fileNameServiceYaml"`
	FileNamePodYaml      string   `json:"fileNamePodYaml"`
	ApiHost              string   `json:"apiHost"`
	ApiServices          string   `json:"apiServices"`
	ApiPod               string   `json:"apiPod"`
	FormatStr1           string   `json:"formatStr1"`
	FormatStr2           string   `json:"formatStr2"`
	TokenPath            string   `json:"tokenPath"`
	HttpHeaderTypeKey    string   `json:"httpHeaderTypeKey"`
	HttpHeaderTypeYaml   string   `json:"httpHeaderTypeYaml"`
	NfsHost              string   `json:"nfsHost"`
	NfsPort              string   `json:"nfsPort"`
	NfsUrl               string   `json:"nfsUrl"`
	RegistryDnsName      string   `json:"registryDnsName"`
	RegistryIntPort      string   `json:"registryIntPort"`
	RegistryExtPort      string   `json:"registryExtPort"`
	RegistryList         []string `json:"registryList"`
	RegistryUrlVer       string   `json:"registryUrlVer"`
	RegistryUrlRepo      string   `json:"registryUrlRepo"`
	RegistryUrlTags      string   `json:"registryUrlTags"`
	RegistryUrlManifests string   `json:"registryUrlManifests"`
}

func main() {
	//init log
	initLogSetting(logging.DEBUG)

	router := gin.Default()
	router.Use(LoadConfiguration())

	//GET Default version
	router.GET("/", showVersion)

	//POST(C)
	router.POST("/kubeGpu/container", createContainerHandler)

	//Single
	router.GET("/kubeGpu/container/:machineId", getContainerHandler)
	router.DELETE("/kubeGpu/container/:machineId", deleteContainerHandler)

	//All
	router.GET("/kubeGpu/containers", getContainersHandler)
	router.DELETE("/kubeGpu/containers", deleteContainersHandler)

	//volumn
	router.POST("/kubeGpu/volumn", createVolumnHandler)
	router.DELETE("/kubeGpu/volumn/:account", deleteVolumnHandler)

	//volumn
	router.GET("/kubeGpu/images", getImagesListHandler)
	router.GET("/kubeGpu/image/:repository", getImageHandler)

	router.Run(":8000")
}

func LoadConfiguration() gin.HandlerFunc {
	return func(c *gin.Context) {
		configFile, err := os.Open(configPath)
		defer configFile.Close()
		if err != nil {
			log.Errorf("Error load config file :", configPath)
		}
		jsonParser := json.NewDecoder(configFile)
		decodeErr := jsonParser.Decode(&configPara)
		if decodeErr != nil {
			log.Errorf("Decode fail : ", decodeErr)
			log.Errorf("JsonParser fail : ", jsonParser)
		}
		log.Info("init config :" + configPara.YamlPath)
		c.Next()
	}
}

func showVersion(c *gin.Context) {
	versionObj := VersionObj{configPara.Version}
	c.JSON(http.StatusOK, versionObj)
	return
}

//init log
func initLogSetting(level logging.Level) {

	//Setup Console format
	var consoleFormat = logging.MustStringFormatter(
		`%{color}%{time:2006-01-02 15:04:05.000} %{level:.4s} > %{shortfile} [%{shortfunc}] %{message}%{color:reset}`,
	)

	consoleBackend := logging.NewLogBackend(os.Stderr, "", 0)
	consoleBackendFormatter := logging.NewBackendFormatter(consoleBackend, consoleFormat)
	consoleBackendLeveled := logging.AddModuleLevel(consoleBackendFormatter)
	consoleBackendLeveled.SetLevel(level, "")

	//Setup file format
	var fileFormat = logging.MustStringFormatter(
		`%{time:2006-01-02 15:04:05.000} %{level:.4s} > %{shortfile} [%{shortfunc}] %{message}`,
	)

	//fileName := filepath.Base(os.Args[0]) + ".log"

	// Write log file
	//createFolder(logPath)
	now := time.Now()
	nowStr := now.Format(configPara.FormatStr2)
	fileName := logPath + "/" + nowStr + ".log"
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	//create log rolling
	f := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
	}
	fileBackend := logging.NewLogBackend(f, "", 0)
	fileBackendFormatter := logging.NewBackendFormatter(fileBackend, fileFormat)
	fileBackendLeveled := logging.AddModuleLevel(fileBackendFormatter)
	fileBackendLeveled.SetLevel(level, "")

	// Set the backends to be used.
	logging.SetBackend(consoleBackendLeveled, fileBackendLeveled)
}
