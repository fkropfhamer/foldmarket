package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"foldmarket/event_stream"
	"io"

	"github.com/kurrent-io/KurrentDB-Client-Go/kurrentdb"
)

func main() {
	es, err := event_stream.GetEventStreamClient()
	stream, err := es.ReadStream(context.Background(), "market-stream", kurrentdb.ReadStreamOptions{}, 10)

	if err != nil {
		panic(err)
	}

	defer stream.Close()

	for {
		event, err := stream.Recv()

		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			panic(err)
		}

		test_event := event_stream.DepositEvent{}
		err = json.Unmarshal(event.Event.Data, &test_event)

		fmt.Println(test_event)

		// Doing something productive with the event
		fmt.Println(event)
	}
}
