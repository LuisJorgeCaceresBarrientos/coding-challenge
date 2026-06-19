import jwt from 'jsonwebtoken';
import { JWT_SECRET } from '../config.js';

/**
 * Middleware para validar el token de autorización JWT en las peticiones.
 * Intercepta la solicitud antes de que llegue al controlador para garantizar
 * que el usuario esté autenticado.
 */
export function authenticateJWT(req, res, next) {
    // Obtenemos la cabecera de autorización de la solicitud HTTP
    const authHeader = req.headers.authorization;

    if (!authHeader) {
        return res.status(401).json({ error: 'Falta la cabecera de Autorización' });
    }

    // Dividimos la cabecera esperando el formato 'Bearer <token>'
    const parts = authHeader.split(' ');
    if (parts.length !== 2 || parts[0].toLowerCase() !== 'bearer') {
        return res.status(401).json({ error: 'Formato de token inválido' });
    }

    const token = parts[1];

    // Verificamos el token JWT contra nuestra clave secreta centralizada
    jwt.verify(token, JWT_SECRET, (err, user) => {
        if (err) {
            return res.status(403).json({ error: 'Token inválido o expirado' });
        }
        
        // Adjuntamos la información del usuario al objeto de la petición (req)
        req.user = user;
        
        // Continuamos con el flujo hacia el siguiente middleware o controlador
        next();
    });
}
