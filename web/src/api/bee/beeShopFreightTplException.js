import service from '@/utils/request'

// @Tags BeeShopFreightTplException
// @Summary 创建运费模版排除项
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeShopFreightTplException true "创建运费模版排除项"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeShopFreightTplException/createBeeShopFreightTplException [post]
export const createBeeShopFreightTplException = (data) => {
  return service({
    url: '/beeShopFreightTplException/createBeeShopFreightTplException',
    method: 'post',
    data
  })
}

// @Tags BeeShopFreightTplException
// @Summary 删除运费模版排除项
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeShopFreightTplException true "删除运费模版排除项"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeShopFreightTplException/deleteBeeShopFreightTplException [delete]
export const deleteBeeShopFreightTplException = (params) => {
  return service({
    url: '/beeShopFreightTplException/deleteBeeShopFreightTplException',
    method: 'delete',
    params
  })
}

// @Tags BeeShopFreightTplException
// @Summary 批量删除运费模版排除项
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除运费模版排除项"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeShopFreightTplException/deleteBeeShopFreightTplException [delete]
export const deleteBeeShopFreightTplExceptionByIds = (params) => {
  return service({
    url: '/beeShopFreightTplException/deleteBeeShopFreightTplExceptionByIds',
    method: 'delete',
    params
  })
}

// @Tags BeeShopFreightTplException
// @Summary 更新运费模版排除项
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeShopFreightTplException true "更新运费模版排除项"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeShopFreightTplException/updateBeeShopFreightTplException [put]
export const updateBeeShopFreightTplException = (data) => {
  return service({
    url: '/beeShopFreightTplException/updateBeeShopFreightTplException',
    method: 'put',
    data
  })
}

// @Tags BeeShopFreightTplException
// @Summary 用id查询运费模版排除项
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeShopFreightTplException true "用id查询运费模版排除项"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeShopFreightTplException/findBeeShopFreightTplException [get]
export const findBeeShopFreightTplException = (params) => {
  return service({
    url: '/beeShopFreightTplException/findBeeShopFreightTplException',
    method: 'get',
    params
  })
}

// @Tags BeeShopFreightTplException
// @Summary 分页获取运费模版排除项列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取运费模版排除项列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeShopFreightTplException/getBeeShopFreightTplExceptionList [get]
export const getBeeShopFreightTplExceptionList = (params) => {
  return service({
    url: '/beeShopFreightTplException/getBeeShopFreightTplExceptionList',
    method: 'get',
    params
  })
}
