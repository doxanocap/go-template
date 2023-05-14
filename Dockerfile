FROM golang:1.20-alpine as builder
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api ./cmd/main.go

FROM alpine as production
WORKDIR /app

COPY --from=builder /build/ ./
CMD ["/app/api"]