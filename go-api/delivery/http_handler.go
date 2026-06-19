package delivery

import (
	"fmt"
	"go-api/domain"
	"go-api/usecase"

	"github.com/gofiber/fiber/v2"
)

type HttpHandler struct {
	authUsecase   *usecase.SimpleAuthUsecase
	matrixUsecase *usecase.MatrixUsecase
}

func NewHttpHandler(au *usecase.SimpleAuthUsecase, mu *usecase.MatrixUsecase) *HttpHandler {
	return &HttpHandler{
		authUsecase:   au,
		matrixUsecase: mu,
	}
}



// Register maneja la petición de registro de un nuevo usuario.
func (h *HttpHandler) Register(c *fiber.Ctx) error {
	var user domain.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Formato JSON inválido",
		})
	}

	err := h.authUsecase.Register(user.Username, user.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Usuario registrado exitosamente",
	})
}

// Login maneja el inicio de sesión y devuelve el token JWT generado.
func (h *HttpHandler) Login(c *fiber.Ctx) error {
	var user domain.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Formato JSON inválido",
		})
	}

	token, err := h.authUsecase.Login(user.Username, user.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

// ProcessMatrix maneja el flujo de recibir, rotar, factorizar la matriz y consultar estadísticas.
func (h *HttpHandler) ProcessMatrix(c *fiber.Ctx) error {
	var req domain.MatrixRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Formato JSON inválido en la matriz",
		})
	}

	// 1. Rota la matriz 90 grados en sentido horario
	rotated, err := h.matrixUsecase.RotateMatrix90Clockwise(req.Matrix)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// 2. Calcula la factorización QR de la matriz original
	q, r, err := h.matrixUsecase.ComputeQRDecomposition(req.Matrix)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Error al realizar la factorización QR: %s", err.Error()),
		})
	}

	// 3. Envía los resultados a la API de Node.js reenviando el token de autorización
	jwtToken := c.Get("Authorization")
	stats, err := h.matrixUsecase.GetStatisticsFromNodeAPI(jwtToken, rotated, q, r)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": fmt.Sprintf("Error al comunicarse con el módulo de estadísticas: %s", err.Error()),
		})
	}

	// 4. Construye y retorna la respuesta final estructurada
	response := domain.ProcessResponse{
		Original: req.Matrix,
		Rotated:  rotated,
		Q:        q,
		R:        r,
		Stats:    *stats,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
