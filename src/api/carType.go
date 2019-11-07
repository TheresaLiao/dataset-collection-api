package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"database/sql"
	"net/http"
	"net/url"
	"time"
	"io/ioutil"
	"encoding/json"
)

const RFC3339 = "2006-01-02T15:04:05Z07:00"
const SEARCH_URL = "https://www.googleapis.com/youtube/v3/search"
const VIEDOS_URL = "https://www.googleapis.com/youtube/v3/videos"
const YOUTUBE_API_KEY = "AIzaSyCcEHKC8RDbGlwY3LFEbhukJE9hXe4oboM" 
const SEARCH_KEYWORD1 = "擦撞 行車" 
const SEARCH_KEYWORD2 = "碰撞 監視器" 
const SEARCH_KEYWORD3 = "行車紀錄 事故" 
const SEARCH_KEYWORD4 = "車禍 行車紀錄" 
const SEARCH_KEYWORD5 = "車禍 行車視角" 

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

	sql_statement := ` SELECT "CarAccidentID", "title", "youtube_id" ,"URL","thumbnail"
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
	var thumbnail string
	var trainTwOrgVo TrainTwOrgVo
	var trainTwOrgVos []TrainTwOrgVo

	for rows.Next() {
		switch err := rows.Scan(&carAccidentID, &title, &youtube_id, &url, &thumbnail); err {
        case sql.ErrNoRows:
			log.Info("No rows were returned")
		case nil:			
			trainTwOrgVo.CarAccidentID = carAccidentID
			trainTwOrgVo.Title = title
			trainTwOrgVo.YoutubeId = youtube_id
			trainTwOrgVo.Url = url
			trainTwOrgVo.Thumbnail = thumbnail
			trainTwOrgVos = append(trainTwOrgVos, trainTwOrgVo)
        default:
           	checkError(err)
        }
	}

	var dataSetVo TrainTwOrgDataSetVo
	dataSetVo.Title = "Car Type dataset"
	dataSetVo.Desc = "Include all type of car video"
	dataSetVo.Data =  trainTwOrgVos

	c.Header("Access-Control-Allow-Origin", "*") 
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": dataSetVo})
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

func getSearchByKeyWord(c *gin.Context){
	log.Info("getSearchByKeyWord")
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

	getYoutubeListSearchByKeyWord(SEARCH_KEYWORD1)
	getYoutubeListSearchByKeyWord(SEARCH_KEYWORD2)
	getYoutubeListSearchByKeyWord(SEARCH_KEYWORD3)
	getYoutubeListSearchByKeyWord(SEARCH_KEYWORD4)
	getYoutubeListSearchByKeyWord(SEARCH_KEYWORD5)

	//insert into sql
	// sql_statement := ` INSERT INTO "train_tw_org" (
	// 					"CarAccidentID", 
	// 					"title" ,
	// 					"URL" ,
	// 					"youtube_id")
	// 				   VALUES ( ? ,? ,? ,"NULL")` 

	// result, err := db.Exec(sql_statement,,,)
}

func getYoutubeInfoById(c *gin.Context){
	log.Info("getYoutubeInfoById")

	youtubeIdStr := c.Param("youtubeId")
	pic := getYoutubeInfoByIdhttp(youtubeIdStr)
	log.Info(pic)

	c.Header("Access-Control-Allow-Origin", "*") 
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": pic})
}

func getTrainTwOrgThumbnail(c *gin.Context){
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

	// Search all youtube_id
	sql_statement := ` SELECT DISTINCT "youtube_id"
					   FROM train_tw_org 
					   WHERE "youtube_id" != 'NULL'
					   AND "thumbnail" = ''
					   ORDER BY "youtube_id"`

	rows, err := db.Query(sql_statement)
    checkError(err)
	defer rows.Close()

	var youtubeId string
	var youtubeIds []string
	for rows.Next() {
		switch err := rows.Scan(&youtubeId); err {
        case sql.ErrNoRows:
			log.Info("No rows were returned")
		case nil:
			youtubeIds = append(youtubeIds, youtubeId)
        default:
           	checkError(err)
        }
	}

	// Get Thumbnail by youtubeId
	thumbnailAry := make(map[string]string)
	for _, youtubrId := range youtubeIds {
		thumbnail := getYoutubeInfoByIdhttp(youtubrId)
		thumbnailAry[youtubrId] = thumbnail
	}

	// Insert thumbnail for each youtubeId
	for _, youtubrId := range youtubeIds {
		
		if thumbnailAry[youtubrId] != ""{
			log.Info(thumbnailAry[youtubrId])
			sql_statement2 := `UPDATE "train_tw_org"
			SET "thumbnail" = $1
			WHERE youtube_id = $2` 
	

			log.Info(thumbnailAry[youtubrId])
			_, err = db.Exec(sql_statement2,thumbnailAry[youtubrId] ,youtubrId)
			if err != nil {
				log.Info(err)
			}
		}
	}

	c.Header("Access-Control-Allow-Origin", "*") 
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": thumbnailAry})
}

func getYoutubeInfoByIdhttp(videoId string)(pic string){
	baseUrl, err := url.Parse(VIEDOS_URL)
	if err != nil {
		log.Fatal(err)
	}
	params := url.Values{}
	params.Add("key", YOUTUBE_API_KEY)
	params.Add("part", "snippet")
	params.Add("id", videoId)
	baseUrl.RawQuery = params.Encode()

	// get responce
	resp, err := http.Get(baseUrl.String())  
	if err != nil {  
		log.Fatal(err)
	}
	defer resp.Body.Close()  
	body, err := ioutil.ReadAll(resp.Body)  
	if err != nil {
		log.Fatal(err)
	}  
	respJson := string(body)	

	// parsing json to struct
	videoVo := VideoVo{}
	json.Unmarshal([]byte(respJson), &videoVo)
	
	var respStr = ""
	if len(videoVo.Items) > 0{
		respStr = videoVo.Items[0].Snippet.Thumbnails.Medium.URL
	}
	return respStr
}

func getYoutubeListSearchByKeyWord(searchKeyWord string){

	now := time.Now()
    weekago := now.AddDate(0, 0, -7)
	strdate := weekago.Format(RFC3339)
	enddate := now.Format(RFC3339)

	baseUrl, err := url.Parse(SEARCH_URL)
	if err != nil {
		log.Fatal(err)
	}
	params := url.Values{}
	params.Add("key", YOUTUBE_API_KEY)
	params.Add("part", "snippet")
	params.Add("q", searchKeyWord)
	params.Add("publishedAfter", strdate)
	params.Add("publishedBefore", enddate)
	baseUrl.RawQuery = params.Encode()

	// get responce
	resp, err := http.Get(baseUrl.String())  
	if err != nil {  
		log.Fatal(err)
	}
	defer resp.Body.Close()  
	body, err := ioutil.ReadAll(resp.Body)  
	if err != nil {
		log.Fatal(err)
	}  
	respJson := string(body)	
	
	// parsing json to struct
	searchVo := SearchVo{}
	json.Unmarshal([]byte(respJson), &searchVo)

	for _, item := range searchVo.Items {
		log.Info(item.Snippet.Title)
		log.Info(item.ID.VideoID)
	}
}