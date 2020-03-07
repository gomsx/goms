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
	Confpath string
	dao      dao.Dao
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
		log.Fatalf("failed to get service config file!: %v", err)
	}
	log.Printf("service config version: %v", sc.Confversion)

	s = &Service{}
	s.Confpath = confpath
	s.dao = dao.New(confpath)
	log.Printf("dao new! addr: %v", &s.dao)

	rand.Seed(time.Now().UnixNano())
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
