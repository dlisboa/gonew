FROM golang:1.22-alpine as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -mod=readonly -o bin/server ./cmd/server

FROM scratch
COPY --from=builder /app/bin/server ./server
ENTRYPOINT ["./server"]
