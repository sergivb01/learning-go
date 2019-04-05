package main

import (
	"fmt"
	"time"
)

func benchFunc(name string, f func([]point, point) point, points []point, c point) point {
	t := time.Now()
	p := f(points, c)
	d := time.Since(t)

	fmt.Printf("Closest point to %+v is %+v\n", c, p)
	fmt.Printf("[#%s] Took %s to run.\n\n", name, d)
	return p
}

func generateRandomPoints(n int) []point {
	var points []point
	fmt.Printf("Started generating %d points...", n)
	for i := 0; i < n; i++ {
		points = append(points, point{
			x: randNumber(),
			y: randNumber(),
			z: randNumber(),
		})
	}
	fmt.Printf(" Finished!\n\n")
	return points
}
