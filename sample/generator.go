package sample

import (
	"great/psm"
	"math/rand"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewKeyBoard() *psm.KeyboardType {
	return &psm.KeyboardType{
		Layout:  keyboardTypeRandomizer(),
		Backlit: boolRandomizer(),
	}
}

func NewMemory() *psm.Memory {
	return &psm.Memory{
		Value: memoryRandomizer(),
		Unit:  psm.Memory_GIGABYTE,
	}
}

func NewCPU() *psm.CPU {
	return &psm.CPU{
		Brand:   "Nvidia",
		Name:    "GTX 1080",
		Cores:   uint32(rand.Intn(8) + 1),
		Threads: uint32(rand.Intn(16) + 1),
		MinGhz:  float64(rand.Intn(3) + 1),
		MaxGhz:  float64(rand.Intn(5) + 1),
	}
}

func NewGPU() *psm.GPU {
	return &psm.GPU{
		Brand:  "Intel",
		Name:   "Iris Pro",
		Memory: NewMemory(),
		MinGhz: float64(rand.Intn(4) + 1),
		MaxGhz: float64(rand.Intn(9) + 1),
	}
}

func NewScreen() *psm.Screen {
	return &psm.Screen{
		Resolution: screenResolutionRandomizer(),
		Panel:      screenPanelRandomizer(),
		SizeInch:   float32(rand.Intn(30) + 1),
		MultiTouch: boolRandomizer(),
	}
}

func NewStorage() *psm.Storage {
	return &psm.Storage{
		Driver: storageDriverRandomizer(),
		Memory: NewMemory(),
	}
}

func NewLaptop() *psm.Laptop {
	return &psm.Laptop{
		Id:          uuid.New().String(),
		Brand:       "Phalke",
		Name:        "Phalke X10",
		Keyboard:    NewKeyBoard(),
		Cpu:         NewCPU(),
		Gpu:         []*psm.GPU{NewGPU(), NewGPU()},
		Ram:         NewMemory(),
		Screen:      NewScreen(),
		Storages:    []*psm.Storage{NewStorage(), NewStorage()},
		Weight:      &psm.Laptop_WeightKg{WeightKg: 1.6},
		ReleaseYear: uint32(rand.Intn(12) + 2010),
		PriceInr:    float64(rand.Intn(10)*100000 + 100000),
		UpdatedAt:   timestamppb.Now(),
	}
}
