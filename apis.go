package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Posts(g *gin.Context)  {

}
func Store(g *gin.Context)  {
	var post Post
	if err :=g.ShouldBindJSON(&post); err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":"something went wrong with your request","data":""})
		return
	}
	post.Status="Active"
	db.Create(&post)
	g.JSON(http.StatusCreated,gin.H{"error":"","data":post})
	return


}
func Update(g *gin.Context)  {

}
func Delete(g *gin.Context)  {

}
func Show(g *gin.Context)  {

}
