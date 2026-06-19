package main

import (
	"log"
	"os"
	"go-api/delivery"
	"go-api/delivery/middleware"
	"go-api/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// 1. Inicializamos la aplicación de Fiber
	app := fiber.New(fiber.Config{
		AppName: "Interseguro Coding Challenge - Go API",
	})

	// 2. Agregamos middleware de registro de peticiones en consola
	app.Use(logger.New())

	// 3. Configuramos CORS para permitir peticiones desde cualquier origen (necesario para Angular)
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, OPTIONS",
	}))

	// 4. Instanciamos las capas de la arquitectura limpia
	authUsecase := usecase.NewSimpleAuthUsecase()
	matrixUsecase := usecase.NewMatrixUsecase()
	handler := delivery.NewHttpHandler(authUsecase, matrixUsecase)

	// 5. Definición de rutas y endpoints del API
	api := app.Group("/api")

	// Rutas públicas de autenticación
	auth := api.Group("/auth")
	auth.Post("/register", handler.Register)
	auth.Post("/login", handler.Login)

	// Rutas protegidas mediante el middleware de JWT
	matrix := api.Group("/matrix")
	matrix.Use(middleware.JWTMiddleware)
	matrix.Post("/process", handler.ProcessMatrix)

	// 6. Iniciamos el servidor en el puerto configurado (8080 por defecto)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor API Go escuchando en el puerto %s...", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Error al iniciar el servidor Fiber: %v", err)
	}
}
