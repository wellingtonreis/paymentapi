FROM golang:1.23.0 AS builder

WORKDIR /cashback

RUN GOBIN=/usr/local/bin go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go mod tidy

EXPOSE 3002

CMD ["air"]