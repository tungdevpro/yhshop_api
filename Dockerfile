# Initial Dockerfile
FROM golang:1.20-alpine as builder

# Setting current folder in Container
WORKDIR /app

# Copy go mod and go sum
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd

FROM golang:1.20-alpine


WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/.env .
ENTRYPOINT ["./main"]