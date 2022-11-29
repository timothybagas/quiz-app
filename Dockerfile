FROM golang:1.18.8-alpine

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN apk update
RUN apk add make

RUN make tidy
RUN make build

EXPOSE 8080
CMD ["/app/main"]