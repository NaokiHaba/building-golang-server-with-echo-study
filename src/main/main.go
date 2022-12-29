package main

import (
	"crypto/subtle"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

func mainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "admin page")
}

func mainCookie(c echo.Context) error {
	return c.String(http.StatusOK, "cookie page")
}

func mainJwt(c echo.Context) error {
	user := c.Get("user")
	token := user.(*jwt.Token)
	claim := token.Claims.(jwt.MapClaims)
	log.Println("user Name", claim["name"].(string), "user ID", claim["jti"])

	return c.String(http.StatusOK, "Jwt page")
}

func login(c echo.Context) error {
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

type JwtClaims struct {
	Name string `json:"name"`

	// JWTのペイロード情報をまとめた構造体
	jwt.StandardClaims
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

// ServerHeader Custom middleware https://echo.labstack.com/guide/context/
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// HTTPレスポンスヘッダーに "Server: BlueBot/1.0" というエントリーを追加
		c.Response().Header().Set(echo.HeaderServer, "BlueBot/1.0")

		// 処理を次のミドルウェアに渡す
		return next(c)
	}
}

func checkCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("username")

		//　strings.Contains は、文字列が特定の文字列を含んでいるかどうかを判断する
		if strings.Contains(err.Error(), "named cookie not present") {
			return c.String(http.StatusUnauthorized, "not cookie")
		}

		if err != nil {
			return err
		}

		if cookie.Value == "jon" {
			return next(c)
		}

		return c.String(http.StatusOK, "check true")
	}
}

func main() {
	e := echo.New()

	e.Use(ServerHeader)

	// グループ化
	adminGroup := e.Group("/admin")
	cookieGroup := e.Group("/cookie")
	jwtGroup := e.Group("/jwt")

	// Echo アプリケーションで発生するリクエストやレスポンスに関する情報をログとして出力
	adminGroup.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		// リクエストの処理時刻、HTTP メソッド、リクエスト URI、HTTP ステータスコードを表す
		Format: "time=${time_rfc3339_nano}, method=${method}, uri=${uri}, status=${status}\n",
	}))

	adminGroup.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// 2つの文字列の長さに関係なく、常に同じ時間で比較
		// 長い文字列を比較するのにより多くの時間がかかる（タイムアタック攻撃を防ぐ）
		if subtle.ConstantTimeCompare([]byte(username), []byte("joe")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("secret")) == 1 {
			return true, nil
		}
		return false, nil
	}))

	//cookieGroup.Use(checkCookie)

	jwtGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		// JWT の署名を行うために使用する秘密鍵を保持
		SigningKey: []byte("AllYourBase"),

		// JWT署名アルゴリズム
		SigningMethod: "HS256",

		// JWT トークンを取得する方法を指定する
		TokenLookup: "cookie:JWTCookie",
	}))

	cookieGroup.GET("/main", mainCookie)
	adminGroup.GET("/main", mainAdmin)
	jwtGroup.GET("/main", mainJwt)

	cookieGroup.GET("/login", login)

	e.GET("/", hello)
	e.GET("/login", login)
	e.GET("/cats/:data", getCats)
	e.POST("/cats", addCat)
	e.POST("/dog", addDog)
	e.POST("/hamster", addHamster)

	// サーバー起動
	e.Logger.Fatal(e.Start(":1323"))
}
