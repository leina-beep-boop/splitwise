package pkg

import (
	"context"
	"log"

	pb "github.com/leina-beep-boop/splitwise/protos/splitter"
)

type SplitterServiceServer struct {
	pb.UnimplementedSplitterServiceServer
}

func (SplitterServiceServer) CalculateDebt(ctx context.Context, req *pb.SplitterRequest) (*pb.SplitterResponse, error) {
	log.Printf("Hello World! %v", req)
	return &pb.SplitterResponse{Status: 0}, nil
}
