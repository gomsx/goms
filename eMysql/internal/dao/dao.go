package dao

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/aivuca/goms/eMysql/internal/model"
	"github.com/aivuca/goms/pkg/conf"
	_ "github.com/go-sql-driver/mysql"
)

// Dao Dao.
type Dao struct {
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

// New new a Dao.
func New(cfgpath string) *Dao {
	//db
	df, err := getDBConfig(cfgpath)
	if err != nil {
		log.Panicf("failed to get config: %v", err)
	}
	mdb, err := sql.Open("mysql", df.DSN)
	if err != nil {
		log.Panicf("open db: %v", err)
	}
	if err := mdb.Ping(); err != nil {
		log.Panicf("ping db: %v", err)
	}
	log.Printf("ping db err=%v", err)
	return &Dao{
		db: mdb,
	}
}

// Close close the resource.
func (d *Dao) Close() {
	d.db.Close()
}

// Ping ping the resource.
func (d *Dao) Ping(c context.Context) (err error) {
	return d.db.PingContext(c)
}
