FROM golang:1.18-alpine

RUN apk update
RUN apk add git

ADD . /go/src/user-service
WORKDIR /go/src/user-service
COPY go.mod . 
COPY go.sum .
COPY .env .

RUN go mod tidy -v
RUN go build -v


EXPOSE 8080

CMD [ "./user-service"]
