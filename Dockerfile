FROM golang:1.14

WORKDIR /app
COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]