FROM golang:1.23-alpine AS backend-builder
WORKDIR /app

# Initialize go.mod manually
RUN go mod init main

COPY main.go .
RUN go build -o main

FROM alpine:latest
WORKDIR /app
COPY --from=backend-builder /app/main .
EXPOSE 8080
CMD ["./main"]
