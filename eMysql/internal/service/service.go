package service

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	"github.com/fuwensun/goms/eMysql/internal/dao"
	"github.com/fuwensun/goms/pkg/conf"
)

// config config of service.
type config struct {
	Name    string `yaml:"name,omitempty"`
	Version string `yaml:"version,omitempty"`
}

// Service service struct.
type Service struct {
	cfg *config
	dao *dao.Dao
}

// getConfig get config from config file.
func getConfig(cfgpath string) (*config, error) {
	cfg := &config{}
	filep := filepath.Join(cfgpath, "app.yaml")
	if err := conf.GetConf(filep, cfg); err != nil {
		err = fmt.Errorf("get config file: %w", err)
		return nil, err
	}
	return cfg, nil
}

// New new service and return.
func New(cfgpath string, d *dao.Dao) *Service {
	cfg, err := getConfig(cfgpath)
	if err != nil {
		log.Panicf("failed to get config file: %v", err)
	}
	return &Service{
		cfg: cfg,
		dao: d,
	}
}

// Ping ping service.
func (s *Service) Ping(c context.Context) (err error) {
	return s.dao.Ping(c)
}

// Close close service.
func (s *Service) Close() {
	s.dao.Close()
}
