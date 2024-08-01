import service from '@/utils/request'

// @Tags BeeOrder
// @Summary 创建用户订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeOrder true "创建用户订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeOrder/createBeeOrder [post]
export const createBeeOrder = (data) => {
  return service({
    url: '/beeOrder/createBeeOrder',
    method: 'post',
    data
  })
}

// @Tags BeeOrder
// @Summary 删除用户订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeOrder true "删除用户订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeOrder/deleteBeeOrder [delete]
export const deleteBeeOrder = (params) => {
  return service({
    url: '/beeOrder/deleteBeeOrder',
    method: 'delete',
    params
  })
}

// @Tags BeeOrder
// @Summary 批量删除用户订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除用户订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeOrder/deleteBeeOrder [delete]
export const deleteBeeOrderByIds = (params) => {
  return service({
    url: '/beeOrder/deleteBeeOrderByIds',
    method: 'delete',
    params
  })
}

// @Tags BeeOrder
// @Summary 更新用户订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeOrder true "更新用户订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeOrder/updateBeeOrder [put]
export const updateBeeOrder = (data) => {
  return service({
    url: '/beeOrder/updateBeeOrder',
    method: 'put',
    data
  })
}

// @Tags BeeOrder
// @Summary 更新用户订单 ext json字段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeOrder true "更新用户订单ext json字段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeOrder/updateBeeOrderExtJsonStr [put]
export const updateBeeOrderExtJsonStr = (data) => {
  return service({
    url: '/beeOrder/updateBeeOrderExtJsonStr',
    method: 'put',
    data
  })
}

// @Tags BeeOrder
// @Summary 设置为已支付订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeOrder true "更新用户订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeOrder/markBeeOrderDone [put]
export const markBeeOrderDone = (data) => {
  return service({
    url: '/beeOrder/markBeeOrderDone',
    method: 'put',
    data
  })
}

// @Tags BeeOrder
// @Summary 设置为已完成订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeOrder true "更新用户订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeOrder/markBeeOrderDone [put]
export const markBeeOrderPaid = (data) => {
  return service({
    url: '/beeOrder/markBeeOrderPaid',
    method: 'put',
    data
  })
}

// @Tags BeeOrder
// @Summary 用id查询用户订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeOrder true "用id查询用户订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeOrder/findBeeOrder [get]
export const findBeeOrder = (params) => {
  return service({
    url: '/beeOrder/findBeeOrder',
    method: 'get',
    params
  })
}

// @Tags BeeOrder
// @Summary 分页获取用户订单列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取用户订单列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeOrder/getBeeOrderList [get]
export const getBeeOrderList = (params) => {
  return service({
    url: '/beeOrder/getBeeOrderList',
    method: 'get',
    params
  })
}
