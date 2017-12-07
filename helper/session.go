package helper

import (
	"fmt"
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

type Orders struct {
	IdProduct int64
}

var cookieHandler = securecookie.New(
	// securecookie.GenerateRandomKey(64),
	// securecookie.GenerateRandomKey(32),
	[]byte("12345"),
	[]byte("1234567890123456"),
)
var keyFlash string = "flash-session"
var store = sessions.NewCookieStore([]byte("secret-password"))

func SetFlash(value interface{}, w http.ResponseWriter, req *http.Request) bool {
	session, err := store.Get(req, keyFlash)
	if err != nil {
		return false
	}
	session.AddFlash(value)
	errr := session.Save(req, w)
	fmt.Println(errr)
	return true
}
func GetFlash(key string, w http.ResponseWriter, req *http.Request) interface{} {
	session, err := store.Get(req, keyFlash)
	if err != nil {
		return []string{}
	}

	flash := session.Flashes()
	session.Save(req, w)
	return flash
	// // fmt.Println(session.Flashes())

	return false
}

func SetSession(key string, val string, r http.ResponseWriter) bool {
	if encoded, err := cookieHandler.Encode(key, val); err == nil {
		cookie := &http.Cookie{
			Name:  key,
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(r, cookie)
		return true
	}
	return true
}

func ClearSession(key string, r http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   key,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(r, cookie)
}

func GetSession(key string, r *http.Request) string {
	// fmt.Println(securecookie.GenerateRandomKey(64))
	// fmt.Println(securecookie.GenerateRandomKey(32))
	var cookieRes string
	if cookie, err := r.Cookie(key); err == nil {
		err := cookieHandler.Decode(key, cookie.Value, &cookieRes)
		if err != nil {
			fmt.Println("GetSession", err)
		}
	}
	return cookieRes
}
