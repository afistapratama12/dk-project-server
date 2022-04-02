package main

import (
	"dk-project-service/config"
	"dk-project-service/entity"
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
	case "create_seed_data":
		ExecuteQueries(db, "./migration/createDataSeed.sql")
	case "create_admin_data":
		ExecuteQueries(db, "./migration/createAdmin.sql")
	case "drop_db":
		// drop tables
		ExecuteQueries(db, "./migration/droptable.sql")
	case "create_user_dl":
		CreateUserAndBA(db)
	case "scenarion_ro_wd":

	default:
		break
	}
}

func Err(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func DBCreateUserAndBA(db *gorm.DB, user entity.User, ba entity.BankAccount) {
	var query = `INSERT INTO users (id_generate, role, fullname, phone_number, username, password, parent_id, position, ro_balance) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	if err := db.Exec(query, user.IdGenerate, user.Role, user.Fullname, user.PhoneNumber, user.Username, user.Password, user.ParentId, user.Position, user.ROBalance).Error; err != nil {
		fmt.Println("create user", err.Error())
	} else {
		fmt.Println("success create user", user.Fullname)
	}

	if err := db.Create(&ba).Error; err != nil {
		fmt.Println("create bank account", err.Error())
	} else {
		fmt.Println("success create bank acc", user.Fullname)
	}
}

func EntityUserAndBankAcc(userId int, position string, parentId int) (entity.User, entity.BankAccount) {
	user := entity.User{
		Role:        "user",
		Fullname:    fmt.Sprintf("user%d", userId),
		PhoneNumber: "082132491234",
		Username:    fmt.Sprintf("DK-%d-user%d", userId, userId),
		Password:    "1234",
		ParentId:    parentId,
		Position:    position,
		ROBalance:   1,
	}
	bankAccount := entity.BankAccount{
		UserId:     userId,
		BankName:   "BCA",
		BankNumber: "123123123",
		NameOnBank: fmt.Sprintf("user%d", userId),
	}

	return user, bankAccount
}

// scenario 5 kedalaman, dan melakukan penarikan bersama ro 1
func CreateUserAndBA(db *gorm.DB) {
	// case 5 kedalaman, ada 3 posisi, (kiri, tengah, kanan)

	// upline
	fmt.Print("create upline\n\n")

	user, bankAcc := EntityUserAndBankAcc(2, "left", 1)

	DBCreateUserAndBA(db, user, bankAcc)

	// create user 5 kedalaman
	// kedalaman 1  2 => 3, 4, 5
	var position = []string{"left", "center", "right"}

	fmt.Print("\nkedalaman 1\n\n")

	for i := 3; i <= 5; i++ {
		u, ba := EntityUserAndBankAcc(i, position[i-3], 2)

		DBCreateUserAndBA(db, u, ba)
	}

	// kedalaman 2  | 3 => 6,7, 8 | 4 => 9, 10, 11 | 5 => 12, 13, 14
	fmt.Print("\nkedalaman 2\n\n")

	var idUser = 6

	for i := 3; i <= 5; i++ {
		var flagStop = 0
		var posId = 0

		for {
			u, ba := EntityUserAndBankAcc(idUser, position[posId], i)

			DBCreateUserAndBA(db, u, ba)

			idUser++
			posId++
			flagStop++
			if flagStop == 3 {
				break
			}
		}
	}

	// kedalaman 3
	// 6,7,8, 9, 19, 11, 12, 13, 14
	fmt.Print("\nkedalaman 3\n\n")
	for i := 6; i <= 14; i++ {
		var flagStop = 0
		var posId = 0

		for {
			u, ba := EntityUserAndBankAcc(idUser, position[posId], i)

			DBCreateUserAndBA(db, u, ba)

			idUser++
			posId++
			flagStop++
			if flagStop == 3 {
				break
			}
		}
	}

	// kedalaman 4
	// fmt.Print("\nkedalaman 4\n\n")
	// for i := 15; i <= 41; i++ {
	// 	var flagStop = 0
	// 	var posId = 0

	// 	for {
	// 		u, ba := EntityUserAndBankAcc(idUser, position[posId], i)

	// 		DBCreateUserAndBA(db, u, ba)

	// 		idUser++
	// 		posId++
	// 		flagStop++
	// 		if flagStop == 3 {
	// 			break
	// 		}
	// 	}
	// }

	// // kedalaman 5
	// fmt.Print("\nkedalaman 5\n\n")
	// for i := 42; i <= 122; i++ {
	// 	var flagStop = 0
	// 	var posId = 0

	// 	for {
	// 		u, ba := EntityUserAndBankAcc(idUser, position[posId], i)

	// 		DBCreateUserAndBA(db, u, ba)

	// 		idUser++
	// 		posId++
	// 		flagStop++
	// 		if flagStop == 3 {
	// 			break
	// 		}
	// 	}
	// }
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
