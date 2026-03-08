package biz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BizModelRouter struct{}

// InitBizModelRouter 初始化 模型管理 路由信息
func (s *BizModelRouter) InitBizModelRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	modelRouter := Router.Group("model").Use(middleware.OperationRecord())
	modelRouterWithoutRecord := Router.Group("model")
	{
		modelRouter.POST("createBizModel", modelApi.CreateBizModel)             // 新建模型管理
		modelRouter.DELETE("deleteBizModel", modelApi.DeleteBizModel)           // 删除模型管理
		modelRouter.DELETE("deleteBizModelByIds", modelApi.DeleteBizModelByIds) // 批量删除模型管理
		modelRouter.PUT("updateBizModel", modelApi.UpdateBizModel)              // 更新模型管理
	}
	{
		modelRouterWithoutRecord.GET("findBizModel", modelApi.FindBizModel)       // 根据ID获取模型管理
		modelRouterWithoutRecord.GET("getBizModelList", modelApi.GetBizModelList) // 获取模型管理列表
	}
}
