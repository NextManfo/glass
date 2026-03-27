package hw

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"log"
	"time"

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
		Rotated: false,
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
func (display *Display) DrawLabel(x, y int, label string) {
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

func (display *Display) EndDraw() {
	display.Device.Draw(display.Device.Bounds(), display.CurrentScreen, image.Point{})
}

func (display *Display) DrawLabelScroll(y int, label string) {
	x := 128 // parte da destra fuori schermo
	img := display.CurrentScreen
	drawer := &font.Drawer{Face: basicfont.Face7x13}
	testoWidth := int(drawer.MeasureString(label) >> 6)

	for {
		// Pulisci
		draw.Draw(img, img.Bounds(), &image.Uniform{color.Black}, image.Point{}, draw.Src)

		// Disegna testo nella posizione corrente
		d := &font.Drawer{
			Dst:  img,
			Src:  &image.Uniform{color.White},
			Face: basicfont.Face7x13,
			Dot:  fixed.P(x, y),
		}
		d.DrawString(label)

		// Invia al display
		if err := display.Device.Draw(display.CurrentScreen.Bounds(), img, image.Point{}); err != nil {
			log.Fatal(err)
		}

		// Muovi a sinistra
		x -= 2 // ← velocità di scorrimento in pixel

		// Quando il testo è uscito completamente a sinistra, ricomincia da destra
		if x < -testoWidth {
			x = 128
		}

		time.Sleep(20 * time.Millisecond) // ← controlla la fluidità
	}
}
