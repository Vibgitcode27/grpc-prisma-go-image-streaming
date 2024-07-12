package service

import (
	"context"
	"fmt"
	"great/prisma/db"
	"great/psm"
	"great/sample"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PrismaLaptopService struct {
	psm.UnimplementedLaptopServiceServer
	Db *db.PrismaClient
}

func NewPrismaLaptopService(db *db.PrismaClient) *PrismaLaptopService {
	return &PrismaLaptopService{Db: db}
}

func (p *PrismaLaptopService) Createlaptop(ctx context.Context, laptop *psm.Laptop) error {
	// save to database
	data, err := p.Db.Laptop.CreateOne(
		db.Laptop.ID.Set(laptop.Id),
		db.Laptop.Brand.Set(laptop.Brand),
		db.Laptop.Name.Set(laptop.Name),
		db.Laptop.CPU.Set(laptop.Cpu.String()),
		db.Laptop.Gpu.Set(laptop.GetCpu().Brand),
		db.Laptop.RAM.Set(laptop.Ram.String()),
		db.Laptop.Keyboard.Set(laptop.Keyboard.String()),
		db.Laptop.Storages.Set(fmt.Sprintf("%s", laptop.Storages)),
		db.Laptop.Screen.Set(laptop.Screen.String()),
		db.Laptop.PriceInr.Set(laptop.PriceInr),
		db.Laptop.ReleaseYear.Set(2014),
	).Exec(ctx)

	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Laptop object creating failed: %v", err)
	}
	fmt.Println(data)
	// return response
	return nil
}

func (p *PrismaLaptopService) FindLaptop(ctx context.Context, laptop *psm.Laptop) (*psm.Laptop, error) {
	data, err := p.Db.Laptop.FindUnique(
		db.Laptop.ID.Equals(laptop.Id),
	).Exec(ctx)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Laptop object finding failed: %v", err)
	}
	fmt.Println(data.ID)

	fetchedLaptop := convertToProtoLaptop(data)

	return fetchedLaptop, nil
	// return nil
}

// convertToProtoLaptop converts a Prisma Laptop model to a gRPC Laptop message
func convertToProtoLaptop(data *db.LaptopModel) *psm.Laptop {
	// Assuming you have appropriate conversion logic from string to enums for the proto messages
	// Adjust as per your application logic

	laptop := sample.NewLaptop()

	return &psm.Laptop{
		Id:          data.ID,
		Brand:       data.Brand,
		Name:        data.Name,
		Cpu:         laptop.Cpu,                         // Will have to extract the CPU from the data string , as this is the practice project I am not doing it
		Gpu:         laptop.Gpu,                         // Will have to extract the CPU from the data string , as this is the practice project I am not doing it
		Ram:         laptop.Ram,                         // Will have to extract the CPU from the data string , as this is the practice project I am not doing it
		Keyboard:    laptop.Keyboard,                    // Will have to extract the CPU from the data string , as this is the practice project I am not doing it
		Storages:    laptop.Storages,                    // Will have to extract the CPU from the data string , as this is the practice project I am not doing it
		Screen:      laptop.Screen,                      // Will have to extract the CPU from the data string , as this is the practice project I am not doing it
		Weight:      &psm.Laptop_WeightKg{WeightKg: 10}, // Adjust as necessary for your weight logic
		PriceInr:    data.PriceInr,
		ReleaseYear: uint32(data.ReleaseYear),
		UpdatedAt:   timestamppb.New(data.UpdatedAt),
	}
}

func (p *PrismaLaptopService) FilterLaptopsasdasdas(ctx context.Context, filter *psm.Filter, found func(laptop *psm.Laptop) error) error {

	log.Printf("searching laptops with filter: %v", filter)
	laptops, err := p.Db.Laptop.FindMany(
		// db.Laptop.PriceInr.Equals(filter.GetMaxPriceInr()),
		// db.Laptop.CPU.Equals(filter.GetCpu()),
		db.Laptop.Gpu.Equals(filter.GetGpu()),
	).Exec(ctx)

	log.Printf("searching laptops with filter LKJSDLKJASLDKJSALKDJASLKDJ: %v", laptops)

	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Laptop object finding failed: %v", err)
	}

	for _, laptop := range laptops {
		err := found(convertToProtoLaptop(&laptop))
		if err != nil {
			return status.Errorf(codes.Internal, "Internal error: %v", err)
		}
	}

	return status.Errorf(codes.Unimplemented, "method FilterLaptops not implemented")
}
