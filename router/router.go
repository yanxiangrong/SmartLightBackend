package router

import (
	"SmartLightBackend/router/api"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default()) // 使用跨域中间件

	rApi := r.Group("/api")
	{
		rApi.GET("/ping", func(context *gin.Context) {
			context.String(http.StatusOK, "pong")
		})
		rApi.GET("/allDevices", api.AllLamps) // 获取所有灯，
		// {"devices":[{"addr":26709,"current":2,"switch":true,"fault":false,"update_at":"1636108091","remark":"办公室电灯"},
		//			   {"addr":23125,"current":2,"switch":true,"fault":false,"update_at":"1634116076","remark":""}]}
		rApi.POST("/setDevices", api.SetLamp)     // 设置灯开或关，json {"addr": 26709, "switch": true}
		rApi.GET("/openAll", api.OpenAll)         // 打开所有灯
		rApi.GET("/closeAll", api.CloseAll)       // 关闭所有灯
		rApi.GET("/open", api.OpenLamp)           // 开灯，url参数 /open?addr=26709
		rApi.GET("/close", api.CloseLamp)         // 关灯，url参数 /open?addr=26709
		rApi.GET("/deleteDevice", api.DeleteLamp) // 删除设备，url参数 /open?addr=26709
		rApi.POST("/setRemark", api.SetRemark)    // 设置备注，json {"addr": 26709, "remark", "办公室"}
	}

	return r
}
