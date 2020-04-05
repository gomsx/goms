package service

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	"github.com/fuwensun/goms/eTest/internal/dao"
	"github.com/fuwensun/goms/eTest/internal/model"
	"github.com/fuwensun/goms/pkg/conf"
)

type Svc interface {
	UpdateHttpPingCount(c context.Context, pingcount model.PingCount) error
	ReadHttpPingCount(c context.Context) (model.PingCount, error)
	UpdateGrpcPingCount(c context.Context, pingcount model.PingCount) error
	ReadGrpcPingCount(c context.Context) (model.PingCount, error)
	CreateUser(c context.Context, user *model.User) error
	UpdateUser(c context.Context, user *model.User) error
	ReadUser(c context.Context, uid int64) (model.User, error)
	DeleteUser(c context.Context, uid int64) error

	Ping(ctx context.Context) (err error)
	Close()
}

// Service service.
type service struct {
	cfg svccfg
	dao dao.Dao
}

// Service conf
type svccfg struct {
	Name    string `yaml:"name,omitempty"`
	Version string `yaml:"version,omitempty"`
}

func getSvcConfig(cfgpath string) (svccfg, error) {
	var sc svccfg
	path := filepath.Join(cfgpath, "app.yml")
	if err := conf.GetConf(path, &sc); err != nil {
		log.Printf("get config file: %v", err)
		err = fmt.Errorf("get config: %w", err)
		return sc, err
	}
	log.Printf("config name: %v,version: %v", sc.Name, sc.Version)
	return sc, nil
}

// New new a service and return.
func New(cfgpath string, d dao.Dao) (Svc, func(), error) {
	sc, err := getSvcConfig(cfgpath)
	if err != nil {
		return &service{}, nil, err
	}
	s := &service{cfg: sc, dao: d}
	initUidGenerator()
	return s, s.Close, nil
}

// Ping ping the resource.
func (s *service) Ping(c context.Context) (err error) {
	return s.dao.Ping(c)
}

// Close close the resource.
//<<**haha**谁 new ,谁 clean. dao 不是 svc new 的,这里不应该 close.>>
func (s *service) Close() {
	// s.dao.Close()
}
