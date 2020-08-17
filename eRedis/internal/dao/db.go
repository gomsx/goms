package dao

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	e "github.com/fuwensun/goms/eRedis/internal/pkg/err"
	"github.com/fuwensun/goms/pkg/conf"

	_ "github.com/go-sql-driver/mysql" // for init()
)

// dbcfg db config.
type dbcfg struct {
	DSN string `yaml:"dsn"`
}

// getDBConfig get db config from file and env.
func getDBConfig(cfgpath string) (*dbcfg, error) {
	var err error
	cfg := &dbcfg{}

	path := filepath.Join(cfgpath, "mysql.yaml")
	if err = conf.GetConf(path, &cfg); err != nil { //file
		log.Printf("get db config file, error: %v", err)
	} else if cfg.DSN != "" {
		log.Printf("get db config file, DSN: ***")
		return cfg, nil
	} else if cfg.DSN = os.Getenv("MYSQL_SVC_DSN"); cfg.DSN == "" { //env
		log.Printf("get db config env, empty")
	} else {
		log.Printf("get db config env, DSN: ***")
		return cfg, nil
	}
	err = fmt.Errorf("get file and env: %w", e.ErrNotFoundData)
	return nil, err
}

// newDB new a database.
func newDB(cfgpath string) (*sql.DB, func(), error) {
	if df, err := getDBConfig(cfgpath); err != nil {
		log.Printf("get db config, error: %v", err)
		return nil, nil, err
	} else if db, err := sql.Open("mysql", df.DSN); err != nil {
		log.Printf("open db, error: %v", err)
		return nil, nil, err
	} else if err := db.Ping(); err != nil {
		log.Printf("ping db, error: %v", err)
		return nil, nil, err
	} else {
		return db, func() {
			db.Close()
		}, nil
	}
}
