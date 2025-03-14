FROM golang:1.22-alpine AS builder

RUN apk add --no-cache gcc musl-dev sqlite-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o todo-app main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/todo-app .
EXPOSE 8080
CMD ["./todo-app"]
