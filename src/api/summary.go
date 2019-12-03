package main

import (
	
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"database/sql"
	"net/http"
)

func GetDatasetList(c *gin.Context){
	log.Info("GetDatasetList")

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

	sql_statement := `	SELECT "title", "desc", "api"
				   		FROM "dataset_summary" `

	rows, err := db.Query(sql_statement)
	checkError(err)
	defer rows.Close()

	var title string
	var desc string
	var api string
	var datasetSummaryVo DatasetSummaryVo
	var datasetSummaryVos []DatasetSummaryVo

	for rows.Next() {
		switch err := rows.Scan(&title, &desc, &api); err {
        case sql.ErrNoRows:
			log.Info("No rows were returned")
		case nil:			
			datasetSummaryVo.Title = title
			datasetSummaryVo.Desc = desc
			datasetSummaryVo.Api = api
			
			datasetSummaryVos = append(datasetSummaryVos, datasetSummaryVo)
        default:
           	checkError(err)
        }
	}

	c.Header("Access-Control-Allow-Origin", "*") 
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": datasetSummaryVos})
}