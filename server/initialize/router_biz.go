package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
)

func holder(routers ...*gin.RouterGroup) {
	_ = routers
	_ = router.RouterGroupApp
}

func initBizRouter(routers ...*gin.RouterGroup) {
	publicGroup := routers[0]
	privateGroup := routers[1]
	{
		beeRouter := router.RouterGroupApp.Bee

		beeRouter.InitBeeUserMapperRouter(privateGroup, publicGroup)
		beeRouter.InitBeeShopGoodsCategoryRouter(privateGroup, publicGroup)
		beeRouter.InitBeeShopInfoRouter(privateGroup, publicGroup)
		beeRouter.InitBeeUserBalanceLogRouter(privateGroup, publicGroup)
		beeRouter.InitBeeCouponRouter(privateGroup, publicGroup)
		beeRouter.InitBeeUserCouponRouter(privateGroup, publicGroup)
		beeRouter.InitBeeCashLogRouter(privateGroup, publicGroup)
		beeRouter.InitBeeConfigRouter(privateGroup, publicGroup)
		beeRouter.InitBeeShopGoodsPropRouter(privateGroup, publicGroup)
		beeRouter.InitBeeShopGoodsAdditionRouter(privateGroup, publicGroup)
		beeRouter.InitBeeShopGoodsRouter(privateGroup, publicGroup)
		beeRouter.InitBeeShopGoodsImagesRouter(privateGroup, publicGroup)

		beeRouter.InitBeeShopGoodsSkuRouter(privateGroup, publicGroup)

		beeRouter.InitBeeShopGoodsContentRouter(privateGroup, publicGroup)
		beeRouter.InitBeeWxConfigRouter(privateGroup, publicGroup)

		beeRouter.InitBeeCommentRouter(privateGroup, publicGroup)
		beeRouter.InitBeeOrderRouter(privateGroup, publicGroup)
		beeRouter.InitBeeOrderReputationRouter(privateGroup, publicGroup)
		beeRouter.InitBeeBannerRouter(privateGroup, publicGroup)
		beeRouter.InitBeeUserAmountRouter(privateGroup, publicGroup)
		beeRouter.InitBeeCyTableRouter(privateGroup, publicGroup)

		beeRouter.InitBeeRegionRouter(privateGroup, publicGroup)
		beeRouter.InitBeeCmsInfoRouter(privateGroup, publicGroup)
		beeRouter.InitBeeUserAddressRouter(privateGroup, publicGroup)

		beeRouter.InitBeePeiSongRouter(privateGroup, publicGroup)
		beeRouter.InitBeeUserRouter(privateGroup, publicGroup)
		beeRouter.InitBeeOrderGoodsRouter(privateGroup, publicGroup)
		beeRouter.InitBeeOrderLogisticsRouter(privateGroup, publicGroup)
		beeRouter.InitBeeUploadFileRouter(privateGroup, publicGroup)
		beeRouter.InitBeePayLogRouter(privateGroup, publicGroup)
		beeRouter.InitBeeWxPayConfigRouter(privateGroup, publicGroup)
		beeRouter.InitBeeNoticeRouter(privateGroup, publicGroup)
		beeRouter.InitBeeLogisticsRouter(privateGroup, publicGroup)

		beeRouter.InitBeeOrderLogRouter(privateGroup, publicGroup)
	}

	holder(publicGroup, privateGroup)
}
