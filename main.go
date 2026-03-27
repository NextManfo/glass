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
	dis.DrawLabel(0, 22, "Inizializzazione display")
	dis.EndDraw()
	time.Sleep(10 * time.Second) // Delay di 2 secondi
	statusBar := display.StatusBar{Wifi: true, Battery: 100, Position: 10}
	var currentScreen display.Screen = &display.HomeScreen{Title: "Home screen"}
	tickDisplay := time.NewTicker(30 * time.Millisecond) // scroll fluido
	tickStatus := time.NewTicker(1 * time.Second)        // statusbar e MQTT

	for {
		select {
		// case newMsg := <-msgChan:
		//	currentScreen = &display.ChatScreen{Message: newMsg}

		case <-tickStatus.C:
			statusBar.Time = time.Now().Format("15:04")
			battery, err := hw.ReadBattery(dis.Bus)
			if err == nil {
				statusBar.Battery = battery
			}
			wifi, err := hw.ReadWifiStrength()
			if err == nil {
				statusBar.WifiStrength = wifi
			} else {
				statusBar.WifiStrength = 0 // nessun segnale
			}

		case <-tickDisplay.C:
			dis.BeginDraw()
			statusBar.Render(dis)
			if scrollable, ok := currentScreen.(display.ScrollableScreen); ok {
				scrollable.RenderScrollHorizontal(dis)
			} else {
				currentScreen.Render(dis)
			}
			dis.EndDraw()
		}
	}
}
