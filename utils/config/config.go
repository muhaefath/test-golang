package config

import (
	"fmt"
	"path/filepath"

	"github.com/Shopify/sarama"
	"github.com/lovoo/goka"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const ProjectContext = "test-golang_ctx"

// Environment is the running application mode (dev, test, production)
var Environment string = "dev"
var Config config
var GoogleCredentialLocation string
var Consumer bool

type config struct {
	PgCfg struct {
		Database    string `mapstructure:"database"`
		Host        string `mapstructure:"host"`
		Username    string `mapstructure:"username"`
		Password    string `mapstructure:"password"`
		MaxConn     int    `mapstructure:"max_conn"`      // max count of connections stored in the connection pool (the connection pool is maintained by go-pg)
		MaxIdleConn int    `mapstructure:"max_idle_conn"` // max count of connection in idle state in the connection pool, when the count of connection in idle state is more than this config, the excess connection will be closed
		Timeout     int    `mapstructure:"timeout"`       // timeout tolerance for dial attempt to database server (in seconds)
		PoolTimeout int    `mapstructure:"pool_timeout"`  // timeout tolerance for fetching a connection from the connection pool (in seconds)
	} `mapstructure:"postgres"`
}

func init() {
	var err error

	viper.SetEnvPrefix("CNX")
	viper.AutomaticEnv()

	configName := "application"

	if viper.IsSet("ENV") {
		Environment = viper.GetString("ENV")
	}

	Consumer = viper.GetBool("CONSUMER")

	log.Info("Consumer :", Consumer)

	if Environment != "production" {
		configName = configName + "." + Environment
	}

	log.Info("CNX_ENV: ", Environment)

	viper.SetConfigName(configName) // name of config file (without extension)
	viper.AddConfigPath(filepath.Join(GetAppBasePath(), "conf"))
	viper.AddConfigPath(".")   // optionally look for config in the working directory
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	//Unmarshal application yml to config
	err = viper.Unmarshal(&Config)

	if err != nil {
		log.Errorf("unable to decode into struct, %v", err)
	}

	cfg := goka.DefaultConfig()
	cfg.Version = sarama.V2_4_0_0
	goka.ReplaceGlobalConfig(cfg)
}

func GetAppBasePath() string {
	basePath, _ := filepath.Abs(".")
	for filepath.Base(basePath) != "test-golang" {
		basePath = filepath.Dir(basePath)
	}
	return basePath
}
