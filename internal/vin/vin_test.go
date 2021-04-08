package vin_test

import (
	"ooptutorial/internal/vin"
	"testing"
)

const euSmallVIN = "W09000051T2123456"

func TestVIN_EU_SmallManufacturer_Polymorphism(t *testing.T) {

	var testVINs []vin.VIN
	testVIN, _ := vin.NewVINEU(euSmallVIN)
	// now there is no need to cast!
	testVINs = append(testVINs, testVIN)

	for _, vin := range testVINs {
		manufacturer := vin.Manufacturer()
		if manufacturer != "W09123" {
			t.Errorf("unexpected manufacturer %s for VIN %s", manufacturer, testVIN)
		}
	}
}
