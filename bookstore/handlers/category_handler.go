package handlers

import (
	"bookstore/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var categories []models.Category
var nextCategoryID = 1

func GetCategories(c *gin.Context) {
	c.JSON(http.StatusOK, categories)
}

func CreateCategory(c *gin.Context) {
	var newCategory models.Category
	if err := c.ShouldBindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newCategory.ID = nextCategoryID
	nextCategoryID++
	categories = append(categories, newCategory)
	c.JSON(http.StatusCreated, newCategory)
}
