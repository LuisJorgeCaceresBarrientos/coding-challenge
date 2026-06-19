import assert from 'assert';
import { isDiagonalMatrix } from './src/utils/matrixUtils.js';

// Caso de prueba 1: Matriz diagonal
const diagonal = [
    [5, 0, 0],
    [0, 8, 0],
    [0, 0, 1]
];
assert.strictEqual(isDiagonalMatrix(diagonal), true, "Debería identificar una matriz diagonal");

// Caso de prueba 2: Matriz no diagonal
const nonDiagonal = [
    [5, 1, 0],
    [0, 8, 0],
    [0, 0, 1]
];
assert.strictEqual(isDiagonalMatrix(nonDiagonal), false, "Debería identificar una matriz que NO es diagonal");

// Caso de prueba 3: Matriz rectangular (no cuadrada)
const rectangular = [
    [5, 0],
    [0, 8],
    [0, 0]
];
assert.strictEqual(isDiagonalMatrix(rectangular), false, "Una matriz no cuadrada no debe considerarse diagonal");

console.log("¡Todas las pruebas unitarias de Node.js pasaron exitosamente!");
