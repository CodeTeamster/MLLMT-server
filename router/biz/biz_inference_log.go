package biz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BizInferenceLogRouter struct{}

// InitBizInferenceLogRouter 初始化 推理详细记录 路由信息
func (s *BizInferenceLogRouter) InitBizInferenceLogRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	inferenceLogRouter := Router.Group("inferenceLog").Use(middleware.OperationRecord())
	inferenceLogRouterWithoutRecord := Router.Group("inferenceLog")
	{
		inferenceLogRouter.POST("createBizInferenceLog", inferenceLogApi.CreateBizInferenceLog) // 新建推理详细记录
	}
	{
		inferenceLogRouterWithoutRecord.GET("getBizInferenceLogSample", inferenceLogApi.GetBizInferenceLogSample) // 分页获取推理详细记录列表
		inferenceLogRouterWithoutRecord.GET("getBizInferenceLogList", inferenceLogApi.GetBizInferenceLogList)     // 获取推理详细记录列表
	}
}
