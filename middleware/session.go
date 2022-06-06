package middleware

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"klik/utils/session"
)

func NewCookieStore() *sessions.CookieStore {
	authKey := []byte("q3t6w9z$")
	encryptionKey := []byte("Qy3RBtseuIXUfBYxveg4YA==")
	s := sessions.NewCookieStore(authKey, encryptionKey)
	s.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 360,
		HttpOnly: true,
	}
	return s
}

func SessionMiddleware(s *session.ConfigSession) echo.MiddlewareFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			result, err := s.Get(context, session.SessionId)
			if err != nil {
				return context.Redirect(302, "/klik/auth/login")
			}
			if result == nil {
				return context.Redirect(302, "/klik/auth/login")
			} else {
				return handlerFunc(context)
			}
		}
	}
}
