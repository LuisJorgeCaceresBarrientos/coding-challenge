# Usamos la imagen oficial de Node.js (versión 22 en su sabor ligero Alpine)
FROM node:22-alpine

# Establecemos el directorio de trabajo dentro del contenedor
WORKDIR /usr/src/app

# Copiamos los archivos de manifiesto de dependencias
COPY package*.json ./

# Instalamos únicamente dependencias necesarias para ejecución (producción)
RUN npm ci --only=production

# Copiamos el código fuente de la aplicación al contenedor
COPY . .

# Exponemos el puerto de escucha (5000 por defecto para Node)
EXPOSE 5000

# Iniciamos el servidor Express
CMD ["node", "src/index.js"]
