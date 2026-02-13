package respondio

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (service *Service) SendMessage() string {
	return "pong"
}
