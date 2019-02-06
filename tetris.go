package main

import (
	"github.com/nsf/termbox-go"
	"math/rand"
	"time"
)

func main() {
	setup()
	eventQueue := listen()
	myPiece := newPiece()

	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		render(board, termbox.ColorWhite)
		render(myPiece.serialize(), myPiece.tetromino.color)

		termbox.Flush()
		time.Sleep(100 * time.Millisecond)

		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey {
				switch ev.Key {
				case termbox.KeyEsc:
					exit()
				case termbox.KeyArrowRight:
					myPiece = myPiece.move([2]int{2, 0})
				case termbox.KeyArrowLeft:
					myPiece = myPiece.move([2]int{-2, 0})
				case termbox.KeyArrowDown:
					myPiece = myPiece.move([2]int{0, 2})
				case termbox.KeyTab:
					myPiece = myPiece.rotate()
				}
			}
		}
	}
}

func setup() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
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

func pick() tetromino {
	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(tetrominos))
	return tetrominos[index]
}

func newPiece() piece {
	return piece{
		tetromino: pick(),
		rot:       0,
		p0:        [2]int{boardX0 + width/2, boardY0 + 1},
	}
}
