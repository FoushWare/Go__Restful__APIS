package main
import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)
func main(){

	//Connecting to the Database
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:root@tcp(127.0.0.1:3306)/crud?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err !=nil{
		log.Fatal("Error connection Database")
	}

	//Routing
	r := gin.Default()
	//POSTS Routing
	r.GET("/posts", Posts )
	r.GET("/posts/:id", Show )
	r.POST("/posts", Store )
	r.PUT("/posts/:id", Update )
	r.DELETE("/posts/:id", Delete )
	r.Run(":9080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

