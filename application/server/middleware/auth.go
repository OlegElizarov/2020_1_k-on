package middleware

import (
	"github.com/go-park-mail-ru/2020_1_k-on/application/session"
	"github.com/labstack/echo"
	"net/http"
	"time"
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
			return WriteErrResponse(ctx, http.StatusUnauthorized, "no cookie")
		}

		uid, err := a.sessions.GetUserId(cookie.Value)
		if err != nil {
			cookie.Expires = time.Now().AddDate(0, 0, -1)
			ctx.SetCookie(cookie)
			return WriteErrResponse(ctx, http.StatusUnauthorized, "session does not exist")
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
			return WriteErrResponse(ctx, http.StatusForbidden, "already login")
		}

		return next(ctx)
	}
}