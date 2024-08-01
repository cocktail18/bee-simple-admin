package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/bee"
    beeReq "github.com/flipped-aurora/gin-vue-admin/server/model/bee/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type BeeOrderGoodsApi struct {}

var beeOrderGoodsService = service.ServiceGroupApp.BeeServiceGroup.BeeOrderGoodsService


// CreateBeeOrderGoods 创建订单商品信息
// @Tags BeeOrderGoods
// @Summary 创建订单商品信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeOrderGoods true "创建订单商品信息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beeOrderGoods/createBeeOrderGoods [post]
func (beeOrderGoodsApi *BeeOrderGoodsApi) CreateBeeOrderGoods(c *gin.Context) {
	var beeOrderGoods bee.BeeOrderGoods
	err := c.ShouldBindJSON(&beeOrderGoods)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    shopUserId := int(utils.GetShopUserID(c))
	beeOrderGoods.UserId = &shopUserId
	if err := beeOrderGoodsService.CreateBeeOrderGoods(&beeOrderGoods); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeeOrderGoods 删除订单商品信息
// @Tags BeeOrderGoods
// @Summary 删除订单商品信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeOrderGoods true "删除订单商品信息"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beeOrderGoods/deleteBeeOrderGoods [delete]
func (beeOrderGoodsApi *BeeOrderGoodsApi) DeleteBeeOrderGoods(c *gin.Context) {
	id := c.Query("id")
    shopUserId := int(utils.GetShopUserID(c))
	if err := beeOrderGoodsService.DeleteBeeOrderGoods(id, shopUserId); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeeOrderGoodsByIds 批量删除订单商品信息
// @Tags BeeOrderGoods
// @Summary 批量删除订单商品信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beeOrderGoods/deleteBeeOrderGoodsByIds [delete]
func (beeOrderGoodsApi *BeeOrderGoodsApi) DeleteBeeOrderGoodsByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
    shopUserId := int(utils.GetShopUserID(c))
	if err := beeOrderGoodsService.DeleteBeeOrderGoodsByIds(ids, shopUserId); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeeOrderGoods 更新订单商品信息
// @Tags BeeOrderGoods
// @Summary 更新订单商品信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeOrderGoods true "更新订单商品信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeOrderGoods/updateBeeOrderGoods [put]
func (beeOrderGoodsApi *BeeOrderGoodsApi) UpdateBeeOrderGoods(c *gin.Context) {
	var beeOrderGoods bee.BeeOrderGoods
	err := c.ShouldBindJSON(&beeOrderGoods)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    shopUserId := int(utils.GetShopUserID(c))
    beeOrderGoods.UserId = &shopUserId
	if err := beeOrderGoodsService.UpdateBeeOrderGoods(beeOrderGoods, shopUserId); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBeeOrderGoods 用id查询订单商品信息
// @Tags BeeOrderGoods
// @Summary 用id查询订单商品信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeeOrderGoods true "用id查询订单商品信息"
// @Success 200 {object} response.Response{data=object{rebeeOrderGoods=bee.BeeOrderGoods},msg=string} "查询成功"
// @Router /beeOrderGoods/findBeeOrderGoods [get]
func (beeOrderGoodsApi *BeeOrderGoodsApi) FindBeeOrderGoods(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if rebeeOrderGoods, err := beeOrderGoodsService.GetBeeOrderGoods(id, shopUserId); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeeOrderGoods, c)
	}
}

// GetBeeOrderGoodsList 分页获取订单商品信息列表
// @Tags BeeOrderGoods
// @Summary 分页获取订单商品信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeOrderGoodsSearch true "分页获取订单商品信息列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeOrderGoods/getBeeOrderGoodsList [get]
func (beeOrderGoodsApi *BeeOrderGoodsApi) GetBeeOrderGoodsList(c *gin.Context) {
	var pageInfo beeReq.BeeOrderGoodsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if list, total, err := beeOrderGoodsService.GetBeeOrderGoodsInfoList(pageInfo, shopUserId); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}

// GetBeeOrderGoodsPublic 不需要鉴权的订单商品信息接口
// @Tags BeeOrderGoods
// @Summary 不需要鉴权的订单商品信息接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeOrderGoodsSearch true "分页获取订单商品信息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beeOrderGoods/getBeeOrderGoodsPublic [get]
func (beeOrderGoodsApi *BeeOrderGoodsApi) GetBeeOrderGoodsPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的订单商品信息接口信息",
    }, "获取成功", c)
}
