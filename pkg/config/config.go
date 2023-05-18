package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/dev-mockingbird/logf"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	Mysql string = "mysql"
)

type DBConfig struct {
	Disable  bool
	DBMS     string
	Host     string
	Database string
	Port     int
	User     string
	Password string
}

type RedisConfig struct {
	Disable  bool
	Addr     string
	Username string
	Password string
	DB       int
}

type KafkaConfig struct {
	Disable  bool
	Brokers  []string
	User     string
	Password string
}

type HttpConfig struct {
	Disable bool
	Port    int
}

type GrpcConfig struct {
	Disable bool
	Port    int
}

type JaegerConfig struct {
	URL string
}

type Config struct {
	Http   HttpConfig
	Grpc   GrpcConfig
	DB     DBConfig
	Redis  RedisConfig
	Kafka  KafkaConfig
	Jaeger JaegerConfig
}

func (cfg DBConfig) MysqlDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
}

func (cfg DBConfig) DSN() string {
	if cfg.DBMS == "" {
		cfg.DBMS = Mysql
	}
	switch cfg.DBMS {
	case Mysql:
		return cfg.MysqlDSN()
	default:
		panic("not support dbms: " + cfg.DBMS)
	}
}

func ReadInput(cfg *Config, logger logf.Logger) error {
	viper.SetEnvPrefix("mb")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()
	var (
		jsonFile string
		yamlFile string
	)
	pflag.StringVar(&jsonFile, "json", "", "json config path file")
	pflag.StringVar(&yamlFile, "yaml", "", "yaml config path file")
	pflag.Bool("db.disable", false, "enable db or not")
	pflag.String("db.dbms", "mysql", "dbms")
	pflag.String("db.host", "127.0.0.1", "db host")
	pflag.String("db.database", "mockingbird", "database to use")
	pflag.Int("db.port", 3306, "db port")
	pflag.String("db.user", "root", "db user")
	pflag.String("db.password", "", "db password")
	pflag.Bool("redis.disable", false, "enable redis or not")
	pflag.String("redis.addr", "127.0.0.1:3379", "redis addr")
	pflag.String("redis.username", "", "redis username")
	pflag.String("redis.password", "", "redis password")
	pflag.Int("redis.db", 0, "redis db")
	pflag.Bool("kafka.disable", false, "enable kafka or not")
	pflag.StringArray("kafka.brokers", []string{"127.0.0.1:9092"}, "kafka brokers")
	pflag.String("kafka.user", "", "kafka user")
	pflag.String("kafka.password", "", "kafka password")
	pflag.String("jaeger.url", "", "jaeger url")
	pflag.String("http.disable", "", "enable http or not")
	pflag.Int("http.port", 7000, "http port")
	pflag.String("grpc.disable", "", "enable grpc or not")
	pflag.Int("grpc.port", 7001, "grpc port")
	pflag.Parse()
	if jsonFile != "" {
		viper.SetConfigType("json")
		f, err := os.Open(jsonFile)
		if err != nil {
			logger.Logf(logf.Error, "open file [%s]: %s", jsonFile, err.Error())
			return err
		}
		if err := viper.ReadConfig(f); err != nil {
			logger.Logf(logf.Error, "read json config: %s", err.Error())
			return err
		}
	}
	if yamlFile != "" {
		viper.SetConfigType("yaml")
		f, err := os.Open(yamlFile)
		if err != nil {
			logger.Logf(logf.Error, "open file [%s]: %s", yamlFile, err.Error())
			return err
		}
		if err := viper.ReadConfig(f); err != nil {
			logger.Logf(logf.Error, "read json config: %s", err.Error())
			return err
		}
	}
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		logger.Logf(logf.Error, "read command line: %s", err.Error())
		return err
	}

	return viper.Unmarshal(cfg)
}
