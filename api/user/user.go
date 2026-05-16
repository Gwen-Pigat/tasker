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

func scanUser(row interface{ Scan(...any) error }) (User, error) {
	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.DateAdd, &user.IsActive, &user.Token)
	return user, err
}

func GetUser(wrapper *initializers.Wrapper) {
	rows, err := SelectUser("id", wrapper.ReturnUser())
	if err != nil {
		wrapper.Error(err.Error())
		return
	}
	defer rows.Close()

	if rows.Next() {
		user, err := scanUser(rows)
		if err != nil {
			wrapper.Error(err.Error())
			return
		}
		wrapper.Render(map[string]any{"data": user})
		return
	}
	wrapper.Error("User not found", http.StatusNotFound)
}

func GetUserConnect(wrapper *initializers.Wrapper) {
	var payload struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := wrapper.ParseJSON(&payload); err != nil {
		wrapper.Error(err.Error(), http.StatusBadRequest)
		return
	}

	rows, err := SelectUser("username", payload.Username)
	if err != nil {
		wrapper.Error(err.Error())
		return
	}
	defer rows.Close()

	if rows.Next() {
		user, err := scanUser(rows)
		if err != nil {
			wrapper.Error(err.Error())
			return
		}
		if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
			wrapper.Error("Invalid credentials", http.StatusUnauthorized)
			return
		}
		wrapper.Render(map[string]any{"data": ReturnUserPayload(user)})
		return
	}
	wrapper.Error("User not found", http.StatusNotFound)
}

func ReturnUserPayload(user User) map[string]string {
	return map[string]string{
		"username": user.Username,
		"dateAdd":  user.DateAdd,
		"token":    user.Token,
	}
}

func GetUserAuth(wrapper *initializers.Wrapper) (int, error) {
	token := strings.Replace(wrapper.Data["token"].(string), "Bearer ", "", -1)
	rows, err := SelectUser("token", token)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	if rows.Next() {
		user, err := scanUser(rows)
		if err != nil {
			return 0, err
		}
		return user.ID, nil
	}
	return 0, fmt.Errorf("invalid token")
}

func SelectUser(column string, value any) (*sql.Rows, error) {
	return initializers.DB.Query("SELECT "+userSetup["payload"]+" FROM "+userSetup["table"]+" WHERE "+column+"=?", value)
}

func CreateUser(wrapper *initializers.Wrapper) {
	var payload struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := wrapper.ParseJSON(&payload); err != nil {
		wrapper.Error(err.Error(), http.StatusBadRequest)
		return
	}

	if len(payload.Password) < 8 {
		wrapper.Error("Password must be at least 8 characters", http.StatusBadRequest)
		return
	}

	// Check if exists
	rows, err := SelectUser("username", payload.Username)
	if err == nil {
		defer rows.Close()
		if rows.Next() {
			wrapper.Error("Username already exists", http.StatusConflict)
			return
		}
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	user := User{
		Username: payload.Username,
		Password: string(hashedPassword),
		DateAdd:  time.Now().UTC().Format(initializers.Format),
		IsActive: true,
		Token:    uuid.New().String(),
	}

	res, err := initializers.DB.Exec(
		"INSERT INTO "+userSetup["table"]+" (username, password, date_add, is_active, token) VALUES (?,?,?,?,?)",
		user.Username, user.Password, user.DateAdd, user.IsActive, user.Token,
	)
	if err != nil {
		wrapper.Error(err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := res.LastInsertId()
	user.ID = int(id)
	wrapper.Render(map[string]any{"data": ReturnUserPayload(user)})
}
