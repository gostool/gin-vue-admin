package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitTestGameRouter(Router *gin.RouterGroup) {
	TestGameRouter := Router.Group("testGame").Use(middleware.OperationRecord())
	{
		TestGameRouter.POST("createTestGame", v1.CreateTestGame)   // 新建TestGame
		TestGameRouter.DELETE("deleteTestGame", v1.DeleteTestGame) // 删除TestGame
		TestGameRouter.DELETE("deleteTestGameByIds", v1.DeleteTestGameByIds) // 批量删除TestGame
		TestGameRouter.PUT("updateTestGame", v1.UpdateTestGame)    // 更新TestGame
		TestGameRouter.GET("findTestGame", v1.FindTestGame)        // 根据ID获取TestGame
		TestGameRouter.GET("getTestGameList", v1.GetTestGameList)  // 获取TestGame列表
	}
}
