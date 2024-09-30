package dataTypes

import "fmt"

type Circle struct {
	X, Y, Radius int
}

type Wheel struct {
	X, Y, Radius, Spokes int
}

func playWithCircleAndWheels() {
	var c Circle
	c.X = 10
	c.Y = 20
	c.Radius = 50
	fmt.Println("Circle : ", c)

	var w Wheel
	w.X = 11
	w.Y = 22
	w.Radius = 55
	w.Spokes = 24
	fmt.Println("Wheel : ", w)
}

type Point1 struct {
	X, Y int
}

type Circle1 struct {
	Center Point1
	Radius int
}

type Wheel1 struct {
	Circle Circle1
	Spokes int
}

func playWithCircleAndWheelsAgain() {
	var w Wheel1
	w.Circle.Center.X = 11
	w.Circle.Center.Y = 22
	w.Circle.Radius = 55
	w.Spokes = 24
	fmt.Println("Wheel1 : ", w)

	// w.X = 110
	// w.Y = 220
	// w.Radius = 550
	// w.Spokes = 240
	// fmt.Println("Wheel : ", w)

}

//________________________________________________________________

type Point2 struct {
	X, Y int
}

type Circle2 struct {
	Point2 //Structure Embedding
	Radius int
}

type Wheel2 struct {
	Circle2 //Structure Embedding
	Spokes  int
}

func playWithCircleAndWheelsOnceAgain() {
	var w Wheel2
	w.Circle2.Point2.X = 11
	w.Circle2.Point2.Y = 22
	w.Circle2.Radius = 55
	w.Spokes = 24
	fmt.Println("Wheel2 : ", w)

	w.X = 110
	w.Y = 220
	w.Radius = 550
	w.Spokes = 240
	fmt.Println("Wheel2.1 : ", w)

	var ww Wheel2
	ww = Wheel2{Circle2{Point2{88, 99}, 100}, 24}
	fmt.Println("Wheel2.2 : ", ww)

	var www Wheel2
	www = Wheel2{ // Labelled Initialser
		Circle2: Circle2{
			Point2: Point2{
				X: 88,
				Y: 99,
			},
			Radius: 100,
		},
		Spokes: 24,
	}
	fmt.Println("Wheel2.3 : ", www)

}

func playWithStruct() {
	playWithCircleAndWheels()
	playWithCircleAndWheelsAgain()
	playWithCircleAndWheelsOnceAgain()
}
