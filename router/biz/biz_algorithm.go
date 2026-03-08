package biz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BizAlgorithmRouter struct{}

// InitBizAlgorithmRouter 初始化 推理加速算法 路由信息
func (s *BizAlgorithmRouter) InitBizAlgorithmRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	algorithmRouter := Router.Group("algorithm").Use(middleware.OperationRecord())
	algorithmRouterWithoutRecord := Router.Group("algorithm")
	{
		algorithmRouter.POST("createBizAlgorithm", algorithmApi.CreateBizAlgorithm)             // 新建推理加速算法
		algorithmRouter.DELETE("deleteBizAlgorithm", algorithmApi.DeleteBizAlgorithm)           // 删除推理加速算法
		algorithmRouter.DELETE("deleteBizAlgorithmByIds", algorithmApi.DeleteBizAlgorithmByIds) // 批量删除推理加速算法
		algorithmRouter.PUT("updateBizAlgorithm", algorithmApi.UpdateBizAlgorithm)              // 更新推理加速算法
	}
	{
		algorithmRouterWithoutRecord.GET("findBizAlgorithm", algorithmApi.FindBizAlgorithm)       // 根据ID获取推理加速算法
		algorithmRouterWithoutRecord.GET("getBizAlgorithmList", algorithmApi.GetBizAlgorithmList) // 获取推理加速算法列表
	}
}
