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
		if farm.Specialization == "Птицеводство (перепела)" {
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

// Масса общего азота в полученном органическом удобрении, т/год
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

// Масса общего фосфора в полученном органическом удобрении, т/год
func PhosphorMassInFertilizer(farm pldb.Farm) float64 {
	if farm.OF_type == "ТОУ" {

		if farm.Specialization == "КРС" {
			MN := float64(farm.HeadsOfCows) * K1 * 0.365
			MV := MN * ((VN - VKS) / (VKS - VP))
			return ((MN * K19 / 100) + MV*K20/100) * (100 - K21) / 100
		}

		if farm.Specialization == "Птицеводство (куры)" {
			MP := float64(farm.HeadsOfAnimals) * K3 * 0.365
			MV := MP / 2
			return ((MP * K25 / 100) + MV*T*K20/100) * (100 - K21) / 100
		}
		if farm.Specialization == "Птицеводство (перепела)" {
			MP := float64(farm.HeadsOfAnimals) * K4 * 0.365
			MV := MP / 2
			return ((MP * K26 / 100) + MV*T*K20/100) * (100 - K9) / 100
		}

	} else {
		if farm.Specialization == "Свиноводство" {
			MN := float64(farm.HeadsOfAnimals) * K2 * 0.365
			return (MN * K24 / 100) * (100 - K23) / 100
		}
		if farm.Specialization == "КРС" {
			MN := float64(farm.HeadsOfCows) * K1 * 0.365
			return (MN * K22 / 100) * (100 - K23) / 100
		}
	}
	return 0
}

// Масса общего азота, которую готовы принять земельные угодья самого предприятия», т/год
func NitrogenMassForSoil(farm pldb.Farm) float64 {
	if farm.OF_type == "ТОУ" {
		return farm.SAL * K15 / 1000
	} else {
		return farm.SAL * K16 / 1000
	}
}

// Масса общего фосфора, которую готовы принять земельные угодья самого предприятия», т/год
func PhosphorMassForSoil(farm pldb.Farm) float64 {
	return farm.SAL * K27 / 1000
}

// Потенциал внесения органического удобрения по азоту, т/год
func FertilizerPotentialByNitrogen(farm pldb.Farm) float64 {
	var PV float64
	if farm.Specialization == "Растениеводство" {
		PV = NitrogenMassForSoil(farm)
	} else {
		PV = NitrogenMassForSoil(farm) - NitrogenMassInFertilizer(farm)
	}
	if PV > 0 {
		if farm.OF_type == "ТОУ" {
			if farm.Specialization == "КРС" {
				return (PV * 1000) / (K28 * 10)
			} else if (farm.Specialization == "Птицеводство (куры)") || (farm.Specialization == "Птицеводство (перепела)") {
				return (PV * 1000) / (K31 * 10)
			}
		} else {
			if farm.Specialization == "КРС" {
				return (PV * 1000) / (K29 * 10)
			} else if farm.Specialization == "Свиноводство" {
				return (PV * 1000) / (K30 * 10)
			}
		}
	}
	return 0
}

// Потенциал внесения органического удобрения по фосфору, т/год
func FertilizerPotentialByPhosphor(farm pldb.Farm) float64 {
	var PV float64
	if farm.Specialization == "Растениеводство" {
		PV = PhosphorMassForSoil(farm)
	} else {
		PV = PhosphorMassForSoil(farm) - PhosphorMassInFertilizer(farm)
	}

	if PV > 0 {
		if farm.OF_type == "ТОУ" {
			if farm.Specialization == "КРС" {
				return (PV * 1000) / (K32 * 10)
			} else if (farm.Specialization == "Птицеводство (куры)") || (farm.Specialization == "Птицеводство (перепела)") {
				return (PV * 1000) / (K35 * 10)
			}
		} else {
			if farm.Specialization == "КРС" {
				return (PV * 1000) / (K33 * 10)
			} else if farm.Specialization == "Свиноводство" {
				return (PV * 1000) / (K34 * 10)
			}
		}
	}
	return 0
}

// Необходимая дополнительная площадь земельных угодий для внесения ОУ, га
func SquareDemandForNitrogen(farm pldb.Farm) float64 {
	var PV float64
	if farm.Specialization == "Растениеводство" {
		PV = NitrogenMassForSoil(farm)
	} else {
		PV = NitrogenMassForSoil(farm) - NitrogenMassInFertilizer(farm)
	}

	if PV < 0 {
		if farm.OF_type == "ТОУ" {
			return (NitrogenMassInFertilizer(farm) - NitrogenMassForSoil(farm)) * (1 / K15) * 1000
		}
	} else {
		return (NitrogenMassInFertilizer(farm) - NitrogenMassForSoil(farm)) * (1 / K16) * 1000
	}

	return 0
}

func SquareDemandForPhosphor(farm pldb.Farm) float64 {
	var PV float64
	if farm.Specialization == "Растениеводство" {
		PV = PhosphorMassForSoil(farm)
	} else {
		PV = PhosphorMassForSoil(farm) - PhosphorMassInFertilizer(farm)
	}

	if PV < 0 {
		return (PhosphorMassInFertilizer(farm) - PhosphorMassForSoil(farm)) * (1 / K27) * 1000
	}

	return 0
}

// Свободная площадь земельных угодий, на которую можно внести органическое удобрение с других предприятий, га
func SquareFreeForNitrogen(farm pldb.Farm) float64 {
	var PV float64
	if farm.Specialization == "Растениеводство" {
		PV = PhosphorMassForSoil(farm)
	} else {
		PV = PhosphorMassForSoil(farm) - PhosphorMassInFertilizer(farm)
	}

	if PV > 0 {
		if farm.OF_type == "ТОУ" {
			return (NitrogenMassForSoil(farm) - NitrogenMassInFertilizer(farm)) * (1 / K15) * 1000
		}
	} else {
		return (NitrogenMassForSoil(farm) - NitrogenMassInFertilizer(farm)) * (1 / K16) * 1000
	}

	return 0
}

func SquareFreeForPhosphor(farm pldb.Farm) float64 {
	var PV float64
	if farm.Specialization == "Растениеводство" {
		PV = PhosphorMassForSoil(farm)
	} else {
		PV = PhosphorMassForSoil(farm) - PhosphorMassInFertilizer(farm)
	}

	if PV > 0 {
		return (PhosphorMassForSoil(farm) - PhosphorMassInFertilizer(farm)) * (1 / K27) * 1000
	} else {
		return 0
	}
}

// Необходимая вместимость навозохранилищ для переработки навоза/помета в органическое удобрение, т
func DemandForOFStorage(farm pldb.Farm) float64 {
	if farm.OF_type == "ТОУ" {
		return ManureMass(farm) * K18
	} else {
		return ManureMass(farm) * K17
	}
}
