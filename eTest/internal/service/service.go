package service

import (
	"context"
	"fmt"

	"github.com/gomsx/goms/eTest/internal/dao"
	m "github.com/gomsx/goms/eTest/internal/model"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

// Svc service interface.
type Svc interface {
	HandPing(ctx context.Context, p *m.Ping) (*m.Ping, error)

	CreateUser(ctx context.Context, user *m.User) error
	ReadUser(ctx context.Context, uid int64) (*m.User, error)
	UpdateUser(ctx context.Context, user *m.User) error
	DeleteUser(ctx context.Context, uid int64) error

	Ping(ctx context.Context) (err error)
	Close()
}

// Service service struct.
type service struct {
	cfg *config
	dao dao.Dao
}

// Service config of service.
type config struct {
	Name    string
	Version string
}

// getConfig get config from config file.
func getConfig() (*config, error) {
	cfg := &config{}
	if err := viper.UnmarshalKey("service", cfg); err != nil {
		err = fmt.Errorf("get config file: %w", err)
		return nil, err
	}
	return cfg, nil
}

// New new service and return.
func New(dao dao.Dao) (Svc, func(), error) {
	cfg, err := getConfig()
	if err != nil {
		log.Errorf("get config error: %v", err)
		return nil, nil, err
	}
	log.Infof("service config version: %v", cfg.Version)
	svc := &service{cfg: cfg, dao: dao}
	log.Info("service ok")
	return svc, svc.Close, nil
}

// Ping ping the resource.
func (s *service) Ping(ctx context.Context) (err error) {
	return s.dao.Ping(ctx)
}

// Close close the resource.
func (s *service) Close() {
}
