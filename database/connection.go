package database

import (
	"fmt"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

var connectionPool map[string]*pg.DB = map[string]*pg.DB{}

func init() {
	orm.SetTableNameInflector(func(s string) string {
		return s
	})
}

func RegisterConnection(
	alias string,
	username string,
	password string,
	host string,
	port int,
	database string,
	timeout int,
	poolTimeout int,
	maxIdleConn int,
	maxConn int,
	debug bool,
) error {
	if prevConn, ok := connectionPool[alias]; ok {
		err := prevConn.Close()
		if err != nil {
			return err
		}
	}
	connectionPool[alias] = pg.Connect(&pg.Options{
		User:         username,
		Password:     password,
		Addr:         fmt.Sprintf("%s:%d", host, port),
		Database:     database,
		DialTimeout:  time.Duration(int64(time.Second) * int64(timeout)),
		PoolSize:     maxConn,
		PoolTimeout:  time.Duration(int64(time.Second) * int64(poolTimeout)),
		MinIdleConns: 5,
		OnConnect: func(c *pg.Conn) error {
			return nil
		},
	})

	connectionPool[alias].AddQueryHook(dbTimeUpdater{})
	return nil
}

func GetOrmer() Ormer {
	return GetOrmerWithAlias("default")
}

func GetOrmerWithAlias(alias string) Ormer {
	conn, _ := connectionPool[alias]
	return NewOrmer(conn)
}
