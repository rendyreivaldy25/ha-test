FROM golang:1.17-alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN cp configs/config.json.dev configs/config.json
RUN go build -o main .
CMD ["/app/main"]