package utils

import (
	"time"

	_ "gopkg.in/goyy/goyy.v0/util/cookies"
)

// func Cookies(r *http.Request, w http.ResponseWriter) {
// 	cookie := http.Cookie{
// 		Name:     "sess",
// 		Value:    "123",
// 		Path:     "/",
// 		HttpOnly: true,
// 		Expires:  time.Now().Add(48 * time.Hour),
// 	}
// 	http.SetCookie(w, &cookie)
// }

func Midnight() time.Time {
	year, month, day := time.Now().Date()
	loc, err := time.LoadLocation("Europe/Kiev")

	if err != nil {
		return time.Now()
	}
	return time.Date(year, month, day+1, 0, 0, 0, 0, loc)
}

// err == nil {
// 	if session, success := jwt.Decode(cookie.Value); success {
// 		ctx := context.WithValue(r.Context(), "session", session)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 		return
// 	}
