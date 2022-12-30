package circle

import (
	"fmt"
	"math"
)

type point struct {
	x float64
	y float64
}

type Circle struct {
	point  point
	radius float64
}

func New(x, y, radius float64) Circle {
	return Circle{
		point:  point{x: x, y: y},
		radius: radius,
	}
}

func (c *Circle) circumferencePoint(radian float64) point {
	return point{
		x: c.radius*math.Cos(radian) + c.point.x,
		y: c.radius*math.Sin(radian) + c.point.y,
	}
}

var fullRadian = 2 * math.Pi

/*
M x y
A rx ry x-axis-rotation large-arc-flag sweep-flag x y  (mid)
A rx ry x-axis-rotation large-arc-flag sweep-flag x y  (end)
L x y
*/
var pathDFormat = `M %.2f %.2f A %.2f %.2f 0 0 1 %.2f %.2f A %.2f %.2f 0 0 1 %.2f %.2f L %.2f %.2f`

// NewSlice creates a slice with start, end in percentages.
func (c *Circle) NewSlice(start, end float64) string {
	// mid point is required because larger slices will take a "shortcut" instead
	mid := start + (end-start)/2
	startPoint := c.circumferencePoint(start * fullRadian)
	midPoint := c.circumferencePoint(mid * fullRadian)
	endPoint := c.circumferencePoint(end * fullRadian)

	origin := c.point

	return fmt.Sprintf(pathDFormat,
		startPoint.x, startPoint.y,
		c.radius, c.radius, midPoint.x, midPoint.y,
		c.radius, c.radius, endPoint.x, endPoint.y,
		origin.x, origin.y)
}

/*
Helpful links:
- https://medium.com/hackernoon/a-simple-pie-chart-in-svg-dbdd653b6936
- https://www.w3.org/TR/SVG11/paths.html#PathData
*/
