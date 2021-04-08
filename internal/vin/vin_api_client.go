package vin

type VINAPIClient interface {
	IsEuropean(code string) bool
}

type vinAPIClient struct {
	apiURL string
	apiKey string
	// .. internals go here ...
}

func NewVinAPIClient(apiURL, apiKey string) *vinAPIClient {

	return &vinAPIClient{apiURL, apiKey}
}

func (client *vinAPIClient) IsEuropean(code string) bool {

	// calls external API and returns correct value
	return true
}
