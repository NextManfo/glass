package display

import "glass/hw"

type Screen interface {
	Render(d *hw.Display)
}

type ScrollableScreen interface {
	Screen
	RenderScrollHorizontal(d *hw.Display)
	RenderScrollVertical(d *hw.Display)
}
