FROM golang:latest

WORKDIR $GOPATH/src/github.com/sysdevguru/checkout-service
COPY . .

RUN apt-get update && apt-get install -y uuid-runtime
RUN go get google.golang.org/grpc && go get gopkg.in/yaml.v2
RUN go build -o checkout-server

EXPOSE 8080

ENTRYPOINT [ "./checkout-server" ]
