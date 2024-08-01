import service from '@/utils/request'

// @Tags BeeShopFreightTplFreeShipping
// @Summary 创建运费模版免运费配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeShopFreightTplFreeShipping true "创建运费模版免运费配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeShopFreightTplFreeShipping/createBeeShopFreightTplFreeShipping [post]
export const createBeeShopFreightTplFreeShipping = (data) => {
  return service({
    url: '/beeShopFreightTplFreeShipping/createBeeShopFreightTplFreeShipping',
    method: 'post',
    data
  })
}

// @Tags BeeShopFreightTplFreeShipping
// @Summary 删除运费模版免运费配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeShopFreightTplFreeShipping true "删除运费模版免运费配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeShopFreightTplFreeShipping/deleteBeeShopFreightTplFreeShipping [delete]
export const deleteBeeShopFreightTplFreeShipping = (params) => {
  return service({
    url: '/beeShopFreightTplFreeShipping/deleteBeeShopFreightTplFreeShipping',
    method: 'delete',
    params
  })
}

// @Tags BeeShopFreightTplFreeShipping
// @Summary 批量删除运费模版免运费配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除运费模版免运费配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeShopFreightTplFreeShipping/deleteBeeShopFreightTplFreeShipping [delete]
export const deleteBeeShopFreightTplFreeShippingByIds = (params) => {
  return service({
    url: '/beeShopFreightTplFreeShipping/deleteBeeShopFreightTplFreeShippingByIds',
    method: 'delete',
    params
  })
}

// @Tags BeeShopFreightTplFreeShipping
// @Summary 更新运费模版免运费配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeShopFreightTplFreeShipping true "更新运费模版免运费配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeShopFreightTplFreeShipping/updateBeeShopFreightTplFreeShipping [put]
export const updateBeeShopFreightTplFreeShipping = (data) => {
  return service({
    url: '/beeShopFreightTplFreeShipping/updateBeeShopFreightTplFreeShipping',
    method: 'put',
    data
  })
}

// @Tags BeeShopFreightTplFreeShipping
// @Summary 用id查询运费模版免运费配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeShopFreightTplFreeShipping true "用id查询运费模版免运费配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeShopFreightTplFreeShipping/findBeeShopFreightTplFreeShipping [get]
export const findBeeShopFreightTplFreeShipping = (params) => {
  return service({
    url: '/beeShopFreightTplFreeShipping/findBeeShopFreightTplFreeShipping',
    method: 'get',
    params
  })
}

// @Tags BeeShopFreightTplFreeShipping
// @Summary 分页获取运费模版免运费配置列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取运费模版免运费配置列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeShopFreightTplFreeShipping/getBeeShopFreightTplFreeShippingList [get]
export const getBeeShopFreightTplFreeShippingList = (params) => {
  return service({
    url: '/beeShopFreightTplFreeShipping/getBeeShopFreightTplFreeShippingList',
    method: 'get',
    params
  })
}
