package engine

type Platform struct {
	ScreenW     int32
	ScreenH     int32
	RenderW     int32
	RenderH     int32
	AspectRatio float32
	Landscape   bool
}

func (p *Platform) NewPlatform(sw int32, sh int32, rw int32, rh int32) {
	p.ScreenW = sw
	p.ScreenH = sh
	p.RenderW = rw
	p.RenderH = rh
	if p.RenderW > p.RenderH {
		p.Landscape = true
		p.AspectRatio = float32(p.RenderW) / float32(p.RenderH)
	} else {
		p.Landscape = false
		p.AspectRatio = float32(p.RenderH) / float32(p.RenderW)
	}
}
