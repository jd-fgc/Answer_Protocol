package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type roomCoord struct {
	row, col int
}

// Positionne chaque salle connue sur une grille, en partant de la salle courante,
// en suivant les liens North/South/East/West (BFS)
func computeRoomGrid(zones Rooms, start string) map[string]roomCoord {
	coords := map[string]roomCoord{start: {0, 0}}
	queue := []string{start}

	roomByName := map[string]Room{}
	for _, r := range zones.Rooms {
		roomByName[r.Name] = r
	}

	for len(queue) > 0 {
		name := queue[0]
		queue = queue[1:]
		r, ok := roomByName[name]
		if !ok {
			continue
		}
		c := coords[name]

		neighbors := map[string][2]int{
			r.North: {c.row - 1, c.col},
			r.South: {c.row + 1, c.col},
			r.East:  {c.row, c.col + 1},
			r.West:  {c.row, c.col - 1},
		}
		for nName, pos := range neighbors {
			if nName == "" {
				continue
			}
			if _, exists := coords[nName]; !exists {
				coords[nName] = roomCoord{pos[0], pos[1]}
				queue = append(queue, nName)
			}
		}
	}
	return coords
}

func buildMiniMap(zones Rooms, currentZone string, visited map[string]bool) fyne.CanvasObject {
	coords := computeRoomGrid(zones, currentZone)

	minRow, maxRow, minCol, maxCol := 0, 0, 0, 0
	for _, c := range coords {
		if c.row < minRow {
			minRow = c.row
		}
		if c.row > maxRow {
			maxRow = c.row
		}
		if c.col < minCol {
			minCol = c.col
		}
		if c.col > maxCol {
			maxCol = c.col
		}
	}
	rows := maxRow - minRow + 1
	cols := maxCol - minCol + 1

	const cellSize float32 = 60

	grid := container.NewWithoutLayout()
	grid.Resize(fyne.NewSize(float32(cols)*cellSize, float32(rows)*cellSize))

	for name, c := range coords {
		x := float32(c.col-minCol) * cellSize
		y := float32(c.row-minRow) * cellSize

		var cell fyne.CanvasObject
		switch {
		case name == currentZone:
			rect := canvas.NewRectangle(theme.Color(theme.ColorNamePrimary))
			rect.StrokeColor = theme.Color(theme.ColorNameForeground)
			rect.StrokeWidth = 2
			cell = rect
		case visited[name]:
			rect := canvas.NewRectangle(color.NRGBA{R: 0x33, G: 0x33, B: 0x33, A: 0xff})
			rect.StrokeColor = theme.Color(theme.ColorNameForeground)
			rect.StrokeWidth = 1
			cell = rect
		default:
			rect := canvas.NewRectangle(color.NRGBA{R: 0x11, G: 0x11, B: 0x11, A: 0xff})
			rect.StrokeColor = theme.Color(theme.ColorNameForeground)
			rect.StrokeWidth = 1
			mark := widget.NewLabelWithStyle("?", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
			cell = container.NewStack(rect, container.NewCenter(mark))
		}

		cell.Resize(fyne.NewSize(cellSize-4, cellSize-4))
		cell.Move(fyne.NewPos(x+2, y+2))
		grid.Add(cell)
	}

	return container.NewCenter(grid)
}
