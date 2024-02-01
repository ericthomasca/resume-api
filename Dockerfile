FROM golang:latest

WORKDIR /app

COPY main .

EXPOSE 7777

CMD ["./main"]