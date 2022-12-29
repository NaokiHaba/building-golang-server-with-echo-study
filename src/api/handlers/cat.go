package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"io/ioutil"
	"log"
	"net/http"
)

type Cat struct {
	Name string `json:"name" query:"name"`
	Type string `json:"type" query:"type"`
}

func GetCats(c echo.Context) error {
	dt := c.Param("data")

	ca := Cat{}

	if err := c.Bind(&ca); err != nil {
		return c.String(
			http.StatusBadRequest,
			fmt.Sprintf("Error is %s", err),
		)
	}

	if dt == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name": ca.Name,
			"type": ca.Type,
		})
	}

	if dt == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("your cat name is: %s\nand his type is: %s\n", ca.Name, ca.Type))
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "you need to lets us know if you want json or string data",
	})
}

func AddCat(c echo.Context) error {
	ca := Cat{}

	// レスポンスに関連するリソースを開放
	defer c.Request().Body.Close()

	// HTTPリクエストのボディを全て読み込む
	// 注意: ioutil.ReadAll関数を使用する前に、必ずdefer c.Request().Body.Close()を呼び出すようにしてください。
	// これは、リクエストのボディを閉じることで、リソースのリークを防ぐために必要です。
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("ReadAll Error：%s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	// 指定されたJSONデータを構造体や別のデータ型に変換
	err = json.Unmarshal(b, &ca)
	if err != nil {
		log.Printf("Unmarshal Error：%s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	log.Printf("animal is %#v", ca)
	return c.String(http.StatusOK, "Success")
}
