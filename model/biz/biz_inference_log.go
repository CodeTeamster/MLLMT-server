// 自动生成模板BizInferenceLog
package biz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/datatypes"
)

// 推理详细记录 结构体  BizInferenceLog
type BizInferenceLog struct {
	global.GVA_MODEL
	TaskHash      *string        `json:"taskHash" form:"taskHash" gorm:"comment:任务编号;column:task_hash;" binding:"required"`                         //任务编号
	DatasetId     *int64         `json:"datasetId" form:"datasetId" gorm:"comment:数据集Id;column:dataset_id;" binding:"required"`                     //数据集Id
	SampleId      *int64         `json:"sampleId" form:"sampleId" gorm:"comment:样本Id;column:sample_id;" binding:"required"`                         //样本Id
	ModelName     *string        `json:"modelName" form:"modelName" gorm:"comment:模型名称;column:model_name;" binding:"required"`                      //模型名称
	ModelId       *int64         `json:"modelId" form:"modelId" gorm:"comment:模型Id;column:model_id;" binding:"required"`                            //模型Id
	AlgorithmName *string        `json:"algorithmName" form:"algorithmName" gorm:"comment:算法名称;column:algorithm_name;" binding:"required"`          //算法名称
	AlgorithmId   *int64         `json:"algorithmId" form:"algorithmId" gorm:"comment:算法Id;column:algorithm_id;" binding:"required"`                //算法Id
	Throughput    *float64       `json:"throughput" form:"throughput" gorm:"comment:吞吐量;column:throughput;" binding:"required"`                     //吞吐量
	Latency       *float64       `json:"latency" form:"latency" gorm:"comment:单token推理延迟;column:latency;" binding:"required"`                       //单token推理延迟
	GpuMemory     *float64       `json:"gpuMemory" form:"gpuMemory" gorm:"comment:GPU显存占用;column:gpu_memory;" binding:"required"`                   //GPU显存占用
	TaskInnerSeq  *int64         `json:"taskInnerSeq" form:"taskInnerSeq" gorm:"comment:任务内执行顺序，序号从1开始;column:task_inner_seq;" binding:"required"`  //任务内执行顺序
	JsonLog       datatypes.JSON `json:"jsonLog" form:"jsonLog" gorm:"comment:json形式推理结果;column:json_log;" swaggertype:"object" binding:"required"` //json形式推理结果
}

// TableName 推理详细记录 BizInferenceLog自定义表名 biz_inference_log
func (BizInferenceLog) TableName() string {
	return "biz_inference_log"
}
