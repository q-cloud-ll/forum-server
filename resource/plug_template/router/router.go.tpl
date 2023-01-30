package router

import (
	"forum-server/plugin/{{ .Snake}}/api"
	"github.com/gin-gonic/gin"
)

type {{ .PlugName}}Router struct {
}

func (s *{{ .PlugName}}Router) Init{{ .PlugName}}Router(Router *gin.RouterGroup) {
	plugRouter := Router
	plugApi := api.ApiGroupApp.{{ .PlugName}}Api
	{
		plugRouter.POST("routerName", plugApi.ApiName)
	}
}
