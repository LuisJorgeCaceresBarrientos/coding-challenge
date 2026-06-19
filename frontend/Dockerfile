# Fase 1: Compilación de la aplicación Angular
FROM node:22-alpine AS builder

WORKDIR /usr/src/app

COPY package*.json ./
RUN npm ci

COPY . .
# Compilamos la aplicación para producción
RUN npm run build -- --configuration=production

# Fase 2: Servidor web ligero Nginx para entregar el frontend
FROM nginx:1.25-alpine

# Copiamos los archivos compilados de Angular al directorio que sirve Nginx.
# Nota: Dependiendo de la versión de Angular, la salida puede estar en dist/frontend/browser o dist/frontend.
# Validaremos la ruta copiando directamente el resultado de dist.
COPY --from=builder /usr/src/app/dist/frontend/browser /usr/share/nginx/html

# Exponemos el puerto 80 para el tráfico web
EXPOSE 80

# Iniciamos Nginx
CMD ["nginx", "-g", "daemon off;"]
