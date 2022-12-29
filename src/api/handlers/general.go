package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

func Hello(c echo.Context) error {
	//　指定されたステータスコードと文字列を返す
	return c.String(http.StatusOK, "Hello world")
}
