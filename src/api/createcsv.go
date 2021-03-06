package main

import (
    "os"
    "encoding/csv"
)

func getcsv(header []string,records [][]string, destFilePath string){
    log.Info("getcsv")
    file, err := os.Create(destFilePath)
    if err != nil {
		  log.Info("Cannot write to file", err)
	  }
    defer file.Close()
	
    writer := csv.NewWriter(file)
    defer writer.Flush()

    writer.Write(header) 
    if err := writer.Error(); err != nil {
		log.Info("error Write csv:", err)
	}
    writer.WriteAll(records) 
    if err := writer.Error(); err != nil {
		log.Info("error WriteAll csv:", err)
	}
}