package database

import "github.com/go-pg/pg"

var ErrNoRows = pg.ErrNoRows
var ErrMultiRows = pg.ErrMultiRows
var ErrStoreIsDeleted = "store is deleted"
var ErrClientLoginTimeout = "client_login_timeout"
var ErrEOF = "EOF"
var ErrorNoSuchHost = "no such host"
