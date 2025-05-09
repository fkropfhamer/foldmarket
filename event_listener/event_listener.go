package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"foldmarket/event_stream"
	"foldmarket/read_model"
	"io"

	"github.com/kurrent-io/KurrentDB-Client-Go/kurrentdb"
)

func main() {
	ctx := context.Background()
	es, err := event_stream.GetEventStreamClient()
	stream, err := es.ReadStream(ctx, "market-stream", kurrentdb.ReadStreamOptions{}, 10)

	if err != nil {
		panic(err)
	}

	defer stream.Close()

	accountBalances := make(map[int32]int64)

	for {
		event, err := stream.Recv()

		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			panic(err)
		}

		if event.Event.EventType == event_stream.DepositEventType {
			depositEvent := event_stream.DepositEvent{}
			err = json.Unmarshal(event.Event.Data, &depositEvent)

			if err != nil {
				panic(err)
			}

			accountBalances[depositEvent.AccountId] += depositEvent.Amount
		} else if event.Event.EventType == event_stream.WithdrawEventType {
			withdrawEvent := event_stream.WithdrawEvent{}
			err = json.Unmarshal(event.Event.Data, &withdrawEvent)

			if err != nil {
				panic(err)
			}

			accountBalances[withdrawEvent.AccountId] -= withdrawEvent.Amount

		} else {
			fmt.Println("Unknown event type")
			continue
		}
	}

	stream.Close()

	conn, err := read_model.GetConnection()
	defer conn.Close(ctx)
	queries := read_model.New(conn)

	queries.DeleteAllAccounts(ctx)
	for accountId, balance := range accountBalances {
		_, err := queries.CreateAccount(ctx, read_model.CreateAccountParams{
			ID:      accountId,
			Balance: balance,
		})

		if err != nil {
			panic(err)
		}
	}
}
