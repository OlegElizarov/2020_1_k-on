package middleware

import (
	"github.com/go-park-mail-ru/2020_1_k-on/internal/session"
	"github.com/labstack/echo"
	"net/http"
)

type Auth struct {
	sessions session.Repository
}

func NewAuth(s session.Repository) Auth {
	return Auth{sessions: s}
}

func (a *Auth) GetSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		cookie, err := ctx.Cookie(session.CookieName)
		if err != nil || cookie == nil {
			return echo.ErrUnauthorized
		}

		uid, err := a.sessions.GetUserId(cookie.Value)
		if err != nil {
			return echo.ErrUnauthorized
		}

		ctx.Set(session.CookieName, cookie.Value)
		ctx.Set(session.UserIdKey, uid)
		return next(ctx)
	}
}

func (a *Auth) AlreadyLoginErr(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		_, err := ctx.Cookie(session.CookieName)
		if err != http.ErrNoCookie {
			return echo.NewHTTPError(http.StatusForbidden, "already login")
		}

		return next(ctx)
	}
}
