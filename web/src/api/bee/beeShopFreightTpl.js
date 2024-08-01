import service from '@/utils/request'

// @Tags BeeShopFreightTpl
// @Summary 创建商城运费模版
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeShopFreightTpl true "创建商城运费模版"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeShopFreightTpl/createBeeShopFreightTpl [post]
export const createBeeShopFreightTpl = (data) => {
  return service({
    url: '/beeShopFreightTpl/createBeeShopFreightTpl',
    method: 'post',
    data
  })
}

// @Tags BeeShopFreightTpl
// @Summary 删除商城运费模版
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeShopFreightTpl true "删除商城运费模版"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeShopFreightTpl/deleteBeeShopFreightTpl [delete]
export const deleteBeeShopFreightTpl = (params) => {
  return service({
    url: '/beeShopFreightTpl/deleteBeeShopFreightTpl',
    method: 'delete',
    params
  })
}

// @Tags BeeShopFreightTpl
// @Summary 批量删除商城运费模版
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除商城运费模版"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeShopFreightTpl/deleteBeeShopFreightTpl [delete]
export const deleteBeeShopFreightTplByIds = (params) => {
  return service({
    url: '/beeShopFreightTpl/deleteBeeShopFreightTplByIds',
    method: 'delete',
    params
  })
}

// @Tags BeeShopFreightTpl
// @Summary 更新商城运费模版
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeShopFreightTpl true "更新商城运费模版"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeShopFreightTpl/updateBeeShopFreightTpl [put]
export const updateBeeShopFreightTpl = (data) => {
  return service({
    url: '/beeShopFreightTpl/updateBeeShopFreightTpl',
    method: 'put',
    data
  })
}

// @Tags BeeShopFreightTpl
// @Summary 用id查询商城运费模版
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeShopFreightTpl true "用id查询商城运费模版"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeShopFreightTpl/findBeeShopFreightTpl [get]
export const findBeeShopFreightTpl = (params) => {
  return service({
    url: '/beeShopFreightTpl/findBeeShopFreightTpl',
    method: 'get',
    params
  })
}

// @Tags BeeShopFreightTpl
// @Summary 分页获取商城运费模版列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取商城运费模版列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeShopFreightTpl/getBeeShopFreightTplList [get]
export const getBeeShopFreightTplList = (params) => {
  return service({
    url: '/beeShopFreightTpl/getBeeShopFreightTplList',
    method: 'get',
    params
  })
}
