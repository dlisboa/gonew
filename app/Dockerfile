FROM golang:1.22-alpine as builder

RUN apk add --no-cache build-base # for gcc

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -mod=readonly -o ./bin/server ./cmd/server

ENTRYPOINT ["./bin/server"]
