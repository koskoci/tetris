package main

import (
	"github.com/nsf/termbox-go"
	"sort"
)

type shape [][2]int

var frozenPixels shape

var board = func() shape {
	result := make(shape, 0, 1)
	x := boardX0
	y := boardY0

	for i := 0; i < width; i++ {
		result = append(result, [2]int{x + i, y})
		result = append(result, [2]int{x + i, y + height - 1})
	}
	for j := 0; j < height; j++ {
		result = append(result, [2]int{x, y + j})
		result = append(result, [2]int{x + 1, y + j})
		result = append(result, [2]int{x + width - 2, y + j})
		result = append(result, [2]int{x + width - 1, y + j})
	}
	return result
}()

func environment() shape {
	return append(board, frozenPixels...)
}

func (s shape) fullRows() shape {
	rowTally := make(map[int][][2]int)
	var fullRows shape

	for _, v := range s {
		rowTally[v[1]] = append(rowTally[v[1]], v)
	}

	for _, v := range rowTally {
		if len(v) == 20 {
			fullRows = append(fullRows, v...)
		}
	}

	return fullRows
}

func (s shape) explode() {
	rowIndices := make([]int, 0)

	for _, pixel := range s {
		y := pixel[1]
		if !includes(rowIndices, y) {
			rowIndices = append(rowIndices, y)
		}
	}
	sort.Ints(rowIndices)

	frozenPixels = frozenPixels.subtractShape(s)

	for _, y := range rowIndices {
		for i, v := range frozenPixels {
			if v[1] < y {
				frozenPixels[i][1]++
			}
		}
	}
}

func (s shape) subtractShape(z shape) shape {
	for _, pixel := range z {
		s = s.subtractPixel(pixel)
	}
	return s
}

func (s shape) subtractPixel(pixel [2]int) shape {
	result := make([][2]int, len(s))
	copy(result, s)

	i := index(result, pixel)
	result = append(result[:i], result[i+1:]...)
	return result
}

func (s shape) overlap(z shape) bool {
	result := false
	for _, v := range s {
		for _, w := range z {
			if v[0] == w[0] && v[1] == w[1] {
				result = true
			}
		}
	}
	return result
}

func (s shape) render(color termbox.Attribute) {
	for _, p := range s {
		termbox.SetCell(p[0], p[1], ' ', termbox.ColorDefault, color)
	}
}
