package main

import (
	"flag"
	"fmt"
	"great/config"
	"great/psm"
	"great/service"
	"net"

	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 0, "the server port")
	flag.Parse()
	fmt.Println("Starting on port", *port)

	prismaClient, err := config.ConnectDb()
	if err != nil {
		panic("Connection to DB failed")
	}

	defer prismaClient.Prisma.Disconnect()

	// laptopServer := service.NewLaptopService(service.NewInMemoryLaptopStore())
	laptopServer := service.NewLaptopService(prismaClient)

	// _, err = laptopServer.CreateLaptop(context.Background(), &psm.CreateLaptopRequest{})

	// go func(){
	// 	fmt.Println(data)
	// }()

	grpcServer := grpc.NewServer()

	psm.RegisterLaptopServiceServer(grpcServer, laptopServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
