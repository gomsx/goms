package service

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/fuwensun/goms/eApi/internal/dao"
	m "github.com/fuwensun/goms/eApi/internal/model"
	"github.com/fuwensun/goms/pkg/conf"

	log "github.com/sirupsen/logrus"
)

// Svc service interface.
type Svc interface {
	Close()
	Ping(c context.Context) (err error)
	// ping
	HandPing(c context.Context, p *m.Ping) (*m.Ping, error)
	// user
	CreateUser(c context.Context, user *m.User) error
	ReadUser(c context.Context, uid int64) (*m.User, error)
	UpdateUser(c context.Context, user *m.User) error
	DeleteUser(c context.Context, uid int64) error
}

// Service service struct.
type service struct {
	cfg *config
	dao dao.Dao
}

// Service config of service.
type config struct {
	Name    string `yaml:"name,omitempty"`
	Version string `yaml:"version,omitempty"`
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
func New(cfgpath string, dao dao.Dao) (Svc, func(), error) {
	cfg, err := getConfig(cfgpath)
	if err != nil {
		log.Errorf("get config error: %v", err)
		return nil, nil, err
	}
	log.Infof("service config version: %v", cfg.Version)
	svc := &service{cfg: cfg, dao: dao}
	return svc, svc.Close, nil
}

// Ping ping the resource.
func (s *service) Ping(c context.Context) (err error) {
	return s.dao.Ping(c)
}

// Close close the resource.
func (s *service) Close() {
}
