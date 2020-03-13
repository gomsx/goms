package service

import (
	"context"
	"log"
	"math/rand"
	"path/filepath"
	"time"

	"github.com/fuwensun/goms/eRedis/internal/dao"
	"github.com/fuwensun/goms/pkg/conf"
)

// Service service.
type Service struct {
	Cfgpath string
	dao     dao.Dao
}

// Service conf
type ServiceConfig struct {
	Cfgversion string `yaml:"cfgversion"`
}

var (
	sc       ServiceConfig
	cfgfile = "app.yml"
)

// New new a service and return.
func New(cfgpath string) (s *Service) {

	pathname := filepath.Join(cfgpath, cfgfile)
	if err := conf.GetConf(pathname, &sc); err != nil {
		log.Fatalf("failed to get service config file!: %v", err)
	}
	log.Printf("service config version: %v", sc.Cfgversion)

	s = &Service{}
	s.Cfgpath = cfgpath
	s.dao = dao.New(cfgpath)
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
