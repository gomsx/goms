package service

import (
	"context"
	"log"
	"math/rand"
	"path/filepath"
	"time"

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
func New(cfgpath string, d dao.Dao) (Svc, func(), error) {

	path := filepath.Join(cfgpath, cfgfile)
	if err := conf.GetConf(path, &sc); err != nil {
		log.Printf("get config file: %v", err)
		// return nil, nil, err
	}
	log.Printf("config version: %v", sc.Version)

	s := &service{dao: d}

	rand.Seed(time.Now().UnixNano())
	return s, s.Close, nil
}

// Ping ping the resource.
func (s *service) Ping(ctx context.Context) (err error) {
	return s.dao.Ping(ctx)
}

// Close close the resource.
//<<**haha**谁 new ,谁 clean. dao 不是 svc new 的,这里不应该 close.>>
func (s *service) Close() {
	// s.dao.Close()
}
