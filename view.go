package main

import (
	"github.com/nsf/termbox-go"
)

func (p piece) serialize() shape {
	result := make(shape, 0, 1)
	config := p.tetromino.configs[p.rot]
	x0 := p.p0[0]
	y0 := p.p0[1]

	for i := 0; i < 8; i++ {
		for j := 0; j < 4; j++ {
			if hasBit(config, uint(32-(8*j+i)-1)) {
				pixel := [2]int{x0 + i, y0 + j}
				result = append(result, pixel)
			}
		}
	}
	return result
}

func overlap(s, z shape) bool {
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

func render(s shape, color termbox.Attribute) {
	for _, p := range s {
		termbox.SetCell(p[0], p[1], ' ', termbox.ColorDefault, color)
	}
}

func hasBit(n uint32, pos uint) bool {
	val := n & (1 << pos)
	return (val > 0)
}
