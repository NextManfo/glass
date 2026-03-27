package display

import "glass/hw"

type StatusBar struct {
	Wifi     bool
	Battery  int
	Time     string
	Position int
}

func (s *StatusBar) Render(d *hw.Display) {
	d.AddLabel(0, s.Position, s.Time)
	d.DrawBitmap(IconWifi, 90, s.Position)
	d.DrawBitmap(IconBattery, 100, s.Position)
}
