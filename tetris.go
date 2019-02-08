package main

import (
	"github.com/nsf/termbox-go"
	"math/rand"
	"time"
)

const boardX0 = 3
const boardY0 = 3
const width = 24
const height = 22

var start = [2]int{boardX0 + width/2 - 2, boardY0 + 1}

var kill bool

func main() {
	setup()
	defer termbox.Close()

	eventQueue := startQueue()
	clock := startClock()

	for {
		if kill || gameOver() {
			return
		}
		drop(newPiece(), clock, eventQueue)
	}
}

func gameOver() (result bool) {
	for _, pixel := range frozenPixels {
		if pixel[1] == boardY0+2 {
			result = true
		} else {
			result = false
		}
	}
	return
}

func drop(p piece, clock chan bool, eventQueue chan termbox.Event) {
	for {
		renderAll(p)

		select {
		case <-clock:
			var err error
			p, err = p.move([2]int{0, 1})
			if err != nil {
				p.freeze()
				return
			}
		case ev := <-eventQueue:
			switch ev.Type {
			case termbox.EventKey:
				switch ev.Key {
				case termbox.KeyEsc:
					kill = true
					return
				case termbox.KeyArrowRight:
					p, _ = p.move([2]int{2, 0})
				case termbox.KeyArrowLeft:
					p, _ = p.move([2]int{-2, 0})
				case termbox.KeyArrowDown:
					var err error
					p, err = p.move([2]int{0, 1})
					if err != nil {
						p.freeze()
						return
					}
				case termbox.KeyTab:
					p = p.rotate()
				}
			case termbox.EventError:
				panic(ev.Err)
			}
		}
	}
}

func renderAll(p piece) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	renderBoard()
	renderFrozenPixels()
	handleFullRows()
	renderPiece(p)

	termbox.Flush()
	time.Sleep(100 * time.Millisecond)
}

func handleFullRows() {
	fullRows := frozenPixels.fullRows()

	if len(fullRows) > 0 {
		fullRows.render(termbox.ColorMagenta)
		time.Sleep(50 * time.Millisecond)
		fullRows.explode()
	}
}

func renderBoard() {
	board.render(termbox.ColorWhite)
}

func renderFrozenPixels() {
	frozenPixels.render(termbox.ColorWhite)
}

func renderPiece(p piece) {
	p.serialize().render(p.tetromino.color)
}

func setup() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
}

func startQueue() chan termbox.Event {
	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	return eventQueue
}

func startClock() chan bool {
	clock := make(chan bool)
	go func() {
		for {
			time.Sleep(300 * time.Millisecond)
			clock <- true
		}
	}()

	return clock
}

func pick() tetromino {
	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(tetrominos))
	return tetrominos[index]
}

func newPiece() piece {
	return piece{
		tetromino: pick(),
		rot:       0,
		p0:        start,
	}
}
