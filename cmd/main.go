package main

import (
	"log"

	"go-api/controler"
	"go-api/db"
	"go-api/internal/cert"
	"go-api/internal/middleware"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	//Middleware de segurança
	server.Use(middleware.SecurityHeaders())

	//Conexão com o banco de dados
	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	//Camada repository
	ProductRepository := repository.NewProductRepository(dbConnection)
	//Camada usecase
	ProductUsecase := usecase.NewProductUsecase(ProductRepository)
	//Camada controller
	ProductController := controler.NewProductController(ProductUsecase)

	server.GET("/ping", func(ctx *gin.Context) {

		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:productId", ProductController.GetProductsById)
	server.DELETE("/product/:productId", ProductController.DeleteProduct)
	server.PUT("/product/:productId", ProductController.UpdateProduct)

	// Ensure TLS certs exist (generate dev certs if missing)
	if err := cert.EnsureCerts("cert.pem", "key.pem"); err != nil {
		log.Println("failed to ensure certs:", err)
	}

	// Run HTTPS server using generated certs
	if err := server.RunTLS(":8080", "cert.pem", "key.pem"); err != nil {
		log.Fatal(err)
	}
}
