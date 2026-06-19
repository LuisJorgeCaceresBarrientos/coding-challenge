import { Router } from 'express';
import { getStatistics } from '../controllers/statisticsController.js';
import { authenticateJWT } from '../middlewares/authMiddleware.js';

const router = Router();

// Definimos la ruta POST protegida para las estadísticas.
// Primero pasa por el middleware de validación JWT, y si es exitoso, llega al controlador.
router.post('/statistics', authenticateJWT, getStatistics);

export default router;
