FROM golang:latest
COPY . /app
WORKDIR /app
EXPOSE 7733
CMD ["/app/server"]