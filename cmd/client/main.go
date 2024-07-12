package main

import (
	"context"
	"flag"
	"fmt"
	"great/psm"
	"great/sample"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CreateLaptop(laptopClient psm.LaptopServiceClient) {
	laptop := sample.NewLaptop()

	req := &psm.CreateLaptopRequest{
		Laptop: laptop,
	}

	res, err := laptopClient.CreateLaptop(context.Background(), req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			fmt.Println("Laptop already exists")
		} else {
			log.Fatal("Cannot create laptop", err)
		}
		return
	}

	log.Printf("Created laptop with id: %s", res.Id)
}

func FindLaptop(laptopClient psm.LaptopServiceClient, laptop *psm.Laptop) {
	req := &psm.FindLaptopRequest{
		Id: laptop.Id,
	}

	res, err := laptopClient.FindLaptop(context.Background(), req)
	if err != nil {
		log.Fatal("Cannot find laptop", err)
	}

	log.Printf("Found laptop: %s", res)

}

func main() {
	serverAddress := flag.String("address", "", "The server address in the format of host:port")
	flag.Parse()
	fmt.Println("Dial server on address", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	laptopClient := psm.NewLaptopServiceClient(conn)

	fmt.Println("laptopClient", laptopClient)

	// CreateLaptop(laptopClient)
	// CreateLaptop(laptopClient)
	// CreateLaptop(laptopClient)

	FindLaptop(laptopClient, &psm.Laptop{Id: "ebc6a11c-4b2e-4e5c-98b3-8b5c70f55747"})
}
