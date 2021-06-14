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
	RegionID       int
}

type Region struct {
	ID           int
	Title        string
	Longtitude   float64
	Latitude     float64
	ApproxSquare float64
}

type FarmCharacteristics struct {
	ManureMass                    float64
	NitrogenMassInFertilizer      float64
	PhosphorMassInFertilizer      float64
	NitrogenMassForSoil           float64
	PhosphorMassForSoil           float64
	FertilizerPotentialByNitrogen float64
	FertilizerPotentialByPhosphor float64
	SquareDemandForNitrogen       float64
	SquareDemandForPhosphor       float64
	SquareFreeForNitrogen         float64
	SquareFreeForPhosphor         float64
	DemandForOFStorage            float64
	FarmID                        int
}

type RegionCharacteristics struct {
	ManureMass                    float64
	NitrogenMassInFertilizer      float64
	PhosphorMassInFertilizer      float64
	NitrogenMassForSoil           float64
	PhosphorMassForSoil           float64
	FertilizerPotentialByNitrogen float64
	FertilizerPotentialByPhosphor float64
	SquareDemandForFertilizer     float64
	SquareFreeForFertilizer       float64
	DemandForOFStorage            float64
	Consumers                     float64
	Producers                     float64
	RegionID                      int
}
