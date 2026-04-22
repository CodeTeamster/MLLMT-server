package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/biz"
	"gorm.io/gorm"
)

type BizInferenceTaskSearch struct {
	ID            uint           `gorm:"primarykey" json:"ID"` // 主键ID
	CreatedAt     time.Time      // 创建时间
	UpdatedAt     time.Time      // 更新时间
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`                                                                                                       // 删除时间
	ModelName     *string        `json:"modelName" form:"modelName" gorm:"comment:模型名称;column:model_name;" binding:"required"`                                 //模型名称
	ModelId       *int64         `json:"modelId" form:"modelId" gorm:"comment:模型Id;column:model_id;" binding:"required"`                                       //模型Id
	AlgorithmName *string        `json:"algorithmName" form:"algorithmName" gorm:"column:algorithm_name;" binding:"required"`                                  //算法名称
	AlgorithmId   *int64         `json:"algorithmId" form:"algorithmId" gorm:"comment:算法Id;column:algorithm_id;" binding:"required"`                           //算法Id
	ModelType     *int64         `json:"modelType" form:"modelType" gorm:"comment:模型类型：0: Local 1: OpenAI 2: Anthropic;column:model_type;" binding:"required"` //模型类型
	DatasetName   *string        `json:"datasetName" form:"datasetName" gorm:"comment:数据集名称;column:dataset_name;" binding:"required"`                          //数据集名称
	DatasetId     *int64         `json:"datasetId" form:"datasetId" gorm:"comment:数据集Id;column:dataset_id;" binding:"required"`                                //数据集Id
	TaskHash      *string        `json:"taskHash" form:"taskHash" gorm:"comment:任务唯一编号;column:task_hash;" binding:"required"`                                  //任务唯一编号
	OperatorName  *string        `json:"operatorName" form:"operatorName" gorm:"comment:执行人用户名;column:operator_name;" binding:"required"`                      //执行人用户名
	OperatorId    *int64         `json:"operatorId" form:"operatorId" gorm:"comment:执行人Id;column:operator_id;" binding:"required"`                             //执行人Id
}

// 推理任务记录 结构体  BizInferenceTask
type BizInferenceCompleteRecord struct {
	InferenceLogs  [][]biz.BizInferenceLog `json:"inferenceLogs" form:"inferenceLogs"`
	InferenceTasks []biz.BizInferenceTask  `json:"inferenceTasks" form:"inferenceTasks"` //推理任务记录
}
