package controllers

import (
	"ecommerce/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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
func RegisterUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input models.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create User
	user := models.Users{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}

	_, err := user.SaveUser(db)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"username": user.Username, "email": user.Email})
}

// LoginUser godoc
// @Summary Login as a user.
// @Description Logging in to get jwt token to access api by roles user.
// @Tags Login User
// @Param type body models.LoginInput true "the body to login a user"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /login [post]
func LoginUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input models.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.Users{}

	u.Username = input.Username
	u.Password = input.Password
	u.Email = input.Email

	token, err := models.LoginCheck(u.Email, u.Username, u.Password, db)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "login success", "token": token})

}
