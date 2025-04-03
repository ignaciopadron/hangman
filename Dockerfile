#ETAPA DE CONSTRUCCIÓN

# 1. Usa una imagen oficial de Go como punto de partida
FROM golang:1.24.2-alpine3.21 AS builder
RUN apk add --no-cache git upx

# Fuerza a Go a usar toolchain remoto si la versión local no alcanza
ENV GOTOOLCHAIN=auto

# 2. Crea un directorio dentro del contenedor

WORKDIR /app

# 3. Copia tu código fuente (archivos) al contenedor
COPY ["go.mod", "go.sum", "./"]

# 4. Compila el binario del juego desde cmd/ahorcado/main.go
RUN go mod download -x

COPY . .

# 4.1 Compila el binario del juego desde cmd/ahorcado/main.go
RUN go build -o ahorcado -v ./cmd/ahorcado

RUN upx ahorcado

# 5. Define el comando para correr el juego
CMD ["./ahorcado"]

#ETAPA FINAL (imagen mínima)
FROM alpine:3.21
LABEL Name=dockerization

RUN apk update
RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/ahorcado .

ENTRYPOINT [ "./ahorcado" ]
EXPOSE 8080
