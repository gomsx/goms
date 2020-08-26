package dao

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	e "github.com/aivuca/goms/eApi/internal/pkg/err"
	"github.com/aivuca/goms/pkg/conf"

	_ "github.com/go-sql-driver/mysql" // for init()
	"github.com/rs/zerolog/log"
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
	err = fmt.Errorf("get file and env: %w", e.ErrNotFoundData)
	return nil, err
}

// newDB new a database.
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
