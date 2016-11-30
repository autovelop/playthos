package gde

type Device struct {
	ScreenW     int32
	ScreenH     int32
	RenderW     int32
	RenderH     int32
	AspectRatio float32
	Landscape   bool
}

func (d *Device) NewDevice(sw int32, sh int32, rw int32, rh int32) {
	d.ScreenW = sw
	d.ScreenH = sh
	d.RenderW = rw
	d.RenderH = rh
	if d.RenderW > d.RenderH {
		d.Landscape = true
		d.AspectRatio = float32(d.RenderW) / float32(d.RenderH)
	} else {
		d.Landscape = false
		d.AspectRatio = float32(d.RenderH) / float32(d.RenderW)
	}
}
