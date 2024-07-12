package service

import (
	"context"
	"fmt"
	"great/prisma/db"
	"great/psm"
	"log"

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

func (Db *LaptopService) FindLaptop(ctx context.Context, p *psm.FindLaptopRequest) (*psm.FindLaptopResponse, error) {

	laptopId := p.GetId()
	if len(laptopId) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Laptop ID cannot be empty")
	}

	db_service := NewPrismaLaptopService(Db.Db)

	lappy, err := db_service.FindLaptop(context.Background(), &psm.Laptop{Id: laptopId})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "cannot find laptop: %v", err)
	}

	return &psm.FindLaptopResponse{
		Laptop: lappy,
	}, nil

}

func (Db *LaptopService) FilterLaptop(req *psm.FilterLaptopRequest, stream psm.LaptopService_FilterLaptopServer) error {
	filter := req.GetFilter()

	db_service := NewPrismaLaptopService(Db.Db)

	log.Printf("searching laptop with filter: %v", filter)
	err := db_service.FilterLaptopsasdasdas(context.Background(), filter, func(laptop *psm.Laptop) error {
		log.Printf("found laptop with id: %s", laptop.Id)
		err := stream.Send(&psm.FilterLaptopResponse{Laptop: laptop})
		if err != nil {
			log.Printf("cannot send laptop: %v", err)
			return err
		}
		return nil
	})

	if err != nil {
		return status.Errorf(codes.Internal, "unexpected error: %v", err)
	}

	return nil

}
