// Package consts contains all values that are not supposed to change (such as URL keys and hostnames).
package consts

const (
	APIHostname  = "0.0.0.0"
	APIPort      = 3000
	DBHostname   = "db"
	DBPort       = 5432
	ServiceCount = 2
)

type Logger string

var (
	ErrLogger Logger = "errLogger"
	OutLogger Logger = "outLogger"
)
