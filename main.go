package main

import (
	"flag"
	"log"
	"net"

	"github.com/BRO3886/go-avro-validation/pkg/event"
	"github.com/BRO3886/go-avro-validation/server"
	"github.com/riferrei/srclient"
	"google.golang.org/grpc"
)

var (
	schemaRegistryUrl = flag.String("reg-url", "http://localhost:8081", "schema registry url")
)

type RegisteredSchema struct {
	Subject string `json:"subject"`
	Version int    `json:"version"`
	ID      int    `json:"id"`
	Schema  string `json:"schema"`
}

func main() {
	flag.Parse()
	client := srclient.CreateSchemaRegistryClient(*schemaRegistryUrl)

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()

	event.RegisterEventServiceServer(srv, server.EventServer{
		SchemaregClient: client,
	})

	log.Println("Starting server on port 9000")
	log.Panic(srv.Serve(lis))
}
