package service

import (
	"context"
	"fmt"
	"great/prisma/db"
	"great/psm"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LaptopService struct {
	psm.UnimplementedLaptopServiceServer
	Db *db.PrismaClient
}

func NewLaptopService(db *db.PrismaClient) *LaptopService {
	return &LaptopService{Db: db}
}

func (Db *LaptopService) CreateLaptop(ctx context.Context, p *psm.CreateLaptopRequest) (*psm.CreateLaptopResponse, error) {
	laptop := p.GetLaptop()
	if len(laptop.Id) > 0 {
		_, err := uuid.Parse(laptop.Id)

		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Laptop ID is not a valid UUID: %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot generate a new laptop ID: %v", err)
		}
		laptop.Id = id.String()
	}

	db_service := NewPrismaLaptopService(Db.Db)

	err := db_service.Createlaptop(context.Background(), laptop)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot create laptop: %v", err)
	}

	fmt.Println("Laptop saved successfully")
	return &psm.CreateLaptopResponse{
		Id: laptop.Id,
	}, nil
}
