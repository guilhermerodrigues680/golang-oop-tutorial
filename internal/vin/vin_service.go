package vin

type VINService struct {
	client VINAPIClient
}

type VINServiceConfig struct{}

func NewVINService(config *VINServiceConfig, apiClient VINAPIClient) *VINService {
	return &VINService{client: apiClient}
}

func (s *VINService) CreateFromCode(code string) (VIN, error) {

	if s.client.IsEuropean(code) {
		return NewVINEU(code)
	}

	return NewVin(code)
}
