package biz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BizSampleRouter struct{}

// InitBizSampleRouter 初始化 数据集样本管理 路由信息
func (s *BizSampleRouter) InitBizSampleRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	sampleRouter := Router.Group("sample").Use(middleware.OperationRecord())
	sampleRouterWithoutRecord := Router.Group("sample")
	{
		sampleRouter.POST("createBizSample", sampleApi.CreateBizSample)             // 新建数据集样本管理
		sampleRouter.DELETE("deleteBizSample", sampleApi.DeleteBizSample)           // 删除数据集样本管理
		sampleRouter.DELETE("deleteBizSampleByIds", sampleApi.DeleteBizSampleByIds) // 批量删除数据集样本管理
		sampleRouter.PUT("updateBizSample", sampleApi.UpdateBizSample)              // 更新数据集样本管理
	}
	{
		sampleRouterWithoutRecord.GET("findBizSample", sampleApi.FindBizSample)       // 根据ID获取数据集样本管理
		sampleRouterWithoutRecord.GET("getBizSampleList", sampleApi.GetBizSampleList) // 获取数据集样本管理列表
	}
}
