package main

import (
	"fmt"
	"github.com/aabitbekov/endterm/avg/avgpb"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
	"io"
	"log"
	"net"
)

type server struct {
	avgpb.UnimplementedAvgServiceServer
}

func (s *server) StreamOfNumber(numberServer avgpb.AvgService_StreamOfNumberServer) error {
	var sum , count int64
	for {
		num, err := numberServer.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		count++
		//sum = num + sum
	}
	//sum = sum/count
	//return sum
}


func main() {
	fmt.Println("start")

	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	avgpb.RegisterAvgServiceServer(s,&server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

//
//var count, sum int64
//for {
//msg,err := stream.Recv()
//if err == io.EOF {
//log.Println("EOF")
//return nil
//}
//if err != nil {
//return err
//}
//count++
//sum += msg.Result
//sum = sum/count
//fmt.Printf(string(sum))
//}