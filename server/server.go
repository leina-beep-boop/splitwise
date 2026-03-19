package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/grpc/grpc-go/reflection"

	"github.com/leina-beep-boop/splitwise/ledgerserver"
	lpb "github.com/leina-beep-boop/splitwise/protos"
)

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

	ls := ledgerserver.New()
	lpb.RegisterLedgerServiceServer(grpcServer, ls)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
