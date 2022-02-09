FROM golang:latest

WORKDIR /app

COPY . .
RUN go mod download

RUN go build .

CMD [ "go", "run", "." ]

EXPOSE 8080
