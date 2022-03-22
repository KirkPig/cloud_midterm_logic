FROM golang:1.17-alpine

WORKDIR /app

COPY ["go.mod", "go.sum", "./"]

RUN go mod download

COPY . .

ENV GIN_MODE=release
RUN go build -o main

EXPOSE 80
CMD ["/app/main"]
