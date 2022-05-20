package routes

import (
	"ecommerce/controllers"
	"ecommerce/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies([]string{"localhost"})

	// CORS Setting
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"}
	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true
	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("OPTIONS")
	r.Use(cors.New(corsConfig))

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	// ROUTES
	// USER routes
	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)

	// user routes with JWT
	jwt := r.Group("/")
	jwt.Use(middlewares.JwtAuthMiddleware())

	jwt.PUT("/profile", controllers.UpdateUserProfile)

	// CART Routes
	jwt.GET("/cart", controllers.GetAllCart)
	jwt.GET("/cart/:ids", controllers.GetCartByUserIdAndStoreId)
	jwt.POST("/cart", controllers.CreateNewCart)
	jwt.PUT("/cart/:id", controllers.UpdateCart)
	jwt.DELETE("/cart/:id", controllers.DeleteCart)

	// SELLER Routes
	r.POST("/store/register", controllers.RegisterStore)
	r.POST("/store/login", controllers.LoginStore)

	jwt.PUT("/store/profile", controllers.UpdateStoreProfile)

	// PRODUCT Routes
	jwt.GET("/store/products", controllers.GetAllProduct)
	jwt.GET("/store/products/:id", controllers.GetProductById)

	jwt.POST("/store/products", controllers.CreateNewProduct)
	jwt.PUT("/store/products/:id", controllers.UpdateProduct)
	jwt.DELETE("/store/products/:id", controllers.DeleteProduct)

	// SUPER USER Routes
	r.GET("/admin/products-category", controllers.GetProductCategory)
	r.POST("/admin/products-category", controllers.InputProductCategory)
	r.GET("/admin/user", controllers.GetAllUser)
	r.GET("/admin/store", controllers.GetAllStore)

	// swagger routes
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
