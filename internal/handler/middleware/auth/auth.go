package auth

import (
	"context"
	"errors"
	"github.com/paulwwyvern/urlshortener/pkg/httphelpers/httperr"
	"github.com/paulwwyvern/urlshortener/pkg/httphelpers/httpuser"
	"net/http"
)

type ErrUserNotFound struct {
	err  error
	desc string
}

func NewErrUserNotFound(err error, desc string) *ErrUserNotFound {
	return &ErrUserNotFound{
		err:  err,
		desc: desc,
	}
}

func (e ErrUserNotFound) Error() string {
	return "auth: " + e.desc + ": " + e.err.Error()
}

func (e ErrUserNotFound) Unwrap() error {
	return e.err
}

type UserService interface {
	CreateUser(ctx context.Context) (int32, error)
}

func WithAuthRequire(key string) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return httperr.Adapt(func(w http.ResponseWriter, r *http.Request) error {
			userId, err := GetUserID(key, r)
			var errUserNotFound *ErrUserNotFound
			if err != nil {
				if !errors.As(err, &errUserNotFound) {
					w.WriteHeader(http.StatusInternalServerError)
					return err
				}
			}

			if errUserNotFound != nil {
				// нет юзера - не пускаем дальше
				w.WriteHeader(http.StatusUnauthorized)
				return err
			}

			httpuser.SetUserID(r, userId)

			h.ServeHTTP(w, r)

			return nil
		})
	}
}

func WithAuth(key string, userService UserService) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return httperr.Adapt(func(w http.ResponseWriter, r *http.Request) error {
			userId, err := GetUserID(key, r)
			var errUserNotFound *ErrUserNotFound
			if err != nil {
				if !errors.As(err, &errUserNotFound) {
					w.WriteHeader(http.StatusInternalServerError)
					return err
				}
			}

			if errUserNotFound != nil {
				// нет юзера - создаём нового
				userId, err = userService.CreateUser(r.Context())
				if err != nil {
					return err
				}

				token, err := CreateJWTToken(key, userId)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return err
				}

				cookie := &http.Cookie{
					Name:     "Token",
					Value:    token,
					Path:     "/",
					HttpOnly: true,
				}

				http.SetCookie(w, cookie)
			}

			httpuser.SetUserID(r, userId)

			h.ServeHTTP(w, r)

			return nil
		})
	}
}

func GetUserID(key string, r *http.Request) (int32, error) {

	cookie, err := r.Cookie("Token")
	if err != nil {
		if !errors.Is(err, http.ErrNoCookie) {
			return 0, err
		}
		return 0, NewErrUserNotFound(err, "cookie not present")
	}
	token := cookie.Value

	userId, err := GetUserIDFromJWTToken(key, token)

	if err != nil {
		return 0, NewErrUserNotFound(err, "JWT token parse error")
	}

	return userId, nil
}
