package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getById(g *gin.Context) Post  {
	id:=g.Param("id")
	var post Post
	db.First(&post,id)
	if post.ID ==0 {
		g.JSON(http.StatusNotFound,gin.H{"message":"We can not find the post","data":""})

	}
	return post
}




func Posts(g *gin.Context)  {

	limit, _ := strconv.Atoi(g.DefaultQuery("limit","10"))
	offset, _ := strconv.Atoi(g.DefaultQuery("offset","0"))
	var posts []Post

	db.Limit(limit).Offset(offset).Find(&posts)
	g.JSON(http.StatusOK,gin.H{
		"message":"",
		"data":posts,
	})

}
func Store(g *gin.Context)  {
	var post Post
	if err :=g.ShouldBindJSON(&post); err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"message":"something went wrong with your request","data":""})
		return
	}
	post.Status="Active"
	db.Create(&post)
	g.JSON(http.StatusCreated,gin.H{"message":"Post has been created ","data":post})
	return


}
func Update(g *gin.Context)  {
	//Get the post from the DataBase
	oldPost := getById(g)
	if oldPost.ID == 0 {
		return
	}

	//Post returned from the request [changed values]
	var requestedpost Post
	if err :=g.ShouldBindJSON(&requestedpost); err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"message":"something went wrong with your request","data":""})
		return
	}

	//update the old post with the new values from the requested post
	oldPost.Title=requestedpost.Title
	oldPost.Des=requestedpost.Des
	if requestedpost.Status !=""{
		oldPost.Status=requestedpost.Status
	}
	db.Save(&oldPost)


	g.JSON(http.StatusOK,gin.H{"message":"Post Has been updated ","data":oldPost})
	return




}
func Delete(g *gin.Context)  {
	post := getById(g)
	if post.ID == 0 {
		return
	}

	//Soft Delete [delete it but it still hase reference in the DB
	//db.Delete(&post)

	//Hard delete
	db.Unscoped().Delete(&post)

	g.JSON(http.StatusOK,gin.H{"message":"Post has been deleted","data":""})

}
func Show(g *gin.Context)  {

	post := getById(g)
	if post.ID == 0 {
		return
	}
  	g.JSON(http.StatusOK,gin.H{"message":"Post has been found ","data":post})

}

