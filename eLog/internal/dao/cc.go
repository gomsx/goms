package dao

import (
	"os"
	"path/filepath"

	e "github.com/fuwensun/goms/eLog/internal/pkg/err"
	"github.com/fuwensun/goms/pkg/conf"

	_ "github.com/go-sql-driver/mysql" // for init()
	"github.com/gomodule/redigo/redis"
	"github.com/rs/zerolog/log"
)

// cccfg config for cache.
type cccfg struct {
	Addr string `yaml:"addr"`
	Pass string `yaml:"pass"`
}

// getCCConfig get cache config from file and env.
func getCCConfig(cfgpath string) (*cccfg, error) {
	cfg := &cccfg{}
	path := filepath.Join(cfgpath, "redis.yaml")
	if err := conf.GetConf(path, &cfg); err != nil { //file
		log.Warn().Msgf("get cc config file error: %v", err)
	} else if cfg.Addr != "" {
		log.Info().Msgf("get cc config file succ, addr: %v", cfg.Addr)
		return cfg, nil
	} else if cfg.Addr = os.Getenv("REDIS_SVC_ADDR"); cfg.Addr == "" { //env
		log.Warn().Msgf("get cc config env error: empty")
	} else {
		log.Info().Msgf("get cc config env succ, addr: %v", cfg.Addr)
		return cfg, nil
	}
	return nil, e.ErrNotFoundData
}

// newCC new cache.
func newCC(cfgpath string) (redis.Conn, func(), error) {
	if cf, err := getCCConfig(cfgpath); err != nil {
		log.Error().Msgf("get cc config error: %v", err)
		return nil, nil, err
	} else if cc, err := redis.Dial("tcp", cf.Addr, redis.DialPassword(cf.Pass)); err != nil {
		log.Error().Msgf("dial cc error: %v", err)
		return nil, nil, err
	} else if _, err = cc.Do("PING"); err != nil {
		log.Error().Msgf("ping cc error: %v", err)
		return nil, nil, err
	} else {
		return cc, func() {
			cc.Close()
		}, nil
	}
}
