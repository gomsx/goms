package service

import (
	"log"
	"path/filepath"

	"github.com/fuwensun/goms/pkg/conf"
)

// Service service.
type Service struct {
	Cfgpath string
}

// Service conf
type ServiceConfig struct {
	Cfgversion string `yaml:"cfgversion"`
}

var (
	sc      ServiceConfig
	cfgfile = "app.yml"
)

// New new a service and return.
func New(cfgpath string) (s *Service) {

	pathname := filepath.Join(cfgpath, cfgfile)
	if err := conf.GetConf(pathname, &sc); err != nil {
		log.Fatalf("failed to get the service config file!: %v", err)
	}
	log.Printf("service config version: %v\n", sc.Cfgversion)

	s = &Service{}
	s.Cfgpath = cfgpath
	return
}

// Close close the resource.
func (s *Service) Close() {
}
