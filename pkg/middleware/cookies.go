package middleware

import (
	"net/http"
)

// SetCookie - SetCookie is a function that sets a cookie.
//
//	@param w - http.ResponseWriter
//	@param name - string
//	@param value - string
//	@param path - string
//	@param domain - string
//	@param maxAge - int
//	@param secure - bool
//	@param httpOnly - bool
//	@param sameSite - http.SameSite
func SetCookie(w http.ResponseWriter, name, value, path, domain string, maxAge int, secure, httpOnly bool, sameSite http.SameSite) {
	// create the cookie
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     path,
		Domain:   domain,
		MaxAge:   maxAge,
		Secure:   secure,
		HttpOnly: httpOnly,
		SameSite: sameSite,
	}

	// set the cookie
	http.SetCookie(w, cookie)
}

// GetCookie - GetCookie is a function that gets a cookie.
//
//	@param r - *http.Request
//	@param name - string
//	@return *http.Cookie
func GetCookie(r *http.Request, name string) *http.Cookie {
	// get the cookie
	cookie, err := r.Cookie(name)
	// check if there is an error
	if err != nil {
		return nil
	}

	// return the cookie
	return cookie
}

// RemoveCookie - RemoveCookie is a function that deletes a cookie.
//
//	@param w - http.ResponseWriter
//	@param name - string
//	@param path - string
//	@param domain - string
func RemoveCookie(w http.ResponseWriter, name, path, domain string) {
	// delete the cookie
	SetCookie(w, name, "", path, domain, -1, false, false, http.SameSiteLaxMode)
}
