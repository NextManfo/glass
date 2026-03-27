package display

import "glass/hw"

type Screen interface {
	Render(d *hw.Display)
}
