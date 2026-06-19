package middleware

import (
	"strings"
	"go-api/domain"
	"go-api/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// JWTMiddleware es un middleware para interceptar peticiones y validar el token JWT.
func JWTMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Falta la cabecera de Autorización",
		})
	}

	// Esperamos el formato estándar: "Bearer <token>"
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Formato de token inválido (debe ser: Bearer <token>)",
		})
	}

	tokenString := parts[1]

	// Parseamos y verificamos la firma del token JWT utilizando la clave secreta
	token, err := jwt.ParseWithClaims(tokenString, &domain.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validamos el método de firma (debe ser HMAC)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}
		return usecase.JWTSecret, nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token inválido o expirado",
		})
	}

	// Almacenamos los claims en la sesión de la petición (Locals) para que otros controladores puedan usarlos.
	claims, ok := token.Claims.(*domain.JwtCustomClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "No se pudieron procesar las credenciales del token",
		})
	}

	c.Locals("username", claims.Username)
	return c.Next()
}
