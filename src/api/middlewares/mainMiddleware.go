package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetMainMiddleware(e *echo.Echo) {
	// 静的ファイルを返す HTTP ハンドラーを作成
	// web アプリケーションから "/static/" パスにアクセスしたときに、
	// "/static/files" ディレクトリ内のファイルを返す HTTP ハンドラーを作成
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root: "../static",
	}))

	e.Use(serverHeader)
}

// ServerHeader Custom middleware https://echo.labstack.com/guide/context/
func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// HTTPレスポンスヘッダーに "Server: BlueBot/1.0" というエントリーを追加
		c.Response().Header().Set(echo.HeaderServer, "BlueBot/1.0")

		// 処理を次のミドルウェアに渡す
		return next(c)
	}
}
