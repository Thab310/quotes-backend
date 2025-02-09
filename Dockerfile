FROM golang:1.21-alpine AS backend-builder
WORKDIR /app
COPY go.mod .
COPY main.go .
RUN go build -o main

FROM alpine:latest
WORKDIR /app
COPY --from=backend-builder /app/main .
EXPOSE 8080
CMD ["./main"]