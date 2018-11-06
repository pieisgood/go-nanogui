package gonanogui

type BoxLayout struct {
	orientation Orientation
	alignment   Alignment
	margin      int
	spacing     int
}

func (b *BoxLayout) Orientation() Orientation {
	return b.orientation
}

func (b *BoxLayout) SetOrientation(o Orientation) {
	b.orientation = o
}

func (b *BoxLayout) Alignment() Alignment {
	return b.alignment
}

func (b *BoxLayout) SetAlignment(a Alignment) {
	b.alignment = a
}

func (b *BoxLayout) Spacing() int {
	return b.spacing
}

func (b *BoxLayout) SetSpacing(s int) {
	b.spacing = s
}

func (b *BoxLayout) PreferredSize() {

}

func (b *BoxLayout) PerformLayout() {

}

func (b *BoxLayout) Margin() int {
	return b.margin
}

func (b *BoxLayout) SetMargin(m int) {
	b.margin = m
}
