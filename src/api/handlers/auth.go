package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"time"
)

type JwtClaims struct {
	Name string `json:"name"`

	// JWTのペイロード情報をまとめた構造体
	jwt.StandardClaims
}

func Login(c echo.Context) error {
	name := c.QueryParam("name")
	password := c.QueryParam("password")

	if name == "name" && password == "password" {
		// Create a Cookie
		cookie := new(http.Cookie)
		cookie.Name = "username"
		cookie.Value = "jon"
		cookie.Expires = time.Now().Add(24 * time.Hour)
		c.SetCookie(cookie)

		token, err := createJwtToken()

		if err != nil {
			log.Println("Error creating jwt token", err)
			return c.String(http.StatusInternalServerError, "something went wrong")
		}

		jwtCookie := &http.Cookie{}
		jwtCookie.Name = "JWTCookie"
		jwtCookie.Value = token
		jwtCookie.Expires = time.Now().Add(48 * time.Hour)
		c.SetCookie(jwtCookie)

		return c.JSON(http.StatusOK, map[string]string{
			"message": "log in success",
			"token":   token,
		})
	}

	return c.String(http.StatusUnauthorized, "filed login")
}

// https://github.com/dgrijalva/jwt-go/blob/master/example_test.go#L31-L53
func createJwtToken() (string, error) {
	mySigningKey := []byte("AllYourBase")

	// Create the Claims
	claims := JwtClaims{
		"bar",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			Id:        "user_id",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return ss, nil
}
