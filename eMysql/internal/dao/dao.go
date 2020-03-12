package dao

import (
	"context"
	"database/sql"
	"log"
	"path/filepath"

	"github.com/fuwensun/goms/eMysql/internal/model"
	"github.com/fuwensun/goms/pkg/conf"
	_ "github.com/go-sql-driver/mysql"
)

// DBConfig mysql config.
type DBConfig struct {
	DSN string `yaml:"dsn"`
}

var (
	conffile = "mysql.yml"
	DSN      = "user:password@/dbname"
)

// Dao dao interface
type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)
	//call
	UpdatePingCount(c context.Context, t model.PingType, v model.PingCount) error
	ReadPingCount(c context.Context, t model.PingType) (model.PingCount, error)
}

// dao dao.
type dao struct {
	db *sql.DB
}

// New new a dao.
func New(confpath string) Dao {

	var dc DBConfig
	pathname := filepath.Join(confpath, conffile)
	if err := conf.GetConf(pathname, &dc); err != nil {
		log.Printf("get db config file: %v", err)
	}

	if dc.DSN != "" {
		DSN = dc.DSN
	}
	log.Printf("db DSN: %v", DSN)

	mdb, err := sql.Open("mysql", DSN)
	if err != nil {
		log.Panicf("failed to open db: %v", err)
	}
	if err := mdb.Ping(); err != nil {
		log.Panicf("failed to ping db: %v", err)
	}
	return &dao{
		db: mdb,
	}
}

// Close close the resource.
func (d *dao) Close() {
	d.db.Close()
}

// Ping ping the resource.
func (d *dao) Ping(ctx context.Context) (err error) {
	return d.db.PingContext(ctx)
}
