package service

import (
	"bytes"
	"context"
	"fmt"
	"great/prisma/db"
	"great/psm"
	"io"
	"log"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const MaxImageSize = 1 << 20 // 1 MB

type LaptopService struct {
	psm.UnimplementedLaptopServiceServer
	imageStore ImageStore
	Db         *db.PrismaClient
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

func (Db *LaptopService) UploadImage(stream psm.LaptopService_UploadImageServer) error {

	req, err := stream.Recv()
	if err != nil {
		return status.Errorf(codes.Unknown, "cannot receive image info: %v", err)
	}

	laptopID := req.GetImage().GetImageId()
	imageType := req.GetImage().GetImageType()

	log.Printf("Receive an upload-image request for laptop %s with image type %s\n", laptopID, imageType)

	db_service := NewPrismaLaptopService(Db.Db)

	laptop, err := db_service.FindLaptop(context.Background(), &psm.Laptop{Id: laptopID})
	if err != nil {
		return status.Errorf(codes.Internal, "cannot find laptop: %v", err)
	}

	if laptop == nil {
		return status.Errorf(codes.InvalidArgument, "laptop %s doesn't exist", laptopID)
	}

	imageDate := bytes.Buffer{}
	imageSize := 0

	for {
		log.Print("Waiting to receive more data")

		req, err := stream.Recv()

		if err == io.EOF {
			log.Print("No more data")
			break
		}

		if err != nil {
			return status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err)
		}

		chunk := req.GetChunkData()
		size := len(chunk)

		imageSize += size
		if imageSize >= MaxImageSize {
			return status.Errorf(codes.InvalidArgument, "image is too large: %d > %d", imageSize, MaxImageSize)
		}

		_, err = imageDate.Write(chunk)
		if err != nil {
			return status.Errorf(codes.Internal, "cannot write chunk data: %v", err)
		}

	}

	imageID, err := Db.imageStore.Save(laptopID, imageType, imageDate)
	if err != nil {
		return status.Errorf(codes.Internal, "cannot save image to store: %v", err)
	}

	res := &psm.UploadImageResponse{
		Id:   imageID,
		Size: uint32(imageSize),
	}

	err = stream.SendAndClose(res)

	if err != nil {
		return status.Errorf(codes.Internal, "cannot send response: %v", err)
	}

	log.Printf("Image with ID %s and size %d has been saved successfully", imageID, imageSize)

	return nil
}
