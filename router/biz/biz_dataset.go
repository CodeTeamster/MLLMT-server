package biz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BizDatasetRouter struct{}

// InitBizDatasetRouter 初始化 数据集管理 路由信息
func (s *BizDatasetRouter) InitBizDatasetRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	datasetRouter := Router.Group("dataset").Use(middleware.OperationRecord())
	datasetRouterWithoutRecord := Router.Group("dataset")
	{
		datasetRouter.POST("createBizDataset", datasetApi.CreateBizDataset)             // 新建数据集管理
		datasetRouter.DELETE("deleteBizDataset", datasetApi.DeleteBizDataset)           // 删除数据集管理
		datasetRouter.DELETE("deleteBizDatasetByIds", datasetApi.DeleteBizDatasetByIds) // 批量删除数据集管理
		datasetRouter.PUT("updateBizDataset", datasetApi.UpdateBizDataset)              // 更新数据集管理
	}
	{
		datasetRouterWithoutRecord.GET("findBizDataset", datasetApi.FindBizDataset)       // 根据ID获取数据集管理
		datasetRouterWithoutRecord.GET("getBizDatasetList", datasetApi.GetBizDatasetList) // 获取数据集管理列表
	}
}
