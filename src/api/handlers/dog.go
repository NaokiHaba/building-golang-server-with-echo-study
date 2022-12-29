package handlers

import (
	"encoding/json"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

type Dog struct {
	Name string `json:"name" query:"name"`
	Type string `json:"type" query:"type"`
}

func AddDog(c echo.Context) error {
	d := Dog{}

	// レスポンスに関連するリソースを開放
	defer c.Request().Body.Close()

	// json.NewDecoder関数を使用してデコーダを作成
	// そのデコーダを使用してJSONデータを読み込み
	// そのデータをaという構造体に変換
	if err := json.NewDecoder(c.Request().Body).Decode(&d); err != nil {
		log.Printf("Decode Error：%s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("animal is %#v", d)
	return c.String(http.StatusOK, "Success")
}
