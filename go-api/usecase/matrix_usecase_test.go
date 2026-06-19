package usecase

import (
	"math"
	"testing"
)

// TestRotateMatrix90Clockwise verifica que la rotación de 90 grados sea correcta.
func TestRotateMatrix90Clockwise(t *testing.T) {
	mu := NewMatrixUsecase()

	// Matriz original de 3 filas x 2 columnas:
	// [ 1  2 ]
	// [ 3  4 ]
	// [ 5  6 ]
	original := [][]float64{
		{1, 2},
		{3, 4},
		{5, 6},
	}

	// Resultado esperado después de rotar 90° horario (2 filas x 3 columnas):
	// [ 5  3  1 ]
	// [ 6  4  2 ]
	expected := [][]float64{
		{5, 3, 1},
		{6, 4, 2},
	}

	rotated, err := mu.RotateMatrix90Clockwise(original)
	if err != nil {
		t.Fatalf("No se esperaba un error al rotar: %v", err)
	}

	// Verificamos que las dimensiones sean las correctas
	if len(rotated) != len(expected) || len(rotated[0]) != len(expected[0]) {
		t.Fatalf("Dimensiones incorrectas. Esperado %dx%d, obtenido %dx%d",
			len(expected), len(expected[0]), len(rotated), len(rotated[0]))
	}

	// Verificamos cada elemento uno por uno
	for i := range expected {
		for j := range expected[i] {
			if rotated[i][j] != expected[i][j] {
				t.Errorf("Error en (%d, %d). Esperado: %f, Obtenido: %f",
					i, j, expected[i][j], rotated[i][j])
			}
		}
	}
}

// TestComputeQRDecomposition verifica que la descomposición QR sea correcta usando Gonum.
//
// La propiedad que validamos es: A ≈ Q × R
// Si reconstruimos la matriz multiplicando Q y R, deberíamos obtener la original.
func TestComputeQRDecomposition(t *testing.T) {
	mu := NewMatrixUsecase()

	// Matriz original de prueba (3x3)
	original := [][]float64{
		{12, -51, 4},
		{6, 167, -68},
		{-4, 24, -41},
	}

	q, r, err := mu.ComputeQRDecomposition(original)
	if err != nil {
		t.Fatalf("No se esperaba un error al calcular QR: %v", err)
	}

	m := len(original)
	n := len(original[0])

	// 1. Verificar que Q tiene las dimensiones correctas (M x N)
	if len(q) != m || len(q[0]) != n {
		t.Errorf("Q tiene dimensiones incorrectas: %dx%d (esperado %dx%d)", len(q), len(q[0]), m, n)
	}

	// 2. Verificar que R tiene las dimensiones correctas (N x N)
	if len(r) != n || len(r[0]) != n {
		t.Errorf("R tiene dimensiones incorrectas: %dx%d (esperado %dx%d)", len(r), len(r[0]), n, n)
	}

	// 3. Verificar que A ≈ Q × R reconstruyendo la matriz original
	tolerance := 0.01 // Tolerancia de 1 centésima para errores de redondeo

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// Calculamos Q[i,:] × R[:,j] (producto punto de fila i de Q con columna j de R)
			sum := 0.0
			for k := 0; k < n; k++ {
				sum += q[i][k] * r[k][j]
			}

			// La diferencia entre lo reconstruido y lo original debe ser muy pequeña
			diff := math.Abs(sum - original[i][j])
			if diff > tolerance {
				t.Errorf("QR reconstruida difiere en (%d,%d): original=%f, reconstruida=%f, diff=%f",
					i, j, original[i][j], sum, diff)
			}
		}
	}
}
