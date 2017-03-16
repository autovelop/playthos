package engine

type Platform struct {
	ScreenW float32
	ScreenH float32
	// RenderW     float32
	// RenderH     float32
	// AspectRatio float32
	// Landscape   bool
}

// func (p *Platform) NewPlatform(sw float32, sh float32, rw float32, rh float32) {
func (p *Platform) NewPlatform(sw float32, sh float32) {
	p.ScreenW = sw
	p.ScreenH = sh
	// p.RenderW = rw
	// p.RenderH = rh
	// if p.RenderW > p.RenderH {
	// 	p.Landscape = true
	// 	p.AspectRatio = p.RenderW / p.RenderH
	// } else {
	// 	p.Landscape = false
	// 	p.AspectRatio = p.RenderH / p.RenderW
	// }
}
