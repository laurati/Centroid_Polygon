package main

import "fmt"

//Point2D Structure
type Point2D struct {
	x float64
	y float64
}

func compute2DPolygonCentroid(vertices []Point2D, vertexCount int) Point2D {
	centroid := Point2D{0, 0}
	signedArea := 0.0
	x0 := 0.0
	y0 := 0.0
	x1 := 0.0
	y1 := 0.0
	a := 0.0

	// For all vertices except last
	i := 0
	for i < vertexCount-1 {
		x0 = vertices[i].x
		y0 = vertices[i].y
		x1 = vertices[i+1].x
		y1 = vertices[i+1].y
		a = x0*y1 - x1*y0
		signedArea += a
		centroid.x += (x0 + x1) * a
		centroid.y += (y0 + y1) * a
		i++
	}

	// Do last vertex separately to avoid performing an expensive
	// modulus operation in each iteration.
	x0 = vertices[i].x
	y0 = vertices[i].y
	x1 = vertices[0].x
	y1 = vertices[0].y
	a = x0*y1 - x1*y0
	signedArea += a
	centroid.x += (x0 + x1) * a
	centroid.y += (y0 + y1) * a

	signedArea *= 0.5
	centroid.x /= (6.0 * signedArea)
	centroid.y /= (6.0 * signedArea)
	return centroid
}

func main() {
	polygon := []Point2D{{25.774, -80.19}, {18.466, -66.118}, {32.321, -64.757}, {25.774, -80.19}}
	vertexCount := len(polygon)
	fmt.Println("Centroid:", compute2DPolygonCentroid(polygon, vertexCount))

}
