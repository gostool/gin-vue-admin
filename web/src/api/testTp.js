import service from '@/utils/request'

// @Tags TestTp
// @Summary 创建TestTp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestTp true "创建TestTp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testTp/createTestTp [post]
export const createTestTp = (data) => {
     return service({
         url: "/testTp/createTestTp",
         method: 'post',
         data
     })
 }


// @Tags TestTp
// @Summary 删除TestTp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestTp true "删除TestTp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /testTp/deleteTestTp [delete]
 export const deleteTestTp = (data) => {
     return service({
         url: "/testTp/deleteTestTp",
         method: 'delete',
         data
     })
 }

// @Tags TestTp
// @Summary 删除TestTp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TestTp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /testTp/deleteTestTp [delete]
 export const deleteTestTpByIds = (data) => {
     return service({
         url: "/testTp/deleteTestTpByIds",
         method: 'delete',
         data
     })
 }

// @Tags TestTp
// @Summary 更新TestTp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestTp true "更新TestTp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /testTp/updateTestTp [put]
 export const updateTestTp = (data) => {
     return service({
         url: "/testTp/updateTestTp",
         method: 'put',
         data
     })
 }


// @Tags TestTp
// @Summary 用id查询TestTp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestTp true "用id查询TestTp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /testTp/findTestTp [get]
 export const findTestTp = (params) => {
     return service({
         url: "/testTp/findTestTp",
         method: 'get',
         params
     })
 }


// @Tags TestTp
// @Summary 分页获取TestTp列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取TestTp列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testTp/getTestTpList [get]
 export const getTestTpList = (params) => {
     return service({
         url: "/testTp/getTestTpList",
         method: 'get',
         params
     })
 }