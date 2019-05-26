package main

import (
	"github.com/sawadeeeen/go-training/app/lib"
	"net/http"

	"github.com/fgrosse/goldi"
	"github.com/labstack/echo"
	"github.com/sawadeeeen/go-training/app/service"
)

// 食べさせる環境変数の接頭語
//var envPrefix string = "TEST_"

func main() {
	registry := goldi.NewTypeRegistry()
	lib.RegisterTypes(registry)
	container := goldi.NewContainer(registry, nil)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		// サービスをとってきて、メソッドを叩いてみる
		service := container.MustGet("service").(service.Service)
		return c.String(http.StatusOK, service.Call("aaa"))
	})
	e.Logger.Fatal(e.Start(":1323"))
}
