package service

import (
	"log"
	"path/filepath"

	"github.com/fuwensun/goms/eConf/internal/pkg/conf"
)

// Service service.
type Service struct {
	Confpath string
}

// Service conf
type ServiceConfig struct {
	Confversion string `yaml:"confversion"`
}

var (
	sc       ServiceConfig
	conffile = "app.yml"
)

// New new a service and return.
func New(confpath string) (s *Service) {

	pathname := filepath.Join(confpath, conffile)
	if err := conf.GetConf(pathname, &sc); err != nil {
		log.Fatalf("failed to get service config file ! err: %v", err)
	}
	log.Printf("service config version: %v\n", sc.Confversion)

	s = &Service{}
	s.Confpath = confpath
	return
}

// Close close the resource.
func (s *Service) Close() {
}
