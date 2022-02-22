package api

import (
	"SmartLightBackend/control"
	"SmartLightBackend/models"
	"SmartLightBackend/pkg/logging"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AllLamps(ctx *gin.Context) {
	lamps := make(map[string]interface{})
	lamps["devices"] = models.GetAllDevices()
	ctx.JSON(http.StatusOK, lamps)
}

// SetLamp 设置电灯状态（开或关）
func SetLamp(ctx *gin.Context) {
	type DeviceStatus struct {
		Addr   int  `json:"addr"`
		Switch bool `json:"switch"`
	}

	var status DeviceStatus
	err := ctx.ShouldBindJSON(&status)
	if err != nil {
		logging.Info("Context Bind Error ", err.Error())
		ctx.String(http.StatusBadRequest, "Request error")
		return
	}
	logging.Debug(fmt.Sprintf("Set %d %s", status.Addr, func() string {
		if status.Switch {
			return "ON"
		} else {
			return "OFF"
		}
	}()))
	control.SetStatus(status.Addr, status.Switch)
	ctx.Status(http.StatusOK)
}

// SetRemark 设置备注
func SetRemark(ctx *gin.Context) {
	type DeviceRemark struct {
		Addr   int    `json:"addr"`
		Remark string `json:"remark"`
	}

	var remark DeviceRemark
	err := ctx.ShouldBindJSON(&remark)
	if err != nil {
		logging.Info("Context Bind Error ", err.Error())
		ctx.String(http.StatusBadRequest, "Request error")
		return
	}
	logging.Debug(fmt.Sprintf("Set %d remark %s", remark.Addr, remark.Remark))

	models.UpdateDeviceWithoutTime(models.Device{Addr: remark.Addr, Remark: &remark.Remark})
	ctx.Status(http.StatusOK)
}

func DeleteLamp(ctx *gin.Context) {
	addrStr := ctx.Query("addr")
	var addr int
	_, err := fmt.Sscanf(addrStr, "%d", &addr)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Request error")
		return
	}

	models.DeleteDevice(models.Device{Addr: addr})
	ctx.String(http.StatusOK, "OK")
}

func OpenLamp(ctx *gin.Context) {
	addrStr := ctx.Query("addr")
	var addr int
	_, err := fmt.Sscanf(addrStr, "%d", &addr)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Request error")
		return
	}

	control.OpenLight(addr)
	ctx.String(http.StatusOK, "OK")
}

func CloseLamp(ctx *gin.Context) {
	addrStr := ctx.Query("addr")
	var addr int
	_, err := fmt.Sscanf(addrStr, "%d", &addr)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Request error")
		return
	}

	control.CloseLight(addr)
	ctx.String(http.StatusOK, "OK")
}

func OpenAll(ctx *gin.Context) {
	control.OpenAllLight()
	ctx.String(http.StatusOK, "OK")
}

func CloseAll(ctx *gin.Context) {
	control.CloseAllLight()
	ctx.String(http.StatusOK, "OK")
}
