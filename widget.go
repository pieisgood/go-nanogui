package gonanogui

import "github.com/go-gl/mathgl/mgl32"

type Widget struct {
	Parent         *Widget
	Theme          Theme
	Layout         Layout
	ID             string
	Pos            mgl32.Vec2
	Size           mgl32.Vec2
	FixedSize      mgl32.Vec2
	Children       []Widget
	Visible        bool
	Enabled        bool
	Focused        bool
	MouseFocus     bool
	ToolTip        string
	FontSize       int
	IconExtraScale float32
	Cursor         Cursor
}

type WidgetNode interface {
}
