FROM golang:1.16-alpine
WORKDIR /ChatService
COPY ./ ./
RUN go mod download
RUN go build -o main ./cmd/main.go
ENTRYPOINT ["./main"]