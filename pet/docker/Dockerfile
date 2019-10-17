FROM golang:1.13

MAINTAINER Jack Butler <jack@jackhas.codes>

WORKDIR /app

COPY ./ .

RUN go mod vendor && go build -o main cmd/pet/main.go

EXPOSE 8081

CMD ["./main"]