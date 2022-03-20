package main

import (
	"dk-project-service/config"
	"fmt"
	"log"
	"os"
	"strings"

	"gorm.io/gorm"
)

func main() {
	db := config.Conn()

	var checkFlag string

	for _, arg := range os.Args[1:] {
		checkFlag += arg
	}

	fmt.Println(checkFlag)

	switch checkFlag {
	case "migrate_db":
		// excute create table
		ExecuteQueries(db, "./migration/createtable.sql")

		//excute seeding data
		ExecuteQueries(db, "./migration/createDataSeed.sql")
	case "drop_db":
		// drop tables
		ExecuteQueries(db, "./migration/droptable.sql")
	default:
		break
	}
}

func Err(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ExecuteQueries(db *gorm.DB, pathFile string) {
	dat, err := os.ReadFile(pathFile)
	Err(err)

	listExecs := strings.Split(string(dat), ";")

	for _, qExec := range listExecs[:len(listExecs)-1] {
		if err := db.Exec(qExec).Error; err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("success execute", qExec)
		}
	}
}
