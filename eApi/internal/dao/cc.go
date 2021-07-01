package dao

import (
	"os"
	"path/filepath"

	"github.com/gomsx/goms/pkg/conf"
	e "github.com/gomsx/goms/pkg/err"

	_ "github.com/go-sql-driver/mysql" // for init()
	"github.com/gomodule/redigo/redis"
	log "github.com/sirupsen/logrus"
)

// cccfg config of cache.
type cccfg struct {
	Addr string `yaml:"addr"`
	Pass string `yaml:"pass"`
}

// getCCConfig get cache config from file and env.
func getCCConfig(cfgpath string) (*cccfg, error) {
	var err error
	cfg := &cccfg{}
	// file
	path := filepath.Join(cfgpath, "redis.yaml")
	if err = conf.GetConf(path, &cfg); err != nil {
		log.Warnf("get cc config file error: %v", err)
	} else if cfg.Addr == "" {
		log.Warnf("get cc config file succeeded, but ADDR IS EMPTY")
	}
	// env
	if addr, exist := os.LookupEnv("REDIS_SVC_ADDR"); exist == false {
		log.Warnf("get cc config env error: %v", e.ErrNotFoundData)
	} else if addr == "" {
		log.Warnf("get cc config env succeeded, but ADDR IS EMPTY")
	} else {
		log.Infof("get cc config env succeeded, Addr: %v", cfg.Addr)
		cfg.Addr = addr
	}

	if pass, exist := os.LookupEnv("REDIS_SVC_PASS"); exist == false {
		log.Warnf("get cc config env error: %v", e.ErrNotFoundData)
	} else if pass == "" {
		log.Warnf("get cc config env succeeded, but PASS IS EMPTY")
	} else {
		log.Infof("get cc config env succeeded, PASS: %v", cfg.Pass)
		cfg.Pass = pass
	}

	if cfg.Addr == "" {
		return nil, e.ErrNotFoundData
	}
	return cfg, nil
}

// newCC new cache and return.
func newCC(cfgpath string) (redis.Conn, func(), error) {
	if cf, err := getCCConfig(cfgpath); err != nil {
		log.Errorf("get cc config error: %v", err)
		return nil, nil, err
	} else if cc, err := redis.Dial("tcp", cf.Addr, redis.DialPassword(cf.Pass)); err != nil {
		log.Errorf("dial cc error: %v", err)
		return nil, nil, err
	} else if _, err = cc.Do("PING"); err != nil {
		log.Errorf("ping cc error: %v", err)
		return nil, nil, err
	} else {
		return cc, func() {
			cc.Close()
		}, nil
	}
}
