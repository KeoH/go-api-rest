FROM golang:1.11.0-alpine3.8

WORKDIR /go/src/app

RUN apk add --no-cache git mercurial \
    && go get github.com/gorilla/mux \
    && go get gopkg.in/mgo.v2 \
    && apk del git mercurial

COPY ./src .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8000

CMD ["app"]