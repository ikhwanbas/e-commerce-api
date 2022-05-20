package controllers

import (
	"ecommerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// RegisterUser godoc
// @Summary Register New User.
// @Description Registering a new User.
// @Tags Register User
// @Param type body models.RegisterInput true "the body to register a new user"
// @Produce json
// @Success 200 {object} models.Users
// @Router /user [post]
func InputProductCategory(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var inputCategory models.ProductCategoryInput
	if err := c.ShouldBindJSON(&inputCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Product Category
	productCategory := models.ProductCategory{
		ID:           uuid.New(),
		CategoryName: inputCategory.CategoryName,
	}

	var err error = db.Create(&productCategory).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"category_name": productCategory.CategoryName})
}

func GetProductCategory(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var productCategory []models.ProductCategory
	db.Find(&productCategory)

	c.JSON(http.StatusOK, gin.H{"data": productCategory})
}
