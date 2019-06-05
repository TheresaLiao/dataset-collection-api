package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


type CarAccidentVo struct {
	Id  int `json:"id"`
	Title string `json:"title"`
	Url string `json:"url"`
}

type CarAccidentTagVo struct {
	Id  int `json:"id"`
	TagName string `json:"tagName"`
}

type CarAccidentTag struct {}

func queryCarAccidentTagHandler(c *gin.Context){
		db, err := gorm.Open("postgres", connStr)
		if err != nil {
			log.Info("failed to connect database")
		}
		log.Info("Connection connection")

		db.SingularTable(true)
	
		// carAccidentTags :=  &CarAccidentTag{}
		// db.Debug().Find(&carAccidentTags)
		// for i := range carAccidentTags {
		// 	fmt.Println(i, &carAccidentTags[i])
		// 	v := &carAccidentTags[i]
		// 	CarAccidentTag[v.PackageName] = &carAccidentTag[i]
		// }

		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "apiserver ready and Summary Connection ",
		})
}

func test(c *gin.Context){
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
	var carAccidentTagVo CarAccidentTagVo
	var carAccidentTagVos []CarAccidentTagVo

	for rows.Next() {
		switch err := rows.Scan(&id, &tagName); err {
        case sql.ErrNoRows:
			log.Info("No rows were returned")
		case nil:
			carAccidentTagVo.Id = id
			carAccidentTagVo.TagName = tagName
			carAccidentTagVos = append(carAccidentTagVos, carAccidentTagVo)
        default:
           checkError(err)
        }
	}

  c.Header("Access-Control-Allow-Origin", "*") 
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": carAccidentTagVos})
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

	var carAccidentVo CarAccidentVo
	var carAccidentVos []CarAccidentVo

	for rows.Next() {
		switch err := rows.Scan(&id, &title, &url); err {
        case sql.ErrNoRows:
			log.Info("No rows were returned")
		case nil:
			log.Info("Data row = (%d, %s, %d)\n", id, title, url)
			carAccidentVo.Id = id
			carAccidentVo.Title = title
			carAccidentVo.Url = url
			log.Info("Data row = (%d, %s, %s)\n", id, title, url)
			carAccidentVos = append(carAccidentVos, carAccidentVo)
		
        default:
           checkError(err)
        }
	}
  c.Header("Access-Control-Allow-Origin", "*") 
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": carAccidentVos})

}