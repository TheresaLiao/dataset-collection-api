package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
	_ "github.com/lib/pq"
)

type CarAccident struct {
	Id  int `json:"id"`
	Title string `json:"title"`
	Url string `json:"url"`
//	CopyRight string `json:"copyRight"`
//	AccidentTime string `json:"accidentTime"`
//	CarType string `json:"carType"`
//	DayTime string `json:"dayTime"`
//	Collision string `json:"collision"`
}

type CarAccidentTag struct {
	Id  int `json:"id"`
	TagName string `json:"tagName"`
}

func queryCarAccidentTagHandler(c *gin.Context){
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

	// select table :car_accident_tag ,all rows data
	sql_statement := "SELECT * FROM  car_accident_tag;"
 	rows, err := db.Query(sql_statement)
 	checkError(err)
	defer rows.Close()

	//parse raw data into json 
	var id int
	var tagName string
	var carAccidentTag CarAccidentTag
	var carAccidentTags []CarAccidentTag

	for rows.Next() {
		switch err := rows.Scan(&id, &tagName); err {
        case sql.ErrNoRows:
			log.Info("No rows were returned")
		case nil:
			carAccidentTag.Id = id
			carAccidentTag.TagName = tagName
			log.Info("Data row = (%d, %s)\n", id, tagName)
			carAccidentTags = append(carAccidentTags, carAccidentTag)
        default:
           checkError(err)
        }
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": carAccidentTags})
}

func queryCarAccidentByCarAccidentTagIdHandler(c *gin.Context){
	carAccidentTagIdStr := c.Param("carAccidentTagId")

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

	sql_statement := `SELECT A.id, A.title, A.url 
	 				  FROM car_accident as A 
	 				  LEFT JOIN car_accident_tag_map AS B ON A.id = B.car_accident_id 
	 				  WHERE B.car_accident_tag_id = $1`
    rows, err := db.Query(sql_statement, carAccidentTagIdStr)
    if err != nil {
		//log.Fatal(err)
		log.Info(err)
	}
	defer rows.Close()

	var (
		id   int
		title string
		url string
	)

	var carAccident CarAccident
	var carAccidents []CarAccident

	for rows.Next() {
		switch err := rows.Scan(&id, &title, &url); err {
        case sql.ErrNoRows:
			log.Info("No rows were returned")
		case nil:
			log.Info("Data row = (%d, %s, %d)\n", id, title, url)
			carAccident.Id = id
			carAccident.Title = title
			carAccident.Url = url
			log.Info("Data row = (%d, %s, %s)\n", id, title, url)
			carAccidents = append(carAccidents, carAccident)
		
        default:
           checkError(err)
        }
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": carAccidents})
}