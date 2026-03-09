package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BizInferenceTaskSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	request.PageInfo
	OperatorId *int64
}

type BizInferenceTaskCreate struct {
	ModelName         *string  `json:"modelName" form:"modelName" binding:"required"`                 //模型名称
	ModelId           *int64   `json:"modelId" form:"modelId" binding:"required"`                     //模型Id
	AlgorithmName     *string  `json:"algorithmName" form:"algorithmName" binding:"required"`         //算法名称
	AlgorithmId       *int64   `json:"algorithmId" form:"algorithmId" binding:"required"`             //算法Id
	ModelType         *int64   `json:"modelType" form:"modelType" binding:"required"`                 //模型类型
	DatasetName       *string  `json:"datasetName" form:"datasetName" binding:"required"`             //数据集名称
	DatasetId         *int64   `json:"datasetId" form:"datasetId" binding:"required"`                 //数据集Id
	AverageThroughput *float64 `json:"averageThroughput" form:"averageThroughput" binding:"required"` //平均吞吐量
	AverageLatency    *float64 `json:"averageLatency" form:"averageLatency" binding:"required"`       //平均生成延迟
	AverageGpuMemory  *float64 `json:"averageGpuMemory" form:"averageGpuMemory" binding:"required"`   //平均GPU显存占用
}

type BizInferenceRank struct {
	request.PageInfo
	DatasetId *int64 `json:"datasetId" form:"datasetId" binding:"required"` //数据集Id
	PerfType  *int64 `json:"perfType" form:"perfType" binding:"required"`   //性能类型
}
