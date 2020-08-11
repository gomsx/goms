package dao

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	m "github.com/aivuca/goms/eApi/internal/model"
	e "github.com/aivuca/goms/eApi/internal/pkg/err"
	lg "github.com/aivuca/goms/eApi/internal/pkg/log"
	"github.com/aivuca/goms/pkg/conf"

	_ "github.com/go-sql-driver/mysql" // for init()
	"github.com/gomodule/redigo/redis"
)

// Dao dao interface.
type Dao interface {
	Close()

	Ping(ctx context.Context) (err error)
	//count
	ReadPing(c context.Context, t string) (*m.Ping, error)
	UpdatePing(c context.Context, p *m.Ping) error
	//user
	CreateUser(c context.Context, user *m.User) error
	ReadUser(c context.Context, uid int64) (*m.User, error)
	UpdateUser(c context.Context, user *m.User) error
	DeleteUser(c context.Context, uid int64) error
}

// dao dao.
type dao struct {
	db    *sql.DB
	redis redis.Conn
}

// dbcfg db config.
type dbcfg struct {
	DSN string `yaml:"dsn"`
}

// cccfg cache config.
type cccfg struct {
	Addr string `yaml:"addr"`
	Pass string `yaml:"pass"`
}

// Log.
var log = lg.Lgd

// getDBConfig get db config from file and env.
func getDBConfig(cfgpath string) (*dbcfg, error) {
	var err error
	cfg := &dbcfg{}

	path := filepath.Join(cfgpath, "mysql.yaml")
	if err = conf.GetConf(path, &cfg); err != nil { //file
		log.Warn().Msgf("get db config file, error: %v", err)
	} else if cfg.DSN != "" {
		log.Info().Msgf("get db config file, DSN: ***")
		return cfg, nil
	} else if cfg.DSN = os.Getenv("MYSQL_SVC_DSN"); cfg.DSN == "" { //env
		log.Warn().Msg("get db config env, empty")
	} else {
		log.Info().Msgf("get db config env, DSN: ***")
		return cfg, nil
	}
	err = fmt.Errorf("get file and env: %w", e.ErrNotFoundData)
	return nil, err
}

// getCCConfig get cache config from file and env.
func getCCConfig(cfgpath string) (*cccfg, error) {
	var err error
	cfg := &cccfg{}

	path := filepath.Join(cfgpath, "redis.yaml")
	if err = conf.GetConf(path, &cfg); err != nil { //file
		log.Warn().Msgf("get cc config file, error: %v", err)
	} else if cfg.Addr != "" {
		log.Info().Msgf("get cc config file, Addr: %v", cfg.Addr)
		return cfg, nil
	} else if cfg.Addr = os.Getenv("REDIS_SVC_ADDR"); cfg.Addr == "" { //env
		log.Warn().Msgf("get cc config env, empty")
	} else {
		log.Info().Msgf("get cc config env, Addr: %v", cfg.Addr)
		return cfg, nil
	}
	err = fmt.Errorf("get file and env: %w", e.ErrNotFoundData)
	return nil, err
}

// New new a Dao.
func New(cfgpath string) (Dao, func(), error) {
	return new(cfgpath)
}

// New new a dao.
func new(cfgpath string) (*dao, func(), error) {
	mdao := &dao{}
	//cc
	if cf, err := getCCConfig(cfgpath); err != nil {
		log.Error().Msgf("get cc config, error: %v", err)
		return nil, nil, err
	} else if mcc, err := redis.Dial("tcp", cf.Addr, redis.DialPassword(cf.Pass)); err != nil {
		log.Error().Msgf("dial cc, error: %v", err)
		return nil, nil, err
	} else if _, err = mcc.Do("PING"); err != nil {
		log.Error().Msgf("ping cc, error: %v", err)
		return nil, nil, err
	} else {
		mdao.redis = mcc
		log.Info().Msg("cc ok")
	}
	//db
	if df, err := getDBConfig(cfgpath); err != nil {
		log.Error().Msgf("get db config, error: %v", err)
		return nil, nil, err
	} else if mdb, err := sql.Open("mysql", df.DSN); err != nil {
		log.Error().Msgf("open db, error: %v", err)
		return nil, nil, err
	} else if err := mdb.Ping(); err != nil {
		log.Error().Msgf("ping db, error: %v", err)
		return nil, nil, err
	} else {
		mdao.db = mdb
		log.Info().Msg("db ok")
	}
	//return
	log.Info().Msg("dao ok")
	return mdao, mdao.Close, nil
}

// Close close the resource.
func (d *dao) Close() {
	d.redis.Close()
	d.db.Close()
}

// Ping ping the resource.
func (d *dao) Ping(ctx context.Context) (err error) {
	if _, err = d.redis.Do("PING"); err != nil {
		return
	}
	return d.db.PingContext(ctx)
}
