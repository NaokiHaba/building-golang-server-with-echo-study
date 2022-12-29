package middlewares

import (
	"crypto/subtle"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetAdminMiddlewares(g *echo.Group) {
	// Logger
	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		// リクエストの処理時刻、HTTP メソッド、リクエスト URI、HTTP ステータスコードを表す
		Format: "time=${time_rfc3339_nano}, method=${method}, uri=${uri}, status=${status}\n",
	}))

	// Basic Auth
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// 2つの文字列の長さに関係なく、常に同じ時間で比較
		// 長い文字列を比較するのにより多くの時間がかかる（タイムアタック攻撃を防ぐ）
		if subtle.ConstantTimeCompare([]byte(username), []byte("joe")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("secret")) == 1 {
			return true, nil
		}
		return false, nil
	}))
}
