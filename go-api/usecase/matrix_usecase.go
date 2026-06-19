package usecase

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"os"
	"time"

	// gonum/mat es la librería de álgebra lineal de Go.
	// Es el equivalente al módulo "numpy.linalg" de Python.
	// Nos da funciones ya programadas para matrices: QR, SVD, inversas, etc.
	"gonum.org/v1/gonum/mat"

	"go-api/domain"
)

// MatrixUsecase es nuestra "caja de herramientas" para todo lo relacionado con matrices.
type MatrixUsecase struct {
	nodeAPIURL string // Guarda la URL de la API de Node.js
}

// NewMatrixUsecase crea una nueva instancia de MatrixUsecase.
// Lee la URL de Node.js de las variables de entorno (configuración de Docker).
func NewMatrixUsecase() *MatrixUsecase {
	url := os.Getenv("NODE_API_URL")
	if url == "" {
		// Si no hay variable de entorno, usamos localhost para desarrollo local
		url = "http://localhost:5000"
	}
	return &MatrixUsecase{nodeAPIURL: url}
}

// RotateMatrix90Clockwise rota una matriz 90 grados en sentido horario.
//
// La fórmula es simple: el elemento en la posición [fila][columna]
// de la original pasa a la posición [columna][filas_totales - 1 - fila] en la nueva.
//
// Ejemplo visual:       Original        →       Rotada
//
//	[ 1  2 ]                      [ 5  3  1 ]
//	[ 3  4 ]            →         [ 6  4  2 ]
//	[ 5  6 ]
func (mu *MatrixUsecase) RotateMatrix90Clockwise(matrix [][]float64) ([][]float64, error) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil, errors.New("la matriz no puede estar vacía")
	}

	rows := len(matrix)    // Cuántas filas tiene la matriz original
	cols := len(matrix[0]) // Cuántas columnas tiene la matriz original

	// La nueva matriz tiene las dimensiones invertidas: cols filas y rows columnas.
	rotated := make([][]float64, cols)
	for i := range rotated {
		rotated[i] = make([]float64, rows)
	}

	// Aplicamos la fórmula de rotación 90° horaria elemento por elemento
	for i := 0; i < rows; i++ {
		if len(matrix[i]) != cols {
			return nil, errors.New("la matriz debe ser rectangular (todas las filas con la misma longitud)")
		}
		for j := 0; j < cols; j++ {
			rotated[j][rows-1-i] = matrix[i][j]
		}
	}

	return rotated, nil
}

// ComputeQRDecomposition calcula la descomposición QR de una matriz usando la librería Gonum.
//
// ¿Qué hace en palabras simples?
//   - Recibe una matriz A.
//   - Devuelve dos matrices Q y R tal que: A = Q × R
//   - Q: tiene columnas perpendiculares entre sí, de longitud 1.
//   - R: es triangular superior (ceros debajo de la diagonal).
//
// Antes, esto lo hacíamos a mano con el algoritmo Gram-Schmidt (~80 líneas).
// Con Gonum, son solo 3 pasos.
func (mu *MatrixUsecase) ComputeQRDecomposition(matrix [][]float64) ([][]float64, [][]float64, error) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil, nil, errors.New("la matriz no puede estar vacía")
	}

	numRows := len(matrix)    // Número de filas (M)
	numCols := len(matrix[0]) // Número de columnas (N)

	// --- PASO 1: Convertir nuestro [][]float64 al formato que entiende Gonum ---
	//
	// mat.NewDense() crea una matriz de Gonum.
	// Necesita una sola lista plana de números (no una lista de listas).
	// Por eso usamos este bucle para "aplanar" nuestra matriz.
	flatData := make([]float64, 0, numRows*numCols)
	for i := 0; i < numRows; i++ {
		flatData = append(flatData, matrix[i]...)
	}
	// Creamos la matriz de Gonum con los datos aplanados
	gonumMatrix := mat.NewDense(numRows, numCols, flatData)

	// --- PASO 2: Calcular la factorización QR con Gonum ---
	//
	// mat.QR es la estructura que Gonum usa para la descomposición QR.
	// qrDecomp.Factorize(gonumMatrix) hace TODO el cálculo internamente.
	// Equivalente a: numpy.linalg.qr(matrix) en Python.
	var qrDecomp mat.QR
	qrDecomp.Factorize(gonumMatrix)

	// --- PASO 3: Extraer Q y R de la descomposición ---
	//
	// QTo() extrae la matriz Q y la guarda en la variable qMatrix.
	// RTo() extrae la matriz R y la guarda en la variable rMatrix.
	var qMatrix, rMatrix mat.Dense
	qrDecomp.QTo(&qMatrix)
	qrDecomp.RTo(&rMatrix)

	// --- PASO 4: Convertir los resultados de vuelta a nuestro formato [][]float64 ---
	//
	// Gonum maneja sus propios tipos de matrices.
	// Necesitamos convertirlos de vuelta a listas de listas para poder enviarlos como JSON.
	q := convertGonumToSlice(&qMatrix, numRows, numCols)
	r := convertGonumToSlice(&rMatrix, numCols, numCols)

	return q, r, nil
}

// convertGonumToSlice convierte una matriz de Gonum a una lista de listas ([][]float64).
//
// Es una función de ayuda (helper) que solo usamos internamente.
// Redondea cada número a 6 decimales para evitar basura de precisión flotante
// como: 2.0000000000000018 → simplificado a: 2.0
func convertGonumToSlice(m *mat.Dense, rows, cols int) [][]float64 {
	result := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		result[i] = make([]float64, cols)
		for j := 0; j < cols; j++ {
			// Redondeamos a 6 decimales para tener resultados limpios en pantalla
			result[i][j] = math.Round(m.At(i, j)*1e6) / 1e6
		}
	}
	return result
}

// GetStatisticsFromNodeAPI envía las tres matrices calculadas a la API de Node.js.
//
// Flujo: Go calcula Q, R y la rotada → los envía a Node.js → Node.js calcula estadísticas.
func (mu *MatrixUsecase) GetStatisticsFromNodeAPI(jwtToken string, rotated, q, r [][]float64) (*domain.StatisticsResponse, error) {
	// Preparamos el cuerpo de la petición con las tres matrices en formato JSON
	payload := map[string][][]float64{
		"rotated": rotated,
		"q":       q,
		"r":       r,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error al serializar matrices para Node.js: %w", err)
	}

	// Construimos la URL del endpoint de estadísticas
	reqURL := fmt.Sprintf("%s/api/statistics", mu.nodeAPIURL)

	// Creamos la petición HTTP POST con los datos JSON
	req, err := http.NewRequest("POST", reqURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error al crear petición HTTP: %w", err)
	}

	// Agregamos las cabeceras necesarias:
	// - Content-Type le dice a Node.js que el cuerpo viene en formato JSON
	// - Authorization reenvía el token JWT del usuario para mantener la seguridad
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", jwtToken)

	// Enviamos la petición con un timeout de 10 segundos (para no esperar eternamente)
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error de comunicación con la API de Node.js: %w", err)
	}
	defer resp.Body.Close() // "defer" asegura que se cierre la conexión al terminar la función

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("la API de Node.js devolvió un error: código %d", resp.StatusCode)
	}

	// Decodificamos la respuesta JSON de Node.js a nuestra estructura de datos
	var statsResponse domain.StatisticsResponse
	if err := json.NewDecoder(resp.Body).Decode(&statsResponse); err != nil {
		return nil, fmt.Errorf("error al leer las estadísticas de Node.js: %w", err)
	}

	return &statsResponse, nil
}
