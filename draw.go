package main

import (
	"github.com/nsf/termbox-go"
	"time"
)

type tetromino struct {
	label   string
	configs []uint16
	color   termbox.Attribute
}

var tetrominos = []tetromino{
	tetromino{`i`, []uint16{0x0F00, 0x2222, 0x00F0, 0x4444}, termbox.ColorCyan},
	tetromino{`j`, []uint16{0x44C0, 0x8E00, 0x6440, 0x0E20}, termbox.ColorBlue},
	tetromino{`l`, []uint16{0x4460, 0x0E80, 0xC440, 0x2E00}, termbox.ColorWhite},
	tetromino{`o`, []uint16{0xCC00, 0xCC00, 0xCC00, 0xCC00}, termbox.ColorYellow},
	tetromino{`s`, []uint16{0x06C0, 0x8C40, 0x6C00, 0x4620}, termbox.ColorGreen},
	tetromino{`t`, []uint16{0x0E40, 0x4C40, 0x4E00, 0x4640}, termbox.ColorMagenta},
	tetromino{`z`, []uint16{0x0C60, 0x4C80, 0xC600, 0x2640}, termbox.ColorRed},
}

// Draw a tetromino
func Draw(x, y, r int, l string) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	drawStatic()
	t := find(l)
	for i := 0; i < 8; i++ {
		for j := 0; j < 4; j++ {
			if hasBit(t.configs[r], uint(16-(4*j+i/2+1))) {
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

func hasBit(n uint16, pos uint) bool {
	val := n & (1 << pos)
	return (val > 0)
}
