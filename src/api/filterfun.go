package main

import (
	"fmt"
	"path/filepath"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	. "github.com/kkdai/youtube"
)

const DOWNLOADS_PATH = "/tmp"
type YoutubeDownloadVo struct {
	FileName string `form:"filename" json:"filename" binding:"required"`
	Url string `form:"url" json:"url" binding:"required"`
}

func url2Localpath(c *gin.Context){
	var youtubeDownloadVo YoutubeDownloadVo
	c.BindJSON(&youtubeDownloadVo)

	fmt.Println("FileName :" + youtubeDownloadVo.FileName)
	fmt.Println("url :" + youtubeDownloadVo.Url)

	// usr, _ := user.Current()
	// currentDir := fmt.Sprintf("%v/Movies/youtubedr", usr.HomeDir)
	currentDir := DOWNLOADS_PATH
	fmt.Println("download to dir=", currentDir)
	
	y := NewYoutube(true)
	arg := youtubeDownloadVo.Url
	if err := y.DecodeURL(arg); err != nil {
		fmt.Println("err:", err)
	}

	if err := y.StartDownload(filepath.Join(currentDir, youtubeDownloadVo.FileName + ".mp4")); err != nil {
		fmt.Println("err:", err)
	}

	fmt.Println("VideoID = "+ y.VideoID)
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "apiserver ready and Summary Connection ",
	})
}

func download(c *gin.Context){
	fmt.Println("=========download=========")
	fileName := c.Param("filename")
    targetPath := filepath.Join(DOWNLOADS_PATH , "/" , fileName)
    //This ckeck is for example, I not sure is it can prevent all possible filename attacks - will be much better if real filename will not come from user side. I not even tryed this code
    if !strings.HasPrefix(filepath.Clean(targetPath), DOWNLOADS_PATH) {
        c.String(403, "Look like you attacking me")
        return
    }
    //Seems this headers needed for some browsers (for example without this headers Chrome will download files as txt)
    c.Header("Content-Description", "File Transfer")
    c.Header("Content-Transfer-Encoding", "binary")
    c.Header("Content-Disposition", "attachment; filename="+fileName )
    c.Header("Content-Type", "application/octet-stream")
    c.File(targetPath)
}


