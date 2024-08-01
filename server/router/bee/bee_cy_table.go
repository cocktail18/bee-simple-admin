package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BeeCyTableRouter struct {}

// InitBeeCyTableRouter 初始化 桌号信息 路由信息
func (s *BeeCyTableRouter) InitBeeCyTableRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	beeCyTableRouter := Router.Group("beeCyTable").Use(middleware.OperationRecord())
	beeCyTableRouterWithoutRecord := Router.Group("beeCyTable")
	beeCyTableRouterWithoutAuth := PublicRouter.Group("beeCyTable")

	var beeCyTableApi = v1.ApiGroupApp.BeeApiGroup.BeeCyTableApi
	{
		beeCyTableRouter.POST("createBeeCyTable", beeCyTableApi.CreateBeeCyTable)   // 新建桌号信息
		beeCyTableRouter.DELETE("deleteBeeCyTable", beeCyTableApi.DeleteBeeCyTable) // 删除桌号信息
		beeCyTableRouter.DELETE("deleteBeeCyTableByIds", beeCyTableApi.DeleteBeeCyTableByIds) // 批量删除桌号信息
		beeCyTableRouter.PUT("updateBeeCyTable", beeCyTableApi.UpdateBeeCyTable)    // 更新桌号信息
	}
	{
		beeCyTableRouterWithoutRecord.GET("findBeeCyTable", beeCyTableApi.FindBeeCyTable)        // 根据ID获取桌号信息
		beeCyTableRouterWithoutRecord.GET("getBeeCyTableList", beeCyTableApi.GetBeeCyTableList)  // 获取桌号信息列表
	}
	{
	    beeCyTableRouterWithoutAuth.GET("getBeeCyTablePublic", beeCyTableApi.GetBeeCyTablePublic)  // 获取桌号信息列表
	}
}
