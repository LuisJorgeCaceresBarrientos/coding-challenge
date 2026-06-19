# Reto Técnico Interseguro: Matrices y Estadísticas

Este proyecto implementa una solución completa para el reto técnico de Interseguro. Está compuesto por dos APIs microservicios comunicados entre sí y una interfaz de usuario interactiva y minimalista en Angular.

---

## Estructura del Proyecto

- **/go-api**: API construida en **Go** utilizando el framework **Fiber**. Realiza la rotación de la matriz (90 grados en sentido horario) y la descomposición QR utilizando el algoritmo Gram-Schmidt.
- **/node-api**: API construida en **Node.js** utilizando **Express**. Calcula métricas globales de las matrices (Max, Min, Promedio, Suma) y comprueba si alguna es diagonal.
- **/frontend**: Aplicación SPA construida en **Angular** que permite interactuar con el sistema de forma intuitiva, registrar usuarios, iniciar sesión (JWT) y visualizar los resultados de las operaciones en tablas ordenadas.

---

## Requisitos Previos

Asegúrate de tener instalados en tu computadora:
- [Docker](https://www.docker.com/) y Docker Compose.
- Node.js (opcional, para desarrollo local sin Docker).
- Go (opcional, para desarrollo local sin Docker).

---

## Cómo Ejecutar la Solución Completa con Docker

La forma más rápida y recomendada de levantar el ecosistema completo es usando Docker Compose:

1. Abre tu terminal de PowerShell o CMD en el directorio raíz de este proyecto.
2. Ejecuta el comando para compilar y levantar los contenedores:
   ```bash
   docker-compose up --build
   ```
3. Una vez termine de levantar:
   - Accede al **Frontend de Angular** en tu navegador: [http://localhost:4200](http://localhost:4200)
   - El **API de Go** estará disponible en: `http://localhost:8080`
   - El **API de Node.js** estará disponible en: `http://localhost:5000`

---

## Cómo Probar la Aplicación

1. **Paso 1: Registro**: Crea un nuevo usuario ingresando a la opción de registro del frontend.
2. **Paso 2: Inicio de Sesión**: Inicia sesión con el usuario creado. El sistema almacenará de forma segura el token JWT para autorizar tus consultas.
3. **Paso 3: Configuración de Matriz**: Define las dimensiones de la matriz rectangular (ej. 3 filas y 3 columnas). Puedes escribir los números a mano o presionar el botón "Llenar con Aleatorios".
4. **Paso 4: Procesar**: Presiona "Calcular Rotación y Factorización QR". El backend procesará los cálculos y verás instantáneamente los resultados y las estadísticas de forma ordenada.
