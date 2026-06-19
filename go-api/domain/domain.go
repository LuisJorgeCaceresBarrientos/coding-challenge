package domain

import "github.com/golang-jwt/jwt/v5"

// User representa la estructura de un usuario registrado en el sistema.
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// JwtCustomClaims define la estructura del token JWT que generamos al iniciar sesión.
type JwtCustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// MatrixRequest es la estructura del cuerpo de la petición para procesar la matriz.
type MatrixRequest struct {
	Matrix [][]float64 `json:"matrix"`
}

// StatisticsResponse representa las estadísticas calculadas por la API de Node.js.
type StatisticsResponse struct {
	MaxValue        float64 `json:"maxValue"`
	MinValue        float64 `json:"minValue"`
	Average         float64 `json:"average"`
	TotalSum        float64 `json:"totalSum"`
	IsQDiagonal     bool    `json:"isQDiagonal"`
	IsRDiagonal     bool    `json:"isRDiagonal"`
	IsRotatedDiagonal bool  `json:"isRotatedDiagonal"`
}

// MatrixResponse es la respuesta completa que se le devuelve al cliente.
type ProcessResponse struct {
	Original [][]float64        `json:"original"`
	Rotated  [][]float64        `json:"rotated"`
	Q        [][]float64        `json:"q"`
	R        [][]float64        `json:"r"`
	Stats    StatisticsResponse `json:"stats"`
}
