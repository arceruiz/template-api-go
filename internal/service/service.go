package service

type DummyService interface {
	DummyAction(requestPayload map[string]interface{}) (*map[string]interface{}, error)
}
