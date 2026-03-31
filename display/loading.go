package display

import (
	"math"

	"glass/hw"
)

type LoadingScreen struct {
	offset float64
}

func (s *LoadingScreen) Render(d *hw.Display) {
	// ampiezza e centro verticale del body (y=11..32)
	centerY := 22.0
	amplitude := 8.0

	for x := 0; x < 128; x++ {
		// calcola y dell'onda con offset per l'animazione
		y := centerY + amplitude*math.Sin(float64(x)*0.2+s.offset)
		d.CurrentScreen.SetBit(x, int(y), true)

		// onda più spessa — disegna anche il pixel sopra e sotto
		d.CurrentScreen.SetBit(x, int(y)+1, true)
	}

	// avanza l'offset per il prossimo frame
	s.offset -= 0.3
}

func (s *LoadingScreen) RenderScrollHorizontal(d *hw.Display) {
	s.Render(d)
}

func (s *LoadingScreen) RenderScrollVertical(d *hw.Display) {}
