package hw

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadWifiStrength() (int, error) {
	f, err := os.Open("/proc/net/wireless")
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	// salta le prime 2 righe di intestazione
	scanner.Scan()
	scanner.Scan()

	if scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 4 {
			return 0, fmt.Errorf("formato non valido")
		}
		// fields[2] è il livello del segnale (link quality)
		val := strings.TrimSuffix(fields[2], ".")
		quality, err := strconv.Atoi(val)
		if err != nil {
			return 0, err
		}
		// normalizza da 0-70 a 0-100%
		percent := quality * 100 / 70
		if percent > 100 {
			percent = 100
		}
		return percent, nil
	}
	return 0, fmt.Errorf("nessuna interfaccia wifi trovata")
}
