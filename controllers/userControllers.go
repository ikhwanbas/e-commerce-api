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
func UpdateUserProfile(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	id, _ := token.ExtractTokenID(c)

	// Get model if exist
	var UserProfiles models.UsersProfile
	if err := db.Where("id = ?", id).First(&UserProfiles).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// Validate input
	var input models.UsersProfileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Update User Profile
	var userProfile models.UsersProfile
	userProfile.Fullname = input.Fullname
	userProfile.Address = input.Address
	userProfile.PhoneNumber = input.PhoneNumber
	userProfile.City = input.City
	userProfile.PostalCode = input.PostalCode

	// input to database
	var err error = db.Model(&UserProfiles).Updates(userProfile).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": UserProfiles})
}

// superuser
func GetAllUser(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var users []models.Users
	db.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}
