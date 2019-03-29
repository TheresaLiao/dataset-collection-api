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

func url2DownloadCaracdnt(c *gin.Context){
	carAccidentTagIdStr := c.Param("carAccidentTagId")

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
	sql_statement := "SELECT * FROM car_accident WHERE id in (SELECT car_accident_id FROM car_accident_tag_map WHERE car_accident_tag_id =" + carAccidentTagIdStr + ") order by id LIMIT 3;"
    rows, err := db.Query(sql_statement)
    checkError(err)
	defer rows.Close()

	var id int
	var title string
	var url string
	var copyRight string
	var accidentTime string
	var carType string
	var dayTime string
	var collision string
	parentFolderName := "caracdnt_"+carAccidentTagIdStr

	for rows.Next() {
		switch err := rows.Scan(&id, &title, &url,
			&copyRight, &accidentTime, &carType,
			&dayTime, &collision); err {
        case sql.ErrNoRows:
			log.Info("No rows were returned")
		case nil:
			// check url
			y := NewYoutube(true)
				if err := y.DecodeURL(url); err != nil {
				log.Error("err ", err)
			}
			log.Info("VideoID1 = "+ y.VideoID)

			// download file to localpath
			targetPath := filepath.Join(DOWNLOADS_PATH , parentFolderName ,  y.VideoID + ".mp4")
			log.Info("targetPath : " + targetPath)
			if err := y.StartDownload(targetPath); err != nil {
				log.Info("Error StartDownload")
				log.Error("err ", err)
			}
			
			
        default:
           	checkError(err)
        }
	}
	srcDirPath := filepath.Join(DOWNLOADS_PATH , parentFolderName)
	destFilePath := filepath.Join(DOWNLOADS_PATH , parentFolderName+".tar.gz")
	tarDir(srcDirPath,destFilePath)

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

func url2file(c *gin.Context){
	//dowmload file from url to server 
	var youtubeInfoVo YoutubeInfoVo
	c.BindJSON(&youtubeInfoVo)
	log.Info("Parameter FileName :" + youtubeInfoVo.FileName)
	log.Info("Parameter url :" + youtubeInfoVo.Url)
	log.Info("default path : ", DOWNLOADS_PATH)
	
	// check url
	y := NewYoutube(true)
	arg := youtubeInfoVo.Url
	if err := y.DecodeURL(arg); err != nil {
		log.Error("err ", err)
	}

	// download file to localpath
	targetPath := filepath.Join(DOWNLOADS_PATH, youtubeInfoVo.FileName + ".mp4")
	log.Info("targetPath : " + targetPath)
	if err := y.StartDownload(targetPath); err != nil {
		log.Info("Error StartDownload")
		log.Error("err ", err)
	}
	log.Info("VideoID = "+ y.VideoID)
	
	// download file from server to client
    if !strings.HasPrefix(filepath.Clean(targetPath), DOWNLOADS_PATH) {
		c.String(403, "Look like you attacking me")
	}
	
    c.Header("Content-Description", "File Transfer")
    c.Header("Content-Transfer-Encoding", "binary")
    c.Header("Content-Disposition", "attachment; targetPath=" + targetPath )
    c.Header("Content-Type", "application/octet-stream")
	c.File(targetPath)

	// Remove download file
	if err := os.Remove(targetPath) ; err != nil {
        log.Info("file remove Error!")
        log.Error("err ", err)
    } else {
        log.Info("file remove OK!")
    }
}