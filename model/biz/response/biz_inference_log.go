package request

type BizInferenceLogSample struct {
	TaskHash     *string `gorm:"column:task_hash"`
	DatasetId    *int    `gorm:"column:dataset_id"`
	SampleId     *int    `gorm:"column:sample_id"`
	TaskInnerSeq *int    `gorm:"column:task_inner_seq"`
	Prompt       *string `gorm:"column:prompt"`
	Img          *string `gorm:"column:img"`
}
