package main

import (
	"image"
	"image/color"
	"log"
	"time"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"

	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/devices/v3/ssd1306"
	"periph.io/x/host/v3"
)

func main() {
	// Init periph
	_, err := host.Init()
	if err != nil {
		log.Fatal(err)
	}

	// Open I2C
	bus, err := i2creg.Open("")
	if err != nil {
		log.Fatal(err)
	}
	defer bus.Close()

	// Init display 128x32
	dev, err := ssd1306.NewI2C(bus, &ssd1306.Opts{
		W:       128,
		H:       32,
		Rotated: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Crea buffer immagine
	img := image.NewGray(image.Rect(0, 0, 128, 32))

	// Colori
	white := color.Gray{Y: 255}

	// --- Disegna rettangolo ---
	for x := 10; x < 118; x++ {
		img.SetGray(x, 5, white)  // top
		img.SetGray(x, 25, white) // bottom
	}
	for y := 5; y < 25; y++ {
		img.SetGray(10, y, white)  // left
		img.SetGray(118, y, white) // right
	}

	// --- Disegna testo ---
	addLabel(img, 30, 18, "CIAO!")

	// Scrivi su display
	err = dev.Draw(img.Bounds(), img, image.Point{})
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(10 * time.Second)
}

// Funzione per disegnare testo
func addLabel(img *image.Gray, x, y int, label string) {
	col := color.Gray{Y: 255}

	point := fixed.Point26_6{
		X: fixed.I(x),
		Y: fixed.I(y),
	}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}

	d.DrawString(label)
}
