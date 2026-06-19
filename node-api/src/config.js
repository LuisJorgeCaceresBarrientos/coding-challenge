// Centralizamos las variables de entorno o constantes de configuración de la aplicación.
// Esto permite que sea fácil de modificar en el futuro, posiblemente leyendo de un archivo .env.

export const JWT_SECRET = 'my-super-secret-key-12345';
export const PORT = process.env.PORT || 5000;
