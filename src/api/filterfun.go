package main

import (
	"os"
	"path/filepath"
	"strings"
	"github.com/gin-gonic/gin"
	. "github.com/kkdai/youtube"
	_ "github.com/lib/pq"
	"database/sql"
	"archive/tar"
	"compress/gzip"
	"path"
	"io"
)

const DOWNLOADS_PATH = "/tmp"

type YoutubeInfoVo struct {
	FileName string `form:"filename" json:"filename" binding:"required"`
	Url string `form:"url" json:"url" binding:"required"`
}
// func url2DownloadSubtitle(c *gin.Context){
// }


func compressTar(c *gin.Context){
	parentFolderName := "caracdnt"
	srcDirPath := filepath.Join(DOWNLOADS_PATH , parentFolderName)
	destFilePath := filepath.Join(DOWNLOADS_PATH , "des")

	log.Info("srcDirPath : " + srcDirPath)
	log.Info("destFilePath : " + destFilePath)

	fw, err := os.Create(destFilePath)
    if err != nil{
		panic(err)
	}
	defer fw.Close()
	
	gw := gzip.NewWriter(fw)
	defer gw.Close()
	
	tw := tar.NewWriter(gw)
    defer tw.Close()
	tarGzDir(srcDirPath, path.Base(srcDirPath), tw)
}

func tarGzDir(srcDirPath string, recPath string, tw *tar.Writer) {
	log.Info("start tarGzDir")

	log.Info("srcDirPath : "+srcDirPath)
	log.Info("recPath : "+recPath)

    // Open source diretory
    dir, err := os.Open(srcDirPath)
	if err != nil{
		panic(err)
	}
    defer dir.Close()
    // Get file info slice
    fis, err := dir.Readdir(0)
    if err != nil{
		panic(err)
	}
    for _, fi := range fis {
        // Append path
		curPath := srcDirPath + "/" + fi.Name()
		log.Info("curPath : "+curPath)
        // Check it is directory or file
        if fi.IsDir() {
            // Directory
            // (Directory won't add unitl all subfiles are added)
            log.Info("Adding path...%s\\n", curPath)
            tarGzDir(curPath, recPath+"/"+fi.Name(), tw)
        } else {
            // File
            log.Info("Adding file...%s\\n", curPath)
        }
        tarGzFile(curPath, recPath+"/"+fi.Name(), tw, fi)
    }
}

func tarGzFile(srcFile string, recPath string, tw *tar.Writer, fi os.FileInfo) {

	log.Info("start tarGzFile")

    if fi.IsDir() {
        // Create tar header
        hdr := new(tar.Header)
        // if last character of header name is '/' it also can be directory
        // but if you don't set Typeflag, error will occur when you untargz
        hdr.Name = recPath + "/"
        hdr.Typeflag = tar.TypeDir
        hdr.Size = 0
        //hdr.Mode = 0755 | c_ISDIR
        hdr.Mode = int64(fi.Mode())
        hdr.ModTime = fi.ModTime()

        // Write hander
        err := tw.WriteHeader(hdr)
		if err != nil{
			panic(err)
		}
    } else {
        // File reader
        fr, err := os.Open(srcFile)
        if err != nil{
			panic(err)
		}
        defer fr.Close()
        // Create tar header
        hdr := new(tar.Header)
        hdr.Name = recPath
        hdr.Size = fi.Size()
        hdr.Mode = int64(fi.Mode())
        hdr.ModTime = fi.ModTime()

        // Write hander
        err = tw.WriteHeader(hdr)
		if err != nil{
			panic(err)
		}
        // Write file data
        _, err = io.Copy(tw, fr)
        if err != nil{
			panic(err)
		}
    }
}
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
	parentFolderName := "caracdnt"

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
			
			// download file from server to client
			if !strings.HasPrefix(filepath.Clean(targetPath), DOWNLOADS_PATH) {
				c.String(403, "Look like you attacking me")
			}   
        default:
           	checkError(err)
        }
	}
}

func url2Download(c *gin.Context){
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