package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"io/ioutil"
	"log"
	"net/http"
)

func hello(c echo.Context) error {
	//　指定されたステータスコードと文字列を返す
	return c.String(http.StatusOK, "Hello world")
}

type Animal struct {
	Name string `json:"name" query:"name"`
	Type string `json:"type" query:"type"`
}

func getCats(c echo.Context) error {
	var a Animal
	dataType := c.Param("data")

	if err := c.Bind(&a); err != nil {
		return c.String(
			http.StatusBadRequest,
			fmt.Sprintf("Error is %s", err),
		)
	}

	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name": a.Name,
			"type": a.Type,
		})
	}

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("name is %s\n type is %s\n", a.Name, a.Type))
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "データ種別が選択されていません",
	})
}

func addCat(c echo.Context) error {
	a := Animal{}

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
	err = json.Unmarshal(b, &a)
	if err != nil {
		log.Printf("Unmarshal Error：%s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	log.Printf("animal is %#v", a)
	return c.String(http.StatusOK, "Success")
}

func addDog(c echo.Context) error {
	a := Animal{}

	// レスポンスに関連するリソースを開放
	defer c.Request().Body.Close()

	// json.NewDecoder関数を使用してデコーダを作成
	// そのデコーダを使用してJSONデータを読み込み
	// そのデータをaという構造体に変換
	if err := json.NewDecoder(c.Request().Body).Decode(&a); err != nil {
		log.Printf("Decode Error：%s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("animal is %#v", a)
	return c.String(http.StatusOK, "Success")
}

func addHamster(c echo.Context) error {
	a := Animal{}

	// HTTPリクエストのボディを構造体や別のデータ型にマッピングする
	if err := c.Bind(&a); err != nil {
		log.Printf("Bind Error：%s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("animal is %#v", a)
	return c.String(http.StatusOK, "Success")
}

func main() {
	e := echo.New()

	e.GET("/", hello)
	e.GET("/cats/:data", getCats)
	e.POST("/cats", addCat)
	e.POST("/dog", addDog)
	e.POST("/hamster", addHamster)

	// サーバー起動
	e.Logger.Fatal(e.Start(":1323"))
}
