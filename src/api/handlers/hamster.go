package handlers

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
)

type Hamster struct {
	Name string `json:"name" query:"name"`
	Type string `json:"type" query:"type"`
}

func AddHamster(c echo.Context) error {
	h := Hamster{}

	// HTTPリクエストのボディを構造体や別のデータ型にマッピングする
	if err := c.Bind(&h); err != nil {
		log.Printf("Bind Error：%s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("animal is %#v", h)
	return c.String(http.StatusOK, "Success")
}
