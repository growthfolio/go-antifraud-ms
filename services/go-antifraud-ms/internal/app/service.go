package app

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Add(a, b int) int {
	return a + b
}
