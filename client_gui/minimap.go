package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const mapCellSize float32 = 54

type roomCoord struct{ row, col int }

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

func makeCell(isCurrent, isVisited bool) fyne.CanvasObject {
	size := fyne.NewSize(mapCellSize-4, mapCellSize-4)
	switch {
	case isCurrent:
		rect := canvas.NewRectangle(theme.Color(theme.ColorNamePrimary))
		rect.StrokeColor = theme.Color(theme.ColorNameForeground)
		rect.StrokeWidth = 2
		rect.SetMinSize(size)
		return rect
	case isVisited:
		rect := canvas.NewRectangle(color.NRGBA{R: 0x33, G: 0x33, B: 0x33, A: 0xff})
		rect.StrokeColor = theme.Color(theme.ColorNameForeground)
		rect.StrokeWidth = 1
		rect.SetMinSize(size)
		return rect
	default:
		rect := canvas.NewRectangle(color.NRGBA{R: 0x11, G: 0x11, B: 0x11, A: 0xff})
		rect.StrokeColor = theme.Color(theme.ColorNameForeground)
		rect.StrokeWidth = 1
		mark := widget.NewLabelWithStyle("?", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
		return container.NewStack(rect, container.NewCenter(mark))
	}
}

func buildMiniMap(zones Rooms, currentZone string, visited map[string]bool) (fyne.CanvasObject, fyne.Size) {
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
	size := fyne.NewSize(float32(cols)*mapCellSize, float32(rows)*mapCellSize)

	grid := container.NewWithoutLayout()
	grid.Resize(size)

	for name, c := range coords {
		x := float32(c.col-minCol) * mapCellSize
		y := float32(c.row-minRow) * mapCellSize

		cell := makeCell(name == currentZone, visited[name])
		cell.Resize(fyne.NewSize(mapCellSize-4, mapCellSize-4))
		cell.Move(fyne.NewPos(x+2, y+2))
		grid.Add(cell)
	}

	return grid, size
}
