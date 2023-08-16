FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o uscis-case-status

CMD ["./uscis-case-status"]