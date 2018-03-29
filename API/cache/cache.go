package cache

import (
	"fmt"

	lediscfg "github.com/siddontang/ledisdb/config"
	"github.com/siddontang/ledisdb/ledis"
)

type DB struct {
	*ledis.DB
}

var DBjwt *DB
var DBdata *DB

func init() {
	fmt.Println("Initializing DBs")
	DBjwt, _ = NewDB(0)
	//DBdata, _ = NewDB(1)
}

func NewDB(databasenum int) (*DB, error) {
	cfg := lediscfg.NewConfigDefault()
	l, _ := ledis.Open(cfg)

	database, err := l.Select(databasenum)

	return &DB{database}, err
}
