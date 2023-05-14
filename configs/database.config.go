package configs

import (
	"os"

	"gorm.io/gorm"
)

var db_name = ""
var db_user = "root"
var db_passwd = ""
var db_host = ""
var db_port = ""

var DB *gorm.DB

func bootDatabase() {
	if dbNameEnv := os.Getenv("DB_NAME"); dbNameEnv != "" {
		db_name = dbNameEnv
	}
	if dbUserEnv := os.Getenv("DB_USER"); dbUserEnv != "" {
		db_user = dbUserEnv
	}
	if dbPassEnv := os.Getenv("DB_PASSWD"); dbPassEnv != "" {
		db_passwd = dbPassEnv
	}
	if dbHostEnv := os.Getenv("DB_HOST"); dbHostEnv != "" {
		db_host = dbHostEnv
	}
	if dbPortEnv := os.Getenv("DB_PORT"); dbPortEnv != "" {
		db_port = dbPortEnv
	}
}
