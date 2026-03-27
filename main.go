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
	dis.DrawLabelScroll(22, "Inizializzazione display")
	dis.EndDraw()
	time.Sleep(10 * time.Second) // Delay di 2 secondi
	statusBar := display.StatusBar{Wifi: true, Battery: 100, Position: 10}
	var currentScreen display.Screen = &display.HomeScreen{Title: "Home screen"}
	for {
		statusBar.Time = time.Now().Format("15:04")
		dis.BeginDraw()
		statusBar.Render(dis)
		currentScreen.Render(dis)
		dis.EndDraw()
		time.Sleep(1 * time.Second)
	}
}
