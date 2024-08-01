import service from '@/utils/request'

// @Tags BeeShopGoodsAdditionItem
// @Summary 创建商品额外信息绑定
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeShopGoodsAdditionItem true "创建商品额外信息绑定"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeShopGoodsAdditionItem/createBeeShopGoodsAdditionItem [post]
export const createBeeShopGoodsAdditionItem = (data) => {
  return service({
    url: '/beeShopGoodsAdditionItem/createBeeShopGoodsAdditionItem',
    method: 'post',
    data
  })
}

// @Tags BeeShopGoodsAdditionItem
// @Summary 删除商品额外信息绑定
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeShopGoodsAdditionItem true "删除商品额外信息绑定"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeShopGoodsAdditionItem/deleteBeeShopGoodsAdditionItem [delete]
export const deleteBeeShopGoodsAdditionItem = (params) => {
  return service({
    url: '/beeShopGoodsAdditionItem/deleteBeeShopGoodsAdditionItem',
    method: 'delete',
    params
  })
}

// @Tags BeeShopGoodsAdditionItem
// @Summary 批量删除商品额外信息绑定
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除商品额外信息绑定"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeShopGoodsAdditionItem/deleteBeeShopGoodsAdditionItem [delete]
export const deleteBeeShopGoodsAdditionItemByIds = (params) => {
  return service({
    url: '/beeShopGoodsAdditionItem/deleteBeeShopGoodsAdditionItemByIds',
    method: 'delete',
    params
  })
}

// @Tags BeeShopGoodsAdditionItem
// @Summary 更新商品额外信息绑定
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeShopGoodsAdditionItem true "更新商品额外信息绑定"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeShopGoodsAdditionItem/updateBeeShopGoodsAdditionItem [put]
export const updateBeeShopGoodsAdditionItem = (data) => {
  return service({
    url: '/beeShopGoodsAdditionItem/updateBeeShopGoodsAdditionItem',
    method: 'put',
    data
  })
}

// @Tags BeeShopGoodsAdditionItem
// @Summary 用id查询商品额外信息绑定
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeShopGoodsAdditionItem true "用id查询商品额外信息绑定"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeShopGoodsAdditionItem/findBeeShopGoodsAdditionItem [get]
export const findBeeShopGoodsAdditionItem = (params) => {
  return service({
    url: '/beeShopGoodsAdditionItem/findBeeShopGoodsAdditionItem',
    method: 'get',
    params
  })
}

// @Tags BeeShopGoodsAdditionItem
// @Summary 分页获取商品额外信息绑定列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取商品额外信息绑定列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeShopGoodsAdditionItem/getBeeShopGoodsAdditionItemList [get]
export const getBeeShopGoodsAdditionItemList = (params) => {
  return service({
    url: '/beeShopGoodsAdditionItem/getBeeShopGoodsAdditionItemList',
    method: 'get',
    params
  })
}
