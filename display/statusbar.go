package display

import "glass/hw"

type StatusBar struct {
	Wifi     bool
	Battery  int
	Time     string
	Position int
}

func (s *StatusBar) Render(d *hw.Display) {
	d.DrawLabel(30, s.Position, s.Time)
	d.DrawBitmap(IconWifi, 80, 0)
	d.DrawBitmap(IconBattery, 90, 0)
}
