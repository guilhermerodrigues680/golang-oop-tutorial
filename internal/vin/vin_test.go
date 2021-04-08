package vin_test

import (
	"ooptutorial/internal/vin"
	"testing"
)

const euSmallVIN = "W09000051T2123456"

type mockAPIClient struct {
	apiCalls int
}

func NewMockAPIClient() *mockAPIClient {
	return &mockAPIClient{}
}

func (client *mockAPIClient) IsEuropean(code string) bool {
	client.apiCalls++
	return true
}

func TestVIN_EU_SmallManufacturer(t *testing.T) {

	apiClient := NewMockAPIClient()
	service := vin.NewVINService(&vin.VINServiceConfig{}, apiClient)
	testVIN, _ := service.CreateFromCode(euSmallVIN)

	manufacturer := testVIN.Manufacturer()
	if manufacturer != "W09123" {
		t.Errorf("unexpected manufacturer %s for VIN %s", manufacturer, testVIN)
	}

	t.Logf("API CALLS: %d", apiClient.apiCalls)
}
