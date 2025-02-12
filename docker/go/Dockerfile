FROM golang:1.23-alpine

WORKDIR /go/src/app

RUN go install github.com/rubenv/sql-migrate/...@latest
RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV TZ Asia/Tokyo

CMD ["air", "-c", ".air.toml"]