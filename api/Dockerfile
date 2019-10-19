FROM golang:1.13.3

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go get github.com/pilu/fresh
# RUN go get -u github.com/gorilla/mux

RUN go install -v ./...