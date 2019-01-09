package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/huynhsamha/gin-gorm-app/models"
	"github.com/huynhsamha/gin-gorm-app/utils"
	funk "github.com/thoas/go-funk"
)

// UserCtrl : Controller for User
type UserCtrl struct{}

// FindAll : Search user with query string
/**
 * Query: { page, sort[], keywords }
 *
 * Example: ?page=2&sort=username&sort=email:asc&sort=name:desc&keywords=alice%20uk
 *
 * Response: { totalRecords, totalPages, perPage, page, offset, records[] }
 */
func (ctrl UserCtrl) FindAll(ctx *gin.Context) {
	page, _ := strconv.ParseUint(ctx.DefaultQuery("page", "1"), 10, 32)
	keywords := ctx.DefaultQuery("keywords", "")
	sort := ctx.QueryArray("sort")

	/** Get order query - ORDER in SQL */
	sortableColumns := []string{"id", "username", "email", "name"}

	if len(sort) > len(sortableColumns) {
		// Perhaps you're trying to hack me :P
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Sort columns are too much with allowed columns.",
		})
		return
	}

	var orderQueries []string
	for _, v := range sort {
		sortPair := strings.Split(v, ":")
		if funk.Contains(sortableColumns, sortPair[0]) {
			order := "asc"
			if len(sortPair) > 1 {
				order = utils.DefaultStringEmpty(sortPair[1], "asc")
			}
			orderQueries = append(orderQueries, sortPair[0]+" "+order)
		}
	}
	orderQueryString := strings.Join(orderQueries, ",")

	/** Paginate matching records - OFFSET/LIMIT in SQL */
	perPage := 10
	offset := (int(page) - 1) * perPage

	/** Create Keywords Query String - LIKE in SQL */
	keywordsColumns := []string{"username", "email", "name", "location", "title", "github", "twitter"}
	keywordsQueryColumns := funk.Map(keywordsColumns, func(i string) string {
		return i + " LIKE ?"
	})
	keywordsQueryString := strings.Join(keywordsQueryColumns.([]string), " OR ")
	// fmt.Println(keywordsQueryString)
	keywords = "%" + keywords + "%"
	keywordsArray := make([]interface{}, len(keywordsColumns))
	for i := range keywordsArray {
		keywordsArray[i] = keywords
	}

	/** Retrieve matching records in Database */
	var totalRecords int
	var records []models.User

	res := db.Model(&models.User{}).
		// Query
		Where(keywordsQueryString, keywordsArray...).
		// Count total records
		Count(&totalRecords).
		// Query records
		Order(orderQueryString).       // order
		Offset(offset).Limit(perPage). // pagination
		Find(&records)

	if res.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": res.Error.Error()})
		return
	}

	totalPages := (totalRecords + perPage - 1) / perPage

	ctx.JSON(http.StatusOK, gin.H{
		"totalRecords": totalRecords,
		"totalPages":   totalPages,
		"perPage":      perPage,
		"page":         page,
		"offset":       offset,
		"records":      records,
	})
}

// FindOneByID : Find user by ID
func (ctrl UserCtrl) FindOneByID(ctx *gin.Context) {
	id := ctx.Param("id")
	user := models.User{}
	res := db.First(&user, id)
	if res.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// FindOneByUsername : Find user by Username
func (ctrl UserCtrl) FindOneByUsername(ctx *gin.Context) {
	username := ctx.Param("username")
	user := models.User{}
	res := db.Where(&models.User{Username: username}).First(&user)
	if res.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
