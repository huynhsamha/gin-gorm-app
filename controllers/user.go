package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserCtrl : Controller for User
type UserCtrl struct{}

// FindAll :
func (ctrl UserCtrl) FindAll(ctx *gin.Context) {
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

// FindOneByID :
func (ctrl UserCtrl) FindOneByID(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"id":      id,
		"message": "Received",
	})
}

// FindOneByUsername :
func (ctrl UserCtrl) FindOneByUsername(ctx *gin.Context) {
	username := ctx.Param("username")
	ctx.JSON(http.StatusOK, gin.H{
		"username": username,
		"message":  "Received",
	})
}
