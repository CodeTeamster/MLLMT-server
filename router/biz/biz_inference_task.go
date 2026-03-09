package biz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BizInferenceTaskRouter struct{}

// InitBizInferenceTaskRouter 初始化 推理任务记录 路由信息
func (s *BizInferenceTaskRouter) InitBizInferenceTaskRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	inferenceTaskRouter := Router.Group("inferenceTask").Use(middleware.OperationRecord())
	inferenceTaskRouterWithoutRecord := Router.Group("inferenceTask")
	{
		// TODO: runBizInferenceTask
		inferenceTaskRouter.POST("createBizInferenceTask", inferenceTaskApi.CreateBizInferenceTask)             // 新建推理任务记录
		inferenceTaskRouter.DELETE("deleteBizInferenceTask", inferenceTaskApi.DeleteBizInferenceTask)           // 删除推理任务记录
		inferenceTaskRouter.DELETE("deleteBizInferenceTaskByIds", inferenceTaskApi.DeleteBizInferenceTaskByIds) // 批量删除推理任务记录
	}
	{
		// TODO: getBizInferenceCompleteRecord
		inferenceTaskRouterWithoutRecord.GET("getBizInferenceTaskList", inferenceTaskApi.GetBizInferenceTaskList) // 获取推理任务记录列表
		inferenceTaskRouterWithoutRecord.GET("getBizInferenceRank", inferenceTaskApi.GetBizInferenceRank)         // 获取推理性能榜单
	}
}
