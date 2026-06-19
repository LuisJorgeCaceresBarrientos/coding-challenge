package usecase

import (
	"errors"
	"sync"
	"time"
	"go-api/domain"

	"github.com/golang-jwt/jwt/v5"
)

// JWTSecret es la clave secreta utilizada para firmar y validar tokens JWT.
// En un entorno de producción real, esto debería provenir de variables de entorno.
var JWTSecret = []byte("my-super-secret-key-12345")

// SimpleAuthUsecase maneja la lógica de autenticación y el almacenamiento en memoria de usuarios.
type SimpleAuthUsecase struct {
	usersMu sync.RWMutex
	users   map[string]string // Guarda usuario -> contraseña (sin cifrar por simplicidad didáctica)
}

// NewSimpleAuthUsecase crea una nueva instancia del caso de uso de autenticación.
func NewSimpleAuthUsecase() *SimpleAuthUsecase {
	return &SimpleAuthUsecase{
		users: make(map[string]string),
	}
}

// Register registra un nuevo usuario en la base de datos en memoria.
func (u *SimpleAuthUsecase) Register(username, password string) error {
	if username == "" || password == "" {
		return errors.New("el usuario y la contraseña no pueden estar vacíos")
	}

	u.usersMu.Lock()
	defer u.usersMu.Unlock()

	// Validamos si el usuario ya existe en nuestra base de datos en memoria.
	if _, exists := u.users[username]; exists {
		return errors.New("el usuario ya está registrado")
	}

	// Guardamos el usuario. En producción, la contraseña SIEMPRE debe ser cifrada usando bcrypt.
	u.users[username] = password
	return nil
}

// Login valida las credenciales y genera un token JWT si son válidas.
func (u *SimpleAuthUsecase) Login(username, password string) (string, error) {
	u.usersMu.RLock()
	dbPassword, exists := u.users[username]
	u.usersMu.RUnlock()

	if !exists || dbPassword != password {
		return "", errors.New("credenciales inválidas")
	}

	// Definimos las reclamaciones (claims) del token, incluyendo el tiempo de expiración (24 horas).
	claims := &domain.JwtCustomClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Creamos el token con el algoritmo de firma HMAC SHA256.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Firmamos el token con nuestra clave secreta.
	tokenString, err := token.SignedString(JWTSecret)
	if err != nil {
		return "", errors.New("error al generar el token de acceso")
	}

	return tokenString, nil
}
