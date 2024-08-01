import service from '@/utils/request'

// @Tags UserModels
// @Summary 创建用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserModels true "创建用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /userModels/createUserModels [post]
export const createUserModels = (data) => {
  return service({
    url: '/userModels/createUserModels',
    method: 'post',
    data
  })
}

// @Tags UserModels
// @Summary 删除用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserModels true "删除用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userModels/deleteUserModels [delete]
export const deleteUserModels = (params) => {
  return service({
    url: '/userModels/deleteUserModels',
    method: 'delete',
    params
  })
}

// @Tags UserModels
// @Summary 批量删除用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userModels/deleteUserModels [delete]
export const deleteUserModelsByIds = (params) => {
  return service({
    url: '/userModels/deleteUserModelsByIds',
    method: 'delete',
    params
  })
}

// @Tags UserModels
// @Summary 更新用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserModels true "更新用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userModels/updateUserModels [put]
export const updateUserModels = (data) => {
  return service({
    url: '/userModels/updateUserModels',
    method: 'put',
    data
  })
}

// @Tags UserModels
// @Summary 用id查询用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.UserModels true "用id查询用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userModels/findUserModels [get]
export const findUserModels = (params) => {
  return service({
    url: '/userModels/findUserModels',
    method: 'get',
    params
  })
}

// @Tags UserModels
// @Summary 分页获取用户信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取用户信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userModels/getUserModelsList [get]
export const getUserModelsList = (params) => {
  return service({
    url: '/userModels/getUserModelsList',
    method: 'get',
    params
  })
}
