package main

import (
    "os"
    "encoding/csv"
)

func getcsv(records [][]string, destFilePath string){
    file, err := os.Create(destFilePath)
    if err != nil {
		  log.Info("Cannot write to file", err)
	  }
    defer file.Close()
	
    writer := csv.NewWriter(file)
    defer writer.Flush()

    writer.WriteAll(records) 
    if err := writer.Error(); err != nil {
		log.Info("error writing csv:", err)
	}
}