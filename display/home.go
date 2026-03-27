package display

import "glass/hw"

type HomeScreen struct {
	Title   string
	scrollX int
}

func (s *HomeScreen) RenderScrollHorizontal(d *hw.Display) {
	if s.scrollX == 0 {
		s.scrollX = 128 // inizia da destra
	}
	d.DrawLabel(s.scrollX, 22, s.Title)
	s.scrollX -= 2
	if s.scrollX < -len(s.Title)*7 {
		s.scrollX = 128
	}
}

func (s *HomeScreen) RenderScrollVertical(d *hw.Display) {}

func (s *HomeScreen) Render(d *hw.Display) {
	d.DrawLabel(0, 22, s.Title)
}
