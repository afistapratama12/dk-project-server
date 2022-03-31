package repository

import (
	"dk-project-service/entity"
	"dk-project-service/utils"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/twilio/twilio-go"
	"gorm.io/gorm"

	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type UserRepository interface {
	GetuserId(id int) (entity.User, error)

	GetAllUsers() ([]entity.User, error)
	GetUserViews(id string) ([]entity.UserView, error)

	CheckUserLogin(username string, pass string) (entity.User, error)

	GetUsersByParentId(parentId string) ([]entity.User, error)

	// for register repo
	CheckUserId(id int) ([]entity.User, error)
	CreateUser(user entity.User) (entity.User, error)
	UpdateUsername(user entity.User) error

	// for transaction
	UpdateBalance(user entity.User) error

	// send WA message credential
	SendWARegister(user entity.User) error
}

type userRepository struct {
	db     *gorm.DB
	client *twilio.RestClient
}

func NewUserRepository(db *gorm.DB, client *twilio.RestClient) *userRepository {
	return &userRepository{db: db, client: client}
}

func (r *userRepository) GetuserId(id int) (entity.User, error) {
	var user entity.User

	idStr := strconv.Itoa(id)

	if err := r.db.Where("id = ?", idStr).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) GetAllUsers() ([]entity.User, error) {
	var users []entity.User

	if err := r.db.Where("role = ?", "user").Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (r *userRepository) GetUserViews(id string) ([]entity.UserView, error) {
	var userViews []entity.UserView
	var users []entity.User

	if err := r.db.Not("id = ?", id).Find(&users).Error; err != nil {
		return userViews, err
	}

	for _, u := range users {
		if u.Role == "user" {
			userViews = append(userViews, u.ToUserView())
		}
	}

	return userViews, nil
}

func (r *userRepository) CheckUserId(id int) ([]entity.User, error) {
	var usersDownline []entity.User

	idStr := strconv.Itoa(id)

	if err := r.db.Raw("SELECT * FROM users WHERE parent_id = ? AND position IN ('left', 'right', 'center')", idStr).Scan(&usersDownline).Error; err != nil {
		return usersDownline, err
	}

	return usersDownline, nil
}

func (r *userRepository) CheckUserLogin(username string, pass string) (entity.User, error) {
	var user entity.User

	if err := r.db.Where("username = ? AND password = ?", username, pass).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) CreateUser(user entity.User) (entity.User, error) {
	var query = `INSERT INTO users (id_generate, role, fullname, phone_number, username, password, parent_id, position) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	if err := r.db.Exec(query, user.IdGenerate, user.Role, user.Fullname, user.PhoneNumber, user.Username, user.Password, user.ParentId, user.Position).Error; err != nil {
		return user, err
	}

	if err := r.db.Where("id_generate = ?", user.IdGenerate).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) UpdateUsername(user entity.User) error {
	if err := r.db.Exec("UPDATE users SET username = ? WHERE id_generate = ?", user.Username, user.IdGenerate).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) UpdateBalance(user entity.User) error {
	if err := r.db.Exec("UPDATE users SET sas_balance = ?, ro_balance = ?, money_balance = ? WHERE id = ?", user.SASBalance, user.ROBalance, user.MoneyBalance, user.Id).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetUsersByParentId(parentId string) ([]entity.User, error) {
	var users []entity.User

	if err := r.db.Where("parent_id = ? && role = ?", parentId, "user").Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (r *userRepository) SendWARegister(user entity.User) error {
	var from = "+14155238886"

	sendTo, ok := utils.NumberSend(user.PhoneNumber)
	if !ok {
		return errors.New("error number user invalid format")
	}

	params := &openapi.CreateMessageParams{}
	params.SetTo(fmt.Sprintf("whatsapp:%s", sendTo))
	params.SetFrom(fmt.Sprintf("whatsapp:%s", from))
	params.SetBody(fmt.Sprintf("\nSelamat bergabung dengan DK, Berikut adalah username dan pin anda. Username : %s, Pin : %s", user.Username, user.Password))

	resp, err := r.client.ApiV2010.CreateMessage(params)
	if err != nil {
		return err
	}

	response, _ := json.MarshalIndent(*resp, "", "\t")
	log.Println("Response send WA :", string(response))

	return nil
}
