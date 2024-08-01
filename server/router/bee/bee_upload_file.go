package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BeeUploadFileRouter struct {}

// InitBeeUploadFileRouter 初始化 用户上传文件 路由信息
func (s *BeeUploadFileRouter) InitBeeUploadFileRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	beeUploadFileRouter := Router.Group("beeUploadFile").Use(middleware.OperationRecord())
	beeUploadFileRouterWithoutRecord := Router.Group("beeUploadFile")
	beeUploadFileRouterWithoutAuth := PublicRouter.Group("beeUploadFile")

	var beeUploadFileApi = v1.ApiGroupApp.BeeApiGroup.BeeUploadFileApi
	{
		beeUploadFileRouter.POST("createBeeUploadFile", beeUploadFileApi.CreateBeeUploadFile)   // 新建用户上传文件
		beeUploadFileRouter.DELETE("deleteBeeUploadFile", beeUploadFileApi.DeleteBeeUploadFile) // 删除用户上传文件
		beeUploadFileRouter.DELETE("deleteBeeUploadFileByIds", beeUploadFileApi.DeleteBeeUploadFileByIds) // 批量删除用户上传文件
		beeUploadFileRouter.PUT("updateBeeUploadFile", beeUploadFileApi.UpdateBeeUploadFile)    // 更新用户上传文件
	}
	{
		beeUploadFileRouterWithoutRecord.GET("findBeeUploadFile", beeUploadFileApi.FindBeeUploadFile)        // 根据ID获取用户上传文件
		beeUploadFileRouterWithoutRecord.GET("getBeeUploadFileList", beeUploadFileApi.GetBeeUploadFileList)  // 获取用户上传文件列表
	}
	{
	    beeUploadFileRouterWithoutAuth.GET("getBeeUploadFilePublic", beeUploadFileApi.GetBeeUploadFilePublic)  // 获取用户上传文件列表
	}
}
