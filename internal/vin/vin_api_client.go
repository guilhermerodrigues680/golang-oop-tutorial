package vin

type VINAPIClient struct {
	apiURL string
	apiKey string
	// ... internals go here ...
}

func NewVINAPIClient(apiURL, apiKey string) *VINAPIClient {

	return &VINAPIClient{apiURL, apiKey}
}

func (client *VINAPIClient) IsEuropean(code string) bool {

	// calls external API and returns correct value
	return true
}
