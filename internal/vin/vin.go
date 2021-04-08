package vin

import "fmt"

type VIN string

func NewVIN(code string) (VIN, error) {

	if len(code) != 17 {
		return "", fmt.Errorf("invalid VIN %s: more or less than 17 characters", code)
	}

	return VIN(code), nil
}

func (vin VIN) Manufacturer() string {

	manufacturer := vin[:3]
	// if the last digit of the manufacturer ID is a 9
	// the digits 12 to 14 are the second part of the ID
	if manufacturer[2] == '9' {
		manufacturer += vin[11:14]
	}

	return string(manufacturer)
}
