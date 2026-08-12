############################
# 1. Build Stage
############################
FROM golang:1.26-alpine AS builder

# Instalar dependencias necesarias para la compilación (ca-certificates, tzdata)
RUN apk add --no-cache git ca-certificates tzdata

WORKDIR /app

# Copiar go.mod y go.sum para aprovechar la caché de capas de Docker
COPY go.mod go.sum ./
RUN go mod download

# Copiar el resto del código fuente del servicio en Go
COPY . .

# Compilar el binario
# CGO_ENABLED=0 genera un binario estático que no depende de librerías C del SO
# -ldflags="-w -s" elimina información de depuración para reducir el tamaño del binario
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s" \
    -o /app/bin/goservice \
    ./main.go

############################
# 2. Final Runtime Stage
############################
FROM alpine:3.19 AS final

# Certificados SSL/TLS (necesarios si te comunicas por HTTPS/gRPC con TLS)
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# Crear un usuario no-root por seguridad
RUN adduser -D -g '' appuser
USER appuser

# Copiar únicamente el binario compilado desde la etapa builder
COPY --from=builder /app/bin/goservice /app/goservice

# Puerto REST de Go (ajusta según tu configuración)
EXPOSE 8080

ENTRYPOINT ["/app/goservice"]