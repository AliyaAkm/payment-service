FROM golang:1.24 AS builder

WORKDIR /ZerdeStudy

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd

FROM alpine:3.20

WORKDIR /ZerdeStudy

COPY --from=builder /ZerdeStudy/main .

RUN chmod +x ./main

CMD ["./main"]
