package models

import (
	"database/sql"
	"encoding/json"
	"goweb2/helper"
	"html"
	"net/http"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	id         uint
	name       string
	email      string
	phone      string
	created_at time.Time
	updated_at time.Time
}

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
	// db, _ := database.ConnectDB()
	var existsUser string
	err := db.QueryRow("SELECT name FROM users WHERE email=?", email).Scan(&existsUser)
	switch {
	// Username is available
	case err == sql.ErrNoRows:
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		_, err = db.Exec("INSERT INTO users(name, email, password, phone) VALUES(?, ?, ?, ?)", fullName, email, hashedPassword, telephone)
		if err != nil {
			return false, "Server error, unable to create your account"
		}
		return true, "User Created!"
	case err != nil:
		return false, "Server error, unable to create your account."
	default:
		return true, "User Created!"
	}
}

func Login(res http.ResponseWriter, req *http.Request) (result bool, error_msg string) {
	email := html.EscapeString(req.FormValue("email"))
	password := req.FormValue("password")
	if email == "" || password == "" {
		return false, "Please inpull all fields"
	}
	if !rxEmail.MatchString(email) {
		return false, "The email must be a valid email address."
	}
	var dbEmail, dbPassword, idDb, nameDb string
	err := db.QueryRow("Select id, name, email, password from users where email =?", email).Scan(&idDb, &nameDb, &dbEmail, &dbPassword)
	if err != nil {
		return false, "Email not exists!"
	}
	err = bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))
	if err != nil {
		return false, "Password incorect!"
	}
	hashEmail, _ := bcrypt.GenerateFromPassword([]byte(email), bcrypt.DefaultCost)
	_, err = db.Exec("Update users set token = ? where id = ?", hashEmail, idDb)
	dataAuth := map[string]string{
		"name":  nameDb,
		"token": string(hashEmail),
	}
	authJson, _ := json.Marshal(dataAuth)
	helper.SetSession("AuthSession", string(authJson), res)
	return true, ""
}

func CheckLoginWithSession(session string) bool {

	var count int
	var authJson = make(map[string]string)
	err := json.Unmarshal([]byte(session), &authJson)
	token := string(authJson["token"])
	if err != nil || token == "" {
		return false
	}

	err = db.QueryRow("Select count(id) from users where token =?", token).Scan(&count)
	if err == nil && count > 0 {
		return true
	}
	return false

}

// func Logout(res http.ResponseWriter, req *http.Request) {
// 	helper.ClearSession("AuthSession", res)
// }
