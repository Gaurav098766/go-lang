package main

import "fmt"

type Motorcycle interface {
	Mileage() float64
}

type BMW struct {
	distance float64
	fuel     float64
	Avgspeed string
}

type Audi struct {
	distance float64
	fuel     float64
}

func (b BMW) Mileage() float64 {
	return b.distance / b.fuel
}

func (a Audi) Mileage() float64 {
	return a.distance / a.fuel
}

func totalMileage(m []Motorcycle) {
	tm := 0.0
	for _, v := range m {
		tm = tm + v.Mileage()
	}
	fmt.Println("Total Mileage per month %f km/l", tm)
}

func main() {
	b1 := BMW{
		distance: 167.9,
		fuel:     30,
		Avgspeed: "58",
	}

	a1 := Audi{
		distance: 157.9,
		fuel:     30,
	}

	person := []Motorcycle{a1, b1}
	totalMileage(person)
}
