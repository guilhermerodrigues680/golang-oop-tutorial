package vin

type VINService struct {
	client *VINAPIClient
}

type VINServiceConfig struct {
	APIURL string
	APIKey string
	// more configuration values
}

func NewVINService(config *VINServiceConfig) *VINService {

	// use config to create the API client
	apiClient := NewVINAPIClient(config.APIURL, config.APIKey)

	return &VINService{apiClient}
}

func (s *VINService) CreateFromCode(code string) (VIN, error) {

	if s.client.IsEuropean(code) {
		return NewVINEU(code)
	}

	return NewVin(code)
}
