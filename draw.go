package main

import (
	"github.com/nsf/termbox-go"
	"time"
)

type tetromino struct {
	label   string
	configs []uint32
	color   termbox.Attribute
}

var tetrominos = []tetromino{
	tetromino{`i`, []uint32{0x00ff0000, 0x0c0c0c0c, 0x0000ff00, 0x30303030}, termbox.ColorCyan},
	tetromino{`j`, []uint32{0x3030f000, 0xc0fc0000, 0x3c303000, 0x00fc0c00}, termbox.ColorBlue},
	tetromino{`l`, []uint32{0x30303c00, 0x00fcc000, 0xf0303000, 0x0cfc0000}, termbox.ColorWhite},
	tetromino{`o`, []uint32{0xf0f00000, 0xf0f00000, 0xf0f00000, 0xf0f00000}, termbox.ColorYellow},
	tetromino{`s`, []uint32{0x003cf000, 0xc0f03000, 0x3cf00000, 0x303c0c00}, termbox.ColorGreen},
	tetromino{`t`, []uint32{0x00fc3000, 0x30f03000, 0x30fc0000, 0x303c3000}, termbox.ColorMagenta},
	tetromino{`z`, []uint32{0x00f03c00, 0x30f0c000, 0xf03c0000, 0x0c3c3000}, termbox.ColorRed},
}

// Draw all the things
func Draw(x, y, r int, l string) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	drawStatic()
	t := find(l)
	for i := 0; i < 8; i++ {
		for j := 0; j < 4; j++ {
			if hasBit(t.configs[r], uint(32-(8*j+i+1))) {
				termbox.SetCell(x+i, y+j, ' ', termbox.ColorDefault, t.color)
			}
		}
	}
	termbox.Flush()
	time.Sleep(100 * time.Millisecond)
}

func drawStatic() {
	drawBoard(4, 4)
	// drawFrozen()
}

func drawBoard(x, y int) {
	const width = 24
	const height = 22

	for i := 0; i < width; i++ {
		termbox.SetCell(x+i, y, ' ', termbox.ColorDefault, termbox.ColorWhite)
		termbox.SetCell(x+i, y+height-1, ' ', termbox.ColorDefault, termbox.ColorWhite)
	}
	for j := 0; j < height; j++ {
		termbox.SetCell(x, y+j, ' ', termbox.ColorDefault, termbox.ColorWhite)
		termbox.SetCell(x+1, y+j, ' ', termbox.ColorDefault, termbox.ColorWhite)
		termbox.SetCell(x+width-2, y+j, ' ', termbox.ColorDefault, termbox.ColorWhite)
		termbox.SetCell(x+width-1, y+j, ' ', termbox.ColorDefault, termbox.ColorWhite)
	}
}

func find(s string) (t tetromino) {
	var result tetromino
	for _, v := range tetrominos {
		if v.label == s {
			result = v
		}
	}
	return result
}

func hasBit(n uint32, pos uint) bool {
	val := n & (1 << pos)
	return (val > 0)
}
