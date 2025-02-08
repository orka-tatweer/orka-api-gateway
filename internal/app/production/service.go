package production

import "github.com/lai0xn/orka/pkg/client"

type ProductionScheduler struct{}

func NewProductionScheduler() *ProductionScheduler {
	return &ProductionScheduler{}
}

func (p *ProductionScheduler) GeneratePlan(data client.ProductionPlanRequest) (*client.ProductionResponse, error) {
	apiClient := client.NewApiClient()
	response, err := apiClient.ProductionPlan(data)
	if err != nil {
		return nil, err
	}
	return response, nil

}
