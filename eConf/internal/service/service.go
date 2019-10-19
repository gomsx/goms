package service

import (
	"log"

	"github.com/fuwensun/goms/eConf/internal/pkg/conf"
)

// Service service.
type Service struct {
}

// Service conf
type ServiceConfig struct {
	confversion string `yaml:"confversion"`
}

var (
	sc ServiceConfig
)

// New new a service and return.
func New() (s *Service) {
	if err := conf.GetConf("../configs/app.yml", &sc); err != nil {
		panic(err)
	}
	log.Printf("ServiceConfig: %+v , confversion: %+v\n", sc, sc.confversion)
	s = &Service{}
	return
}

// Close close the resource.
func (s *Service) Close() {
}
