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

// QuestionCtrl : Controller for Question
type QuestionCtrl struct{}

// FindAll : Search question with query string
/**
 * Query: { page, sort[], keywords }
 *
 * Example: ?page=2&sort=questionname&sort=email:asc&sort=name:desc&keywords=alice%20uk
 *
 * Response: { totalRecords, totalPages, perPage, page, offset, records[] }
 */
func (ctrl QuestionCtrl) FindAll(ctx *gin.Context) {
	page, _ := strconv.ParseUint(ctx.DefaultQuery("page", "1"), 10, 32)
	keywords := ctx.DefaultQuery("keywords", "")
	sort := ctx.QueryArray("sort")

	/** Get order query - ORDER in SQL */
	sortableColumns := []string{"id", "title", "votes"}

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
	keywordsColumns := []string{"questionname", "email", "name", "location", "title", "github", "twitter"}
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
	var records []models.Question

	res := db.Model(&models.Question{}).
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

// FindOneByID : Find question by ID
func (ctrl QuestionCtrl) FindOneByID(ctx *gin.Context) {
	id := ctx.Param("id")
	question := models.Question{}
	res := db.First(&question, id)
	if res.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	}
	ctx.JSON(http.StatusOK, question)
}

type formEditQuestion struct {
	Title   string `form:"title" json:"title" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
}

// Create : validate form, create question
func (ctrl QuestionCtrl) Create(ctx *gin.Context) {
	payload, _ := AuthCtrl{}.getPayload(ctx)

	var form formEditQuestion
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	question := models.Question{
		Title:   form.Title,
		Content: form.Content,
		OwnerID: payload.UserID,
	}

	if err := db.Create(&question).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Create question successfully",
		"question": question,
	})
}

// type formEditQuestion struct {
// 	Title   string `form:"title" json:"title"`
// 	Content string `form:"content" json:"content"`
// }

// // EditQuestion : edit my question
// func (ctrl QuestionCtrl) EditQuestion(ctx *gin.Context) {
// 	payload, _ := AuthCtrl{}.getPayload(ctx)
// 	id := ctx.Param("id")

// 	var form formEditQuestion
// 	if err := ctx.ShouldBind(&form); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	res := db.Model(&models.Question{}).
// 		Where("id = ?", id).
// 		Updates(models.Question{
// 			Title:   form.Title,
// 			Content: form.Content,
// 		})

// 	if res.Error != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": res.Error.Error()})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"message": "Edit question successfully",
// 	})
// }
