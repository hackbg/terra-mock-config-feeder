
FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go get -d -v ./...

RUN go build -o /feeder

EXPOSE 4000

CMD [ "/feeder" ]