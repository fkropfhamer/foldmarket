package main

import (
	context "context"
	"encoding/json"
	"fmt"
	"net"

	"github.com/kurrent-io/KurrentDB-Client-Go/kurrentdb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"foldmarket/event_stream"
	pb "foldmarket/market"
	"foldmarket/read_model"
)

type marketServer struct {
	pb.UnimplementedMarketServer
	es      *kurrentdb.Client
	queries *read_model.Queries
}

func (s *marketServer) GetBalance(ctx context.Context, req *pb.GetBalanceRequest) (*pb.GetBalanceResponse, error) {
	acc, err := s.queries.GetAccount(ctx, req.AccountId)

	if err != nil {
		return nil, err
	}

	return &pb.GetBalanceResponse{
		AccountId: acc.ID,
		Balance:   acc.Balance,
	}, nil
}

func (s *marketServer) Deposit(ctx context.Context, req *pb.DepositRequest) (*pb.DepositResponse, error) {
	depositEvent := event_stream.DepositEvent{
		AccountId: req.AccountId,
		Amount:    req.Amount,
	}

	data, err := json.Marshal(depositEvent)

	if err != nil {
		panic(err)
	}

	eventData := kurrentdb.EventData{
		ContentType: kurrentdb.ContentTypeJson,
		EventType:   event_stream.DepositEventType,
		Data:        data,
	}

	_, err = s.es.AppendToStream(context.Background(), "market-stream", kurrentdb.AppendToStreamOptions{}, eventData)

	return &pb.DepositResponse{
		AccountId:  req.AccountId,
		NewBalance: req.Amount,
	}, nil
}

func (s *marketServer) Withdraw(ctx context.Context, req *pb.WithdrawRequest) (*pb.WithdrawResponse, error) {
	withdrawEvent := event_stream.WithdrawEvent{
		AccountId: req.AccountId,
		Amount:    req.Amount,
	}

	data, err := json.Marshal(withdrawEvent)

	if err != nil {
		panic(err)
	}

	eventData := kurrentdb.EventData{
		ContentType: kurrentdb.ContentTypeJson,
		EventType:   event_stream.WithdrawEventType,
		Data:        data,
	}

	_, err = s.es.AppendToStream(context.Background(), "market-stream", kurrentdb.AppendToStreamOptions{}, eventData)

	return &pb.WithdrawResponse{
		AccountId:  req.AccountId,
		NewBalance: req.Amount,
	}, nil
}

func newServer(es *kurrentdb.Client, queries *read_model.Queries) *marketServer {
	s := &marketServer{es: es, queries: queries}
	return s
}

func main() {
	fmt.Println("Hello, World!")
	ctx := context.Background()

	es, _ := event_stream.GetEventStreamClient()
	defer es.Close()
	conn, _ := read_model.GetConnection()
	defer conn.Close(ctx)
	queries := read_model.New(conn)

	lis, _ := net.Listen("tcp", "localhost:50051")
	grpcServer := grpc.NewServer()
	pb.RegisterMarketServer(grpcServer, newServer(es, queries))

	reflection.Register(grpcServer)

	grpcServer.Serve(lis)
}
