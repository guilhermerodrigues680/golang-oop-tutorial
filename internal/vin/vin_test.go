package vin_test

import (
	"ooptutorial/internal/vin"
	"testing"
)

const testVIN = "W09000051T2123456"

// const testVIN = "W0"

func TestVIN_Manufacturer_1(t *testing.T) {
	t.Logf("testVIN = %s", testVIN)

	manufacturer := vin.Manufacturer(testVIN)
	t.Logf("manufacturer = %s", manufacturer)

	if manufacturer != "W09123" {
		t.Errorf("unexpected manufacturer %s for VIN %s", manufacturer, testVIN)
	}
}
