package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
	_ "github.com/lib/pq"
)

type SubtitleTag struct {
	Id  int `json:"id"`
	TagName string `json:"tagName"`
}

type Subtitle struct {
	Id  int `json:"id"`
	Title string `json:"title"`
	Url string `json:"url"`
}

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
	var subtitleTag SubtitleTag
	var subtitleTags []SubtitleTag

	for rows.Next() {
		switch err := rows.Scan(&id, &tagName); err {
        case sql.ErrNoRows:
			log.Info("No rows were returned")
		case nil:
			subtitleTag.Id = id
			subtitleTag.TagName = tagName
			log.Info("Data row = (%d, %s)\n", id, tagName)
			subtitleTags = append(subtitleTags, subtitleTag)
        default:
           checkError(err)
        }
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": subtitleTags})
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
	sql_statement := "SELECT *  FROM subtitle WHERE id in (SELECT subtitle_id FROM subtitle_tag_map WHERE subtitle_tag_id =" + subtitleTagIdStr + ");"
    rows, err := db.Query(sql_statement)
    checkError(err)
	defer rows.Close()

	var id int
	var title string
	var videoLanguage string
	var subtitleLanguage string
	var copyRight string
	var url string
	var videoId string
	var youtubeId string
	var embedded bool
	var plugIn bool
	var videoLength int

	var subtitle Subtitle
	var subtitles []Subtitle

	for rows.Next() {
		switch err := rows.Scan(&id, &title, &videoLanguage,
			&subtitleLanguage, &copyRight, &url,
			&videoId, &youtubeId, &embedded, &plugIn, &videoLength); err {
        case sql.ErrNoRows:
			log.Info("No rows were returned")
		case nil:
			log.Info("Data row = (%d, %s, %d)\n", id, title, url)
			subtitle.Id = id
			subtitle.Title = title
			subtitle.Url = url
			log.Info("Data row = (%d, %s, %s)\n", id, title, url)
			subtitles = append(subtitles, subtitle)
			   
        default:
           checkError(err)
        }
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": subtitles})
}

