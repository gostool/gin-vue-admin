import service from '@/utils/request'

// @Tags TestTv
// @Summary 创建TestTv
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestTv true "创建TestTv"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testTv/createTestTv [post]
export const createTestTv = (data) => {
     return service({
         url: "/testTv/createTestTv",
         method: 'post',
         data
     })
 }


// @Tags TestTv
// @Summary 删除TestTv
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestTv true "删除TestTv"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /testTv/deleteTestTv [delete]
 export const deleteTestTv = (data) => {
     return service({
         url: "/testTv/deleteTestTv",
         method: 'delete',
         data
     })
 }

// @Tags TestTv
// @Summary 删除TestTv
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TestTv"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /testTv/deleteTestTv [delete]
 export const deleteTestTvByIds = (data) => {
     return service({
         url: "/testTv/deleteTestTvByIds",
         method: 'delete',
         data
     })
 }

// @Tags TestTv
// @Summary 更新TestTv
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestTv true "更新TestTv"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /testTv/updateTestTv [put]
 export const updateTestTv = (data) => {
     return service({
         url: "/testTv/updateTestTv",
         method: 'put',
         data
     })
 }


// @Tags TestTv
// @Summary 用id查询TestTv
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestTv true "用id查询TestTv"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /testTv/findTestTv [get]
 export const findTestTv = (params) => {
     return service({
         url: "/testTv/findTestTv",
         method: 'get',
         params
     })
 }


// @Tags TestTv
// @Summary 分页获取TestTv列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取TestTv列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testTv/getTestTvList [get]
 export const getTestTvList = (params) => {
     return service({
         url: "/testTv/getTestTvList",
         method: 'get',
         params
     })
 }