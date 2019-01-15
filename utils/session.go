package utils

import (
	// "context"
	// "encoding/gob"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"

	_ "github.com/gorilla/sessions"
	// _"github.com/go-martini/martini"
)

const (
	COOKIE_NAME = "sessionId"
)

type Session struct {
	id           string
	Username     string
	IsAuthorized bool
	Values       map[interface{}]interface{}
}

type SessionStore struct {
	data map[string]*Session
}

func NewSessionStore(c echo.Context) *SessionStore {
	s := new(SessionStore)
	s.data = make(map[string]*Session)
	fmt.Println(s)
	return s
}

func (store *SessionStore) Get(sessionId string) *Session {
	session := store.data[sessionId]
	if session == nil {
		return &Session{id: sessionId}
	}
	return session
}

func (store *SessionStore) Set(session *Session) {
	store.data[session.id] = session
}

func ensureCookie(r *http.Request, w http.ResponseWriter) string {
	cookie, _ := r.Cookie(COOKIE_NAME)
	if cookie != nil {
		return cookie.Value
	}
	sessionId := GenerateId()

	cookie = &http.Cookie{
		Name:    COOKIE_NAME,
		Value:   sessionId,
		Expires: time.Now().Add(365 * 24 * time.Hour),
	}
	http.SetCookie(w, cookie)

	return sessionId
}

// var sessionStore = NewSessionStore()

// func Middleware(r *http.Request, w http.ResponseWriter) {
// 	sessionId := ensureCookie(r, w)
// 	session := sessionStore.Get(sessionId)

// 	Map(session)

// 	Next()

// 	sessionStore.Set(session)
// }

//

// func (s *Session) Save(r *http.Request, w http.ResponseWriter) error {
// 	return s.store.Save(r, w, s)
// }

// func (s *Session) Name() string {
// 	return s.name
// }

// func (s *Session) Store() Store {
// 	return s.store
// }

// //

// type sessionInfo struct {
// 	s *Session
// 	e error
// }

// type contextKey int

// const registryKey contextKey = 0

// func GetRegistry(r *http.Request) *Registry {
// 	var ctx = r.Context()
// 	registry := ctx.Value(registryKey)
// 	if registry != nil {
// 		return registry.(*Registry)
// 	}
// 	newRegistry := &Registry{
// 		request:  r,
// 		sessions: make(map[string]sessionInfo),
// 	}
// 	*r = *r.WithContext(context.WithValue(ctx, registryKey, newRegistry))
// 	return newRegistry
// }

// type Registry struct {
// 	request  *http.Request
// 	sessions map[string]sessionInfo
// }

// func (s *Registry) Get(store Store, name string) (session *Session, err error) {
// 	if !isCookieNameValid(name) {
// 		return nil, fmt.Errorf("sessions: invalid character in cookie name: %s", name)
// 	}
// 	if info, ok := s.sessions[name]; ok {
// 		session, err = info.s, info.e
// 	} else {
// 		session, err = store.New(s.request, name)
// 		session.name = name
// 		s.sessions[name] = sessionInfo{s: session, e: err}
// 	}
// 	session.store = store
// 	return
// }

// func (s *Registry) Save(w http.ResponseWriter) error {
// 	var errMulti MultiError
// 	for name, info := range s.sessions {
// 		session := info.s
// 		if session.store == nil {
// 			errMulti = append(errMulti, fmt.Errorf(
// 				"sessions: missing store for session %q", name))
// 		} else if err := session.store.Save(s.request, w, session); err != nil {
// 			errMulti = append(errMulti, fmt.Errorf(
// 				"sessions: error saving session %q -- %v", name, err))
// 		}
// 	}
// 	if errMulti != nil {
// 		return errMulti
// 	}
// 	return nil
// }

// //

// func init() {
// 	gob.Register([]interface{}{})
// }

// func Save(r *http.Request, w http.ResponseWriter) error {
// 	return GetRegistry(r).Save(w)
// }

// func NewCookie(name, value string, options *Options) *http.Cookie {
// 	cookie := newCookieFromOptions(name, value, options)
// 	if options.MaxAge > 0 {
// 		d := time.Duration(options.MaxAge) * time.Second
// 		cookie.Expires = time.Now().Add(d)
// 	} else if options.MaxAge < 0 {
// 		cookie.Expires = time.Unix(1, 0)
// 	}
// 	return cookie
// }
