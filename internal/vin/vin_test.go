package vin_test

import (
	"ooptutorial/internal/vin"
	"testing"
)

const (
	validVIN   = vin.VIN("W0L000051T2123456")
	invalidVIN = vin.VIN("W0")
)

func TestVIN_Manufacturer(t *testing.T) {

	t.Logf("testVIN = %s", validVIN)
	manufacturer := validVIN.Manufacturer()
	t.Logf("manufacturer = %s", manufacturer)

	if manufacturer != "W0L" {
		t.Errorf("unexpected manufacturer %s for VIN %s", manufacturer, validVIN)
	}
}

func TestVIN_Manufacturer_2(t *testing.T) {

	t.Logf("testVIN = %s", invalidVIN)
	manufacturer := invalidVIN.Manufacturer()
	t.Logf("manufacturer = %s", manufacturer)
}
