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
	Url string `json:"url"`
}

type TrainTwTagVo struct {
	Id int `json:"id"`
	YoutubeId  string `json:"youtubeId"`
	Object string `json:object"`
	Filename string `json:"filename"`
	XNum int `json:"x_num"`
	YNum int `json:"y_num"`
	Width int `json:"width"`
	Height int `json:"height"`
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

	sql_statement := ` SELECT "CarAccidentID", "title", "youtube_id" ,"URL"
					   FROM train_tw_org 
					   WHERE "youtube_id" != 'NULL'
					   ORDER BY "youtube_id"`

	rows, err := db.Query(sql_statement)
    checkError(err)
	defer rows.Close()

	var carAccidentID string
	var title string
	var youtube_id string
	var url string
	var trainTwOrgVo TrainTwOrgVo
	var trainTwOrgVos []TrainTwOrgVo

	for rows.Next() {
		switch err := rows.Scan(&carAccidentID, &title, &youtube_id, &url); err {
        case sql.ErrNoRows:
			log.Info("No rows were returned")
		case nil:			
			trainTwOrgVo.CarAccidentID = carAccidentID
			trainTwOrgVo.Title = title
			trainTwOrgVo.YoutubeId = youtube_id
			trainTwOrgVo.Url = url
			trainTwOrgVos = append(trainTwOrgVos, trainTwOrgVo)
        default:
           	checkError(err)
        }
	}
	c.Header("Access-Control-Allow-Origin", "*") 
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": trainTwOrgVos})
}

func queryTrainYoloTagByYoutubeIdHandler(c *gin.Context){
	log.Info("queryTrainYoloTagByYoutubeIdHandler")
	youtubeIdStr := c.Param("youtubeId")

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

	sql_statement := ` SELECT "youtube_id", "object", "filename" , "id", "x_num", "y_num", "width", "height"
					   FROM train_yolo_tag 
					   WHERE "youtube_id" = $1 
					   ORDER BY "filename"`

	rows, err := db.Query(sql_statement,youtubeIdStr)
    checkError(err)
	defer rows.Close()

	var id int
	var youtube_id string
	var object string
	var filename string
	var trainTwTagVo TrainTwTagVo
	var trainTwTagVos []TrainTwTagVo

	var xNum int
	var yNum int 
	var width int 
	var height int 

	for rows.Next() {
		switch err := rows.Scan(&youtube_id, &object, &filename, &id, &xNum, &yNum, &width, &height); err {
        case sql.ErrNoRows:
			log.Info("No rows were returned")
		case nil:			
			trainTwTagVo.YoutubeId = youtube_id
			trainTwTagVo.Object = object
			trainTwTagVo.Filename = filename
			trainTwTagVo.Id = id
			trainTwTagVo.XNum = xNum
			trainTwTagVo.YNum = yNum
			trainTwTagVo.Width = width
			trainTwTagVo.Height = height
			trainTwTagVos = append(trainTwTagVos, trainTwTagVo)
        default:
           	checkError(err)
        }
	}
	c.Header("Access-Control-Allow-Origin", "*") 
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": trainTwTagVos})
}

func queryTrainLprTagByYoutubeIdHandler(c *gin.Context){
	log.Info("queryTrainLprTagByYoutubeIdHandler")
	youtubeIdStr := c.Param("youtubeId")

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

	sql_statement := ` SELECT "youtube_id", "plateNumber", "filename" , "id", "x_num", "y_num", "width", "height"
					   FROM train_lpr_tag 
					   WHERE "youtube_id" = $1 
					   ORDER BY "filename"`

	rows, err := db.Query(sql_statement,youtubeIdStr)
    checkError(err)
	defer rows.Close()

	var id int
	var youtube_id string
	var object string
	var filename string
	var trainTwTagVo TrainTwTagVo
	var trainTwTagVos []TrainTwTagVo

	var xNum int
	var yNum int 
	var width int 
	var height int 

	for rows.Next() {
		switch err := rows.Scan(&youtube_id, &object, &filename, &id, &xNum, &yNum, &width, &height); err {
        case sql.ErrNoRows:
			log.Info("No rows were returned")
		case nil:			
			trainTwTagVo.YoutubeId = youtube_id
			trainTwTagVo.Object = object
			trainTwTagVo.Filename = filename
			trainTwTagVo.Id = id
			trainTwTagVo.XNum = xNum
			trainTwTagVo.YNum = yNum
			trainTwTagVo.Width = width
			trainTwTagVo.Height = height
			trainTwTagVos = append(trainTwTagVos, trainTwTagVo)
        default:
           	checkError(err)
        }
	}
	c.Header("Access-Control-Allow-Origin", "*") 
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": trainTwTagVos})
}