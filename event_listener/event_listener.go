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

		var testEvent any = nil
		if event.Event.EventType == event_stream.DepositEventType {
			testEvent = event_stream.DepositEvent{}
			err = json.Unmarshal(event.Event.Data, &testEvent)
		} else if event.Event.EventType == event_stream.WithdrawEventType {
			testEvent = event_stream.WithdrawEvent{}
			err = json.Unmarshal(event.Event.Data, &testEvent)
		} else {
			fmt.Println("Unknown event type")
			continue
		}

		fmt.Println(testEvent)
	}
}
