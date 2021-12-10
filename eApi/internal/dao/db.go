package dao

import (
	"database/sql"
	"os"

	e "github.com/gomsx/goms/pkg/err"
	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql" // for init()
	log "github.com/sirupsen/logrus"
)

// dbcfg config of db.
type dbcfg struct {
	DSN string
}

// getDBConfig get db config from file and env.
func getDBConfig() (*dbcfg, error) {
	cfg := &dbcfg{}
	// file
	if err := viper.UnmarshalKey("data.database", cfg); err != nil {
		log.Warnf("get db config file error: %v", err)
	} else if cfg.DSN == "" {
		log.Warnf("get db config file succeeded, but DSN IS EMPTY")
	}
	// env
	if dsn, exist := os.LookupEnv("MYSQL_SVC_DSN"); exist == false {
		log.Warnf("get db config env error: %v", e.ErrNotFoundData)
	} else if dsn == "" {
		log.Warnf("get db config env succeeded, but DSN IS EMPTY")
	} else {
		log.Infof("get db config env succeeded, DSN: %v", "***")
		cfg.DSN = dsn
	}

	if cfg.DSN == "" {
		return nil, e.ErrNotFoundData
	}
	return cfg, nil
}

// newDB new database and return.
func newDB() (*sql.DB, func(), error) {
	if df, err := getDBConfig(); err != nil {
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
