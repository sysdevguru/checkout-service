# Checkout service
Simple Golang checkout server/client based on gRPC

## How to get
```sh
go get github.com/sysdevguru/checkout-service
```

## How to run server
```sh
docker login
docker run -it -p 8080:8080 sysdevguru/checkout-server:1.1.1
```

## How to run client
```sh
cd client/
make all
./checkout-client
```

### Integration Testing of client
```sh
cd client/
make test
```
