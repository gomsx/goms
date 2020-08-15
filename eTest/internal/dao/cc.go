package dao

import (
	"fmt"
	"os"
	"path/filepath"

	e "github.com/aivuca/goms/eTest/internal/pkg/err"
	"github.com/aivuca/goms/pkg/conf"

	_ "github.com/go-sql-driver/mysql" // for init()
	"github.com/gomodule/redigo/redis"
	"github.com/rs/zerolog/log"
)

// cccfg cache config.
type cccfg struct {
	Addr string `yaml:"addr"`
	Pass string `yaml:"pass"`
}

// getCCConfig get cache config from file and env.
func getCCConfig(cfgpath string) (*cccfg, error) {
	var err error
	cfg := &cccfg{}

	path := filepath.Join(cfgpath, "redis.yaml")
	if err = conf.GetConf(path, &cfg); err != nil { //file
		log.Warn().Msgf("get cc config file, error: %v", err)
	} else if cfg.Addr != "" {
		log.Info().Msgf("get cc config file, Addr: %v", cfg.Addr)
		return cfg, nil
	} else if cfg.Addr = os.Getenv("REDIS_SVC_ADDR"); cfg.Addr == "" { //env
		log.Warn().Msgf("get cc config env, empty")
	} else {
		log.Info().Msgf("get cc config env, Addr: %v", cfg.Addr)
		return cfg, nil
	}
	err = fmt.Errorf("get file and env: %w", e.ErrNotFoundData)
	return nil, err
}

// newCC new a cache.
func newCC(cfgpath string) (redis.Conn, func(), error) {
	if cf, err := getCCConfig(cfgpath); err != nil {
		log.Error().Msgf("get cc config, error: %v", err)
		return nil, nil, err
	} else if cc, err := redis.Dial("tcp", cf.Addr, redis.DialPassword(cf.Pass)); err != nil {
		log.Error().Msgf("dial cc, error: %v", err)
		return nil, nil, err
	} else if _, err = cc.Do("PING"); err != nil {
		log.Error().Msgf("ping cc, error: %v", err)
		return nil, nil, err
	} else {
		return cc, func() {
			cc.Close()
		}, nil
	}
}
