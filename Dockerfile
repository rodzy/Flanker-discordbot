FROM golang:latest
LABEL maintainer="Isaac Rodríguez irod2899@gmail.com"
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main
CMD ["/app/main"]