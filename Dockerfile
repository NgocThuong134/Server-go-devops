FROM golang:1.23.2-alpine3.20
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o server-go-api
EXPOSE 4080
CMD ["./server-go-api"]