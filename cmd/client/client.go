package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github/tiagonevestia/go-grpc/pb"

	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50055", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	// AddUser(client)
	// AddUserVerbose(client)
	// AddUsers(client)
	AddUserStreamBoth(client)

}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Tiago Neves",
		Email: "tiagoneves.tia@gmail.com",
	}
	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}
	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {

	req := &pb.User{
		Id:    "0",
		Name:  "Tiago Neves",
		Email: "tiagoneves.tia@gmail.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not receive the msg: %v", err)
		}
		fmt.Println("Status:", stream.Status, " - ", stream.GetUser())
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "100",
			Name:  "Mara Jeannie",
			Email: "mara@gmail.com",
		},
		&pb.User{
			Id:    "110",
			Name:  "Larisa Manuela",
			Email: "lalamanu@gmail.com",
		},
		&pb.User{
			Id:    "120",
			Name:  "André Lisboa",
			Email: "dede12@gmail.com",
		},
		&pb.User{
			Id:    "130",
			Name:  "Carlos Wagner",
			Email: "crw@gmail.com",
		},
		&pb.User{
			Id:    "140",
			Name:  "José Armengo",
			Email: "zearmengo@gmail.com",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println(res)
}

func AddUserStreamBoth(client pb.UserServiceClient) {

	stream, err := client.AddUserStreamBoth(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	reqs := []*pb.User{
		&pb.User{
			Id:    "100",
			Name:  "Mara Jeannie",
			Email: "mara@gmail.com",
		},
		&pb.User{
			Id:    "110",
			Name:  "Larisa Manuela",
			Email: "lalamanu@gmail.com",
		},
		&pb.User{
			Id:    "120",
			Name:  "André Lisboa",
			Email: "dede12@gmail.com",
		},
		&pb.User{
			Id:    "130",
			Name:  "Carlos Wagner",
			Email: "crw@gmail.com",
		},
		&pb.User{
			Id:    "140",
			Name:  "José Armengo",
			Email: "zearmengo@gmail.com",
		},
	}

	wait := make(chan int)

	go func() {
		for _, req := range reqs {
			fmt.Println("Sending user: ", req.Name)
			stream.Send(req)
			time.Sleep(time.Second * 2)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error receiving data: %v", err)
				break
			}
			fmt.Printf("Recebendo user %v com status: %v\n", res.GetUser().GetName(), res.GetStatus())
		}
		close(wait)
	}()

	<-wait

}
