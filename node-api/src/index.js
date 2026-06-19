import express from 'express';
import cors from 'cors';
import { PORT } from './config.js';
import statisticsRoutes from './routes/statisticsRoutes.js';

const app = express();

// Habilitamos CORS para permitir peticiones desde el frontend en Angular o la API en Go.
app.use(cors());

// Habilitamos el análisis de cuerpos JSON en las peticiones HTTP.
app.use(express.json());

// Montamos todas las rutas relacionadas con las estadísticas bajo el prefijo /api.
app.use('/api', statisticsRoutes);

// Iniciamos el servidor Express en el puerto configurado.
app.listen(PORT, () => {
    console.log(`Node.js API running on port ${PORT}...`);
});
