package dao

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Dao Dao.
type Dao struct {
	db *sql.DB
}

// New new a dao.
func New(cfgpath string) *Dao {
	mdb, _, err := newDB(cfgpath)
	if err != nil {
		log.Panicf("failed to get config: %v", err)
	}
	log.Printf("db ok")
	mdao := &Dao{db: mdb}
	return mdao
}

// Close close the resource.
func (d *Dao) Close() {
	d.db.Close()
}

// Ping ping the resource.
func (d *Dao) Ping(c context.Context) (err error) {
	return d.db.PingContext(c)
}
