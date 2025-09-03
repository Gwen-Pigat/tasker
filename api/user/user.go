package user

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"tasker/initializers"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	DateAdd  string `json:"dateAdd"`
	IsActive bool   `json:"isActive"`
	Token    string `json:"token"`
}

var userSetup = map[string]string{
	"payload": "id,username,password,date_add,is_active,token",
	"table":   "user",
}

func GetUser(wrapper *initializers.Wrapper) {
	rows, err := SelectUser("id", wrapper.ReturnUser())
	if err != nil {
		wrapper.Error(err.Error())
		return
	}
	defer rows.Close()
	var user User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.DateAdd, &user.IsActive, &user.Token); err != nil {
			wrapper.Error(err.Error())
			return
		}
	}
	if user.ID == 0 {
		wrapper.Error("User cannot be found", http.StatusNotFound)
		return
	}
	wrapper.Render(map[string]any{
		"data": user,
	})
}

func GetUserConnect(wrapper *initializers.Wrapper) {
	keys := []string{
		"username", "password",
	}
	for _, key := range keys {
		if err := wrapper.WrapData(key); err != nil {
			wrapper.Error(err.Error())
			return
		}
	}
	rows, err := SelectUser("username", wrapper.Data["username"])
	if err != nil {
		wrapper.Error(err.Error())
		return
	}
	defer rows.Close()

	var user User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.DateAdd, &user.IsActive, &user.Token); err != nil {
			wrapper.Error(err.Error())
			return
		}
	}
	if user.ID == 0 {
		wrapper.Error("User cannot be found", http.StatusNotFound)
		return
	}
	if user.Password == "" {
		wrapper.Error("Password is empty in database", http.StatusConflict)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(wrapper.Data["password"].(string)))
	if err != nil {
		wrapper.Error("Password or username is not valid", http.StatusNotFound)
		return
	}
	wrapper.Render(map[string]any{
		"data": ReturnUserPayload(user),
	})
}

func ReturnUserPayload(user User) map[string]string {
	return map[string]string{
		"username": user.Username,
		"dateAdd":  user.DateAdd,
		"token":    user.Token,
	}
}

func GetUserAuth(wrapper *initializers.Wrapper) (userID int, error error) {
	if err := wrapper.WrapData("token"); err != nil {
		return 0, err
	}
	token := strings.Replace(wrapper.Data["token"].(string), "Bearer ", "", -1)
	rows, err := SelectUser("token", token)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	var user User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.DateAdd, &user.IsActive, &user.Token); err != nil {
			return 0, err
		}
	}
	if user.ID == 0 {
		return 0, fmt.Errorf("user cannot be found, token = %v", token)
	}
	return user.ID, err
}

func SelectUser(column string, value any) (*sql.Rows, error) {
	rows, err := initializers.DB.Query("SELECT "+userSetup["payload"]+" FROM "+userSetup["table"]+" WHERE "+column+"=?", value)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func CreateUser(wrapper *initializers.Wrapper) {
	rows, err := SelectUser("username", wrapper.Data["username"].(string))
	if err != nil {
		wrapper.Error(err.Error(), http.StatusBadGateway)
		return
	}
	defer rows.Close()
	if rows.Next() {
		wrapper.Error("Username already exist", http.StatusBadGateway)
		return
	}
	keys := []string{
		"username", "password",
	}
	for _, key := range keys {
		if err := wrapper.WrapData(key); err != nil {
			wrapper.Error(err.Error(), http.StatusBadRequest)
			return
		}
	}
	if len(wrapper.Data["password"].(string)) < 8 {
		wrapper.Error("Password must be 8 caracters long minimum", http.StatusBadRequest)
		return
	}
	smtp, err := initializers.DB.Prepare("INSERT INTO " + userSetup["table"] + "(username,password,date_add,is_active,token) VALUES(?,?,?,?,?)")
	if err != nil {
		wrapper.Error(err.Error(), http.StatusBadGateway)
		return
	}
	defer smtp.Close()

	password, _ := bcrypt.GenerateFromPassword([]byte(wrapper.Data["password"].(string)), bcrypt.DefaultCost)

	user := User{
		Username: wrapper.Data["username"].(string),
		DateAdd:  time.Now().UTC().Truncate(time.Second).Format(initializers.Format),
		IsActive: true,
		Password: string(password),
		Token:    uuid.New().String(),
	}
	result, err := smtp.Exec(user.Username, user.Password, user.DateAdd, user.IsActive, user.Token)
	if err != nil {
		wrapper.Error(err.Error(), http.StatusBadGateway)
		return
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		wrapper.Error(err.Error(), http.StatusBadGateway)
		return
	}
	user.ID = int(lastInsertID)
	wrapper.Render(map[string]any{
		"data": ReturnUserPayload(user),
	})
}
