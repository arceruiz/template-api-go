package dummyIntegration

import (
	"github.com/arceruiz/template-api-go/internal/service"
)

type dummyIntegrationService struct{}

func New() service.DummyService {
	return &dummyIntegrationService{}
}

func (c *dummyIntegrationService) DummyAction(requestPayload map[string]interface{}) (*map[string]interface{}, error) {
	dummyResult := map[string]interface{}{
		"foo1": "bar1",
		"foo2": struct {
			name string
			age  int
		}{"John Doe", 44},
		"foo3": true,
	}
	return &dummyResult, nil
}
