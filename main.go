package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)
var db *gorm.DB =nil
var err error
func main(){

	//Connecting to the Database
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:root@tcp(127.0.0.1:3306)/crud?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,         // Disable color
		},
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger: newLogger,

	})
	if err !=nil{
		log.Fatal("Error connection Database  ")
	}
	db.AutoMigrate(&Post{})





	//Routing
	r := gin.Default()
	//POSTS Routing
	r.GET("/posts", Posts )
	r.GET("/posts/:id", Show )
	r.POST("/posts", Store )
	r.PATCH("/posts/:id", Update )
	r.DELETE("/posts/:id", Delete )
	r.Run(":9080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

