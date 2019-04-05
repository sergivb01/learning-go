package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	// pointsToGenerate = 100000000
	pointsToGenerate = 500000
)

var (
	cnt = 0
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	c := point{
		x: randNumber(),
		y: randNumber(),
		z: randNumber(),
	}

	points := generateRandomPoints(pointsToGenerate)

	p1 := benchFunc("Traditional", calcClosestPoint, points, c)
	p2 := benchFunc("Sergi", sergiMethod, points, c)
	if p1 != p2 {
		fmt.Printf("Different! C=%+v	| T=%+v	| S=%+v\n", c, p1, p2)
	}
}

type point struct {
	x, y, z float64
}

func randNumber() float64 {
	return rand.Float64() * 500
}
