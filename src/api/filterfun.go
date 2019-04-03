package main

import (
	"os"
	"path/filepath"
	"strings"
	"github.com/gin-gonic/gin"
	. "github.com/kkdai/youtube"
	_ "github.com/lib/pq"
	"database/sql"
)
type YoutubeInfoVo struct {
	FileName string `form:"filename" json:"filename" binding:"required"`
	Url string `form:"url" json:"url" binding:"required"`
}
const DOWNLOADS_PATH = "/tmp"
const SUBTITLE_PATH = "subtitle_"
const CARACDNT_PATH = "caracdnt_"
const TEMP_PATH = "temp"
const VIEDO_PATH = "viedo"
const FILE_EXTENTION_TAR = ".tar.gz"
const FILE_EXTENTION_MP4 = ".mp4"


func url2DownloadSubtitle(c *gin.Context){
	subtitleTagIdStr := c.Param("subtitleTagId")
	// parentFolderName : subtitle_N , ex. subtitle_1,subtitle_2...
	parentFolderName := SUBTITLE_PATH + subtitleTagIdStr
	// srcDirPath : /tmp/subtitle_N
	srcDirPath := filepath.Join(DOWNLOADS_PATH, parentFolderName)
	// srcDirPathViedo : /tmp/subtitle_N/viedo
	srcDirPathViedo := filepath.Join(srcDirPath, VIEDO_PATH)
	// srcDirPath : /tmp/subtitle_N.tar.gz
	destFilePath := filepath.Join(DOWNLOADS_PATH , parentFolderName + FILE_EXTENTION_TAR)

	if checkFileIsExit(destFilePath) == false{
		if checkFileIsExit(srcDirPathViedo) == false{
			// query data from sql, than download file
			querySubtitle(subtitleTagIdStr,srcDirPathViedo)
		}
		// tar download folder
		tarDir(srcDirPath,destFilePath)
	}
	// download file from server to client
	respFile2Client(c,destFilePath)
}

func url2DownloadCaracdnt(c *gin.Context){
	carAccidentTagIdStr := c.Param("carAccidentTagId")
	// parentFolderName : caracdnt_N , ex. caracdnt_1,caracdnt_2...
	parentFolderName := CARACDNT_PATH + carAccidentTagIdStr
	// srcDirPath : /tmp/caracdnt_N
	srcDirPath := filepath.Join(DOWNLOADS_PATH, parentFolderName)
	// srcDirPathViedo : /tmp/caracdnt_N/viedo
	srcDirPathViedo := filepath.Join(srcDirPath, VIEDO_PATH)
	// srcDirPath : /tmp/caracdnt_N.tar.gz
	destFilePath := filepath.Join(DOWNLOADS_PATH , parentFolderName + FILE_EXTENTION_TAR)

	if checkFileIsExit(destFilePath) == false{
		if checkFileIsExit(srcDirPathViedo) == false{
			// query data from sql, than download file
			queryCaracdnt(carAccidentTagIdStr,srcDirPathViedo)
		}
		// tar download folder
		tarDir(srcDirPath,destFilePath)
	}
	// download file from server to client
	respFile2Client(c,destFilePath)
}

func url2file(c *gin.Context){
	//dowmload file from url to server 
	var youtubeInfoVo YoutubeInfoVo
	c.BindJSON(&youtubeInfoVo)
	log.Info("Parameter FileName :" + youtubeInfoVo.FileName)
	log.Info("Parameter url :" + youtubeInfoVo.Url)
	
	srcDirPath := filepath.Join(DOWNLOADS_PATH, TEMP_PATH)
	videoID := checkUrlAndDownload(youtubeInfoVo.Url, srcDirPath)
	
	// download file from server to client
	destFilePath := filepath.Join(srcDirPath, videoID + FILE_EXTENTION_MP4)
	respFile2Client(c,destFilePath)

	// Remove download file
	if err := os.Remove(destFilePath) ; err != nil {
        log.Info("file remove Error!")
        log.Error("err ", err)
    } else {
        log.Info("file remove OK!")
    }
}

func checkFileIsExit(filepath string)(resp bool){ 
	log.Info("filepath : "+filepath)

	if _, err := os.Stat(filepath); err == nil {
		log.Info(filepath + " is exeit")
		return true
	}
	return false
}

func checkUrlAndDownload(url string,srcDirPath string)(videoID string){
	log.Info("url : "+url)
	log.Info("srcDirPath : "+ srcDirPath)

	// check url
	y := NewYoutube(true)
	arg := url
	if err := y.DecodeURL(arg); err != nil {
		log.Error("err ", err)
	}

	// download file to localpath
	filename := filepath.Join(srcDirPath, y.VideoID + FILE_EXTENTION_MP4)
	log.Info("filename : " + filename)
	if err := y.StartDownload(filename); err != nil {
		log.Info("Error StartDownload")
		log.Error("err ", err)
	}
	log.Info("VideoID = "+ y.VideoID)
	return y.VideoID
}

func querySubtitle(subtitleTagIdStr string,srcDirPath string){
	log.Info("subtitleTagIdStr : "+subtitleTagIdStr)
	log.Info("srcDirPath : "+ srcDirPath)

	// connect db
	db, err := sql.Open("postgres",connStr)
	if err != nil{
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil{
		panic(err)
	}
	log.Info("success connection")

	// select table :subtitle_tag ,all rows data
	sql_statement := "SELECT id, title, url  FROM subtitle WHERE id in (SELECT subtitle_id FROM subtitle_tag_map WHERE subtitle_tag_id =" + subtitleTagIdStr + ")  limit 3;"
    rows, err := db.Query(sql_statement)
    checkError(err)
	defer rows.Close()

	var id int
	var title string
	var url string

	for rows.Next() {
		switch err := rows.Scan(&id, &title,  &url); err {
        case sql.ErrNoRows:
			log.Info("No rows were returned")
		case nil:			
			checkUrlAndDownload(url, srcDirPath)
        default:
           	checkError(err)
        }
	}
}

func queryCaracdnt(carAccidentTagIdStr string,srcDirPath string){
	log.Info("carAccidentTagIdStr : "+carAccidentTagIdStr)
	log.Info("srcDirPath : "+ srcDirPath)

	// connect db
	db, err := sql.Open("postgres",connStr)
	if err != nil{
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil{
		panic(err)
	}
	log.Info("success connection")

	// select table :subtitle_tag ,all rows data
	sql_statement := "SELECT id, title, url FROM car_accident WHERE id in (SELECT car_accident_id FROM car_accident_tag_map WHERE car_accident_tag_id =" + carAccidentTagIdStr + ")  limit 3;"
    rows, err := db.Query(sql_statement)
    checkError(err)
	defer rows.Close()

	var id int
	var title string
	var url string

	for rows.Next() {
		switch err := rows.Scan(&id, &title, &url); err {
        case sql.ErrNoRows:
			log.Info("No rows were returned")
		case nil:
			checkUrlAndDownload(url, srcDirPath)
        default:
           	checkError(err)
        }
	}
}

func respFile2Client(c *gin.Context,destFilePath string){
	log.Info("destFilePath : "+ destFilePath)

	// download file from server to client
	if !strings.HasPrefix(filepath.Clean(destFilePath), DOWNLOADS_PATH) {
		c.String(403, "Look like you attacking me")
   	}   
   	c.Header("Content-Description", "File Transfer")
   	c.Header("Content-Transfer-Encoding", "binary")
   	c.Header("Content-Disposition", "attachment; destFilePath=" + destFilePath )
   	c.Header("Content-Type", "application/octet-stream")
   	c.File(destFilePath)
}
