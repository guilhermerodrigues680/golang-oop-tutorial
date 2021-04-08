package vin

import "fmt"

type VIN interface {
	Manufacturer() string
}

type vin string
type vinEU vin

func NewVin(code string) (vin, error) {

	if len(code) != 17 {
		return "", fmt.Errorf("invalid VIN %s: more or less than 17 characters", code)
	}

	return vin(code), nil
}

func (v vin) Manufacturer() string {
	return string(v[:3])
}

func NewVINEU(code string) (vinEU, error) {
	v, err := NewVin(code)
	return vinEU(v), err
}

func (v vinEU) Manufacturer() string {

	// call manufacturer on supertype
	manufacturer := vin(v).Manufacturer()

	// if the last digit of the manufacturer ID is a 9
	// the digits 12 to 14 are the second part of the ID
	if manufacturer[2] == '9' {
		manufacturer += string(v[11:14])
	}

	return string(manufacturer)
}
