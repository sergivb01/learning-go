package main

import "math"

func sergiMethod(points []point, c point) point {
	var (
		smallestDist float64
		closestPoint point
	)

	for _, p := range points {
		//dist := math.Abs(p.x-c.x) + math.Abs(p.y-c.y) + math.Abs(p.z-c.z)
		dist := math.Pow(p.x-c.x, 2) + math.Pow(p.y-c.y, 2) + math.Pow(p.z-c.z, 2)

		if smallestDist == 0 || dist < smallestDist {
			smallestDist = dist
			closestPoint = p
		}
	}
	return closestPoint
}
