package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
	_ "github.com/lib/pq"
	"path/filepath"
)

// QuerySubtitleTagHandler example
// @Summary get subtitle summary
// @Description 
// @Tags dataset subtitle 
// @ID query-subtitle-tag-handler
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Router /dataset/subtitle [get]
// curl http://localhost:port/dataset/subtitle
func QuerySubtitleTagHandler(c *gin.Context){
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
	sql_statement := "SELECT * FROM subtitle_tag;"
 	rows, err := db.Query(sql_statement)
 	checkError(err)
	defer rows.Close()
	 
	//parse raw data into json 
	var id int
	var tagName string
	var thumbnail string
	var subtitleTag SubtitleTag
	var subtitleTags []SubtitleTag

	for rows.Next() {
		switch err := rows.Scan(&id, &tagName, &thumbnail); err {
        case sql.ErrNoRows:
			log.Info("No rows were returned")
		case nil:
			subtitleTag.Id = id
			subtitleTag.TagName = tagName
			subtitleTag.Thumbnail = thumbnail
			subtitleTags = append(subtitleTags, subtitleTag)
        default:
           checkError(err)
        }
	}

	var dataSetVo SubtitleTagDataSetVo
	dataSetVo.Title = "Subtitle dataset"
	dataSetVo.Desc = "Include all type of Subtitle videos"
	dataSetVo.Data =  subtitleTags

	c.Header("Access-Control-Allow-Origin", "*") 
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK,  "message": dataSetVo})
}

// QuerySubtitleBySubtitleTagIdHandler example
// @Summary get this subtitleTagId detail list
// @Description 
// @Tags dataset subtitle 
// @ID query-subtitle-by-subtitle-tagId-handler
// @Accept  json
// @Produce  json
// @Param	subtitleTagId	path	string	true	"subtitleTag id"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} string "We need subtitleTag ID!!"
// @Router /dataset/subtitle/{subtitleTagId} [get]
// curl http://localhost:port/dataset/subtitle/${subtitleTagId}
func QuerySubtitleBySubtitleTagIdHandler(c *gin.Context){
	subtitleTagIdStr := c.Param("subtitleTagId")

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
	sql_statement := `SELECT A.id, A.title, A.url, A.thumbnail, A.youtube_id
					  FROM subtitle AS A
					  LEFT JOIN subtitle_tag_map AS B ON A.id=B.subtitle_id
					  WHERE B.subtitle_tag_id = $1`
	rows, err := db.Query(sql_statement, subtitleTagIdStr)
    checkError(err)
	defer rows.Close()

	var id int
	var title string
	var url string
	var thumbnail string
	var youtubeId string

	var subtitle Subtitle
	var subtitles []Subtitle

	for rows.Next() {
		switch err := rows.Scan(&id, &title, &url, &thumbnail, &youtubeId); err {
        case sql.ErrNoRows:
			log.Info("No rows were returned")
		case nil:
			subtitle.Id = id
			subtitle.Title = title
			subtitle.Url = url
			subtitle.Thumbnail = thumbnail
			subtitle.SrtUrl = "https://www.youtube.com/api/timedtext?v="+youtubeId+"&lang=zh-TW"
			subtitles = append(subtitles, subtitle)
			   
        default:
           checkError(err)
        }
	}
	
	var dataSetVo SubtitleVo
	dataSetVo.Title = "Subtitle dataset"
	dataSetVo.Desc = "Include all type of Subtitle videos"
	dataSetVo.Data =  subtitles

	c.Header("Access-Control-Allow-Origin", "*") 
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": dataSetVo})
}

// GetSubTitleThumbnail example
// @Summary update Thumbnail info for all subtitle data
// @Description 
// @Tags dataset subtitle 
// @ID get-subTitle-thumbnail
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Router /dataset/subTitleThumbnail [get]
// curl http://localhost:port/dataset/subTitleThumbnail
func GetSubTitleThumbnail(c *gin.Context){
	log.Info("getSubTitleThumbnail")

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
					   FROM  subtitle
					   WHERE "youtube_id" != ''
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
	for _, youtubeId := range youtubeIds {
		thumbnail := getYoutubeInfoByIdhttp(youtubeId)
		thumbnailAry[youtubeId] = thumbnail
	}

	// Insert thumbnail for each youtubeId
	for _, youtubeId := range youtubeIds {
		
		if thumbnailAry[youtubeId] != ""{
			log.Info(thumbnailAry[youtubeId])
			sql_statement2 := `UPDATE "subtitle"
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


// Url2DownloadSubtitleTag example
// @Summary download dataset by subtitleTagId
// @Description 
// @Tags dataset subtitle 
// @ID url-to-download-subtitle-tag
// @Accept  json
// @Produce  json
// @Param	subtitleTagId	path	string	true	"Subtitle Tag ID"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} string "We need Subtitle Tag ID!!"
// @Router /dataset/youtubeUrl/subtitle/{subtitleTagId} [get]
// curl --request GET \
// http://localhost:port/dataset/youtubeUrl/subtitle/${subtitleTagId} \
// --output ${filename}
// download this dataset all videos into one zip
func Url2DownloadSubtitleTag(c *gin.Context){
	log.Info("start Url2DownloadSubtitleTag")
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

	// check /tmp/subtitle_N.tar.gz is exist
	if checkFileIsExist(destFilePath) == false{
		// query data from sql
		records := querySubtitle(subtitleTagIdStr)
		if len(records) == 0 {
			log.Info("row data empty")
		}else{
			createDirectory(srcDirPath)
			// check /tmp/subtitle_N/map.csv, than search & download
			if checkFileIsExist(srcDirPathCsv) == false{
				title := []string{"id","youtube_id","srt_id","url"}
				getcsv(title ,records, srcDirPathCsv)	
			}
			
			// check /tmp/subtitle_N/viedo, than search & download
			if checkFileIsExist(srcDirPathViedo) == false{
				createDirectory(srcDirPathViedo)
				for _, item := range records {
					youtubeId := checkUrlAndDownload(item[3], srcDirPathViedo)
					updateSubtitleIdByUrl(youtubeId,item[3])
				}
			}
		} 		

		// check /tmp/subtitle_N/viedo, than create tar file
		if checkFileIsExist(srcDirPathViedo) {
			// tar download folder
			tarDir(srcDirPath,destFilePath)
		}
	}
	// check /tmp/subtitle_N.tar.gz is exist?, if exist than return to client
	if checkFileIsExist(destFilePath) {
		// download file from server to client
		respFile2Client(c,destFilePath)
	}
}

// Url2DownloadSubtitleId example
// @Summary download video by subtitleId
// @Description
// @Tags dataset subtitle 
// @ID url-to-download-subtitle-id
// @Accept  json
// @Produce  json
// @Param	subtitleId	path	string	true	"Subtitle ID"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} string "We need Subtitle ID!!"
// @Router /dataset/youtubeUrl/subtitleById/{subtitleId} [get]
// curl --request GET \
// http://localhost:port/dataset/youtubeUrl/subtitleById/${subtitleId} \
// --output filename.mp4
func Url2DownloadSubtitleId(c *gin.Context){
	log.Info("start url2DownloadSubtitleId")
	subtitleIdStr := c.Param("subtitleId")

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
	sql_statement := ` SELECT B.subtitle_tag_id, A.youtube_id
					   FROM subtitle AS A
					   LEFT JOIN subtitle_tag_map AS B ON A.id=B.subtitle_id
					   WHERE A.id = $1`

	rows, err := db.Query(sql_statement, subtitleIdStr)
	checkError(err)
	defer rows.Close()

	var subtitleTagId string
	var youtubeId string
	for rows.Next() {
		switch err := rows.Scan(&subtitleTagId,&youtubeId); err {
        case sql.ErrNoRows:
			log.Info("No rows were returned")
		case nil:
			// parentFolderName : subtitle_N , ex. subtitle_1,subtitle_2...
			parentFolderName := SUBTITLE_PATH + subtitleTagId
			// srcDirPath : /tmp/subtitle_N
			srcDirPath := filepath.Join(DOWNLOADS_PATH, parentFolderName)
			// srcDirPathViedo : /tmp/subtitle_N/viedo
			srcDirPathViedo := filepath.Join(srcDirPath, VIEDO_PATH)
			// srcDirPathViedo : /tmp/subtitle_N/viedo/youtubeId.mp4
			filePath :=  filepath.Join(srcDirPathViedo, youtubeId + FILE_EXTENTION_MP4)

			// check /tmp/subtitle_N/viedo/youtubeId.mp4 is exist?, if exist than return to client
			if checkFileIsExist(filePath) {
				// download file from server to client
				respFile2Client(c,filePath)
			}
        default:
           	checkError(err)
        }
	}
}