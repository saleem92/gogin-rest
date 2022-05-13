FROM golang:1.18-buster

WORKDIR /app

COPY ./src/go.mod .
COPY ./src/go.sum . 

COPY ./src/ .
RUN go mod download

EXPOSE 8080

ENTRYPOINT ["go", "run", "."]
