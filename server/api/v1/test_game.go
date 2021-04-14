package v1

import (
	"gin-vue-admin/global"
    "gin-vue-admin/model"
    "gin-vue-admin/model/request"
    "gin-vue-admin/model/response"
    "gin-vue-admin/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

// @Tags TestGame
// @Summary 创建TestGame
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestGame true "创建TestGame"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testGame/createTestGame [post]
func CreateTestGame(c *gin.Context) {
	var testGame model.TestGame
	_ = c.ShouldBindJSON(&testGame)
	if err := service.CreateTestGame(testGame); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags TestGame
// @Summary 删除TestGame
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestGame true "删除TestGame"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /testGame/deleteTestGame [delete]
func DeleteTestGame(c *gin.Context) {
	var testGame model.TestGame
	_ = c.ShouldBindJSON(&testGame)
	if err := service.DeleteTestGame(testGame); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags TestGame
// @Summary 批量删除TestGame
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TestGame"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /testGame/deleteTestGameByIds [delete]
func DeleteTestGameByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteTestGameByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags TestGame
// @Summary 更新TestGame
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestGame true "更新TestGame"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /testGame/updateTestGame [put]
func UpdateTestGame(c *gin.Context) {
	var testGame model.TestGame
	_ = c.ShouldBindJSON(&testGame)
	if err := service.UpdateTestGame(testGame); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags TestGame
// @Summary 用id查询TestGame
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestGame true "用id查询TestGame"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /testGame/findTestGame [get]
func FindTestGame(c *gin.Context) {
	var testGame model.TestGame
	_ = c.ShouldBindQuery(&testGame)
	if err, retestGame := service.GetTestGame(testGame.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retestGame": retestGame}, c)
	}
}

// @Tags TestGame
// @Summary 分页获取TestGame列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.TestGameSearch true "分页获取TestGame列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testGame/getTestGameList [get]
func GetTestGameList(c *gin.Context) {
	var pageInfo request.TestGameSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetTestGameInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
