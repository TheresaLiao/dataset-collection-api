package main

import (
	"os"
	"path/filepath"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"archive/tar"
	"io"
	"net/http"
	"compress/gzip"
)

func compressTar(c *gin.Context){
	parentFolderName := "11"
	srcDirPath := filepath.Join(DOWNLOADS_PATH , parentFolderName)
	destFilePath := filepath.Join(DOWNLOADS_PATH , parentFolderName+".tar.gz")

	log.Info("srcDirPath : " + srcDirPath)
	log.Info("destFilePath : " + destFilePath)

	tarDir(srcDirPath,destFilePath)

	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "apiserver ready and Summary Connection ",
	})
}

func tarDir(srcDirPath string,destFilePath string){
	 // file write
	 fw, err := os.Create(destFilePath)
	 if err != nil {
		log.Error("err ", err)
	 }
	 defer fw.Close()
 
	 gw := gzip.NewWriter(fw)
	 defer gw.Close()
 
	 // tar write
	 tw := tar.NewWriter(gw)
	 defer tw.Close()
 
	 dir, err := os.Open(srcDirPath)
     if err != nil {
		log.Error("err ", err)
	 }
	 defer dir.Close()
	 files, err := dir.Readdir(0)
	 if err != nil {
		log.Error("err ", err)
	 }

	 for _, file := range files {	
	 	if file.IsDir() {
			 continue
	 	}
		 //open file
		filepath:= filepath.Join(dir.Name()  , file.Name())
		fr, err := os.Open(filepath)
		if err != nil {
			log.Error("err ", err)
	 	}
	 	defer fr.Close()
 
	 	tarHeader := new(tar.Header)
	 	tarHeader.Name = file.Name()
	 	tarHeader.Size = file.Size()
	 	tarHeader.Mode = int64(file.Mode())
	 	tarHeader.ModTime = file.ModTime()
 
	 	err = tw.WriteHeader(tarHeader)
	 	if err != nil {
			log.Error("err ", err)
	 	}
	 	if _, err = io.Copy(tw, fr); err != nil {	 	
			log.Error("err ", err)
	 	}
	 }
	 log.Info("tar.gz ok")
}
