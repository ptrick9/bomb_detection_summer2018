package main

import (
	"fmt"
	"math"
)

type NodeParent interface{
	sensorReport() float32
	distance(b bomb) float32
	row(div int) int
	col(div int) int
}


type NodeImpl struct {
	id int
	x int
	y int
}

type NodeMovement interface{
	NodeParent
	move()
}

type bn struct {
	*NodeImpl
	x_speed int
	y_speed int

}

type wn struct {
	*NodeImpl
	speed int
	dir int
}

type rn struct {
	*NodeImpl
}

func (n *NodeImpl) row(div int) int {
	return n.x/div
}

func (n *NodeImpl) col(div int) int {
	return n.y/div
}

func (n *NodeImpl) distance(b bomb) float32 {
	dist := float32(math.Sqrt(math.Pow(float64(math.Abs(float64(n.x)-float64(b.x))), 2) + math.Pow(float64(math.Abs(float64(n.y)-float64(b.y))), 2)))
	if dist < 1.0 {
		return 10.0
	} else if dist < 2.0 {
		return 5.0
	} else if dist < 3.0 {
		return 1.0
	} else {
		return 0.0
	}

}

func (n *NodeImpl) sensorReport() float32 {
	return 0.1
}

func (n bn) String() string {
	return fmt.Sprintf("x: %v y: %v Xspeed: %v Yspeed: %v id: %v", n.x, n.y, n.x_speed, n.y_speed, n.id)
}

func (n wn) String() string {
	return fmt.Sprintf("x: %v y: %v speed: %v dir: %v id: %v", n.x, n.y, n.speed, n.dir, n.id)
}

func (n rn) String() string {
	return fmt.Sprintf("x: %v y: %v id: %v", n.x, n.y, n.id)
}

func (n* bn) move() {
	if n.x + n.x_speed < maxX && n.x +n.x_speed >= 0 {
		n.x = n.x + n.x_speed
	} else {
		if n.x + n.x_speed >= maxX {
			n.x = n.x - (n.x_speed - (maxX-1-n.x))
			n.x_speed = n.x_speed * -1
		} else {
			n.x = (n.x_speed + n.x)*-1
			n.x_speed = n.x_speed * -1
		}
	}
	if n.y + n.y_speed < maxY && n.y +n.y_speed >= 0 {
		n.y = n.y + n.y_speed
	} else {
		if n.y + n.y_speed >= maxY {
			n.y = n.y - (n.y_speed - (maxY-1-n.y))
			n.y_speed = n.y_speed * -1
		} else {
			n.y = (n.y_speed + n.y)*-1
			n.y_speed = n.y_speed * -1
		}
	}
}

func (n* wn) move() {
	if n.dir == 0 { //x axis
		if n.x + n.speed < maxX && n.x +n.speed >= 0 {
			n.x = n.x + n.speed
		} else {
			if n.x + n.speed >= maxX {
				n.x = n.x - (n.speed - (maxX-1-n.x))
				n.speed = n.speed * -1
			} else {
				n.x = (n.speed + n.x)*-1
				n.speed = n.speed * -1
			}
		}
	} else { //y axis
		if n.y + n.speed < maxY && n.y +n.speed >= 0 {
			n.y = n.y + n.speed
		} else {
			if n.y + n.speed >= maxY {
				n.y = n.y - (n.speed - (maxY-1-n.y))
				n.speed = n.speed * -1
			} else {
				n.y = (n.speed + n.y)*-1
				n.speed = n.speed * -1
			}
		}
	}
}

func (n* rn) move() {
	x_speed := rangeInt(-3,3)
	y_speed := rangeInt(-3, 3)
	if n.x + x_speed < maxX && n.x +x_speed >= 0 {
		n.x = n.x + x_speed
	} else {
		if n.x + x_speed >= maxX {
			n.x = n.x - (x_speed - (maxX-1-n.x))
		} else {
			n.x = (x_speed + n.x)*-1
		}
	}
	if n.y + y_speed < maxY && n.y +y_speed >= 0 {
		n.y = n.y + y_speed
	} else {
		if n.y + y_speed >= maxY {
			n.y = n.y - (y_speed - (maxY-1-n.y))
		} else {
			n.y = (y_speed + n.y)*-1
		}
	}
}