package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
	_ "github.com/lib/pq"
)

func querySubtitleTagHandler(c *gin.Context){
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
	dataSetVo.Desc = "SubTitle dataset"
	dataSetVo.Data =  subtitleTags

	c.Header("Access-Control-Allow-Origin", "*") 
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK,  "message": dataSetVo})
}

func querySubtitleBySubtitleTagIdHandler(c *gin.Context){
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
	sql_statement := `SELECT A.id, A.title, A.url, A.thumbnail
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

	var subtitle Subtitle
	var subtitles []Subtitle

	for rows.Next() {
		switch err := rows.Scan(&id, &title, &url, &thumbnail); err {
        case sql.ErrNoRows:
			log.Info("No rows were returned")
		case nil:
			subtitle.Id = id
			subtitle.Title = title
			subtitle.Url = url
			subtitle.Thumbnail = thumbnail
			subtitles = append(subtitles, subtitle)
			   
        default:
           checkError(err)
        }
	}
	
	c.Header("Access-Control-Allow-Origin", "*") 
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": subtitles})
}

