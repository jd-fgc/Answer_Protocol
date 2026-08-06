package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type draggable struct {
	widget.BaseWidget
	content fyne.CanvasObject
	bounds  fyne.Size
}

func newDraggable(content fyne.CanvasObject, bounds fyne.Size) *draggable {
	d := &draggable{content: content, bounds: bounds}
	d.ExtendBaseWidget(d)
	return d
}

func (d *draggable) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(d.content)
}

func minF(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

func maxF(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}

func (d *draggable) clamp(pos fyne.Position) fyne.Position {
	rangeX := d.bounds.Width - d.Size().Width
	rangeY := d.bounds.Height - d.Size().Height

	lowX, highX := minF(0, rangeX), maxF(0, rangeX)
	lowY, highY := minF(0, rangeY), maxF(0, rangeY)

	if pos.X < lowX {
		pos.X = lowX
	}
	if pos.X > highX {
		pos.X = highX
	}
	if pos.Y < lowY {
		pos.Y = lowY
	}
	if pos.Y > highY {
		pos.Y = highY
	}
	return pos
}

func (d *draggable) Dragged(e *fyne.DragEvent) {
	newPos := d.Position().Add(fyne.NewPos(e.Dragged.DX, e.Dragged.DY))
	d.Move(d.clamp(newPos))
}

func (d *draggable) DragEnd() {}
