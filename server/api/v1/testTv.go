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

// @Tags TestTv
// @Summary 创建TestTv
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestTv true "创建TestTv"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testTv/createTestTv [post]
func CreateTestTv(c *gin.Context) {
	var testTv model.TestTv
	_ = c.ShouldBindJSON(&testTv)
	if err := service.CreateTestTv(testTv); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags TestTv
// @Summary 删除TestTv
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestTv true "删除TestTv"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /testTv/deleteTestTv [delete]
func DeleteTestTv(c *gin.Context) {
	var testTv model.TestTv
	_ = c.ShouldBindJSON(&testTv)
	if err := service.DeleteTestTv(testTv); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags TestTv
// @Summary 批量删除TestTv
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TestTv"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /testTv/deleteTestTvByIds [delete]
func DeleteTestTvByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteTestTvByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags TestTv
// @Summary 更新TestTv
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestTv true "更新TestTv"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /testTv/updateTestTv [put]
func UpdateTestTv(c *gin.Context) {
	var testTv model.TestTv
	_ = c.ShouldBindJSON(&testTv)
	if err := service.UpdateTestTv(testTv); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags TestTv
// @Summary 用id查询TestTv
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestTv true "用id查询TestTv"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /testTv/findTestTv [get]
func FindTestTv(c *gin.Context) {
	var testTv model.TestTv
	_ = c.ShouldBindQuery(&testTv)
	if err, retestTv := service.GetTestTv(testTv.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retestTv": retestTv}, c)
	}
}

// @Tags TestTv
// @Summary 分页获取TestTv列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.TestTvSearch true "分页获取TestTv列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testTv/getTestTvList [get]
func GetTestTvList(c *gin.Context) {
	var pageInfo request.TestTvSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetTestTvInfoList(pageInfo); err != nil {
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
