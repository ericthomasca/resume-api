FROM golang:1.21

WORKDIR /app

COPY go.mod ./

RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /resume-api

EXPOSE 7777

CMD ["/resume-api"]
