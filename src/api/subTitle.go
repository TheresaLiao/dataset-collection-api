package main

import (
	// "bytes"
	// "encoding/json"
	"net/http"
	// "strconv"
	"fmt"
	"github.com/gin-gonic/gin"
	"database/sql"
	_ "github.com/lib/pq"
	
)

const connStr = "postgres://admin:12345@140.96.29.202/pqgotest?sslmode=verify-full"

// type Subtitle struct{
// 	Id					int64
// 	Title				string
// 	VideoLanguage		string
// 	SubtitleLanguage	string
// 	CopyRight			string
// 	Url					string
// 	VideoId				string
// 	YoutubeId			string
// 	Embedded			bool
// 	PlugIn				bool
// 	VideoLength			int
// }
// type SubtitleTag struct{
// 	Id		int64
// 	TagName	string 
// }
// type SubtitleTagMap struct{
// 	Id 				int64
// 	SubtitleId 		int64
// 	SubtitleTagId 	int64
// }



func subTitleHandler(c *gin.Context) {
	connectPostgrepsql()
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "apiserver ready and Summary Connection ",
	})
}

func connectPostgrepsql(){
	//psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",host, port, user, password, dbname)
	db, err := sql.Open("postgres",connStr)
	if err != nil{
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil{
		panic(err)
	}

	fmt.Println("success connection")
	

	sql_statement := "SELECT *  FROM subtitle WHERE id in (SELECT subtitle_id FROM subtitle_tag_map WHERE subtitle_tag_id =3);"
    rows, err := db.Query(sql_statement)
    checkError(err)
	defer rows.Close()


	var id int
    var title string
	var url string
	
	for rows.Next() {
		switch err := rows.Scan(&id, &title, &url); err {
        case sql.ErrNoRows:
           	fmt.Println("No rows were returned")
        case nil:
           	fmt.Printf("Data row = (%d, %s, %d)\n", id, title, url)
        default:
           checkError(err)
        }
    }
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}












