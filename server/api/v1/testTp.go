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

// @Tags TestTp
// @Summary 创建TestTp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestTp true "创建TestTp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testTp/createTestTp [post]
func CreateTestTp(c *gin.Context) {
	var testTp model.TestTp
	_ = c.ShouldBindJSON(&testTp)
	if err := service.CreateTestTp(testTp); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags TestTp
// @Summary 删除TestTp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestTp true "删除TestTp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /testTp/deleteTestTp [delete]
func DeleteTestTp(c *gin.Context) {
	var testTp model.TestTp
	_ = c.ShouldBindJSON(&testTp)
	if err := service.DeleteTestTp(testTp); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags TestTp
// @Summary 批量删除TestTp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TestTp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /testTp/deleteTestTpByIds [delete]
func DeleteTestTpByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteTestTpByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags TestTp
// @Summary 更新TestTp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestTp true "更新TestTp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /testTp/updateTestTp [put]
func UpdateTestTp(c *gin.Context) {
	var testTp model.TestTp
	_ = c.ShouldBindJSON(&testTp)
	if err := service.UpdateTestTp(testTp); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags TestTp
// @Summary 用id查询TestTp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestTp true "用id查询TestTp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /testTp/findTestTp [get]
func FindTestTp(c *gin.Context) {
	var testTp model.TestTp
	_ = c.ShouldBindQuery(&testTp)
	if err, retestTp := service.GetTestTp(testTp.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retestTp": retestTp}, c)
	}
}

// @Tags TestTp
// @Summary 分页获取TestTp列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.TestTpSearch true "分页获取TestTp列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testTp/getTestTpList [get]
func GetTestTpList(c *gin.Context) {
	var pageInfo request.TestTpSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetTestTpInfoList(pageInfo); err != nil {
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
