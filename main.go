package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/sirupsen/logrus"
	grpc "google.golang.org/grpc"

	pb "github.com/dasbeerboot/go-everlast-storage/hello"
)

type server struct {
	pb.UnimplementedGreeterServer // struct embedding (similar as inheritance)
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	logrus.Infof("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "fuck you"}, nil
}

func main() {
	log.Println("this is a log")
	logrus.Info("test")

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		logrus.Info("you failed")
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	logrus.Info("not reachable")

	// fmt.Println("input 줘")

	// var userInput string
	// fmt.Scanln(&userInput)

	// err := writeFile(userInput)
	// if err != nil {
	// 	logrus.Error(err)
	// 	panic(err)
	// }
}

func writeFile(inputText string) error {
	myFile, err := os.Create("/Users/juwon/Desktop/go-storage/test1")
	if err != nil {
		return err
	}
	defer myFile.Close()

	myFile.Write([]byte(inputText + "\n"))

	myFile.Sync() //syscall

	return nil
}
