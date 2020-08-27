package dao

import (
	"log"
	"os"
	"path/filepath"

	e "github.com/aivuca/goms/eRedis/internal/pkg/err"
	"github.com/aivuca/goms/pkg/conf"

	_ "github.com/go-sql-driver/mysql" // for init()
	"github.com/gomodule/redigo/redis"
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

	path := filepath.Join(cfgpath, "redis.yaml")
	if err = conf.GetConf(path, &cfg); err != nil { //file
		log.Printf("get cc config file error: %v", err)
	} else if cfg.Addr != "" {
		log.Printf("get cc config file succ, Addr: %v", cfg.Addr)
		return cfg, nil
	} else if cfg.Addr = os.Getenv("REDIS_SVC_ADDR"); cfg.Addr == "" { //env
		log.Printf("get cc config env error: %v", e.ErrNotFoundData)
	} else {
		log.Printf("get cc config env succ, Addr: %v", cfg.Addr)
		return cfg, nil
	}
	return nil, e.ErrNotFoundData
}

// newCC new cache and return.
func newCC(cfgpath string) (redis.Conn, func(), error) {
	if cf, err := getCCConfig(cfgpath); err != nil {
		log.Printf("get cc config error: %v", err)
		return nil, nil, err
	} else if cc, err := redis.Dial("tcp", cf.Addr, redis.DialPassword(cf.Pass)); err != nil {
		log.Printf("dial cc error: %v", err)
		return nil, nil, err
	} else if _, err = cc.Do("PING"); err != nil {
		log.Printf("ping cc error: %v", err)
		return nil, nil, err
	} else {
		return cc, func() {
			cc.Close()
		}, nil
	}
}
