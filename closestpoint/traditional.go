package main

import (
	"math"
)

func calcClosestPoint(points []point, c point) point {
	var (
		closestDist  float64
		closestPoint point
	)

	for _, p := range points {
		dist := math.Sqrt(math.Pow(p.x-c.x, 2) +
			math.Pow(p.y-c.y, 2) +
			math.Pow(p.z-c.z, 2))

		if closestDist == 0 || dist < closestDist {
			closestDist = dist
			closestPoint = p
		}
	}
	return closestPoint
}
