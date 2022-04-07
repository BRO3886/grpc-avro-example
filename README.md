# Go - Avro schema registry

Basic app which makes an HTTP call to a schema registry, returns the schema for a given subject and version, and validates it against a hard-coded json data.

The docker-compose file will be used to run the schema registry and its required dependencies. Then run
```bash
go run main.go
```