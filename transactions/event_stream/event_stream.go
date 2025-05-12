package event_stream

import "github.com/kurrent-io/KurrentDB-Client-Go/kurrentdb"

func GetEventStreamClient() (*kurrentdb.Client, error) {
	settings, err := kurrentdb.ParseConnectionString("kurrentdb://localhost:2113?tls=false")

	if err != nil {
		panic(err)
	}

	return kurrentdb.NewClient(settings)
}
