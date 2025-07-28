FROM golang:1.24 AS builder

WORKDIR /usr/src/app

COPY . .
RUN go build -o test -ldflags "-s -w" .

FROM golang:1.24 AS runtime

WORKDIR /usr/local/bin

COPY --from=builder /usr/src/app/test .
CMD ["./test"]
