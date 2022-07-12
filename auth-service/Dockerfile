FROM golang:1.18-alpine3.16

RUN apk update
RUN apk add git

ADD . /go/src/auth-service
WORKDIR /go/src/auth-service
COPY go.mod . 
COPY go.sum .
COPY .env .

RUN go mod tidy -v
RUN go build -v


EXPOSE 8081

CMD [ "./auth-service"]
