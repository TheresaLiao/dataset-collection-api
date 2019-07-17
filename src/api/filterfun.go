package main

import (
	"os"
	"path/filepath"
	"strings"
	"github.com/gin-gonic/gin"
	. "github.com/kkdai/youtube"
	_ "github.com/lib/pq"
	"database/sql"
	"context"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"net/http"
	"os/exec"
)
type YoutubeInfoVo struct {
	FileName string `form:"filename" json:"filename" binding:"required"`
	Url string `form:"url" json:"url" binding:"required"`
}

type Item struct {
    Tag 		[]Tag 	`json:"tag"`
    Filename 	string 	`json:"filename"`
}
type Tag struct {
    Confidences  []int    `json:"confidences"`
    ObjectHeight int      `json:"objectHeight"`
    ObjectPicY   int      `json:"objectPicY"`
    ObjectTypes  []string `json:"objectTypes"`
    ObjectWidth  int      `json:"objectWidth"`
    ObjectPicX   int      `json:"objectPicX"`
}

const DOWNLOADS_PATH = "/tmp"
const SUBTITLE_PATH = "subtitle_"
const CARACDNT_PATH = "caracdnt_"
const TRAINTWORG_PATH = "traintworg"
const VIEDO_PATH = "video"
const FILE_EXTENTION_TAR = ".tar.gz"
const FILE_EXTENTION_MP4 = ".mp4"
const MAP_CSV_NAME = "map.csv"
const YOLO_URL = "http://task5-4-1:8080/yolo_coco_image"
const YOLO_FOLDER = "_yolo"

func getYoloImg(c *gin.Context){
	youtubeIdStr := c.Param("youtubeId")
	filenameStr := c.Param("filename")
	
	// srcDirPath : /tmp/traintworg
	srcDirPath := filepath.Join(DOWNLOADS_PATH, TRAINTWORG_PATH)
	// srcDirPathViedo : /tmp/traintworg/viedo
	srcDirPathViedo := filepath.Join(srcDirPath, VIEDO_PATH)
	srcFilePathViedo := filepath.Join(srcDirPathViedo, youtubeIdStr + YOLO_FOLDER, filenameStr)
	
	content, err := ioutil.ReadFile(srcFilePathViedo)
	if err != nil{
		log.Info(err)
	}

	c.Header("Access-Control-Allow-Origin", "*") 
	c.Data(http.StatusOK, "text/html; charset=utf-8", content)
}

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
		updateTrainTwOrgUrlByUrl(youtubeId,url)
	
		// TO-DO trigger training data
		// curl -d '{"videonames":"./traintworg/video/Rf9MxTLfdik.mp4", "dirname": "./traintworg/video/Rf9MxTLfdik"}' \
		// -X POST http://localhost:8080/yolo_coco_image
		// videonames := "./traintworg/video/" + youtubeId + ".mp4"
		// dirname := "./traintworg/video/" + youtubeId + YOLO_FOLDER
		videonames := filepath.Join(".", TRAINTWORG_PATH, VIEDO_PATH, youtubeId+".mp4")
		dirname := filepath.Join(".", TRAINTWORG_PATH, VIEDO_PATH, youtubeId+YOLO_FOLDER)

		
		triggerYoloApi(YOLO_URL,videonames,dirname)
	}
}

func parsingTrainingResult(c *gin.Context){
	// Read file and mapping viedoId.txt and jpg file into train_tw_tag
	// /tmp/traintworg/video
	log.Info("parsingTrainingResult")
	srcDirPath := filepath.Join(DOWNLOADS_PATH, TRAINTWORG_PATH, VIEDO_PATH)

	youtubeIds := queryTrainTwOrgUrlFilterByTag()
	for _, youtubeId := range youtubeIds {
		log.Info(youtubeId)
		// parding json file 
		// /tmp/traintworg/video/XwJa0bhThzI/XwJa0bhThzI.txt
		youtubeIdfile := filepath.Join(srcDirPath, youtubeId+YOLO_FOLDER, youtubeId+".txt")
		
		byteValue , _ := ioutil.ReadFile(youtubeIdfile)
		strValue := string(byteValue)
		
		// filter item
		strItems := strings.Split(strValue,"\n")
		for cntItem := range strItems {
			var item Item 
			var jsonBlob = []byte(strItems[cntItem])
    		err := json.Unmarshal(jsonBlob, &item )
			if err != nil {
        		log.Error("error:", err, strItems[cntItem])
			}
			if len(item.Tag) > 0 {
				for cntTag  := range item.Tag {
					if isTraffic(item.Tag[cntTag].ObjectTypes[0]){
						// write into db 
						insertTrainTwTagItem(youtubeId,
						item.Tag[cntTag].ObjectPicX,item.Tag[cntTag].ObjectPicY,
						item.Tag[cntTag].ObjectWidth,item.Tag[cntTag].ObjectHeight,
						item.Tag[cntTag].ObjectTypes[0],item.Filename)	
					}
				}
			}
		}
	}
}

func triggerYoloApi(urlStr string, videonames string, dirname string) {
	log.Info("start triggerYoloApi")
 
    post := "{\"videonames\":\"" + videonames + "\", \"dirname\": \"" + dirname + "\"}"
	log.Info(urlStr, "post", post)

    var jsonStr = []byte(post)
    log.Info("jsonStr", urlStr)
    log.Info("new_str", bytes.NewBuffer(jsonStr))

    req, err := http.NewRequest("POST", urlStr, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

	log.Info("response Status:", resp.Status)
	log.Info("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
	log.Info("response Body:", string(body))
}

func isTraffic(objectTypes string)(resp bool){
	if (objectTypes == "motorcycle" || objectTypes == "bicycle" || 
		objectTypes == "bus" || objectTypes == "train"|| 
		objectTypes == "truck" || objectTypes == "boat" || 
		objectTypes == "airplane"|| objectTypes == "car"){
		return true
	}
	return false
}

// records(youtube_id, x_num, y_num, weigtht, hight, object,filename)
func insertTrainTwTagItem( youtube_id string, x_num int, y_num int,
						   weigtht int,hight int, object string,
						   filename string){
	log.Info("insertTrainTwTagItem")
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
	
	// update columns youtube_id
	ctx := context.Background()
	sql_statement := ` INSERT INTO 
	train_tw_tag("youtube_id","x_num","y_num","weigtht","hight","object","filename") 
	VALUES ($1 ,$2 ,$3 ,$4 ,$5 ,$6 ,$7 )`

	result, err := db.ExecContext(ctx, sql_statement, youtube_id, x_num, y_num,
		weigtht, hight, object,filename)
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

func url2DownloadSubtitle(c *gin.Context){
	subtitleTagIdStr := c.Param("subtitleTagId")
	// parentFolderName : subtitle_N , ex. subtitle_1,subtitle_2...
	parentFolderName := SUBTITLE_PATH + subtitleTagIdStr
	// srcDirPath : /tmp/subtitle_N
	srcDirPath := filepath.Join(DOWNLOADS_PATH, parentFolderName)
	// srcDirPathViedo : /tmp/subtitle_N/viedo
	srcDirPathViedo := filepath.Join(srcDirPath, VIEDO_PATH)
	// srcDirPathCsv :/tmp/subtitle_N/map.csv
	srcDirPathCsv := filepath.Join(srcDirPath, MAP_CSV_NAME)
	// srcDirPath : /tmp/subtitle_N.tar.gz
	destFilePath := filepath.Join(DOWNLOADS_PATH , parentFolderName + FILE_EXTENTION_TAR)

	// check /tmp/subtitle_N.tar.gz is exit
	if checkFileIsExit(destFilePath) == false{
		// query data from sql
		records := querySubtitle(subtitleTagIdStr)
		if len(records) == 0 {
			log.Info("row data empty")
		}else{
			createDirectory(srcDirPath)
			// check /tmp/subtitle_N/map.csv, than search & download
			if checkFileIsExit(srcDirPathCsv) == false{
				title := []string{"id","youtube_id","srt_id","url"}
				getcsv(title,records, srcDirPathCsv)	
			}
			// check /tmp/subtitle_N/viedo, than search & download
			if checkFileIsExit(srcDirPathViedo) == false{
				for _, item := range records {
					checkUrlAndDownload(item[3], srcDirPathViedo)
				}
			}
		} 		

		// check /tmp/subtitle_N/viedo, than create tar file
		if checkFileIsExit(srcDirPathViedo) {
			// tar download folder
			tarDir(srcDirPath,destFilePath)
		}
	}
	// check /tmp/subtitle_N.tar.gz is exit?, if exit than retur to client
	if checkFileIsExit(destFilePath) {
		// download file from server to client
		respFile2Client(c,destFilePath)
	}
}

func url2DownloadCaracdnt(c *gin.Context){
	carAccidentTagIdStr := c.Param("carAccidentTagId")
	// parentFolderName : caracdnt_N , ex. caracdnt_1,caracdnt_2...
	parentFolderName := CARACDNT_PATH + carAccidentTagIdStr
	// srcDirPath : /tmp/caracdnt_N
	srcDirPath := filepath.Join(DOWNLOADS_PATH, parentFolderName)
	// srcDirPathViedo : /tmp/caracdnt_N/viedo
	srcDirPathViedo := filepath.Join(srcDirPath, VIEDO_PATH)
	// srcDirPathCsv :/tmp/caracdnt_N/map.csv
	srcDirPathCsv := filepath.Join(srcDirPath, MAP_CSV_NAME)
	// srcDirPath : /tmp/caracdnt_N.tar.gz
	destFilePath := filepath.Join(DOWNLOADS_PATH , parentFolderName + FILE_EXTENTION_TAR)


	// check /tmp/caracdnt_N.tar.gz is exit
	if checkFileIsExit(destFilePath) == false{
		// query data from sql
		records := queryCaracdnt(carAccidentTagIdStr)
		if len(records) == 0 {
			log.Info("row data empty")
		}else{

			createDirectory(srcDirPath)

			// check /tmp/caracdnt_N/map.csv, than search & download
			if checkFileIsExit(srcDirPathCsv) == false{
				title := []string{"id","youtube_id","collision_time","url"}
				getcsv(title,records, srcDirPathCsv)	
			}

			// check /tmp/caracdnt_N/viedo, than search & download
			if checkFileIsExit(srcDirPathViedo) == false{
				for _, item := range records {
					checkUrlAndDownload(item[3], srcDirPathViedo)
				}
			}
		}

		// check /tmp/caracdnt_N/viedo, than create tar file
		if checkFileIsExit(srcDirPathViedo) {
			// tar download folder
			tarDir(srcDirPath,destFilePath)
		}
	}

	// check /tmp/caracdnt_N.tar.gz is exit?, if exit than retur to client
	if checkFileIsExit(destFilePath) {
		// download file from server to client
		respFile2Client(c,destFilePath)
	}
}

func url2file(c *gin.Context){
	//dowmload file from url to server 
	var youtubeInfoVo YoutubeInfoVo
	c.BindJSON(&youtubeInfoVo)
	log.Info("Parameter FileName :" + youtubeInfoVo.FileName)
	log.Info("Parameter url :" + youtubeInfoVo.Url)
	
	srcDirPath := filepath.Join(DOWNLOADS_PATH, TRAINTWORG_PATH, VIEDO_PATH)
	videoID := checkUrlAndDownload(youtubeInfoVo.Url, srcDirPath)
	
	// download file from server to client
	destFilePath := filepath.Join(srcDirPath, videoID + FILE_EXTENTION_MP4)
	respFile2Client(c,destFilePath)

	// Remove download file
	if err := os.Remove(destFilePath) ; err != nil {
        log.Info("file remove Error!")
        log.Error("err ", err)
    } else {
        log.Info("file remove OK!")
    }
}


func checkFileIsExit(filepath string)(resp bool){ 
	log.Info("checkFileIsExit : "+filepath)

	if _, err := os.Stat(filepath); err == nil {
		log.Info(filepath + " is exeit")
		return true
	}
	return false
}

func checkUrlAndDownload(url string,srcDirPath string)(videoID string){
	log.Info("url : "+url)
	log.Info("srcDirPath : "+ srcDirPath)

	// // download file to localpath
	// filename := filepath.Join(srcDirPath, y.VideoID + FILE_EXTENTION_MP4)
	// log.Info("filename : " + filename)
	// if err := y.StartDownload(filename); err != nil {
	// 	log.Info("Error StartDownload")
	// 	log.Error("err ", err)
	// }
	// log.Info("VideoID = "+ y.VideoID)

	var youtubeId string
	// check url
	y := NewYoutube(true)
	if err := y.DecodeURL(url); err != nil {
		log.Error("err ", err)
		youtubeId = ""
	}else{
		youtubeId = y.VideoID
	}


	log.Info("annie -o "+srcDirPath+" -O "+y.VideoID+" "+url)
	cmd := exec.Command("annie","-o",srcDirPath,"-O",y.VideoID,url)
	err := cmd.Run()
	if err != nil {
		log.Error("err ", err)
	}
	// out, err := cmd.CombinedOutput()
	// if err != nil {
	// 	log.Info("cmd.Run() failed with %s\n", err)
	// }
	// log.Info("combined out:\n%s\n", string(out))

	return youtubeId
}

func querySubtitle(subtitleTagIdStr string)(resp [][]string){
	records := [][]string{}
	log.Info("subtitleTagIdStr : "+subtitleTagIdStr)

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
	sql_statement := `SELECT A.id, A.title, A.url, A.youtube_id, A.video_id 
					  FROM subtitle AS A
					  LEFT JOIN subtitle_tag_map AS B ON A.id=B.subtitle_id
					  WHERE B.subtitle_tag_id = $1
					  ORDER BY id LIMIT 30`
    rows, err := db.Query(sql_statement,subtitleTagIdStr)
    checkError(err)
	defer rows.Close()

	var id string
	var title string
	var url string
	var video_id string
	var youtube_id string
	
	for rows.Next() {
		switch err := rows.Scan(&id, &title, &url, &youtube_id, &video_id); err {
        case sql.ErrNoRows:
			log.Info("No rows were returned")
		case nil:			
			item := []string{id,youtube_id+".mp4",video_id+".srt",url}
			records =  append(records, item)
        default:
           	checkError(err)
        }
	}
	return records
}

func queryCaracdnt(carAccidentTagIdStr string)(resp [][]string){
	records := [][]string{}
	log.Info("carAccidentTagIdStr : "+carAccidentTagIdStr)
	

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
	sql_statement := ` SELECT DISTINCT A.id, A.title, A.url, A.youtube_id, C.collision_time
					   FROM car_accident AS A
					   LEFT JOIN car_accident_tag_map AS B ON A.id = B.car_accident_id
					   LEFT JOIN car_accident_collision_time C ON A.id = C.car_accident_id
					   WHERE B.car_accident_tag_id = $1
					   ORDER BY id`

    rows, err := db.Query(sql_statement, carAccidentTagIdStr)
    checkError(err)
	defer rows.Close()

	var id string
	var title string
	var url string
	var youtube_id string
	var collision_time string
	
	for rows.Next() {
		switch err := rows.Scan(&id, &title, &url, &youtube_id, &collision_time); err {
        case sql.ErrNoRows:
			log.Info("No rows were returned")
		case nil:
			item := []string{id,youtube_id+".mp4",collision_time,url}
			records =  append(records, item)
        default:
           	checkError(err)
        }
	}
	return records
}


// records(youtube_id,url)
func updateTrainTwOrgUrlByUrl(youtubeId string,url string){
	log.Info("updateTrainTwOrgUrlByUrl")
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

	// update columns youtube_id
	ctx := context.Background()
	sql_statement := ` UPDATE train_tw_org SET "youtube_id" =$1 WHERE "URL" =$2 `
	result, err := db.ExecContext(ctx, sql_statement, youtubeId, url)
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

func queryTrainTwOrgUrlFilterByTag()([]string){
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
	sql_statement := `  SELECT "youtube_id" 
					    FROM train_tw_org 
					    WHERE "youtube_id" NOT IN (
						   SELECT DISTINCT "youtube_id" 
						   FROM train_tw_tag 
						)
					    ORDER BY "youtube_id"`

	rows, err := db.Query(sql_statement)
    checkError(err)
	defer rows.Close()

	var youtubeId string

	for rows.Next() {
		switch err := rows.Scan(&youtubeId); err {
        case sql.ErrNoRows:
			log.Info("No rows were returned")
		case nil:
			records =  append(records, youtubeId)
        default:
           	checkError(err)
        }
	}
	return records
}

func respFile2Client(c *gin.Context,destFilePath string){
	log.Info("start respFile2Client")
	log.Info("destFilePath : "+ destFilePath)

	// download file from server to client
	if !strings.HasPrefix(filepath.Clean(destFilePath), DOWNLOADS_PATH) {
		c.String(403, "Look like you attacking me")
	   }   
	c.Header("Access-Control-Allow-Origin", "*") 
   	c.Header("Content-Description", "File Transfer")
   	c.Header("Content-Transfer-Encoding", "binary")
   	c.Header("Content-Disposition", "attachment; destFilePath=" + destFilePath )
   	c.Header("Content-Type", "application/octet-stream")
   	c.File(destFilePath)
}

func createDirectory(dirName string) bool {
	src, err := os.Stat(dirName)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(dirName, 0755)
		if errDir != nil {
			log.Info(errDir)
		}
		return true
	}

	if src.Mode().IsRegular() {
		log.Info(dirName, "already exist as a file!")
		return false
	}
	return false
}
