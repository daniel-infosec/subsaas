FROM golang:1.10

WORKDIR $GOPATH/src/subsaas

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
CMD ["subsaas"]