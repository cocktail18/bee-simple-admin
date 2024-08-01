import service from '@/utils/request'

// @Tags BeeLogisticsDetail
// @Summary 创建运费模版详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeLogisticsDetail true "创建运费模版详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeLogisticsDetail/createBeeLogisticsDetail [post]
export const createBeeLogisticsDetail = (data) => {
  return service({
    url: '/beeLogisticsDetail/createBeeLogisticsDetail',
    method: 'post',
    data
  })
}

// @Tags BeeLogisticsDetail
// @Summary 删除运费模版详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeLogisticsDetail true "删除运费模版详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeLogisticsDetail/deleteBeeLogisticsDetail [delete]
export const deleteBeeLogisticsDetail = (params) => {
  return service({
    url: '/beeLogisticsDetail/deleteBeeLogisticsDetail',
    method: 'delete',
    params
  })
}

// @Tags BeeLogisticsDetail
// @Summary 批量删除运费模版详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除运费模版详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeLogisticsDetail/deleteBeeLogisticsDetail [delete]
export const deleteBeeLogisticsDetailByIds = (params) => {
  return service({
    url: '/beeLogisticsDetail/deleteBeeLogisticsDetailByIds',
    method: 'delete',
    params
  })
}

// @Tags BeeLogisticsDetail
// @Summary 更新运费模版详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeLogisticsDetail true "更新运费模版详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeLogisticsDetail/updateBeeLogisticsDetail [put]
export const updateBeeLogisticsDetail = (data) => {
  return service({
    url: '/beeLogisticsDetail/updateBeeLogisticsDetail',
    method: 'put',
    data
  })
}

// @Tags BeeLogisticsDetail
// @Summary 用id查询运费模版详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeLogisticsDetail true "用id查询运费模版详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeLogisticsDetail/findBeeLogisticsDetail [get]
export const findBeeLogisticsDetail = (params) => {
  return service({
    url: '/beeLogisticsDetail/findBeeLogisticsDetail',
    method: 'get',
    params
  })
}

// @Tags BeeLogisticsDetail
// @Summary 分页获取运费模版详情列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取运费模版详情列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeLogisticsDetail/getBeeLogisticsDetailList [get]
export const getBeeLogisticsDetailList = (params) => {
  return service({
    url: '/beeLogisticsDetail/getBeeLogisticsDetailList',
    method: 'get',
    params
  })
}
