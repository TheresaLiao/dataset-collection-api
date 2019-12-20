package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
	_ "github.com/lib/pq"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type CarAccidentTag struct {}

// curl http://localhost:port/dataset/caracdnt
func queryCarAccidentTagHandler(c *gin.Context){
	log.Info("queryCarAccidentTagHandler")
		
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

	sql_statement := ` SELECT "id", "KeyWord" FROM car_accident `
	rows, err := db.Query(sql_statement)
   	checkError(err)
	defer rows.Close()
	var id int
	var keyWord string
	var carAccidentVo CarAccidentVo
	var carAccidentVos []CarAccidentVo

	for rows.Next() {
		switch err := rows.Scan(&id, &keyWord); err {
				case sql.ErrNoRows:
				log.Info("No rows were returned")
		case nil:			
				carAccidentVo.Id = id
				carAccidentVo.KeyWord = keyWord
				carAccidentVos = append(carAccidentVos, carAccidentVo)
				default:
						 checkError(err)
		}
	}
	
	var dataSetVo CarAccitVo
	dataSetVo.Title = "Car Accident dataset"
	dataSetVo.Desc = "Include all type of Car Accident videos"
	dataSetVo.Data =  carAccidentVos
	
	c.Header("Access-Control-Allow-Origin", "*") 
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK,  "message": dataSetVo})
}

//  curl http://localhost:port/dataset/caracdnt/:{car_accident.id}
func queryCarAccidentByIdHandler(c *gin.Context){
		log.Info("queryCarAccidentByIdHandler")
		carAccidentIdStr := c.Param("carAccidentId")

		// connect dbㄠ
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

		sql_statement := ` SELECT B."CarAccidentID", B."title", B."youtube_id", B."URL", B."thumbnail", B."KeyWord"
											 FROM "car_accident" AS A
											 LEFT JOIN train_tw_org as B ON A."KeyWord" = B."KeyWord"
											 WHERE A."id" = $1`
	
		rows, err := db.Query(sql_statement, carAccidentIdStr)
		checkError(err)
		defer rows.Close()

		var carAccidentID string
		var title string
		var youtube_id string
		var url string
		var thumbnail string
		var keyWord string
		var trainTwOrgVo TrainTwOrgVo
		var trainTwOrgVos []TrainTwOrgVo

		for rows.Next() {
			switch err := rows.Scan(&carAccidentID, &title, &youtube_id, &url, &thumbnail, &keyWord); err {
					case sql.ErrNoRows:
					log.Info("No rows were returned")
			case nil:			
					trainTwOrgVo.CarAccidentID = carAccidentID
					trainTwOrgVo.Title = title
					trainTwOrgVo.YoutubeId = youtube_id
					trainTwOrgVo.Url = url
					trainTwOrgVo.Thumbnail = thumbnail
					trainTwOrgVo.KeyWord = keyWord
					trainTwOrgVos = append(trainTwOrgVos, trainTwOrgVo)
			default:
					checkError(err)
			}
		}
		c.Header("Access-Control-Allow-Origin", "*") 
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK,  "message": trainTwOrgVos})
}