package service

import "great/psm"

type LaptopService struct {
	psm.UnimplementedLaptopServiceServer
	Laptop *psm.Laptop
}

func NewLaptopService() *LaptopService {
	return &LaptopService{}
}
