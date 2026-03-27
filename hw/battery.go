package hw

import (
	"periph.io/x/conn/v3/i2c"
)

const pisugarAddr = 0x57

func ReadBattery(bus i2c.Bus) (int, error) {
	// leggi il registro della percentuale
	buf := make([]byte, 1)
	if err := bus.Tx(pisugarAddr, []byte{0x2a}, buf); err != nil {
		return 0, err
	}
	percent := int(buf[0])
	if percent > 100 {
		percent = 100
	}
	return percent, nil
}
