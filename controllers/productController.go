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
func CreateNewProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// validate token as store id
	id, _ := token.ExtractTokenID(c)
	s := models.Stores{}
	err := db.Model(s).Where("ID = ?", id).Take(&s).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate Product Input
	var inputProduct models.ProductInput
	if err := c.ShouldBindJSON(&inputProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Product
	product := models.Products{
		ID:                uuid.New(),
		ProductName:       inputProduct.ProductName,
		Stock:             inputProduct.Stock,
		Price:             inputProduct.Price,
		Weight:            inputProduct.Weight,
		Description:       inputProduct.Description,
		ImageURL:          inputProduct.ImageURL,
		ProductCategoryID: inputProduct.ProductCategoryID,
		StoresID:          id,
	}

	err = db.Create(&product).Error

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
func UpdateProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var product models.Products
	if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// Validate input
	var update models.ProductInput
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update Product
	var newProduct models.Products
	newProduct.ProductName = update.ProductName
	newProduct.Stock = update.Stock
	newProduct.Price = update.Price
	newProduct.Weight = update.Weight
	newProduct.Description = update.Description
	newProduct.ImageURL = update.ImageURL
	newProduct.ProductCategoryID = update.ProductCategoryID

	// input to database
	var err error = db.Model(&product).Updates(newProduct).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
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
func DeleteProduct(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var product models.Products
	if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&product)

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
func GetAllProduct(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	// validate token as store id
	id, _ := token.ExtractTokenID(c)

	var product []models.Products
	err := db.Find(&product).Where("stores_id = ?", id).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func GetProductById(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	var product []models.Products
	err := db.First(&product, "id = ?", c.Param("id")).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}
