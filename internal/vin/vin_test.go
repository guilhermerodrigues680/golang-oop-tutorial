package vin_test

import (
	"ooptutorial/internal/vin"
	"testing"
)

const (
	validVIN   = "W0L000051T2123456"
	invalidVIN = "W0"
)

func TestVIN_New(t *testing.T) {

	_, err := vin.NewVIN(validVIN)
	if err != nil {
		t.Errorf("creating valid VIN returned an error: %s", err.Error())
	}
	t.Logf("ok! não retornou erro")

	_, err = vin.NewVIN(invalidVIN)
	if err == nil {
		t.Error("creating invalid VIN did not return an error")
	}
	t.Logf("ok! retornou o erro: %s", err)
}

func TestVIN_Manufacturer(t *testing.T) {

	testVIN, _ := vin.NewVIN(validVIN)
	manufacturer := testVIN.Manufacturer()
	if manufacturer != "W0L" {
		t.Errorf("unexpected manufacturer %s for VIN %s", manufacturer, testVIN)
	}
}
