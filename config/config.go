package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type APPConfig struct {
	AppPort int
	DBConf  DBConfig
	RdsConf RedisConfig
}

type DBConfig struct {
	DbHost string
	DbPort int
	DbUser string
	DbPass string
	DbName string
}

type RedisConfig struct {
	RdsHost string
	RdsPort int
	RdsPass string
	RdsName string
}

func loadConfig() *APPConfig {
	var res = new(APPConfig)
	_, err := os.Stat(".env")
	if err == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Failed to fetch .env file")
		}
	}

	if v, f := os.LookupEnv("APP_PORT"); f {
		port, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal("Config : invalid server port", err.Error())
			return nil
		}
		res.AppPort = port
	}

	if value, found := os.LookupEnv("DBHOST"); found {
		res.DBConf.DbHost = value
	}

	if value, found := os.LookupEnv("DBPORT"); found {
		port, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal("Config : invalid db port", err.Error())
			return nil
		}
		res.DBConf.DbPort = port
	}

	if value, found := os.LookupEnv("DBUSER"); found {
		res.DBConf.DbUser = value
	}

	if value, found := os.LookupEnv("DBPASS"); found {
		res.DBConf.DbPass = value
	}

	if value, found := os.LookupEnv("DBNAME"); found {
		res.DBConf.DbName = value
	}
	if value, found := os.LookupEnv("RDSHOST"); found {
		res.RdsConf.RdsHost = value
	}
	if value, found := os.LookupEnv("RDSPORT"); found {
		port, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal("Config : invalid db port", err.Error())
			return nil
		}
		res.RdsConf.RdsPort = port
	}
	if value, found := os.LookupEnv("RDSPASS"); found {
		res.RdsConf.RdsPass = value
	}
	if value, found := os.LookupEnv("RDSNAME"); found {
		res.RdsConf.RdsName = value
	}
	return res
}

func RunConfig() *APPConfig {
	return loadConfig()
}