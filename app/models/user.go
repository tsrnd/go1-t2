package models

import (
	"encoding/json"
	"fmt"
	"goweb2/helper"
	"html"
	"net/http"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id         string
	Name       string
	Email      string
	Password   string
	Phone      string
	Token      string    `gorm:"default:null"`
	Created_at time.Time `gorm:"default:null"`
	Updated_at time.Time `gorm:"default:null"`
}

var CurrentUser User

const _EXP_EMAIL = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`

var rxEmail = regexp.MustCompile("^(((([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|((\\x22)((((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(([\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(\\([\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(\\x22)))@((([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$")

func StoreUser(req *http.Request) (result bool, error_msg string) {
	fullName := html.EscapeString(req.FormValue("full_name"))
	email := html.EscapeString(req.FormValue("email"))
	telephone := html.EscapeString(req.FormValue("telephone"))
	password := req.FormValue("password")
	passwordConf := req.FormValue("password_conf")
	if fullName == "" || email == "" || telephone == "" || password == "" || passwordConf == "" {
		return false, "Please inpull all fields"
	}
	if !rxEmail.MatchString(email) {
		return false, "The email must be a valid email address."
	}
	if password != passwordConf {
		return false, "The password confirmation does not match."
	}
	var existsUser User
	res := db.Select("email").Where("email = ?", email).First(&existsUser)
	if existsUser.Email != "" {
		return false, "Email has exists"
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	var user = User{Name: fullName, Email: email, Phone: telephone, Password: string(hashedPassword)}
	creat := db.Create(&user)
	if creat.Error != nil {
		return false, "Server error, unable to create your account"
	}
	return true, "User Created!"
}

func Login(res http.ResponseWriter, req *http.Request) (bool, string) {
	email := html.EscapeString(req.FormValue("email"))
	password := req.FormValue("password")
	if email == "" || password == "" {
		return false, "Please inpull all fields"
	}
	if !rxEmail.MatchString(email) {
		return false, "The email must be a valid email address."
	}
	var user User
	result := db.Where("email = ?", email).First(&user)

	// err := db.QueryRow("Select id, name, email, password from users where email =?", email).Scan(&idDb, &nameDb, &dbEmail, &dbPassword)
	if result != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return false, "Login fail"
	}
	user.Token = user.Password
	db.Save(&user)
	dataAuth := map[string]string{
		"Id":    user.Id,
		"Email": user.Email,
		"Name":  user.Name,
		"Token": user.Token,
	}
	authJson, _ := json.Marshal(dataAuth)
	helper.SetSession("AuthSession", string(authJson), res)
	return true, ""
}

func CheckLoginWithSession(session string) bool {

	var authJson = make(map[string]string)
	var user User
	err := json.Unmarshal([]byte(session), &authJson)
	token := string(authJson["Token"])
	if err != nil || token == "" {
		return false
	}
	result := db.Where("token = ?", token).First(&user)
	if result.Error == nil && user.Name != "" {
		return true
	}
	return false

}

func CheckAuth(r *http.Request) bool {
	authSS := helper.GetSession("AuthSession", r)
	if authSS != "" {
		return CheckLoginWithSession(authSS)
	}
	return false
}

func (auth User) IsLogin() bool {
	if auth.Id != "" {
		return true
	}
	return false
}

func GetAuth(r *http.Request) User {
	authData := User{}
	authSS := helper.GetSession("AuthSession", r)
	var authJson = make(map[string]string)
	err := json.Unmarshal([]byte(authSS), &authJson)
	if err != nil {
		fmt.Println("Helper View: GetAuther", err)
	} else {
		authData.Id = authJson["Id"]
		authData.Name = authJson["Name"]
		authData.Email = authJson["Email"]
		authData.Token = authJson["Token"]
	}
	return authData
}
