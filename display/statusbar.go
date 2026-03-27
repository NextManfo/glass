package display

import (
	"glass/hw"
)

type StatusBar struct {
	Wifi         bool
	WifiStrength int
	Battery      int
	Time         string
	Position     int
}

func (s *StatusBar) Render(d *hw.Display) {
	d.DrawLabel(30, s.Position, s.Time)
	d.DrawBitmap(IconWifi, 75, 1)

	switch {
	case s.Battery > 60:
		d.DrawBitmap(IconBattery, 85, 1)
	case s.Battery > 20:
		d.DrawBitmap(IconBatteryMid, 85, 1)
	case s.Battery <= 20:
		d.DrawBitmap(IconBatteryLow, 85, 1)
	default:
		d.DrawBitmap(IconBattery, 85, 1)
	}

	switch {
	case s.WifiStrength > 60:
		d.DrawBitmap(IconWifi, 100, 1)
	case s.WifiStrength > 30:
		d.DrawBitmap(IconWifiMid, 100, 1)
	case s.WifiStrength > 0:
		d.DrawBitmap(IconWifiLow, 100, 1)
	default:
		d.DrawBitmap(IconWifiOff, 100, 1)
	}
}
