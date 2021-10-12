# Building the binary of the App
FROM golang:1.17 AS builder
WORKDIR /fiber-versioning-boilerplate
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o fvb-server api/app.go

# Moving the binary to the 'final Image' to make it smaller
FROM alpine
WORKDIR /app
COPY --from=builder /fiber-versioning-boilerplate/.env.docker .env
COPY --from=builder /fiber-versioning-boilerplate/fvb-server fvb-server

# Exposes port 3000 because our program listens on that port
EXPOSE 3000
CMD ["./fvb-server"]