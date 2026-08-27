FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s" \
    -o /app/bin/gojob \
    ./main.go

FROM alpine:3.19 AS final

WORKDIR /app

RUN adduser -D -g '' appuser
USER appuser

COPY --from=builder /app/bin/gojob /app/gojob

ENTRYPOINT [ "/app/gojob" ]