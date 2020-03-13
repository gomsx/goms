package service

import (
	"context"
	"log"
	"math/rand"
	"path/filepath"
	"time"

	"github.com/fuwensun/goms/eTest/internal/dao"
	"github.com/fuwensun/goms/pkg/conf"
)

// Service service.
type Service struct {
	dao dao.Dao
}

// Service conf
type ServiceCfg struct {
	Version string `yaml:"version"`
}

var (
	sc      ServiceCfg
	cfgfile = "app.yml"
)

// New new a service and return.
func New(cfgpath string, dao dao.Dao) (*Service, func(), error) {

	pathname := filepath.Join(cfgpath, cfgfile)
	if err := conf.GetConf(pathname, &sc); err != nil {
		log.Printf("get service config file: %v", err)
		return nil, nil, err
	}
	log.Printf("service config version: %v", sc.Version)

	s := &Service{}
	s.dao = dao

	rand.Seed(time.Now().UnixNano())
	return s, s.Close, nil
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context) (err error) {
	return s.dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
	s.dao.Close()
}
