package main

import (
	"context"
	"log"
	"net"

	pb "drpcf/fold" //package_name/folder(if folder is present)

	"storj.io/drpc/drpcmux"
	"storj.io/drpc/drpcserver"
)

type CookieServer struct {
	pb.DRPCGreeterUnimplementedServer
}

func (s *CookieServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello World"}, nil
}

func main() {
	err := Serve(context.Background())
	if err != nil {
		panic(err)
	}
}

func Serve(ctx context.Context) error {
	// create an RPC server
	cookieMonster := &CookieServer{}

	// create a drpc RPC mux
	m := drpcmux.New()

	// register the proto-specific methods on the mux
	err := pb.DRPCRegisterGreeter(m, cookieMonster)
	if err != nil {
		return err
	}

	// create a drpc server
	s := drpcserver.New(m)

	// listen on a tcp socket
	lis, err := net.Listen("tcp", ":8080")
	log.Printf("server listening at %v", lis.Addr())
	// log.Printf("server listening at %v", lis.Addr())
	if err != nil {
		return err
	}

	// run the server
	// N.B.: if you want TLS, you need to wrap the net.Listener with
	// TLS before passing to Serve here.
	return s.Serve(ctx, lis)
}
