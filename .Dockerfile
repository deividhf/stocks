FROM golang

WORKDIR /app/src/stocks

ENV GOPATH=/app

COPY . /app/src/stocks

RUN go build server.go

ENTRYPOINT ./server

EXPOSE 8080