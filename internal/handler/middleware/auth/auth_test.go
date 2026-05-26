package auth

import (
	"github.com/paulwwyvern/urlshortener/pkg/httphelpers/httpuser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

//go:generate mockgen -source=auth.go -destination=mock_auth.go -package=auth

func echoUser(w http.ResponseWriter, r *http.Request) {

	user := httpuser.GetUserID(r)

	w.WriteHeader(200)
	w.Write([]byte(strconv.Itoa(int(user))))
}

func TestAuth_Success(t *testing.T) {
	tests := []struct {
		name string
		key  string
		want int32
	}{
		{
			name: "Test #1 With Token",
			key:  "9oII0UJVsCAe2mRDGJ27",
			want: 123456,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			userService := NewMockUserService(ctrl)

			handler := WithAuth(tt.key, userService)(http.HandlerFunc(echoUser))

			r := httptest.NewRequest(http.MethodPost, "/", nil)
			w := httptest.NewRecorder()

			token, _ := CreateJWTToken(tt.key, tt.want)
			r.AddCookie(&http.Cookie{
				Name:  "Token",
				Value: token,
			})

			handler.ServeHTTP(w, r)

			assert.Equal(t, strconv.Itoa(int(tt.want)), w.Body.String())
		})
	}
}

func TestAuth_WithoutToken(t *testing.T) {
	tests := []struct {
		name string
		key  string
		want int32
	}{
		{
			name: "Test #1 Without Token",
			key:  "9oII0UJVsCAe2mRDGJ27",
			want: 123456,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			userService := NewMockUserService(ctrl)
			userService.EXPECT().CreateUser(gomock.Any()).Return(tt.want, nil)

			handler := WithAuth(tt.key, userService)(http.HandlerFunc(echoUser))

			r := httptest.NewRequest(http.MethodPost, "/", nil)
			w := httptest.NewRecorder()

			handler.ServeHTTP(w, r)

			cookies := w.Result().Cookies()
			var cookie *http.Cookie
			for _, c := range cookies {
				if c.Name == "Token" {
					cookie = c
				}
			}

			require.NotNil(t, cookie)

			token, _ := CreateJWTToken(tt.key, tt.want)
			assert.Equal(t, token, cookie.Value)
		})
	}
}

func TestAuthRequire_Success(t *testing.T) {
	tests := []struct {
		name string
		key  string
		want int32
	}{
		{
			name: "Test #1 With Token",
			key:  "9oII0UJVsCAe2mRDGJ27",
			want: 123456,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			handler := WithAuthRequire(tt.key)(http.HandlerFunc(echoUser))

			r := httptest.NewRequest(http.MethodPost, "/", nil)
			w := httptest.NewRecorder()

			token, _ := CreateJWTToken(tt.key, tt.want)
			r.AddCookie(&http.Cookie{
				Name:  "Token",
				Value: token,
			})

			handler.ServeHTTP(w, r)

			assert.Equal(t, strconv.Itoa(int(tt.want)), w.Body.String())
		})
	}
}

func TestAuthRequire_WithoutToken(t *testing.T) {
	tests := []struct {
		name string
		key  string
	}{
		{
			name: "Test #1 Without Token",
			key:  "9oII0UJVsCAe2mRDGJ27",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			handler := WithAuthRequire(tt.key)(http.HandlerFunc(echoUser))

			r := httptest.NewRequest(http.MethodPost, "/", nil)
			w := httptest.NewRecorder()

			handler.ServeHTTP(w, r)

			assert.Equal(t, http.StatusUnauthorized, w.Result().StatusCode)
		})
	}
}
