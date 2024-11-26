package env

import (
	"github.com/devanfer02/filkom-canteen/internal/pkg/log"
	"github.com/spf13/viper"
)

type Env struct {
	AppEnv        string `mapstructure:"APP_ENV"`
	AppPort       string `mapstructure:"APP_PORT"`
	AppUrl        string `mapstructure:"APP_URL"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPass        string `mapstructure:"DB_PASS"`
	DBName        string `mapstructure:"DB_NAME"`
	JWTKey        string `mapstructure:"JWT_SECRET_KEY"`
	JWTUserRole   string `mapstructure:"JWT_USER_ROLE"`
	JWTAdminRole  string `mapstructure:"JWT_ADMIN_ROLE"`
	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPort     string `mapstructure:"REDIS_PORT"`
	RedisPassword string `mapstructure:"REDIS_PASS"`
	ApiKey        string `mapstructure:"API_KEY"`
}

var AppEnv = getEnv()

func getEnv() *Env {
	env := &Env{}

	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(log.LogInfo{
			"error": err.Error(),
		}, "[ENV][getEnv] failed to read config")
	}

	if err := viper.Unmarshal(env); err != nil {
		log.Fatal(log.LogInfo{
			"error": err.Error(),
		}, "[ENV][getEnv] failed to unmarshal struct")
	}

	return env
}
