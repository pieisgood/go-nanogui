package gonanogui

import "github.com/go-gl/mathgl/mgl32"

type Alignment int

const (
	Minimum Alignment = iota + 1
	Middle
	Maximum
	Fill
)

type Orientation int

const (
	Horizontal Orientation = iota + 1
	Vertical
)

type Layout interface {
	PerformLayout()
	PreferredSize() mgl32.Vec2
}
