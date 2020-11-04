package main

import (
	"bufio"
	pb "client/file_manager"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8088", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewFileManagerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	fmt.Println("enter service number \n1 => write file \n2 => read file")

	var userInput string
	fmt.Scanln(&userInput)

	if userInput == "1" {
		var name string
		var context string

		fmt.Println("input file name : ")
		fmt.Scanln(&name)

		fmt.Println("input text : ")

		in := bufio.NewReader(os.Stdin)

		line, err := in.ReadString('\n')
		if err != nil {
			log.Fatalf("wrong input: %v", err)
		}

		context = line

		r, err := c.PutFile(ctx, &pb.PutFileRequest{FileName: name, FileContext: context})
		if err != nil {
			log.Fatalf("could not save file: %v", err)
		}

		log.Printf("Save file: %v", r)
	} else if userInput == "2" {
		var name string

		fmt.Println("input file name you intend to search : ")
		fmt.Scanln(&name)

		r, err := c.GetFile(ctx, &pb.GetFileRequest{FileName: name})
		if err != nil {
			log.Fatalf("could not read file: %v", err)
		}
		log.Printf("Read file %v: ", r)
	}
}
