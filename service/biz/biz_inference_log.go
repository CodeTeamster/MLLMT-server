package biz

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz"
	bizReq "github.com/flipped-aurora/gin-vue-admin/server/model/biz/request"
	bizRes "github.com/flipped-aurora/gin-vue-admin/server/model/biz/response"
)

type BizInferenceLogService struct{}

// CreateBizInferenceLog 创建推理详细记录记录
// Author [yourname](https://github.com/yourname)
func (inferenceLogService *BizInferenceLogService) CreateBizInferenceLog(ctx context.Context, inferenceLog *biz.BizInferenceLog) (err error) {
	err = global.GVA_DB.Create(inferenceLog).Error
	return err
}

// DeleteBizInferenceLog 删除推理详细记录记录
// Author [yourname](https://github.com/yourname)
func (inferenceLogService *BizInferenceLogService) DeleteBizInferenceLog(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&biz.BizInferenceLog{}, "id = ?", ID).Error
	return err
}

// DeleteBizInferenceLogByIds 批量删除推理详细记录记录
// Author [yourname](https://github.com/yourname)
func (inferenceLogService *BizInferenceLogService) DeleteBizInferenceLogByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]biz.BizInferenceLog{}, "id in ?", IDs).Error
	return err
}

// UpdateBizInferenceLog 更新推理详细记录记录
// Author [yourname](https://github.com/yourname)
func (inferenceLogService *BizInferenceLogService) UpdateBizInferenceLog(ctx context.Context, inferenceLog biz.BizInferenceLog) (err error) {
	err = global.GVA_DB.Model(&biz.BizInferenceLog{}).Where("id = ?", inferenceLog.ID).Updates(&inferenceLog).Error
	return err
}

// GetBizInferenceLog 根据ID获取推理详细记录记录
// Author [yourname](https://github.com/yourname)
func (inferenceLogService *BizInferenceLogService) GetBizInferenceLog(ctx context.Context, ID string) (inferenceLog biz.BizInferenceLog, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&inferenceLog).Error
	return
}

// GetBizInferenceLogSample 分页获取当前任务所有样本内容
// Author [yourname](https://github.com/yourname)
func (inferenceLogService *BizInferenceLogService) GetBizInferenceLogSample(ctx context.Context, info bizReq.BizInferenceLogSample) (list []bizRes.BizInferenceLogSample, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	var inferenceLogSamples []bizRes.BizInferenceLogSample

	db := global.GVA_DB.Table("biz_inference_log AS log").
		Select(
			"ANY_VALUE(log.task_hash) AS task_hash",
			"ANY_VALUE(log.dataset_id) AS dataset_id",
			"log.sample_id",
			"ANY_VALUE(log.task_inner_seq) AS task_inner_seq",
			"s.prompt",
			"s.img",
		).Order("ANY_VALUE(log.task_inner_seq) ASC").
		Joins("LEFT JOIN biz_sample AS s ON log.sample_id = s.id").
		Where("log.task_hash = ?", info.TaskHash).
		Group("log.sample_id")

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&inferenceLogSamples).Error
	return inferenceLogSamples, total, err
}

// GetBizInferenceLogInfoList 分页获取推理详细记录记录
// Author [yourname](https://github.com/yourname)
func (inferenceLogService *BizInferenceLogService) GetBizInferenceLogInfoList(ctx context.Context, info bizReq.BizInferenceLogSearch) (list []biz.BizInferenceLog, total int64, err error) {
	// 创建db
	db := global.GVA_DB.Model(&biz.BizInferenceLog{})
	var inferenceLogs []biz.BizInferenceLog

	db = db.Where("task_hash = ?", info.TaskHash)
	db = db.Where("dataset_id = ?", info.DatasetId)
	db = db.Where("sample_id = ?", info.SampleId)

	if info.AlgorithmId != nil {
		db = db.Where("algorithm_id = ?", info.AlgorithmId)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Find(&inferenceLogs).Error
	return inferenceLogs, total, err
}

func (inferenceLogService *BizInferenceLogService) GetBizInferenceLogPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
