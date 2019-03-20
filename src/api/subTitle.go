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
// const (
// 	host     = "140.96.29.202"
// 	port     = 5566
// 	user     = "admin"
// 	password = "12345"
// 	dbname   = "postgres"
//   )

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
}