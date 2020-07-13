package dao

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	. "github.com/aivuca/goms/eTest/internal/model"
	"github.com/aivuca/goms/pkg/conf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"github.com/rs/zerolog/log"
)

// Dao dao interface
type Dao interface {
	Close()

	Ping(ctx context.Context) (err error)
	//ping
	ReadPing(c context.Context, t string) (*Ping, error)
	UpdatePing(c context.Context, p *Ping) error
	//user-cc
	ExistUserCC(c context.Context, uid int64) (bool, error)
	SetUserCC(c context.Context, user *User) error
	GetUserCC(c context.Context, uid int64) (*User, error)
	DelUserCC(c context.Context, uid int64) error
	//user-db
	CreateUserDB(c context.Context, user *User) error
	ReadUserDB(c context.Context, uid int64) (*User, error)
	UpdateUserDB(c context.Context, user *User) error
	DeleteUserDB(c context.Context, uid int64) error
	//user
	CreateUser(c context.Context, user *User) error
	ReadUser(c context.Context, uid int64) (*User, error)
	UpdateUser(c context.Context, user *User) error
	DeleteUser(c context.Context, uid int64) error
}

// dao dao.
type dao struct {
	db    *sql.DB
	redis redis.Conn
}

// dbcfg mysql config.
type dbcfg struct {
	DSN string `yaml:"dsn"`
}

//
type cccfg struct {
	Addr string `yaml:"addr"`
	Pass string `yaml:"pass"`
}

func getDBConfig(cfgpath string) (dbcfg, error) {
	var cfg dbcfg
	var err error

	//file
	path := filepath.Join(cfgpath, "mysql.yaml")
	if err = conf.GetConf(path, &cfg); err != nil {
		log.Warn().Msgf("get db config file, %v", err)
	}
	if cfg.DSN != "" {
		log.Info().Msgf("get db config file, DSN: %v", cfg.DSN)
		return cfg, nil
	}

	//env
	dsn := os.Getenv("MYSQL_SVC_DSN")
	if dsn == "" {
		log.Warn().Msg("get db config env, empty")
		err = fmt.Errorf("get env: %w", ErrNotFoundData)
	} else {
		cfg.DSN = dsn
		log.Info().Msgf("get db config env, DSN: %v", cfg.DSN)
		return cfg, nil
	}

	return cfg, err
}
func getCCConfig(cfgpath string) (cccfg, error) {
	var cfg cccfg
	var err error

	//file
	path := filepath.Join(cfgpath, "redis.yaml")
	if err = conf.GetConf(path, &cfg); err != nil {
		log.Warn().Msgf("get cc config file, error")
	}
	if cfg.Addr != "" {
		log.Info().Msgf("get cc config file, Addr: %v", cfg.Addr)
		return cfg, nil
	}

	// env
	addr := os.Getenv("REDIS_SVC_ADDR")
	if addr == "" {
		log.Warn().Msgf("get cc config env, empty")
		err = fmt.Errorf("get env: %w", ErrNotFoundData)
	} else {
		cfg.Addr = addr
		log.Info().Msgf("get cc config env, Addr: %v", cfg.Addr)
		return cfg, nil
	}

	return cfg, err
}

// New new a dao.
func New(cfgpath string) (Dao, func(), error) {
	//cc
	cf, err := getCCConfig(cfgpath)
	if err != nil {
		log.Error().Msg("get cc config, error")
		return nil, nil, err
	}
	mcc, err := redis.Dial("tcp", cf.Addr,
		redis.DialPassword(cf.Pass),
	)
	if err != nil {
		log.Error().Msg("dial cc error")
		return nil, nil, err
	}
	if _, err = mcc.Do("PING"); err != nil {
		log.Error().Msg("ping cc error")
		return nil, nil, err
	}
	log.Info().Msg("cc ok")

	//db
	df, err := getDBConfig(cfgpath)
	if err != nil {
		log.Error().Msg("get db config, error")
		return nil, nil, err
	}
	mdb, err := sql.Open("mysql", df.DSN)
	if err != nil {
		log.Error().Msgf("open db error")
		return nil, nil, err
	}
	if err := mdb.Ping(); err != nil {
		log.Error().Msgf("ping db error")
		return nil, nil, err
	}
	log.Info().Msg("db ok")

	//
	mdao := &dao{
		db:    mdb,
		redis: mcc,
	}
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
