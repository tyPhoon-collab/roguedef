package system

import (
	"image/color"
	"math"
	"roguedef/ds"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type VectorDrawer struct {
	*Transform
	ops []VectorDrawerOptions
}

type VectorOpType int

const (
	VectorOpMove VectorOpType = iota
	VectorOpLine
)

type VectorDrawerOptions struct {
	OpType VectorOpType
	Data   any
}

func NewMoveOp(v vec) VectorDrawerOptions {
	return VectorDrawerOptions{
		OpType: VectorOpMove,
		Data:   v,
	}
}

func NewLineOp(v vec) VectorDrawerOptions {
	return VectorDrawerOptions{
		OpType: VectorOpLine,
		Data:   v,
	}
}

func (v *VectorDrawer) Draw(screen *ebiten.Image) {
	path := vector.Path{}

	// offset. This is Transform.Pos
	ox, oy := ds.Unpack[float32](v.Pos)
	path.MoveTo(ox, oy)
	for _, op := range v.ops {
		switch op.OpType {
		case VectorOpMove:
			x, y := ds.Unpack[float32](op.Data.(vec))
			path.MoveTo(x+ox, y+oy)
		case VectorOpLine:
			x, y := ds.Unpack[float32](op.Data.(vec))
			path.LineTo(x+ox, y+oy)
		}
	}
	path.Close()

	vs, is := path.AppendVerticesAndIndicesForFilling(nil, nil)

	for i := range vs {
		vs[i].SrcX = 0
		vs[i].SrcY = 0
		vs[i].ColorR = 1
		vs[i].ColorG = 1
		vs[i].ColorB = 1
		vs[i].ColorA = 1
	}

	op := &ebiten.DrawTrianglesOptions{}
	op.AntiAlias = true
	op.FillRule = ebiten.EvenOdd

	img := ebiten.NewImage(1, 1)
	img.Fill(color.White)
	screen.DrawTriangles(vs, is, img, op)
}

func (v *VectorDrawer) WithTransform(transform *Transform) *VectorDrawer {
	if transform == nil {
		transform = NewTransform()
	}
	v.Transform = transform
	return v
}

func NewVectorDrawer(ops ...VectorDrawerOptions) *VectorDrawer {
	return (&VectorDrawer{ops: ops}).WithTransform(nil)
}

func NewVectorDrawerTriangle(size float64) *VectorDrawer {
	// triangle height
	h := size * math.Sqrt(3.0) / 2.0
	return NewVectorDrawer(
		NewMoveOp(vec{X: 0, Y: 2 * h / 3.0}),
		NewLineOp(vec{X: size / 2.0, Y: -h / 3.0}),
		NewLineOp(vec{X: -size / 2.0, Y: -h / 3.0}),
	)
}

func NewVectorDrawerSquare(size float64) *VectorDrawer {
	// half size
	hSize := size / 2.0
	return NewVectorDrawer(
		NewMoveOp(vec{X: hSize, Y: hSize}),
		NewLineOp(vec{X: hSize, Y: -hSize}),
		NewLineOp(vec{X: -hSize, Y: -hSize}),
		NewLineOp(vec{X: -hSize, Y: hSize}),
	)
}
