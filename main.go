package main

import (
	"fmt"
	"log"
	"net"
	proto "github.com/Egorich42/app_proto"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
)


type Server struct {
	proto.UnimplementedChatServiceServer
	// i don't know what this shit means...
}

func (s *Server) SayHello(ctx context.Context, in *proto.Message) (*proto.Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &proto.Message{Body: "Hello From the Server!"}, nil
}

func main() {
	log.Println("Go gRPC server with prto in different repo started!")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := Server{}

	grpcServer := grpc.NewServer()

	proto.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
