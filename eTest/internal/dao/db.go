package dao

import (
	"database/sql"
	"os"
	"path/filepath"

	e "github.com/aivuca/goms/eTest/internal/pkg/err"
	"github.com/aivuca/goms/pkg/conf"

	_ "github.com/go-sql-driver/mysql" // for init()
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

	path := filepath.Join(cfgpath, "mysql.yaml")
	if err = conf.GetConf(path, &cfg); err != nil { //file
		log.Warn().Msgf("get db config file error: %v", err)
	} else if cfg.DSN != "" {
		log.Info().Msgf("get db config file succ, DSN: ***")
		return cfg, nil
	} else if cfg.DSN = os.Getenv("MYSQL_SVC_DSN"); cfg.DSN == "" { //env
		log.Warn().Msgf("get db config env error: %v", e.ErrNotFoundData)
	} else {
		log.Info().Msgf("get db config env succ, DSN: ***")
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
