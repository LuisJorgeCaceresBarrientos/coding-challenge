# Fase 1: Compilación de la aplicación
FROM golang:1.24-alpine AS builder

# Establecemos el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiamos los archivos de dependencias de Go
COPY go.mod go.sum ./

# Descargamos las dependencias
RUN go mod download

# Copiamos el resto del código fuente del proyecto
COPY . .

# Compilamos la aplicación de forma estática
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Fase 2: Imagen final ligera para ejecución
FROM alpine:latest

# Instalamos certificados de seguridad (útiles para peticiones HTTPS salientes)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiamos el binario compilado desde la fase anterior
COPY --from=builder /app/main .

# Exponemos el puerto de escucha
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]
