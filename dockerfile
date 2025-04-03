# 1. Usa una imagen oficial de Go como punto de partida

FROM golang:1.21

# 2. Crea un directorio dentro del contenedor

WORKDIR /app

# 3. Copia tu código fuente (archivos) al contenedor
COPY . .

# 4. Compila el binario del juego desde cmd/ahorcado/main.go
RUN go build -o ahorcado ./cmd/ahorcado

# 5. Define el comando para correr el juego
CMD ["./ahorcado"]