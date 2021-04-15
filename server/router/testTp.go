package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitTestTpRouter(Router *gin.RouterGroup) {
	TestTpRouter := Router.Group("testTp").Use(middleware.OperationRecord())
	{
		TestTpRouter.POST("createTestTp", v1.CreateTestTp)   // 新建TestTp
		TestTpRouter.DELETE("deleteTestTp", v1.DeleteTestTp) // 删除TestTp
		TestTpRouter.DELETE("deleteTestTpByIds", v1.DeleteTestTpByIds) // 批量删除TestTp
		TestTpRouter.PUT("updateTestTp", v1.UpdateTestTp)    // 更新TestTp
		TestTpRouter.GET("findTestTp", v1.FindTestTp)        // 根据ID获取TestTp
		TestTpRouter.GET("getTestTpList", v1.GetTestTpList)  // 获取TestTp列表
	}
}
