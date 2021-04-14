import service from '@/utils/request'

// @Tags TestGame
// @Summary 创建TestGame
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestGame true "创建TestGame"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testGame/createTestGame [post]
export const createTestGame = (data) => {
     return service({
         url: "/testGame/createTestGame",
         method: 'post',
         data
     })
 }


// @Tags TestGame
// @Summary 删除TestGame
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestGame true "删除TestGame"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /testGame/deleteTestGame [delete]
 export const deleteTestGame = (data) => {
     return service({
         url: "/testGame/deleteTestGame",
         method: 'delete',
         data
     })
 }

// @Tags TestGame
// @Summary 删除TestGame
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TestGame"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /testGame/deleteTestGame [delete]
 export const deleteTestGameByIds = (data) => {
     return service({
         url: "/testGame/deleteTestGameByIds",
         method: 'delete',
         data
     })
 }

// @Tags TestGame
// @Summary 更新TestGame
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestGame true "更新TestGame"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /testGame/updateTestGame [put]
 export const updateTestGame = (data) => {
     return service({
         url: "/testGame/updateTestGame",
         method: 'put',
         data
     })
 }


// @Tags TestGame
// @Summary 用id查询TestGame
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestGame true "用id查询TestGame"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /testGame/findTestGame [get]
 export const findTestGame = (params) => {
     return service({
         url: "/testGame/findTestGame",
         method: 'get',
         params
     })
 }


// @Tags TestGame
// @Summary 分页获取TestGame列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取TestGame列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testGame/getTestGameList [get]
 export const getTestGameList = (params) => {
     return service({
         url: "/testGame/getTestGameList",
         method: 'get',
         params
     })
 }