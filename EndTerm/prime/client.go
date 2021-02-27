package main

import (
	"fmt"
	"log"
	"context"
	"github.com/aabitbekov/endterm/prime/primepb"
	"google.golang.org/grpc"
	"io"
)

func main() {
	fmt.Println("Hello from client")
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect:%v", err)
	}

	defer conn.Close()

	c := primepb.NewPrimeServiceClient(conn)
	// fmt.Printf("created client: %f", c)

	doServerStreaming(c, 120)
	doServerStreaming(c, 251)

}

func doServerStreaming(c primepb.PrimeServiceClient, number int64) {
	fmt.Println("Starting to do a server streaimng rpc....")

	req := &primepb.PrimeRequest{
		Num: number,
	}
	resStream, err := c.PrimeNumber(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GRPC: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)

		}
		fmt.Println(msg.GetResult())
	}
}


