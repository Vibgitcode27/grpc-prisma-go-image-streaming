package service

import (
	"context"
	"fmt"
	"great/prisma/db"
	"great/psm"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
