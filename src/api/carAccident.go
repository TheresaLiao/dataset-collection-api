package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
	_ "github.com/lib/pq"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type CarAccidentTag struct {}

// QueryCarAccidentTagHandler example
// @Summary get car accident summary
// @Description
// @Tags dataset car accident
// @ID query-carAccident-tag-handler
// @Accept  json
// @Produce  json
// @Success 200	{object}	main.CarAccitVo	"ok"
// @Router /dataset/caracdnt [get]
// curl http://localhost:port/dataset/caracdnt
func QueryCarAccidentTagHandler(c *gin.Context){
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

// QueryCarAccidentByIdHandler example
// @Summary get list by carAccidentId
// @Description 
// @Tags dataset car accident
// @ID query-carAccident-by-id-handler
// @Accept  json
// @Produce  json
// @Param	carAccidentId	path	int	true	"Car Accident Id"
// @Success 200	{array}	main.TrainTwOrgVo	"ok"
// @Failure 400 {object} string "We need CarAccidentId!!"
// @Router /dataset/caracdnt/{carAccidentId} [get]
// curl http://localhost:port/dataset/caracdnt/${carAccidentId}
func QueryCarAccidentByIdHandler(c *gin.Context){
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

		sql_statement := ` SELECT B."CarAccidentID", B."title", B."youtube_id", B."URL", B."thumbnail", B."KeyWord", B."collision_time", B."video_length", B."car_type"
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
		var collisionTime string
		var videoLength string
		var carType string
		var trainTwOrgVo TrainTwOrgVo
		var trainTwOrgVos []TrainTwOrgVo

		for rows.Next() {
			switch err := rows.Scan(&carAccidentID, &title, &youtube_id, &url, &thumbnail, &keyWord, &collisionTime, &videoLength, &carType); err {
					case sql.ErrNoRows:
					log.Info("No rows were returned")
			case nil:			
					trainTwOrgVo.CarAccidentID = carAccidentID
					trainTwOrgVo.Title = title
					trainTwOrgVo.YoutubeId = youtube_id
					trainTwOrgVo.Url = url
					trainTwOrgVo.Thumbnail = thumbnail
					trainTwOrgVo.KeyWord = keyWord
					trainTwOrgVo.CollisionTime = collisionTime
					trainTwOrgVo.VideoLength = videoLength
					trainTwOrgVo.CarType = carType
					trainTwOrgVos = append(trainTwOrgVos, trainTwOrgVo)
			default:
					checkError(err)
			}
		}
		c.Header("Access-Control-Allow-Origin", "*") 
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK,  "message": trainTwOrgVos})
}