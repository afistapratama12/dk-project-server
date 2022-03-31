package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Username string
	Password string
	Host     string
	DBName   string
}

func Conn() *gorm.DB {

	var cred Config
	//TODO: godotenv disable
	// err := godotenv.Load()

	cred.Username = "u1656216_dk_project_admin"
	cred.Password = "dk_project_admin_2022"
	cred.Host = "srv143.niagahoster.com"
	cred.DBName = "u1656216_dk_database_project"

	var dns = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", cred.Username, cred.Password, cred.Host, cred.DBName)

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		PrepareStmt: true,
	})

	FailOnError(err, 36, "config/database.go")

	return db
}
