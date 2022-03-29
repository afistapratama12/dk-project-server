package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
	SSLMode  string
}

func Conn() *gorm.DB {

	var cred Config
	//TODO: godotenv disable
	err := godotenv.Load()

	FailOnError(err, 24, "config/database.go")
	cred.Username = os.Getenv("DB_USER")
	cred.Password = os.Getenv("DB_PASS")
	cred.Host = os.Getenv("DB_HOST")
	cred.DBName = os.Getenv("DB_NAME")
	cred.Port = os.Getenv("DB_PORT")
	cred.SSLMode = os.Getenv("SSL")

	var dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cred.Username, cred.Password, cred.Host, cred.Port, cred.DBName)

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		PrepareStmt: true,
	})

	FailOnError(err, 36, "config/database.go")

	return db
}
