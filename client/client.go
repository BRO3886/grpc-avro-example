package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/BRO3886/go-avro-validation/pkg/event"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/protobuf/types/known/structpb"
)

// var (
// 	eventJSON string = *flag.String("event", "", "event name")
// )

func main() {
	// flag.Parse()
	// if eventJSON == "" {
	// 	log.Fatal("event is required")
	// }

	e := &event.Event{
		EventName: "STOCKS_ORDER_BUY_EXECUTED",
		Source:    "STOCKS",
		EntityId:  "GSDJDFS87DFTS7GSDSDF78",
		Tags: map[string]string{
			"product":   "HDFC",
			"segment":   "CASH",
			"orderType": "STOP_LOSS",
		},
		EventMeta: &structpb.Struct{
			Fields: map[string]*structpb.Value{
				"price":              structpb.NewNumberValue(float64(1888)),
				"qty":                structpb.NewNumberValue(float64(2)),
				"instrumentFullName": structpb.NewStringValue("HDFC Bank"),
				"isin":               structpb.NewStringValue("IN1221ESD"),
				"orderPlacedAt":      structpb.NewNumberValue(float64(time.Now().UnixNano())),
			},
		},
		UserAccountId: "ACC2387E2873Y7823",
		Timestamp:     time.Now().Unix(),
		SchemaVersion: 2,
	}

	// err := json.Unmarshal([]byte(eventJSON), e)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := grpc.Dial(
		"localhost:9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithInitialConnWindowSize(256*1024),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                20,
			Timeout:             10,
			PermitWithoutStream: true,
		}),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := event.NewEventServiceClient(conn)

	var wg sync.WaitGroup
	for i := 1; i <= 500; i++ {
		wg.Add(1)
		log.Println("starting producer ", i)
		go func(idx int) {
			for j := 1; j <= 10; j++ {
				err = postEvent(e, client, ctx)
				log.Println("sent event")
				if err != nil {
					log.Println("err", err)
				}
				log.Println("producer ", idx, "event", j)
			}
		}(i)
	}
	wg.Wait()

}

func postEvent(event *event.Event, client event.EventServiceClient, ctx context.Context) error {
	stream, err := client.PostEvent(ctx)
	if err != nil {
		return err
	}

	for {
		err := stream.Send(event)
		if err != nil {
			return err
		}
		// time.Sleep(time.Second)
	}
}
