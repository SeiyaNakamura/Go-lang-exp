FROM golang:1.11

ENV GO111MODULE=on

WORKDIR /go/src/github.com/KaoruOhbayashi/golang_echo

COPY ./src .
RUN go build .

EXPOSE 8080

RUN go get github.com/pilu/fresh

CMD ["fresh"]