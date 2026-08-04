package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type tappableImage struct {
	widget.BaseWidget
	image  *canvas.Image
	border *canvas.Rectangle
	onTap  func()
}

func newTappableImage(path string, size fyne.Size, onTap func()) *tappableImage {
	img := canvas.NewImageFromFile(path)
	img.FillMode = canvas.ImageFillContain
	img.ScaleMode = canvas.ImageScalePixels
	img.SetMinSize(size)

	border := canvas.NewRectangle(color.Transparent)
	border.StrokeColor = color.Transparent
	border.StrokeWidth = 3

	t := &tappableImage{image: img, border: border, onTap: onTap}
	t.ExtendBaseWidget(t)
	return t
}

func (t *tappableImage) CreateRenderer() fyne.WidgetRenderer {
	content := container.NewStack(t.image, t.border)
	return widget.NewSimpleRenderer(content)
}

func (t *tappableImage) Tapped(_ *fyne.PointEvent) {
	if t.onTap != nil {
		t.onTap()
	}
}

func (t *tappableImage) MouseIn(_ *desktop.MouseEvent) {
	t.border.StrokeColor = theme.Color(theme.ColorNamePrimary)
	t.border.Refresh()
}

func (t *tappableImage) MouseOut() {
	t.border.StrokeColor = color.Transparent
	t.border.Refresh()
}

func (t *tappableImage) MouseMoved(_ *desktop.MouseEvent) {}
