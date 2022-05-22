package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	MySQL     *MySQL
	Validator Validator
	Fiber     Fiber
}

func SetArgs() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
	return &Config{
		MySQL: &MySQL{
			host:        os.Getenv("MYSQL_HOST"),
			port:        os.Getenv("MYSQL_PORT"),
			username:    os.Getenv("MYSQL_USER"),
			password:    os.Getenv("MYSQL_PASSWORD"),
			dbName:      os.Getenv("MYSQL_DBNAME"),
			maxIdlePool: getEnvInt("MYSQL_MAX_IDLE_POOL"),
			maxIdleTime: getEnvInt("MYSQL_MAX_IDLE_TIME"),
		},
	}
}

func getEnvInt(envStr string) (envInt int) {
	envInt, err := strconv.Atoi(os.Getenv(envStr))
	if err != nil {
		log.Fatalf("[CONFIG] Cannot get %s | %s\n", envStr, err.Error())
	}
	return
}
