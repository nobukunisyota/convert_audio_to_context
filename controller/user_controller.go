package controller

import (
	"net/http"
	"sample-api/model"
)

func Login() {
	user := model.Users{}
	return c.JSON(http.StatusOK, user)
}
