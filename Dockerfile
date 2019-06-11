FROM golang

WORKDIR /go/src/github.com/BojanKomazec/go-demo

COPY ./cmd/ ./cmd/
COPY ./internal/ ./internal/
COPY ./.env ./.env

# Enable for debugging only.
# Output snippet:
#   GOPATH="/go"
#   GOROOT="/usr/local/go"
RUN go env

RUN go get -d -v ./...

# No need to set GOOS=linux GOARCH=amd64 as these env vars are already present
# and set to these values.
RUN CGO_ENABLED=0 go build -o ./go-demo ./cmd/main.go

CMD ["/go/src/github.com/BojanKomazec/go-demo/go-demo"]
