package main

import (
	"time"

	"glass/display"
	"glass/hw"
)

func main() {
	// Init periph
	dis := hw.NewDisplay()
	dis.BeginDraw()
	dis.AddLabel(10, 30, "Inizializzazione display")
	dis.EndDraw()
	time.Sleep(10 * time.Second) // Delay di 2 secondi
	dis.BeginDraw()
	dis.ClearDisplay()
	dis.AddLabel(10, 30, "Ciao Paolo")
	dis.DrawBitmap(display.IconBattery, 110, 8)
	dis.EndDraw()
}
