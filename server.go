package main
import (
	"log"
	"net"
	"google.golang.org/grpc"
	"github.com/NBLANCHE/go-grpc/chat/github.com/NBLANCHE/go-grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("Failed to listen on port 5000")
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err!= nil {
		log.Fatalf("Failed to serve gPRC server over port 9000: %v", err)
	}

}

