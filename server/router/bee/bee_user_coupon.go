package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BeeUserCouponRouter struct {}

// InitBeeUserCouponRouter 初始化 用户优惠券 路由信息
func (s *BeeUserCouponRouter) InitBeeUserCouponRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	beeUserCouponRouter := Router.Group("beeUserCoupon").Use(middleware.OperationRecord())
	beeUserCouponRouterWithoutRecord := Router.Group("beeUserCoupon")
	beeUserCouponRouterWithoutAuth := PublicRouter.Group("beeUserCoupon")

	var beeUserCouponApi = v1.ApiGroupApp.BeeApiGroup.BeeUserCouponApi
	{
		beeUserCouponRouter.POST("createBeeUserCoupon", beeUserCouponApi.CreateBeeUserCoupon)   // 新建用户优惠券
		beeUserCouponRouter.DELETE("deleteBeeUserCoupon", beeUserCouponApi.DeleteBeeUserCoupon) // 删除用户优惠券
		beeUserCouponRouter.DELETE("deleteBeeUserCouponByIds", beeUserCouponApi.DeleteBeeUserCouponByIds) // 批量删除用户优惠券
		beeUserCouponRouter.PUT("updateBeeUserCoupon", beeUserCouponApi.UpdateBeeUserCoupon)    // 更新用户优惠券
	}
	{
		beeUserCouponRouterWithoutRecord.GET("findBeeUserCoupon", beeUserCouponApi.FindBeeUserCoupon)        // 根据ID获取用户优惠券
		beeUserCouponRouterWithoutRecord.GET("getBeeUserCouponList", beeUserCouponApi.GetBeeUserCouponList)  // 获取用户优惠券列表
	}
	{
	    beeUserCouponRouterWithoutAuth.GET("getBeeUserCouponPublic", beeUserCouponApi.GetBeeUserCouponPublic)  // 获取用户优惠券列表
	}
}
