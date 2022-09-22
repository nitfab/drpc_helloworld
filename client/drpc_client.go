package main

import (
	"context"
	"log"
	"net"

	pb "drpcf/fold" //package_name/folder(if folder is present)

	"storj.io/drpc/drpcconn"
)

func main() {
	err := CookieTime(context.Background())
	if err != nil {
		panic(err)
	}
}

func CookieTime(ctx context.Context) error {
	rawconn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		return err
	}
	conn := drpcconn.New(rawconn)
	defer conn.Close()

	client := pb.NewDRPCGreeterClient(conn)
	r, err := client.SayHello(ctx, &pb.HelloRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
	return err
}
