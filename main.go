package main

import (
	"log"
	"os"
	"product-api/controller"
	"product-api/db"
	"product-api/middleware"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	r := gin.Default()

	db.SetupDatabaseConnection()

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "product-api",
		Key:             []byte("secret key"),
		Timeout:         time.Hour,
		MaxRefresh:      time.Hour,
		IdentityKey:     middleware.IdentityKey,
		PayloadFunc:     middleware.PayloadFunc,
		IdentityHandler: middleware.IdentityHandler,
		Authenticator:   middleware.Authenticator,
		Authorizator:    middleware.Authorizator,
		Unauthorized:    middleware.Unauthorized,
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		SendCookie:      true,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	errInit := authMiddleware.MiddlewareInit()
	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	r.Use(CORSMiddleware())

	r.POST("/login", authMiddleware.LoginHandler)
	admin := r.Group("/admin", controller.LoginAdmin)
	admin.POST("/login", authMiddleware.LoginHandler)
	r.POST("/register", controller.Register)
	r.GET("/logout", authMiddleware.LogoutHandler)

	auth := r.Group("")
	auth.Use(authMiddleware.MiddlewareFunc())
	{

		auth.GET("/self/:username", controller.GetUserByUsername)

		auth.POST("/product", controller.PostProduct)
		auth.GET("/products", controller.GetAllProducts)
		auth.GET("/product", controller.GetProductByID)
		auth.PATCH("/product/patch", controller.UpdateProductByID)
		auth.DELETE("/product/delete", controller.DeleteProductByID)

		auth.POST("/order", controller.PostOrder)
		auth.GET("/orders", controller.GetAllOrders)
		auth.GET("/orders_by_id", controller.GetAllOrderByID)
		auth.GET("/order", controller.GetOrderByID)
		auth.PATCH("/order/patch", controller.UpdateOrderByID)
		auth.DELETE("/order/delete", controller.DeleteOrderByID)

		auth.POST("/order_item", controller.PostOrderItem)
		auth.GET("/order_items", controller.GetAllOrderItems)
		auth.GET("/order_item", controller.GetOrderItemByID)
		auth.PATCH("/order_item/patch", controller.UpdateOrderItemByID)
		auth.DELETE("/order_item/delete", controller.DeleteOrderItemByID)

		auth.POST("/user", controller.PostUser)
		auth.GET("/users", controller.GetAllUsers)
		auth.GET("/user", controller.GetUserByID)
		auth.PATCH("/user/patch", controller.UpdateUserByID)
		auth.DELETE("/user/delete", controller.DeleteUserByID)

		auth.POST("/available", controller.PostAvailable)
		auth.GET("/availables", controller.GetAllAvailables)
		auth.GET("/available", controller.GetAvailableByID)
		auth.GET("/available_product", controller.GetAvailableByProductID)
		auth.GET("/product_by_day", controller.GetProductByDay)
		auth.PATCH("/available/patch", controller.UpdateAvailableByID)
		auth.DELETE("/available/delete", controller.DeleteAvailableByID)

		auth.POST("/order_food", controller.PostOrderFood)

		auth.POST("/upload", controller.FileUpload())

	}

	log.Fatal(r.Run(":" + port))
}


func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
