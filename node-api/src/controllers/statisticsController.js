import { calculateStatistics } from '../services/statisticsService.js';

/**
 * Controlador web para la ruta de estadísticas.
 * Su responsabilidad es manejar el requerimiento HTTP (req), extraer los datos del cuerpo (body),
 * pasárselos al servicio que contiene la lógica de negocio, y devolver la respuesta (res).
 */
export function getStatistics(req, res) {
    const { rotated, q, r } = req.body;

    // Validamos que la solicitud contenga los datos requeridos
    if (!rotated || !q || !r) {
        return res.status(400).json({ error: 'Se requieren las matrices rotated, q y r' });
    }

    try {
        // Llamamos al caso de uso / servicio para procesar la información
        const stats = calculateStatistics(rotated, q, r);
        
        // Respondemos con éxito retornando los cálculos en formato JSON
        res.json(stats);
    } catch (error) {
        // Capturamos cualquier error de negocio lanzado por el servicio (ej. matrices vacías)
        res.status(400).json({ error: error.message });
    }
}
