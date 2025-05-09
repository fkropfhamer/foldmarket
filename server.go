package main

import (
	context "context"
	"fmt"
	"net"
	sync "sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "foldmarket/market"
)

type marketServer struct {
	pb.UnimplementedMarketServer
	mu sync.Mutex
}

func (s *marketServer) GetBalance(ctx context.Context, req *pb.GetBalanceRequest) (*pb.GetBalanceResponse, error) {
	return &pb.GetBalanceResponse{
		AccountId: req.AccountId,
		Balance:   100.0,
	}, nil
}

func (s *marketServer) Deposit(ctx context.Context, req *pb.DepositRequest) (*pb.DepositResponse, error) {
	return &pb.DepositResponse{
		AccountId:  req.AccountId,
		NewBalance: req.Amount,
	}, nil
}

func newServer() *marketServer {
	s := &marketServer{}
	return s
}

func main() {
	fmt.Println("Hello, World!")

	lis, _ := net.Listen("tcp", "localhost:50051")
	grpcServer := grpc.NewServer()
	pb.RegisterMarketServer(grpcServer, newServer())

	reflection.Register(grpcServer)

	grpcServer.Serve(lis)
}
