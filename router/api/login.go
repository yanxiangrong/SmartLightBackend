package api

import (
	"SmartLightBackend/models"
	"SmartLightBackend/pkg/logging"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(ctx *gin.Context) {
	type User struct {
		Id     string `json:"ip"`
		Passwd string `json:"passwd"`
	}

	var user User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		logging.Info("Context Bind Error ", err.Error())
		ctx.String(http.StatusBadRequest, "Request error")
		return
	}

	result := make(map[string]interface{})
	if models.Verify(models.User{Id: user.Id, Passwd: user.Passwd}) {
		result["status"] = "ok"
	} else {
		result["status"] = "fail"
	}

	ctx.JSON(http.StatusOK, result)
}
