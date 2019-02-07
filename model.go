package main

import (
	"errors"
	"github.com/nsf/termbox-go"
)

const boardX0 = 3
const boardY0 = 3
const width = 24
const height = 22

var tetrominos = []tetromino{
	tetromino{`i`, []uint32{0x00ff0000, 0x0c0c0c0c, 0x0000ff00, 0x30303030}, termbox.ColorCyan},
	tetromino{`j`, []uint32{0x3030f000, 0xc0fc0000, 0x3c303000, 0x00fc0c00}, termbox.ColorBlue},
	tetromino{`l`, []uint32{0x30303c00, 0x00fcc000, 0xf0303000, 0x0cfc0000}, termbox.ColorWhite},
	tetromino{`o`, []uint32{0xf0f00000, 0xf0f00000, 0xf0f00000, 0xf0f00000}, termbox.ColorYellow},
	tetromino{`s`, []uint32{0x003cf000, 0xc0f03000, 0x3cf00000, 0x303c0c00}, termbox.ColorGreen},
	tetromino{`t`, []uint32{0x00fc3000, 0x30f03000, 0x30fc0000, 0x303c3000}, termbox.ColorMagenta},
	tetromino{`z`, []uint32{0x00f03c00, 0x30f0c000, 0xf03c0000, 0x0c3c3000}, termbox.ColorRed},
}

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

var frozenPixels shape

type tetromino struct {
	label   string
	configs []uint32
	color   termbox.Attribute
}

type shape [][2]int

type piece struct {
	tetromino tetromino
	rot       int
	p0        [2]int
}

func environment() shape {
	return append(board, frozenPixels...)
}

func (p piece) move(vector [2]int) (piece, error) {
	var result piece
	var err error

	desiredPiece := piece{
		tetromino: p.tetromino,
		rot:       p.rot,
		p0: [2]int{
			p.p0[0] + vector[0],
			p.p0[1] + vector[1],
		},
	}

	if overlap(environment(), desiredPiece.serialize()) {
		result, err = p, errors.New("cannot go there")
	} else {
		result, err = desiredPiece, nil
	}

	return result, err
}

func (p piece) rotate() piece {
	var result piece

	desiredPiece := piece{
		tetromino: p.tetromino,
		rot:       (p.rot + 1) % 4,
		p0:        p.p0,
	}

	if overlap(environment(), desiredPiece.serialize()) {
		result = p
	} else {
		result = desiredPiece
	}

	return result
}

func (p piece) freeze() {
	myShape := p.serialize()
	for _, pixel := range myShape {
		frozenPixels = append(frozenPixels, pixel)
	}
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
	rowIndices := make(map[int]bool, 0)

	for _, v := range s {
		rowIndices[v[1]] = true
	}

	frozenPixels = frozenPixels.subtractShape(s)

	for y, _ := range rowIndices {
		for i, v := range frozenPixels {
			if v[1] < y {
				frozenPixels[i][1]++
			}
		}
	}
}

func index(s shape, p [2]int) int {
	for i, v := range s {
		if v == p {
			return i
		}
	}
	return -1
}

func (s shape) subtractShape(z shape) shape {
	for _, v := range z {
		s = s.subtractPixel(v)
	}
	return s
}

func (s shape) subtractPixel(p [2]int) shape {
	result := make([][2]int, len(s))
	copy(result, s)

	i := index(result, p)
	result = append(result[:i], result[i+1:]...)
	return result
}
