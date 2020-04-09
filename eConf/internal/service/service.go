package service

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/fuwensun/goms/pkg/conf"
)

// Service.
type Service struct {
	cfg config
}

// Config.
type config struct {
	Name    string `yaml:"name,omitempty"`
	Version string `yaml:"version,omitempty"`
}

// getConfig.
func getConfig(cfgpath string) (config, error) {
	var cfg config
	filep := filepath.Join(cfgpath, "app.yml")
	if err := conf.GetConf(filep, &cfg); err != nil {
		log.Printf("get config file: %v", err)
		err = fmt.Errorf("get config: %w", err)
		return cfg, err
	}
	log.Printf("config name: %v,version: %v", cfg.Name, cfg.Version)
	return cfg, nil
}

// New new a service.
func New(cfgpath string) (s *Service) {
	cfg, err := getConfig(cfgpath)
	if err != nil {
		log.Panic(err)
	}
	s = &Service{cfg: cfg}
	return
}

// Close close the resource.
func (s *Service) Close() {
}
