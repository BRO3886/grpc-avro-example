package server

import (
	"context"
	"io"
	"log"
	"sync"

	"github.com/BRO3886/go-avro-validation/pkg/event"
	"github.com/riferrei/srclient"
)

type EventServer struct {
	event.UnimplementedEventServiceServer
	SchemaregClient *srclient.SchemaRegistryClient
}

var inMemoryStore = make(map[string]*srclient.Schema)

func (e EventServer) PostEvent(stream event.EventService_PostEventServer) error {
	client := e.SchemaregClient
	var mutex = &sync.RWMutex{}
	for {
		event, err := stream.Recv()

		if err != nil {
			if err == io.EOF {
				// log.Println("client closed stream")
				return nil
			}
			log.Println("err", err)
			return err
		}

		log.Println("received event")
		key := event.EventName + "_" + event.Source

		var schema *srclient.Schema

		mutex.Lock()
		if _, ok := inMemoryStore[key]; !ok {
			schema, err := client.GetSchemaByVersion(key, int(event.SchemaVersion))
			if err != nil {
				return err
			}
			inMemoryStore[key] = schema
		}
		mutex.Unlock()

		mutex.RLock()
		schema = inMemoryStore[key]
		codec := schema.Codec()
		bytes, err := event.EventMeta.MarshalJSON()
		mutex.RUnlock()

		if err != nil {
			return err
		}

		_, _, err = codec.NativeFromTextual(bytes)
		if err != nil {
			log.Println("invalid event", err)
			return err
		}
	}
}

func (e EventServer) PostEventBatch(ctx context.Context, payload *event.EventRequest) (*event.EventResponse, error) {
	client := e.SchemaregClient
	var mutex = &sync.RWMutex{}
	log.Println("received events")

	for _, event := range payload.Events {
		var schema *srclient.Schema
		key := event.EventName + "_" + event.Source
		mutex.Lock()
		if _, ok := inMemoryStore[key]; !ok {
			schema, err := client.GetSchemaByVersion(key, int(event.SchemaVersion))
			if err != nil {
				return nil, err
			}
			inMemoryStore[key] = schema
		}
		mutex.Unlock()

		mutex.RLock()
		schema = inMemoryStore[key]
		codec := schema.Codec()
		bytes, err := event.EventMeta.MarshalJSON()
		mutex.RUnlock()

		if err != nil {
			return nil, err
		}

		_, _, err = codec.NativeFromTextual(bytes)
		if err != nil {
			log.Println("invalid event", err)
			return nil, err
		}
	}
	log.Println("processed events")

	return &event.EventResponse{
		Message: "success",
	}, nil
}
