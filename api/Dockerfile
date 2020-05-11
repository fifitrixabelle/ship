FROM golang:1.13

RUN mkdir /app
WORKDIR /app

COPY go.* ./
RUN go mod download
COPY main.go main.go
RUN go build -o ship main.go

EXPOSE 8080
CMD ./ship
