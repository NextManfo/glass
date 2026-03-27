package display

import "glass/hw"

type HomeScreen struct {
	Title string
}

func (s *HomeScreen) Render(d *hw.Display) {
	d.DrawLabelScroll(22, s.Title) // ha Render → implementa Screen
}
