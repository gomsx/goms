package dao

import (
	"database/sql"
	"os"
	"path/filepath"

	e "github.com/aivuca/goms/eApi/internal/pkg/err"
	"github.com/aivuca/goms/pkg/conf"

	_ "github.com/go-sql-driver/mysql" // for init
	"github.com/rs/zerolog/log"
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
		log.Warn().Msgf("get db config file error: %v", err)
	} else if cfg.DSN == "" {
		log.Warn().Msgf("get db config file succ, but DSN IS EMPTY")
	} else {
		log.Info().Msgf("get db config file succ, DSN: %v", "***")
		return cfg, nil
	}
	// env
	if dsn, exist := os.LookupEnv("MYSQL_SVC_DSN"); exist == false {
		log.Warn().Msgf("get db config env error: %v", e.ErrNotFoundData)
	} else if dsn == "" {
		log.Warn().Msgf("get db config env succ, but DSN IS EMPTY")
	} else {
		log.Info().Msgf("get db config env succ, DSN: %v", "***")
		cfg.DSN = dsn
		return cfg, nil
	}
	return nil, e.ErrNotFoundData
}

// newDB new database and return.
func newDB(cfgpath string) (*sql.DB, func(), error) {
	if df, err := getDBConfig(cfgpath); err != nil {
		log.Error().Msgf("get db config error: %v", err)
		return nil, nil, err
	} else if db, err := sql.Open("mysql", df.DSN); err != nil {
		log.Error().Msgf("open db error: %v", err)
		return nil, nil, err
	} else if err := db.Ping(); err != nil {
		log.Error().Msgf("ping db error: %v", err)
		return nil, nil, err
	} else {
		return db, func() {
			db.Close()
		}, nil
	}
}
