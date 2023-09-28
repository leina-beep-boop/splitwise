package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/grpc/grpc-go/reflection"

	pb "github.com/leina-beep-boop/splitwise/protos/ledger"
)

type LedgerServiceServer struct {
	pb.UnimplementedLedgerServiceServer
}

func (LedgerServiceServer) AddExpense(ctx context.Context, req *pb.LedgerRequest) (*pb.LedgerResponse, error) {
	log.Printf("add expense %v", req.Expense)
	return &pb.LedgerResponse{}, nil
}

// should eventually listen to both servers
func main() {
	// establish local connection?
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// connect grpc server to the local connection
	grpcServer := grpc.NewServer()
	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)
	// s := ledger.LedgerServiceServer{}
	s := LedgerServiceServer{}

	pb.RegisterLedgerServiceServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
