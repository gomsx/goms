package dao

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fuwensun/goms/eMysql/internal/model"
	"github.com/fuwensun/goms/pkg/conf"
	_ "github.com/go-sql-driver/mysql"
)

// Dao dao interface
type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)
	//ping
	UpdatePingCount(c context.Context, t model.PingType, v model.PingCount) error
	ReadPingCount(c context.Context, t model.PingType) (model.PingCount, error)
}

// dao dao.
type dao struct {
	db *sql.DB
}

// dbcfg
type dbcfg struct {
	DSN string `yaml:"dsn"`
}

// getDBConfig
func getDBConfig(cfgpath string) (dbcfg, error) {
	var cfg dbcfg
	path := filepath.Join(cfgpath, "mysql.yml")
	if err := conf.GetConf(path, &cfg); err != nil {
		log.Printf("get db config file: %v", err)
	}
	if cfg.DSN != "" {
		log.Printf("get config db DSN: %v", cfg.DSN)
		return cfg, nil
	}
	if dsn := os.Getenv("MYSQL_SVC_DSN"); dsn != "" {
		cfg.DSN = dsn
		log.Printf("get env db DSN: %v", cfg.DSN)
		return cfg, nil
	}
	err := fmt.Errorf("get db DSN: %w", model.ErrNotFoundData)
	return cfg, err
}

// New new a dao.
func New(cfgpath string) Dao {
	//db
	df, err := getDBConfig(cfgpath)
	if err != nil {
		log.Panic(err) //?
	}
	mdb, err := sql.Open("mysql", df.DSN)
	if err != nil {
		log.Panicf("open db: %v", err)
	}
	if err := mdb.Ping(); err != nil {
		log.Panicf("ping db: %v", err)
	}
	log.Printf("ping db err=%v", err)
	return &dao{
		db: mdb,
	}
}

// Close close the resource.
func (d *dao) Close() {
	d.db.Close()
}

// Ping ping the resource.
func (d *dao) Ping(c context.Context) (err error) {
	return d.db.PingContext(c)
}
