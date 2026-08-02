package main

import (
	_ "embed"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

//go:embed assets/IsaacFont.ttf
var fontIsaac []byte

var fontResource = fyne.NewStaticResource("IsaacFont.ttf", fontIsaac)

type tapTheme struct{}

func (t *tapTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff}
	case theme.ColorNameForeground:
		return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	case theme.ColorNamePrimary:
		return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	case theme.ColorNameForegroundOnPrimary:
		return color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff}
	case theme.ColorNameButton:
		return color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff}
	case theme.ColorNameHover:
		return color.NRGBA{R: 0x22, G: 0x22, B: 0x22, A: 0xff}
	case theme.ColorNameInputBackground:
		return color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff}
	case theme.ColorNamePlaceHolder:
		return color.NRGBA{R: 0x77, G: 0x77, B: 0x77, A: 0xff}
	case theme.ColorNameDisabled:
		return color.NRGBA{R: 0x44, G: 0x44, B: 0x44, A: 0xff}
	}
	return theme.DefaultTheme().Color(name, variant)
}

func (t *tapTheme) Font(style fyne.TextStyle) fyne.Resource {
	return fontResource
}

func (t *tapTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (t *tapTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNameText:
		return 15
	case theme.SizeNamePadding:
		return 6
	}
	return theme.DefaultTheme().Size(name)
}
