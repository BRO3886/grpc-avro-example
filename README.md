# Go, gRPC - Avro schema registry

This app is a POC for using gRPC to send events in batches or by using gRPC streaming, and validating some metadata using schema from Schema registry, where the data is stored in avro format. The event message is an arbitrary event which also includes some metadata what we want to validate, based on the schema from the registry.

I used a tool called `ghz` to test both the streaming and batching performance.


Use the docker-compose file run the schema registry and its required dependencies. 
```bash
$ docker-compose up
```

Then run
```bash
$ go run main.go
```

Proto gen:
```bash
$ protoc --go_out=. --go-grpc_out=. ./proto/event.proto
```

Ghz commands
```bash
$ ghz --config=config.json -d "$(cat data.json)"
```