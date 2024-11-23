FROM golang:1.23.0 AS builder

WORKDIR /payments

RUN GOBIN=/usr/local/bin go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go mod tidy

EXPOSE 3000

CMD ["air"]