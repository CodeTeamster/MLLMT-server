// 自动生成模板BizInferenceTask
package biz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 推理任务记录 结构体  BizInferenceTask
type BizInferenceTask struct {
	global.GVA_MODEL
	ModelName         *string  `json:"modelName" form:"modelName" gorm:"comment:模型名称;column:model_name;" binding:"required"`                                 //模型名称
	ModelId           *int64   `json:"modelId" form:"modelId" gorm:"comment:模型Id;column:model_id;" binding:"required"`                                       //模型Id
	AlgorithmName     *string  `json:"algorithmName" form:"algorithmName" gorm:"column:algorithm_name;" binding:"required"`                                  //算法名称
	AlgorithmId       *int64   `json:"algorithmId" form:"algorithmId" gorm:"comment:算法Id;column:algorithm_id;" binding:"required"`                           //算法Id
	ModelType         *int64   `json:"modelType" form:"modelType" gorm:"comment:模型类型：0: Local 1: OpenAI 2: Anthropic;column:model_type;" binding:"required"` //模型类型
	DatasetName       *string  `json:"datasetName" form:"datasetName" gorm:"comment:数据集名称;column:dataset_name;" binding:"required"`                          //数据集名称
	DatasetId         *int64   `json:"datasetId" form:"datasetId" gorm:"comment:数据集Id;column:dataset_id;" binding:"required"`                                //数据集Id
	TaskHash          *string  `json:"taskHash" form:"taskHash" gorm:"comment:任务唯一编号;column:task_hash;" binding:"required"`                                  //任务唯一编号
	AverageThroughput *float64 `json:"averageThroughput" form:"averageThroughput" gorm:"comment:平均吞吐量;column:average_throughput;" binding:"required"`        //平均吞吐量
	AverageLatency    *float64 `json:"averageLatency" form:"averageLatency" gorm:"comment:平均生成延迟;column:average_latency;" binding:"required"`                //平均生成延迟
	AverageGpuMemory  *float64 `json:"averageGpuMemory" form:"averageGpuMemory" gorm:"comment:平均GPU显存占用;column:average_gpu_memory;" binding:"required"`      //平均GPU显存占用
	OperatorName      *string  `json:"operatorName" form:"operatorName" gorm:"comment:执行人用户名;column:operator_name;" binding:"required"`                      //执行人用户名
	OperatorId        *int64   `json:"operatorId" form:"operatorId" gorm:"comment:执行人Id;column:operator_id;" binding:"required"`                             //执行人Id
}

// 执行推理任务结构体
type BizRunInferenceTask struct {
	TaskInnerSeq *int64  `json:"taskInnerSeq" form:"taskInnerSeq" binding:"required"` //任务内执行顺序
	ModelName    *string `json:"modelName" form:"modelName" binding:"required"`       //模型名称
	ModelId      *int64  `json:"modelId" form:"modelId" binding:"required"`           //模型Id
	AlgorithmIds []int64 `json:"algorithmIds" form:"algorithmIds" binding:"required"` //算法Id
	ModelType    *int64  `json:"modelType" form:"modelType" binding:"required"`       //模型类型
	DatasetId    *int64  `json:"datasetId" form:"datasetId" binding:"required"`       //数据集Id
	SampleId     *int64  `json:"sampleId" form:"sampleId" binding:"required"`         //样本Id
	TaskHash     *string `json:"taskHash" form:"taskHash" binding:"required"`         //任务唯一编号
}

// TableName 推理任务记录 BizInferenceTask自定义表名 biz_inference_task
func (BizInferenceTask) TableName() string {
	return "biz_inference_task"
}
