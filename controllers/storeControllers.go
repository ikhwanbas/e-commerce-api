package controllers

import (
	"ecommerce/models"
	"ecommerce/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UpdateUserProfile godoc
// @Summary Update User Profile.
// @Description Updating a User Profile data.
// @Tags Update User Profile
// @Param type body models.UserProfile true "the body to update user profile"
// @Produce json
// @Success 200 {object} models.UserProfiles
// @Router /user [post]
func UpdateStoreProfile(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	id, _ := token.ExtractTokenID(c)

	// Get model if exist
	var storeProfile models.StoresProfile
	if err := db.Where("id = ?", id).First(&storeProfile).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// Validate input
	var input models.StoreProfileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Update Store Profile
	var newProfile models.StoresProfile
	newProfile.Address = input.Address
	newProfile.PhoneNumber = input.PhoneNumber
	newProfile.City = input.City
	newProfile.PostalCode = input.PostalCode

	// input to database
	var err error = db.Model(&storeProfile).Updates(newProfile).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": storeProfile})
}

// superuser
func GetAllStore(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var stores []models.Stores
	db.Find(&stores)

	c.JSON(http.StatusOK, gin.H{"data": stores})
}
