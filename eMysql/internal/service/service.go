package service

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	"github.com/fuwensun/goms/eMysql/internal/dao"
	"github.com/fuwensun/goms/pkg/conf"
)

type Service struct {
	cfg config
	dao dao.Dao
}

// Service conf
type config struct {
	Name    string `yaml:"name,omitempty"`
	Version string `yaml:"version,omitempty"`
}

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

// New new a service and return.
func New(cfgpath string) (s *Service) {
	cfg, err := getConfig(cfgpath)
	if err != nil {
		log.Panic(err)
	}
	dao := dao.New(cfgpath)
	s = &Service{cfg: cfg, dao: dao}
	return
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context) (err error) {
	return s.dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
	s.dao.Close()
}
