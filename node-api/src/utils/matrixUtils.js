/**
 * Verifica si una matriz es diagonal.
 * Una matriz es diagonal si es cuadrada y todos los elementos fuera de la diagonal principal son cero (o cercanos a cero).
 * 
 * @param {number[][]} matrix - La matriz a verificar.
 * @returns {boolean} Verdadero si es diagonal, falso en caso contrario.
 */
export function isDiagonalMatrix(matrix) {
    // Validamos que la matriz exista y no esté vacía
    if (!matrix || matrix.length === 0 || matrix[0].length === 0) {
        return false;
    }

    const rows = matrix.length;
    const cols = matrix[0].length;

    // Una matriz diagonal por definición matemática debe ser cuadrada
    if (rows !== cols) {
        return false;
    }

    // Margen de tolerancia para errores de precisión de coma flotante
    const tolerance = 1e-9; 

    for (let i = 0; i < rows; i++) {
        for (let j = 0; j < cols; j++) {
            // Si no estamos en la diagonal principal (i !== j) y el valor excede la tolerancia, no es diagonal
            if (i !== j && Math.abs(matrix[i][j]) > tolerance) {
                return false;
            }
        }
    }

    return true;
}
