package main

import (
	"github.com/nsf/termbox-go"
	"math/rand"
	"time"
)

func main() {
	setup()
	eventQueue := startQueue()
	clock := startClock()
	defer close(eventQueue)
	defer close(clock)

	myPiece := newPiece()

mainloop:
	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		render(board, termbox.ColorWhite)
		render(myPiece.serialize(), myPiece.tetromino.color)

		termbox.Flush()
		time.Sleep(100 * time.Millisecond)

		select {
		case <-clock:
			myPiece = myPiece.move([2]int{0, 1})
		case ev := <-eventQueue:
			switch ev.Type {
			case termbox.EventKey:
				switch ev.Key {
				case termbox.KeyEsc:
					termbox.Interrupt()
				case termbox.KeyArrowRight:
					myPiece = myPiece.move([2]int{2, 0})
				case termbox.KeyArrowLeft:
					myPiece = myPiece.move([2]int{-2, 0})
				case termbox.KeyArrowDown:
					myPiece = myPiece.move([2]int{0, 2})
				case termbox.KeyTab:
					myPiece = myPiece.rotate()
				}
			case termbox.EventInterrupt:
				break mainloop
			}
		}
	}
	termbox.Close()
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
		p0:        [2]int{boardX0 + width/2, boardY0 + 1},
	}
}
