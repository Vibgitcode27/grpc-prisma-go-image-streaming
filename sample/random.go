package sample

import (
	"great/psm"
	"math/rand"
)

func boolRandomizer() bool {
	return rand.Intn(2) == 0
}

// Keyboard randomizer functions

func keyboardTypeRandomizer() psm.KeyboardType_KeyboardLayout {
	if rand.Intn(4) == 0 {
		return psm.KeyboardType_QWERTY
	} else if rand.Intn(4) == 1 {
		return psm.KeyboardType_AZERTY
	} else if rand.Intn(4) == 2 {
		return psm.KeyboardType_QWERTZ
	} else {
		return psm.KeyboardType_QWERTY
	}
}

// Memory randomizer functions
func memoryRandomizer() uint64 {
	return uint64(rand.Intn(999) + 1)
}

// Screen randomizer functions

func screenResolutionRandomizer() *psm.Screen_Resolution {
	return &psm.Screen_Resolution{
		Width:  uint32(rand.Intn(1920) + 1),
		Height: uint32(rand.Intn(1080) + 1),
	}
}

func screenPanelRandomizer() psm.Screen_Panel {
	if rand.Intn(3) == 0 {
		return psm.Screen_IPS
	} else if rand.Intn(3) == 1 {
		return psm.Screen_LCD
	} else if rand.Intn(3) == 2 {
		return psm.Screen_OLED
	} else {
		return psm.Screen_UNKNOWN
	}
}
func storageDriverRandomizer() psm.Storage_Driver {
	if rand.Intn(2) == 0 {
		return psm.Storage_SSD
	} else {
		return psm.Storage_HDD
	}
}
