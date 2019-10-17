FROM golang:1.13

MAINTAINER Jack Butler <jack@jackhas.codes>

WORKDIR /app

COPY ./ .

RUN go mod vendor && go build -o main

# TODO: EXPOSE 8081

CMD ["./main"]