package std

// VEB combines vertex and element array buffers
type VEB struct {
	VB []float32
	EB []uint8
}

var QuadMesh *VEB
var CircleMesh *VEB

func init() {
	QuadMesh = &VEB{[]float32{
		1.0, 1.0, 0.0, 1.0, 1.0, 1.0, 1.0, 0.0,
		1.0, 0.0, 0.0, 1.0, 1.0, 1.0, 1.0, 1.0,
		0.0, 0.0, 0.0, 1.0, 1.0, 1.0, 0.0, 1.0,
		0.0, 1.0, 0.0, 1.0, 1.0, 1.0, 0.0, 0.0,
	},
		[]uint8{
			0, 1, 3,
			1, 2, 3,
		},
	}
}
