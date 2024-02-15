package main

import (
	"net/http"
	"sample-api/model"

	"github.com/labstack/echo"
)

func index(c echo.Context) error {
	db, _ := model.DB.DB()
	err := db.Ping()
	if err != nil {
		return c.String(http.StatusInternalServerError, "db connection error")
	} else {
		return c.String(http.StatusOK, "db connection success")
	}
}
func main() {
	router := newRouter()
	router.GET("/", index)
	router.Logger.Fatal(router.Start(":8001"))
}
