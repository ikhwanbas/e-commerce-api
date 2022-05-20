package controllers

import (
	"ecommerce/models"
	"ecommerce/utils/token"
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
func CreateNewCart(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// extract token as user id
	id, _ := token.ExtractTokenID(c)
	s := models.Users{}
	err := db.Model(s).Where("ID = ?", id).Take(&s).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate Cart Input
	var inputCart models.CartsInput
	if err := c.ShouldBindJSON(&inputCart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Cart
	cart := models.Carts{
		ID:            uuid.New(),
		UsersID:       id,
		PaymentStatus: inputCart.PaymentStatus,
		Quantity:      inputCart.Quantity,
		ProductsID:    inputCart.ProductID,
	}

	err = db.Create(&cart).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// RegisterUser godoc
// @Summary Register New User.
// @Description Registering a new User.
// @Tags Register User
// @Param type body models.RegisterInput true "the body to register a new user"
// @Produce json
// @Success 200 {object} models.Users
// @Router /user [post]
func UpdateCart(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// extract token as user id
	id, _ := token.ExtractTokenID(c)
	s := models.Users{}
	err := db.Model(s).Where("ID = ?", id).Take(&s).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var cart models.Carts
	if err = db.Where("id = ?", c.Param("id")).First(&cart).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// Validate input
	var update models.CartsInput
	if err = c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update Cart
	var newCart models.Carts
	newCart.PaymentStatus = update.PaymentStatus
	newCart.Quantity = update.Quantity

	// input to database
	err = db.Model(&cart).Updates(newCart).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cart})
}

// DeleteMovie godoc
// @Summary Delete one movie.
// @Description Delete a movie by id.
// @Tags Movie
// @Produce json
// @Param id path string true "movie id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /movies/{id} [delete]
func DeleteCart(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var cart models.Carts
	if err := db.Where("id = ?", c.Param("id")).First(&cart).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&cart)

	c.JSON(http.StatusOK, gin.H{"message": "delete success"})
}

// DeleteMovie godoc
// @Summary Delete one movie.
// @Description Delete a movie by id.
// @Tags Movie
// @Produce json
// @Param id path string true "movie id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /movies/{id} [delete]
func GetAllCart(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	// validate token as user id
	id, _ := token.ExtractTokenID(c)

	var cart []models.Carts
	err := db.Find(&cart).Where("user_id = ?", id).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cart})
}

func GetCartByUserIdAndStoreId(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	// extract token as user id
	id, _ := token.ExtractTokenID(c)
	var cart []models.Carts

	err := db.Table("carts").Select("carts.payment_status, carts.quantity, carts.users_id, products.stores_id").Joins("left join products on products.id = carts.products_id").Where("users_id = ? AND stores_id = ?", id, c.Param("ids")).Scan(&cart).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cart})
}
