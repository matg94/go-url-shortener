
FROM golang:1.16-alpine

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -o go-url-shortener

EXPOSE 8080

CMD [ "/app/go-url-shortener" ]