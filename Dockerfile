FROM golang:latest

COPY . /app

WORKDIR /app

RUN go build -o server ./main

EXPOSE 7733

CMD ["/app/server"]