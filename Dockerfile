FROM golang

WORKDIR /go/src/go-demo
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["go-demo"]
