package pldb

import "fmt"

func Hello() {
	fmt.Println("hello")
}

type Farm struct {
	ID             int
	Title          string
	Specialization string
	HeadsOfAnimals int
	HeadsOfCows    int
	Longtitude     float64
	Latitude       float64
	Address        string
	OF_type        string
	SAL            float64
	RegionID       float64
}

type Region struct {
	ID           int
	Title        string
	Longtitude   float64
	Latitude     float64
	ApproxSquare float64
}
