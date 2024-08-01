import service from '@/utils/request'

// @Tags BeeCmsInfo
// @Summary 创建cms信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeCmsInfo true "创建cms信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeCmsInfo/createBeeCmsInfo [post]
export const createBeeCmsInfo = (data) => {
  return service({
    url: '/beeCmsInfo/createBeeCmsInfo',
    method: 'post',
    data
  })
}

// @Tags BeeCmsInfo
// @Summary 删除cms信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeCmsInfo true "删除cms信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeCmsInfo/deleteBeeCmsInfo [delete]
export const deleteBeeCmsInfo = (params) => {
  return service({
    url: '/beeCmsInfo/deleteBeeCmsInfo',
    method: 'delete',
    params
  })
}

// @Tags BeeCmsInfo
// @Summary 批量删除cms信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除cms信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeCmsInfo/deleteBeeCmsInfo [delete]
export const deleteBeeCmsInfoByIds = (params) => {
  return service({
    url: '/beeCmsInfo/deleteBeeCmsInfoByIds',
    method: 'delete',
    params
  })
}

// @Tags BeeCmsInfo
// @Summary 更新cms信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeCmsInfo true "更新cms信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeCmsInfo/updateBeeCmsInfo [put]
export const updateBeeCmsInfo = (data) => {
  return service({
    url: '/beeCmsInfo/updateBeeCmsInfo',
    method: 'put',
    data
  })
}

// @Tags BeeCmsInfo
// @Summary 用id查询cms信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeCmsInfo true "用id查询cms信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeCmsInfo/findBeeCmsInfo [get]
export const findBeeCmsInfo = (params) => {
  return service({
    url: '/beeCmsInfo/findBeeCmsInfo',
    method: 'get',
    params
  })
}

// @Tags BeeCmsInfo
// @Summary 分页获取cms信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取cms信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeCmsInfo/getBeeCmsInfoList [get]
export const getBeeCmsInfoList = (params) => {
  return service({
    url: '/beeCmsInfo/getBeeCmsInfoList',
    method: 'get',
    params
  })
}
