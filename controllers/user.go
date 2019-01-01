package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserCtrlType : Type controller for User
type UserCtrlType struct {
	FindAll           gin.HandlerFunc
	FindOneByID       gin.HandlerFunc
	FindOneByUsername gin.HandlerFunc
}

// UserCtrl : Controller for User
var UserCtrl = UserCtrlType{
	findAll,
	findOneByID,
	findOneByUsername,
}

func findAll(ctx *gin.Context) {
	email := ctx.DefaultQuery("email", "")
	page := ctx.DefaultQuery("page", "1")
	keywords := ctx.DefaultQuery("keywords", "")

	log.Println(email, page, keywords)

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Received",
		"email":    email,
		"page":     page,
		"keywords": keywords,
	})
}

func findOneByID(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"id":      id,
		"message": "Received",
	})
}

func findOneByUsername(ctx *gin.Context) {
	username := ctx.Param("username")
	ctx.JSON(http.StatusOK, gin.H{
		"username": username,
		"message":  "Received",
	})
}
