package main

import (
	"backend/db"
	"backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}

//var DB *sql.DB
//
//func InitDB() {
//	var err error
//	DB, err = sql.Open("sqlite3", "api.db")
//
//	if err != nil {
//		panic("Could not connect to database.")
//	}
//
//	DB.SetMaxOpenConns(10)
//	DB.SetMaxIdleConns(5)
//
//	createTables()
//}
