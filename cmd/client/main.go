package main

import (
	"context"
	"flag"
	"fmt"
	"great/psm"
	"great/sample"
	"io"
	"log"
	"time"

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

func FilterLaptop(laptopClient psm.LaptopServiceClient) {

	// laptop := sample.NewLaptop()

	filter := &psm.Filter{
		MaxPriceInr: 900000,
		Cpu:         `brand:Nvidia  name:GTX 1080  cores:8  threads:2  min_ghz:2  max_ghz:3`,
		Gpu:         "Nvidia",
	}

	fmt.Println("Search laptop with filter", filter)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &psm.FilterLaptopRequest{
		Filter: filter,
	}

	stream, err := laptopClient.FilterLaptop(ctx, req)
	if err != nil {
		log.Fatal("Cannot filter laptop", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			fmt.Printf("")
			fmt.Printf("")
			log.Printf("End of the file")
			return
		}

		log.Printf("Received laptop with id: %s", res.GetLaptop())

		if err != nil {
			log.Fatal("Cannot receive laptop", err)
			break
		}
		laptop := res.GetLaptop()

		log.Printf("Found laptop with id: %s", laptop.GetId())
		log.Printf("Laptop brand: %s", laptop.GetBrand())
		log.Printf("Laptop name: %s", laptop.GetName())
		log.Printf("Laptop CPU cores: %d", laptop.GetCpu().GetCores())
		log.Printf("Laptop CPU ghz: %f", laptop.GetCpu().GetMinGhz())
		log.Printf("Laptop RAM: %d %s", laptop.GetRam().GetValue(), laptop.GetRam().GetUnit())
	}
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

	// FindLaptop(laptopClient, &psm.Laptop{Id: "ebc6a11c-4b2e-4e5c-98b3-8b5c70f55747"})
	FilterLaptop(laptopClient)
}
