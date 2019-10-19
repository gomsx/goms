package service

import (
	"log"
	"path/filepath"

	"github.com/fuwensun/goms/eConf/internal/pkg/conf"
)

// Service service.
type Service struct {
}

// Service conf
type ServiceConfig struct {
	Confversion string `yaml:"confversion"`
}

var (
	sc      ServiceConfig
	confile = "app.yml"
)

// New new a service and return.
func New(confpath string) (s *Service) {
	pathname := filepath.Join(confpath, confile)
	if err := conf.GetConf(pathname, &sc); err != nil {
		panic(err)
	}
	log.Printf("ServiceConfig: %+v , confversion: %+v\n", sc, sc.Confversion)

	s = &Service{}
	return
}

// Close close the resource.
func (s *Service) Close() {
}
