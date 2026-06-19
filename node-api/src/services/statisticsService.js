import { isDiagonalMatrix } from '../utils/matrixUtils.js';

/**
 * Lógica de negocio para calcular las estadísticas basadas en las tres matrices.
 * Extrae todos los valores, calcula máximos, mínimos, sumas, promedios y verifica
 * las condiciones matemáticas (si son matrices diagonales).
 * 
 * @param {number[][]} rotated - Matriz rotada.
 * @param {number[][]} q - Factorización Q.
 * @param {number[][]} r - Factorización R.
 * @returns {Object} Objeto con las estadísticas calculadas.
 */
export function calculateStatistics(rotated, q, r) {
    // Unimos todos los valores de las tres matrices en un único arreglo para calcular las estadísticas globales.
    const allValues = [];
    
    // Función auxiliar interna para extraer los valores de una matriz específica
    const extractValues = (matrix) => {
        for (let i = 0; i < matrix.length; i++) {
            for (let j = 0; j < matrix[i].length; j++) {
                allValues.push(matrix[i][j]);
            }
        }
    };

    extractValues(rotated);
    extractValues(q);
    extractValues(r);

    // Si las matrices estaban completamente vacías, lanzamos un error de negocio
    if (allValues.length === 0) {
        throw new Error('Las matrices no contienen elementos');
    }

    // Inicializamos variables para los cálculos estadísticos
    let maxValue = -Infinity;
    let minValue = Infinity;
    let totalSum = 0;

    // Recorremos los valores extraídos para calcular máximo, mínimo y sumatoria total
    for (let i = 0; i < allValues.length; i++) {
        const val = allValues[i];
        if (val > maxValue) {
            maxValue = val;
        }
        if (val < minValue) {
            minValue = val;
        }
        totalSum += val;
    }

    // Calculamos el promedio
    const average = totalSum / allValues.length;

    // Verificamos de forma independiente si alguna de las tres matrices es diagonal
    const isRotatedDiagonal = isDiagonalMatrix(rotated);
    const isQDiagonal = isDiagonalMatrix(q);
    const isRDiagonal = isDiagonalMatrix(r);

    // Devolvemos el resultado procesado
    return {
        maxValue,
        minValue,
        average,
        totalSum,
        isRotatedDiagonal,
        isQDiagonal,
        isRDiagonal
    };
}
