package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func hello(c echo.Context) error {
	//　指定されたステータスコードと文字列を返す
	return c.String(http.StatusOK, "Hello world")
}

type Cat struct {
	Name string `json:"name" query:"name"`
	Type string `json:"type" query:"type"`
}

func getCats(c echo.Context) error {
	var cat Cat
	dataType := c.Param("data")

	if err := c.Bind(&cat); err != nil {
		return c.String(
			http.StatusBadRequest,
			fmt.Sprintf("Error is %s", err),
		)
	}

	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name": cat.Name,
			"type": cat.Type,
		})
	}

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("name is %s\n type is %s\n", cat.Name, cat.Type))
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "データ種別が選択されていません",
	})
}

func main() {
	e := echo.New()

	e.GET("/", hello)
	e.GET("/cats/:data", getCats)

	// サーバー起動
	e.Logger.Fatal(e.Start(":1323"))
}
