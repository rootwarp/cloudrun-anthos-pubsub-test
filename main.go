package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.POST("/pubsub", pubsubHandler)
	e.Logger.Fatal(e.Start(":8080"))
}

func pubsubHandler(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "OK")
}
