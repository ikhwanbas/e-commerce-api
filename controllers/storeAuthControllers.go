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
func RegisterStore(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var inputStore models.StoreRegisterInput
	if err := c.ShouldBindJSON(&inputStore); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create User
	store := models.Stores{
		StoreName: inputStore.StoreName,
		Email:     inputStore.Email,
		Password:  inputStore.Password,
	}

	_, err := store.SaveStore(db)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"store_name": store.StoreName, "email": store.Email})
}

// LoginUser godoc
// @Summary Login as a user.
// @Description Logging in to get jwt token to access api by roles user.
// @Tags Login User
// @Param type body models.LoginInput true "the body to login a user"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /login [post]
func LoginStore(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input models.StoreLoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	s := models.Stores{}

	s.Password = input.Password
	s.Email = input.Email

	token, err := models.StoreLoginCheck(s.Email, s.Password, db)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "email or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "login success", "token": token})

}
