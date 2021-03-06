package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	pb "server/file_manager"

	"google.golang.org/grpc"
)

type fileServer struct {
	pb.UnimplementedFileManagerServer

	conf *config
}

func (s *fileServer) PutFile(ctx context.Context, in *pb.PutFileRequest) (*pb.PutFileResponse, error) {
	var fileName = in.GetFileName()
	var fileContext = in.GetFileContext()

	log.Printf("Received fileName and fileContext as : %v", fileName+" | "+fileContext)

	writeFile(s.conf.basePath, fileName, fileContext)

	// writeFile 에서 string을 리턴하고, 현재 함수에서 return값을 받아와서 PutFileResponse로 리턴할 수 없으면 결과 어떻게 로그찍음??
	return &pb.PutFileResponse{Result: "File saved successfully"}, nil
}

func writeFile(basePath string, fileName string, fileContext string) {
	file, err := os.Create(fmt.Sprintf("%s/%s.txt", basePath, fileName))
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	file.Write([]byte(fileContext + "\n"))

	file.Sync() //syscall

	log.Printf("File saved successfully")
}

func (s *fileServer) GetFile(ctx context.Context, in *pb.GetFileRequest) (*pb.GetFileResponse, error) {
	var fileName = in.GetFileName()
	log.Printf("Received fileName : " + fileName)

	data, err := os.Open(fmt.Sprintf("%s/%s.txt", s.conf.basePath, fileName))
	if err != nil {
		log.Fatal(err)
	}

	defer data.Close()

	scanner := bufio.NewScanner(data)

	var readText string
	for scanner.Scan() {
		readText += scanner.Text()
	}

	log.Printf(readText)

	return &pb.GetFileResponse{Result: readText}, nil
}

func main() {
	config, err := newConfig()
	if err != nil {
		log.Fatalf("failed to get config: %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fileSvr := &fileServer{conf: config}

	grpcServer := grpc.NewServer()
	pb.RegisterFileManagerServer(grpcServer, fileSvr)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
