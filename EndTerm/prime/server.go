package main

import (
	"fmt"
	"log"
	"net"

	"github.com/aabitbekov/endterm/prime/primepb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) PrimeNumber(req *primepb.PrimeRequest, stream primepb.PrimeService_PrimeNumberServer) error {
	fmt.Printf("primenum function was invoked with %v", req)

	number := req.GetNum()
	divisor := int64(2)
	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&primepb.PrimeResponse{
				Result: divisor,
			})
			number = number / divisor
		} else {
			divisor++
		}
	}
	return nil

}

func main() {
	fmt.Println("start")

	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	primepb.RegisterPrimeServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

