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
	"strings"
	"io/ioutil"
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

func tarDir(srcDirPath string, destFilePath string){
	 log.Info("start tarDir")
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
 
	 //srcDirPath: /tmp/caracdnt_1
	 walk(srcDirPath, srcDirPath, tw)
	 log.Info("tar.gz ok")
}


func walk(path string, prefix string, tw *tar.Writer) error {
	log.Info("path :" + path)

	// get file
    files, err := os.Open(path)
    if err != nil {
        return err
    }
    defer files.Close()
	
	file, err := files.Stat()
    if err != nil {
        return err
    }
	
	tarHeader, err := tar.FileInfoHeader(file, "")
    if err != nil {
        return err
    } 
	
	// remove Prefix path
	tarHeader.Name = strings.TrimPrefix(path, prefix) 
	log.Info("file name :" + tarHeader.Name)
    if err = tw.WriteHeader(tarHeader); err != nil {
        return err
    }
	
	// if file is folder ,search son file systems
	if file.IsDir() {                
		log.Info("Is Directory")
		// load dir files
		sonFiles, err := ioutil.ReadDir(path)
        if err != nil {
            return err
		}
		for _, sonFile := range sonFiles{
			// get filePath
			sonFilePath := filepath.Join(path , sonFile.Name())
            if err = walk(sonFilePath, prefix, tw); err != nil {
                return err
            }
        }
        return nil
	}

	// else , copy file into tar(tw)
	log.Info("Is File")
	_, err = io.Copy(tw, files) 
	
    return err
}