FROM golang:alpine3.18

RUN mkdir /app
ADD . /app 
WORKDIR /app

COPY go.mod go.sum  ./app/

RUN go mod download

RUN go build -o go-ms cmd/main.go

EXPOSE 3000

CMD ["./go-ms"]