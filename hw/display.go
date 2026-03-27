package hw

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"periph.io/x/conn/v3/i2c"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/devices/v3/ssd1306"
	"periph.io/x/devices/v3/ssd1306/image1bit"
	"periph.io/x/host/v3"
)

type Display struct {
	Device        *ssd1306.Dev
	CurrentScreen *image1bit.VerticalLSB
	Bus           i2c.BusCloser
}

func NewDisplay() *Display {
	newDisplay := Display{}
	// Init periph
	_, err := host.Init()
	if err != nil {
		log.Fatal(err)
	}

	// Open I2C
	bus, err := i2creg.Open("1")
	if err != nil {
		log.Fatal(err)
	}
	newDisplay.Bus = bus

	// Init display 128x32
	dev, err := ssd1306.NewI2C(bus, &ssd1306.Opts{
		W:       128,
		H:       32,
		Rotated: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inizializzato nuovo display: %v ", dev.Bounds().Size())
	newDisplay.Device = dev
	return &newDisplay
}

// ClearDisplay Pulisce il dispay impostandolo tutto nero
func (display *Display) ClearDisplay() {
	draw.Draw(display.CurrentScreen, display.CurrentScreen.Bounds(), &image.Uniform{color.Black}, image.Point{}, draw.Src)
}

// Disegna le icone
func (display *Display) DrawBitmap(bmp [][]int, offsetX, offsetY int) {
	for y, row := range bmp {
		for x, px := range row {
			if px == 1 {
				display.CurrentScreen.SetBit(offsetX+x, offsetY+y, true)
			}
		}
	}
}

// Scrive del testo
func (display *Display) AddLabel(x, y int, label string) {
	col := color.Gray{Y: 255}

	point := fixed.Point26_6{
		X: fixed.I(x),
		Y: fixed.I(y),
	}

	d := &font.Drawer{
		Dst:  display.CurrentScreen,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}

	d.DrawString(label)
}

func (display *Display) BeginDraw() {
	display.CurrentScreen = image1bit.NewVerticalLSB(display.Device.Bounds())
	display.ClearDisplay() // Pulisco gia il display se rinizio a disegnare
}

func (display *Display) Reset() {
	display.Bus.Close()
	display = NewDisplay()
}
