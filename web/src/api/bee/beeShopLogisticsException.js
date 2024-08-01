import service from '@/utils/request'

// @Tags BeeShopLogisticsException
// @Summary 创建运费排除项
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeShopLogisticsException true "创建运费排除项"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeShopLogisticsException/createBeeShopLogisticsException [post]
export const createBeeShopLogisticsException = (data) => {
  return service({
    url: '/beeShopLogisticsException/createBeeShopLogisticsException',
    method: 'post',
    data
  })
}

// @Tags BeeShopLogisticsException
// @Summary 删除运费排除项
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeShopLogisticsException true "删除运费排除项"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeShopLogisticsException/deleteBeeShopLogisticsException [delete]
export const deleteBeeShopLogisticsException = (params) => {
  return service({
    url: '/beeShopLogisticsException/deleteBeeShopLogisticsException',
    method: 'delete',
    params
  })
}

// @Tags BeeShopLogisticsException
// @Summary 批量删除运费排除项
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除运费排除项"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeShopLogisticsException/deleteBeeShopLogisticsException [delete]
export const deleteBeeShopLogisticsExceptionByIds = (params) => {
  return service({
    url: '/beeShopLogisticsException/deleteBeeShopLogisticsExceptionByIds',
    method: 'delete',
    params
  })
}

// @Tags BeeShopLogisticsException
// @Summary 更新运费排除项
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeShopLogisticsException true "更新运费排除项"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeShopLogisticsException/updateBeeShopLogisticsException [put]
export const updateBeeShopLogisticsException = (data) => {
  return service({
    url: '/beeShopLogisticsException/updateBeeShopLogisticsException',
    method: 'put',
    data
  })
}

// @Tags BeeShopLogisticsException
// @Summary 用id查询运费排除项
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeShopLogisticsException true "用id查询运费排除项"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeShopLogisticsException/findBeeShopLogisticsException [get]
export const findBeeShopLogisticsException = (params) => {
  return service({
    url: '/beeShopLogisticsException/findBeeShopLogisticsException',
    method: 'get',
    params
  })
}

// @Tags BeeShopLogisticsException
// @Summary 分页获取运费排除项列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取运费排除项列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeShopLogisticsException/getBeeShopLogisticsExceptionList [get]
export const getBeeShopLogisticsExceptionList = (params) => {
  return service({
    url: '/beeShopLogisticsException/getBeeShopLogisticsExceptionList',
    method: 'get',
    params
  })
}
