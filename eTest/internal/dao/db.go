package dao

import (
	"database/sql"
	"os"
	"path/filepath"

	"github.com/fuwensun/goms/pkg/conf"
	e "github.com/fuwensun/goms/pkg/err"

	_ "github.com/go-sql-driver/mysql" // for init()
	log "github.com/sirupsen/logrus"
)

// dbcfg config of db.
type dbcfg struct {
	DSN string `yaml:"dsn"`
}

// getDBConfig get db config from file and env.
func getDBConfig(cfgpath string) (*dbcfg, error) {
	var err error
	cfg := &dbcfg{}
	// file
	path := filepath.Join(cfgpath, "mysql.yaml")
	if err = conf.GetConf(path, &cfg); err != nil {
		log.Warnf("get db config file error: %v", err)
	} else if cfg.DSN == "" {
		log.Warnf("get db config file succ, but DSN IS EMPTY")
	} else {
		log.Infof("get db config file succ, DSN: %v", "***")
		return cfg, nil
	}
	// env
	if dsn, exist := os.LookupEnv("MYSQL_SVC_DSN"); exist == false {
		log.Warnf("get db config env error: %v", e.ErrNotFoundData)
	} else if dsn == "" {
		log.Warnf("get db config env succ, but DSN IS EMPTY")
	} else {
		log.Infof("get db config env succ, DSN: %v", "***")
		cfg.DSN = dsn
		return cfg, nil
	}
	return nil, e.ErrNotFoundData
}

// newDB new database and return.
func newDB(cfgpath string) (*sql.DB, func(), error) {
	if df, err := getDBConfig(cfgpath); err != nil {
		log.Errorf("get db config error: %v", err)
		return nil, nil, err
	} else if db, err := sql.Open("mysql", df.DSN); err != nil {
		log.Errorf("open db error: %v", err)
		return nil, nil, err
	} else if err := db.Ping(); err != nil {
		log.Errorf("ping db error: %v", err)
		return nil, nil, err
	} else {
		return db, func() {
			db.Close()
		}, nil
	}
}
