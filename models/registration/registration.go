package registration

import (
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	"test-golang/database"
	"test-golang/utils/config"
)

func init() {
	cfg := config.Config.PgCfg
	connString := fmt.Sprintf("postgres://%s:%s@%s/%s?connect_timeout=%s&sslmode=disable", cfg.Username, cfg.Password, cfg.Host, cfg.Database, strconv.Itoa(cfg.Timeout))
	fmt.Println("connString: ", connString)

	err := database.RegisterConnection(
		"default",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		5432,
		cfg.Database,
		cfg.Timeout,
		cfg.PoolTimeout,
		cfg.MaxIdleConn,
		cfg.MaxConn,
		true,
	)

	if err != nil {
		fmt.Println("connString: ", err.Error())
	}
}
