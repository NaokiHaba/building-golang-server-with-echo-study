package main

import (
	"github.com/naokis-practice-project/practice_echo_example/src/router"
)

func main() {

	e := router.New()

	// サーバー起動
	e.Logger.Fatal(e.Start(":1323"))
}
