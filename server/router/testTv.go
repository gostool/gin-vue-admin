package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitTestTvRouter(Router *gin.RouterGroup) {
	TestTvRouter := Router.Group("testTv").Use(middleware.OperationRecord())
	{
		TestTvRouter.POST("createTestTv", v1.CreateTestTv)   // 新建TestTv
		TestTvRouter.DELETE("deleteTestTv", v1.DeleteTestTv) // 删除TestTv
		TestTvRouter.DELETE("deleteTestTvByIds", v1.DeleteTestTvByIds) // 批量删除TestTv
		TestTvRouter.PUT("updateTestTv", v1.UpdateTestTv)    // 更新TestTv
		TestTvRouter.GET("findTestTv", v1.FindTestTv)        // 根据ID获取TestTv
		TestTvRouter.GET("getTestTvList", v1.GetTestTvList)  // 获取TestTv列表
	}
}
