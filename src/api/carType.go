package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"database/sql"
	"net/http"
)

type TrainTwOrgVo struct {
	CarAccidentID  string `json:"carAccidentID"`
	Title string `json:"title"`
	YoutubeId string `json:"youtubeId"`
}

type TrainTwTagVo struct {
	YoutubeId  string `json:"youtubeId"`
	Object string `json:object"`
	Filename string `json:"filename"`
}

func queryTrainTwOrgHandler(c *gin.Context){
	log.Info("queryTrainTwOrgHandler")

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

	sql_statement := ` SELECT "CarAccidentID", "title", "youtube_id" 
					   FROM train_tw_org 
					   WHERE "youtube_id" != 'NULL'`

	rows, err := db.Query(sql_statement)
    checkError(err)
	defer rows.Close()

	var carAccidentID string
	var title string
	var youtube_id string
	var trainTwOrgVo TrainTwOrgVo
	var trainTwOrgVos []TrainTwOrgVo

	for rows.Next() {
		switch err := rows.Scan(&carAccidentID, &title, &youtube_id); err {
        case sql.ErrNoRows:
			log.Info("No rows were returned")
		case nil:			
			trainTwOrgVo.CarAccidentID = carAccidentID
			trainTwOrgVo.Title = title
			trainTwOrgVo.YoutubeId = youtube_id
			trainTwOrgVos = append(trainTwOrgVos, trainTwOrgVo)
        default:
           	checkError(err)
        }
	}
	c.Header("Access-Control-Allow-Origin", "*") 
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": trainTwOrgVos})
}

func queryTrainTwTagByYoutubeIdHandler(c *gin.Context){
	youtubeIdStr := c.Param("youtubeId")
	log.Info("queryTrainTwTagHandler")

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

	sql_statement := ` SELECT "youtube_id", "object", "filename" 
					   FROM train_tw_tag 
					   WHERE "youtube_id" = $1 `

	rows, err := db.Query(sql_statement,youtubeIdStr)
    checkError(err)
	defer rows.Close()

	var youtube_id string
	var object string
	var filename string
	var trainTwTagVo TrainTwTagVo
	var trainTwTagVos []TrainTwTagVo

	for rows.Next() {
		switch err := rows.Scan(&youtube_id, &object, &filename); err {
        case sql.ErrNoRows:
			log.Info("No rows were returned")
		case nil:			
			trainTwTagVo.YoutubeId = youtube_id
			trainTwTagVo.Object = object
			trainTwTagVo.Filename = filename
			trainTwTagVos = append(trainTwTagVos, trainTwTagVo)
        default:
           	checkError(err)
        }
	}
	c.Header("Access-Control-Allow-Origin", "*") 
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": trainTwTagVos})
}