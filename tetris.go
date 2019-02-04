package main

import (
	"github.com/nsf/termbox-go"
	"math/rand"
	"time"
)

var width, height int
var x0, y0 = 5, 5
var pieces = []string{`i`, `j`, `l`, `o`, `s`, `t`, `z`}

func main() {
	setup()
	eventQueue := listen()

	x := x0 + 9
	r := 0
	for y := y0; y < y0+20; y++ {
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey {
				switch ev.Key {
				case termbox.KeyEsc:
					exit()
				case termbox.KeyArrowRight:
					x = x + 2
					Draw(x, y, r, pick())
				case termbox.KeyArrowLeft:
					x = x - 2
					Draw(x, y, r, pick())
				case termbox.KeyTab:
					r = (r + 1) % 4
					Draw(x, y, r, pick())
				}
			}
		default:
			Draw(x, y, r, pick())
		}
	}
	exit()
}

func setup() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	width, height = termbox.Size()
}

func listen() (eventQueue chan termbox.Event) {
	eventQueue = make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	return eventQueue
}

func exit() {
	termbox.Close()
}

func pick() (s string) {
	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(pieces))
	return pieces[index]
}
