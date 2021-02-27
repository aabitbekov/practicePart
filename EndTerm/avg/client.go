package main

import (
	"fmt"
	"github.com/aabitbekov/endterm/avg/avgpb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	fmt.Println("Hello from client")
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect:%v", err)
	}

	defer conn.Close()

	c := avgpb.NewAvgServiceClient(conn)
	fmt.Printf("created client: %f", c)
	var  numbers [] int64
	numbers = append(numbers, 1)
	numbers = append(numbers, 5)
	doServerStreaming(c ,numbers)

}

func doServerStreaming(c avgpb.AvgServiceClient, numbers []int64) {
	fmt.Println("Starting to do a server rpc....")
	size:=0
	for len(numbers) == size {
		req := &avgpb.AvgRequest{
			Result: numbers[size],
		}
		stream , err := c.StreamOfNumber(context.Background())
		if err != nil {
			log.Fatalf("error while calling GRPC: %v", err)
		}
		size++
		msg, err := stream.RecvMsg()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)

		}
		fmt.Println(msg)
	}
}

