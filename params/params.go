package params

import (
	"fmt"

	"github.com/PhilLar/proj/pldb"
)

const (
	K1  = 100
	K2  = 10.25
	K3  = 0.12
	K4  = 0.03
	K5  = 10
	K6  = 5
	K7  = 0.48
	K8  = 0.002
	K9  = 25
	K10 = 0.35
	K11 = 10
	K12 = 0.24
	K13 = 1.4
	K14 = 2
	K15 = 170
	K16 = 300
	K17 = 1.1
	K18 = 0.3
	K19 = 0.17
	K20 = 0.002
	K21 = 5
	K22 = 0.15
	K23 = 1
	K24 = 0.16
	K25 = 0.8
	K26 = 0.9
	K27 = 25
	K28 = 0.3
	K29 = 0.2
	K30 = 0.1
	K31 = 0.5
	K32 = 0.2
	K33 = 0.15
	K34 = 0.1
	K35 = 0.5

	VN  = 83
	VKS = 75
	VP  = 22
	T   = 0.5
)

// Расчет массы образуемого навоза/помета, т/год
func ManureMass(farm pldb.Farm) float64 {
	if farm.OF_type == "ТОУ" {

		if farm.Specialization == "КРС" {
			MN := float64(farm.HeadsOfCows) * K1 * 0.365
			MV := MN * ((VN - VKS) / (VKS - VP))
			fmt.Println(MN)
			fmt.Println(MV)
			return (MN + MV) * (100 - K5) / 100
		}

		if farm.Specialization == "Птицеводство (куры)" {
			MP := float64(farm.HeadsOfAnimals) * K3 * 0.365
			MV := MP / 2
			fmt.Println(MP)
			fmt.Println(MV)
			return (MP + MV*T) * (100 - K5) / 100
		}
		if farm.Specialization == "Птицеводство (перепелки)" {
			MP := float64(farm.HeadsOfAnimals) * K4 * 0.365
			MV := MP / 2
			fmt.Println(MP)
			fmt.Println(MV)
			return (MP + MV*T) * (100 - K5) / 100
		}

	} else {
		if farm.Specialization == "Свиноводство" {
			MN := float64(farm.HeadsOfAnimals) * K2 * 0.365
			return MN * (100 - K6) / 100
		}
		if farm.Specialization == "КРС" {
			MN := float64(farm.HeadsOfCows) * K1 * 0.365
			return MN * (100 - K6) / 100
		}
	}
	return 0
}

func NitrogenMassInFertilizer(farm pldb.Farm) float64 {
	if farm.OF_type == "ТОУ" {

		if farm.Specialization == "КРС" {
			MN := float64(farm.HeadsOfCows) * K1 * 0.365
			MV := MN * ((VN - VKS) / (VKS - VP))
			return ((MN * K7 / 100) + MV*K8/100) * (100 - K9) / 100
		}

		if farm.Specialization == "Птицеводство (куры)" {
			MP := float64(farm.HeadsOfAnimals) * K3 * 0.365
			MV := MP / 2
			return ((MP * K13 / 100) + MV*T*K8/100) * (100 - K9) / 100
		}
		if farm.Specialization == "Птицеводство (перепела)" {
			MP := float64(farm.HeadsOfAnimals) * K4 * 0.365
			MV := MP / 2
			return ((MP * K14 / 100) + MV*T*K8/100) * (100 - K9) / 100
		}

	} else {
		if farm.Specialization == "Свиноводство" {
			MN := float64(farm.HeadsOfAnimals) * K2 * 0.365
			return (MN * K12 / 100) * (100 - K11) / 100
		}
		if farm.Specialization == "КРС" {
			MN := float64(farm.HeadsOfCows) * K1 * 0.365
			return (MN * K10 / 100) * (100 - K11) / 100
		}
	}
	return 0
}

func PhosphorMassInFertilizer(farm pldb.Farm) float64 {
	if farm.OF_type == "ТОУ" {

		if farm.Specialization == "КРС" {
			MN := float64(farm.HeadsOfCows) * K1 * 0.365
			MV := MN * ((VN - VKS) / (VKS - VP))
			return (MN * K19 / 100)
		}

		if farm.Specialization == "Птицеводство (куры)" {
			MP := float64(farm.HeadsOfAnimals) * K3 * 0.365
			MV := MP / 2
			return ((MP * K13 / 100) + MV*T*K8/100) * (100 - K9) / 100
		}
		if farm.Specialization == "Птицеводство (перепела)" {
			MP := float64(farm.HeadsOfAnimals) * K4 * 0.365
			MV := MP / 2
			return ((MP * K14 / 100) + MV*T*K8/100) * (100 - K9) / 100
		}

	} else {
		if farm.Specialization == "Свиноводство" {
			MN := float64(farm.HeadsOfAnimals) * K2 * 0.365
			return (MN * K12 / 100) * (100 - K11) / 100
		}
		if farm.Specialization == "КРС" {
			MN := float64(farm.HeadsOfCows) * K1 * 0.365
			return (MN * K10 / 100) * (100 - K11) / 100
		}
	}
	return 0
}
