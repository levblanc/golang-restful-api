package auth

import (
	"net/http"
)

const cookieName = "mstream-session"

// CreateCookie creates a new cookie
func CreateCookie(w http.ResponseWriter, sid string) {
	cookie := &http.Cookie{
		Name:     cookieName,
		Value:    sid,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   30,
	}

	http.SetCookie(w, cookie)
}

// ReadCookie reads target cookie
// if cookie exist, returns the cookie
// else, returns the error
func ReadCookie(req *http.Request) (*http.Cookie, error) {
	cookie, err := req.Cookie(cookieName)

	if err != nil {
		return nil, err
	}

	return cookie, nil
}

// ExpireCookie expires target cookie
func ExpireCookie(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   cookieName,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}

	http.SetCookie(w, cookie)
}
