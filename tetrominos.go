package main

import (
	"github.com/nsf/termbox-go"
)

var tetrominos = []tetromino{
	tetromino{`i`, []uint32{0x00ff0000, 0x0c0c0c0c, 0x0000ff00, 0x30303030}, termbox.ColorCyan},
	tetromino{`j`, []uint32{0x3030f000, 0xc0fc0000, 0x3c303000, 0x00fc0c00}, termbox.ColorBlue},
	tetromino{`l`, []uint32{0x30303c00, 0x00fcc000, 0xf0303000, 0x0cfc0000}, termbox.ColorWhite},
	tetromino{`o`, []uint32{0xf0f00000, 0xf0f00000, 0xf0f00000, 0xf0f00000}, termbox.ColorYellow},
	tetromino{`s`, []uint32{0x003cf000, 0xc0f03000, 0x3cf00000, 0x303c0c00}, termbox.ColorGreen},
	tetromino{`t`, []uint32{0x00fc3000, 0x30f03000, 0x30fc0000, 0x303c3000}, termbox.ColorMagenta},
	tetromino{`z`, []uint32{0x00f03c00, 0x30f0c000, 0xf03c0000, 0x0c3c3000}, termbox.ColorRed},
}

type tetromino struct {
	label   string
	configs []uint32
	color   termbox.Attribute
}
