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
	"path/filepath"
	"context"
)

const RFC3339 = "2006-01-02T15:04:05Z07:00"
const SEARCH_URL = "https://www.googleapis.com/youtube/v3/search"
const VIEDOS_URL = "https://www.googleapis.com/youtube/v3/videos"
const YOUTUBE_API_KEY = "AIzaSyCcEHKC8RDbGlwY3LFEbhukJE9hXe4oboM" 
const SEARCH_KEYWORD1 = "擦撞+行車" 
const SEARCH_KEYWORD2 = "碰撞+監視器" 
const SEARCH_KEYWORD3 = "行車紀錄+事故" 
const SEARCH_KEYWORD4 = "車禍+行車紀錄" 
const SEARCH_KEYWORD5 = "車禍+行車視角" 


// curl http://localhost:port/dataset/queryTrainTwOrg
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

// curl http://localhost:port/dataset/queryTrainYoloTag/0-7_nvNNdcM
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


//curl http://localhost:port/dataset/queryTrainLprTag/:{youtubeId}

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

// curl http://localhost:port/filterfun/youtubeUrl/getSearchByKeyWord
func getSearchByKeyWord(c *gin.Context){
	log.Info("getSearchByKeyWord")
	
	keyWords := getKeyWordList()
	var trainTwOrgVos []TrainTwOrgVo
	for _, keyWord := range keyWords {
		ary := getYoutubeListSearchByKeyWord(keyWord)
		trainTwOrgVos = append(trainTwOrgVos, ary...)
	}

	youtubeIdSet := make(map[string]bool) // New empty set
	getYoutubeIdList(youtubeIdSet)

	// srcDirPath : /tmp/traintworg
	srcDirPath := filepath.Join(DOWNLOADS_PATH, TRAINTWORG_PATH)
	// srcDirPathViedo : /tmp/traintworg/viedo
	srcDirPathViedo := filepath.Join(srcDirPath, VIEDO_PATH)

	for _, item := range trainTwOrgVos {
		if youtubeIdSet[item.YoutubeId] != true{
			log.Info(item.KeyWord)
			checkUrlAndDownload(item.Url, srcDirPathViedo)
			insertTrainTwOrgItem(item)
		}
	}	
}

func getYoutubeIdList(youtubeIdSet map[string]bool){
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

	// get all search KeyWord
	sql_statement := ` SELECT DISTINCT youtube_id 
						FROM "train_tw_org" 
						WHERE "youtube_id" != '' 
						ORDER BY youtube_id`

	rows, err := db.Query(sql_statement)
   	checkError(err)
	defer rows.Close()

	var youtubeId string
	for rows.Next() {
		switch err := rows.Scan(&youtubeId); err {
			case sql.ErrNoRows:
			log.Info("No rows were returned")
		case nil:		
			youtubeIdSet[youtubeId] = true    
			default:
				checkError(err)
		}
	}
}

func getKeyWordList()(keyWords []string){
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

	// get all search KeyWord
	sql_statement := ` SELECT "KeyWord" FROM car_accident order by "KeyWord"`
	rows, err := db.Query(sql_statement)
   	checkError(err)
	defer rows.Close()

	var keyWord string
	for rows.Next() {
		switch err := rows.Scan(&keyWord); err {
			case sql.ErrNoRows:
			log.Info("No rows were returned")
		case nil:		
			keyWords = append(keyWords,keyWord)
			default:
				checkError(err)
		}
	}
	return keyWords
}

func insertTrainTwOrgItem(item TrainTwOrgVo){
	log.Info("start insertTrainTwOrgItem")
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

	// insert into sql
	ctx := context.Background()
	sql_statement := ` INSERT INTO "train_tw_org" ("title" ,"URL" ,"KeyWord", "youtube_id")VALUES ( $1 ,$2 ,$3, $4)` 

	result, err := db.ExecContext(ctx, sql_statement, item.Title, item.Url, item.KeyWord, item.YoutubeId)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if rows != 1 {
		log.Fatalf("expected to affect 1 row, affected %d", rows)
	}
}

func getYoutubeInfoById(c *gin.Context){
	log.Info("getYoutubeInfoById")

	youtubeIdStr := c.Param("youtubeId")
	pic := getYoutubeInfoByIdhttp(youtubeIdStr)
	log.Info(pic)

	c.Header("Access-Control-Allow-Origin", "*") 
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": pic})
}

// UPDATE "train_tw_org"."thumbnail" 
// WHERE "youtube_id" != 'NULL' AND "thumbnail" = ''
// curl http://localhost:port/dataset/queryTrainTwOrg/getThumbnail
func getTrainTwOrgThumbnail(c *gin.Context){
	log.Info("getTrainTwOrgThumbnail")

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
	thumbnailAry := make(map[string]string)
	for rows.Next() {
		switch err := rows.Scan(&youtubeId); err {
        case sql.ErrNoRows:
			log.Info("No rows were returned")
		case nil:
			youtubeIds = append(youtubeIds, youtubeId)
			thumbnailAry[youtubeId] = getYoutubeInfoByIdhttp(youtubeId)
        default:
           	checkError(err)
        }
	}

	// Insert thumbnail for each youtubeId
	for _, youtubeId := range youtubeIds {
		if thumbnailAry[youtubeId] != ""{
			log.Info(thumbnailAry[youtubeId])
			sql_statement2 := `UPDATE "train_tw_org"
			SET "thumbnail" = $1
			WHERE youtube_id = $2` 
	
			log.Info(thumbnailAry[youtubeId])
			_, err = db.Exec(sql_statement2,thumbnailAry[youtubeId] ,youtubeId)
			if err != nil {
				log.Info(err)
			}
		}
	}

	c.Header("Access-Control-Allow-Origin", "*") 
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": thumbnailAry})
}

func getYoutubeInfoByIdhttp(videoId string)(pic string){
	log.Info("getYoutubeInfoByIdhttp")
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

func getYoutubeListSearchByKeyWord(searchKeyWord string)(trainTwOrgVos []TrainTwOrgVo){
	log.Info("getYoutubeListSearchByKeyWord")
	// Last week video List
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
		// log.Info(searchKeyWord+"/ "+ item.Snippet.Title +"/ "+item.ID.VideoID)
		var trainTwOrgVo TrainTwOrgVo
		trainTwOrgVo.Title = item.Snippet.Title
		trainTwOrgVo.Url = "https://www.youtube.com/watch?v="+item.ID.VideoID
		trainTwOrgVo.KeyWord = searchKeyWord

		trainTwOrgVo.YoutubeId = item.ID.VideoID
		trainTwOrgVo.Thumbnail = item.Snippet.Thumbnails.Medium.URL		
		trainTwOrgVos = append(trainTwOrgVos, trainTwOrgVo)
	}
	return trainTwOrgVos
}

// curl http://localhost:port/filterfun/url2DownloadTrainTwOrg
// auto download all video and update table
func url2DownloadTrainTwOrg(c *gin.Context){
	log.Info("url2DownloadTrainTwOrg")
	// srcDirPath : /tmp/traintworg
	srcDirPath := filepath.Join(DOWNLOADS_PATH, TRAINTWORG_PATH)
	// srcDirPathViedo : /tmp/traintworg/viedo
	srcDirPathViedo := filepath.Join(srcDirPath, VIEDO_PATH)
	log.Info("srcDirPath:"+srcDirPath)
	log.Info("srcDirPathViedo:"+srcDirPathViedo)

	urls := queryTrainTwOrgUrl()

	for _, url := range urls {
		log.Info(url)
		youtubeId := checkUrlAndDownload(url, srcDirPathViedo)
		
		// insert youtubeId into train_tw_org where url = srcDirPathViedo
		log.Info(youtubeId)
		updateTrainTwOrgIdByUrl(youtubeId,url)
	
		videonames := filepath.Join(".", TRAINTWORG_PATH, VIEDO_PATH, youtubeId + FILE_EXTENTION_MP4)
		dirnameYolo := filepath.Join(".", TRAINTWORG_PATH, VIEDO_PATH, youtubeId+YOLO_FOLDER)
		dirnameLpr := filepath.Join(".", TRAINTWORG_PATH, VIEDO_PATH, youtubeId+LPR_FOLDER)
		
		triggerYoloApi(YOLO_URL,videonames,dirnameYolo)
		triggerLprApi(LPR_URL,videonames,dirnameLpr) 
	}
}

func queryTrainTwOrgUrl()([]string){
	log.Info("queryTrainTwOrgUrl")
	records := []string{}

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
	sql_statement := ` SELECT "URL" FROM train_tw_org 
					   WHERE "youtube_id" = 'NULL'
					   ORDER BY "URL"`
					   //  AND "CarAccidentID" = '3203'

	rows, err := db.Query(sql_statement)
    checkError(err)
	defer rows.Close()

	var url string

	for rows.Next() {
		switch err := rows.Scan(&url); err {
        case sql.ErrNoRows:
			log.Info("No rows were returned")
		case nil:
			records =  append(records, url)
        default:
           	checkError(err)
        }
	}
	return records
}

// curl --request GET \
// http://localhost:port/dataset/youtubeUrl/cartype/0-7_nvNNdcM \
// --output 555.mp4
// download from client, by youtubeId
func url2DownloadCarType(c *gin.Context){
	log.Info("start url2DownloadCarType")
	youtubeIdStr := c.Param("youtubeId")

	// srcDirPath: /tmp/traintworg/video
	srcDirPath := filepath.Join(DOWNLOADS_PATH, TRAINTWORG_PATH, VIEDO_PATH)
	destFilePath := filepath.Join(srcDirPath, youtubeIdStr + FILE_EXTENTION_MP4)
	respFile2Client(c,destFilePath)
}