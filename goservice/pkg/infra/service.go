package infra

type Service interface {
	Liveness() (LivenessResponse, error)
	Readiness() (ReadinessResponse, error)
}

type infraService struct{}

func NewInfraService() Service {
	return &infraService{}
}

func (s *infraService) Liveness() (LivenessResponse, error) {
	return LivenessResponse{
		Message: "ok",
	}, nil
}

func (s *infraService) Readiness() (ReadinessResponse, error) {
	return ReadinessResponse{
		Message: "ok",
	}, nil
}
