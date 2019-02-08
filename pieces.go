package main

import "errors"

type piece struct {
	tetromino tetromino
	rot       int
	p0        [2]int
}

func (p piece) move(vector [2]int) (result piece, err error) {
	desiredPiece := piece{
		tetromino: p.tetromino,
		rot:       p.rot,
		p0: [2]int{
			p.p0[0] + vector[0],
			p.p0[1] + vector[1],
		},
	}

	desiredShape := desiredPiece.serialize()

	if desiredShape.overlap(environment()) {
		result, err = p, errors.New("cannot go there")
	} else {
		result, err = desiredPiece, nil
	}

	return
}

func (p piece) rotate() (result piece) {
	desiredPiece := piece{
		tetromino: p.tetromino,
		rot:       (p.rot + 1) % 4,
		p0:        p.p0,
	}

	desiredShape := desiredPiece.serialize()

	if desiredShape.overlap(environment()) {
		result = p
	} else {
		result = desiredPiece
	}

	return
}

func (p piece) freeze() {
	myShape := p.serialize()
	for _, pixel := range myShape {
		frozenPixels = append(frozenPixels, pixel)
	}
}

func (p piece) serialize() (result shape) {
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
	return
}
