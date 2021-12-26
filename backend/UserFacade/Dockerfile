FROM golang:1.16-alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY cmd/main.go ./
RUN go build -o /main
EXPOSE 8080
ENTRYPOINT ["/main"]