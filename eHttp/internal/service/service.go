package service

// Service service.
type Service struct {
}

// New new a service.
func New() (s *Service) {
	s = &Service{}
	return
}

// Close close the resource
func (s *Service) Close() {
}
