FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
COPY . .
RUN go build -o ./bin/main ./cmd

CMD ./bin/main
EXPOSE 3000