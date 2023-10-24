package bootstrap

import (
	"github.com/spf13/viper"
	"log"
)

type Env struct {
	AppEnv         string `mapstructure:"APP_ENV"`
	ServerAddress  string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout int    `mapstructure:"CONTEXT_TIMEOUT"`

	DBMigrationURL string `mapstructure:"DB_MIGRATION_URL"`
	DBSource       string `mapstructure:"DB_SOURCE"`

	DBHost string `mapstructure:"DB_HOST"`
	DBPort string `mapstructure:"DB_PORT"`
	DBUser string `mapstructure:"DB_USER"`
	DBPass string `mapstructure:"DB_PASS"`
	DBName string `mapstructure:"DB_NAME"`

	DBMysqlHost string `mapstructure:"DB_MYSQL_HOST"`
	DBMysqlPort string `mapstructure:"DB_MYSQL_PORT"`
	DBMysqlUser string `mapstructure:"DB_MYSQL_USER"`
	DBMysqlPass string `mapstructure:"DB_MYSQL_PASS"`
	DBMysqlName string `mapstructure:"DB_MYSQL_NAME"`

	DBMysqlClusterHost string `mapstructure:"DB_MYSQL_CLUSTER_HOST"`
	DBMysqlClusterPort string `mapstructure:"DB_MYSQL_CLUSTER_PORT"`
	DBMysqlClusterUser string `mapstructure:"DB_MYSQL_CLUSTER_USER"`
	DBMysqlClusterPass string `mapstructure:"DB_MYSQL_CLUSTER_PASS"`
	DBMysqlClusterName string `mapstructure:"DB_MYSQL_CLUSTER_NAME"`

	DBRedisSource      string `mapstructure:"DB_REDIS_SOURCE"`
	CacheRedisHost     string `mapstructure:"DB_REDIS_HOST"`
	CacheRedisPort     string `mapstructure:"DB_REDIS_PORT"`
	CacheRedisPassword string `mapstructure:"DB_REDIS_PASSWORD"`
	CacheRedisDB       int    `mapstructure:"DB_REDIS_DB"`

	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`

	OTPURL  string `mapstructure:"OTP_URL"`
	OTPFLAG bool   `mapstructure:"OTP_FLAG"`

	TOKEN_KEY    string `mapstructure:"TOKEN_KEY"`
	HTTP_ADDRESS string `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPC_ADDRESS string `mapstructure:"GRPC_SERVER_ADDRESS"`
}

func NewEnv() *Env {
	env := Env{}

	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	//viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}
