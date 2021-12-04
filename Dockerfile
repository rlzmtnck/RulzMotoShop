FROM golang:1.16-alpine

WORKDIR /app

COPY . .

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o mainfile

EXPOSE 8080

CMD ./mainfile
