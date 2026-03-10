package biz

import (
	"context"
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz"
	bizReq "github.com/flipped-aurora/gin-vue-admin/server/model/biz/request"
)

type BizInferenceTaskService struct{}

// CreateBizInferenceTask 创建推理任务记录记录
// Author [yourname](https://github.com/yourname)
func (inferenceTaskService *BizInferenceTaskService) CreateBizInferenceTask(ctx context.Context, inferenceTask *biz.BizInferenceTask) (err error) {
	err = global.GVA_DB.Create(inferenceTask).Error
	return err
}

// DeleteBizInferenceTask 删除推理任务记录记录
// Author [yourname](https://github.com/yourname)
func (inferenceTaskService *BizInferenceTaskService) DeleteBizInferenceTask(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&biz.BizInferenceTask{}, "task_hash = ?", ID).Error
	if err != nil {
		return err
	}
	err = global.GVA_DB.Delete(&biz.BizInferenceLog{}, "task_hash = ?", ID).Error
	return err
}

// DeleteBizInferenceTaskByIds 批量删除推理任务记录记录
// Author [yourname](https://github.com/yourname)
func (inferenceTaskService *BizInferenceTaskService) DeleteBizInferenceTaskByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]biz.BizInferenceTask{}, "task_hash in ?", IDs).Error
	if err != nil {
		return err
	}
	err = global.GVA_DB.Delete(&[]biz.BizInferenceLog{}, "task_hash in ?", IDs).Error
	return err
}

// UpdateBizInferenceTask 更新推理任务记录记录
// Author [yourname](https://github.com/yourname)
func (inferenceTaskService *BizInferenceTaskService) UpdateBizInferenceTask(ctx context.Context, inferenceTask biz.BizInferenceTask) (err error) {
	err = global.GVA_DB.Model(&biz.BizInferenceTask{}).Where("id = ?", inferenceTask.ID).Updates(&inferenceTask).Error
	return err
}

// GetBizInferenceTask 根据ID获取推理任务记录记录
// Author [yourname](https://github.com/yourname)
func (inferenceTaskService *BizInferenceTaskService) GetBizInferenceTask(ctx context.Context, ID string) (inferenceTask biz.BizInferenceTask, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&inferenceTask).Error
	return
}

// GetBizInferenceTaskInfoList 分页获取推理任务记录记录
// Author [yourname](https://github.com/yourname)
func (inferenceTaskService *BizInferenceTaskService) GetBizInferenceTaskInfoList(ctx context.Context, info bizReq.BizInferenceTaskSearch) (list []biz.BizInferenceTask, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&biz.BizInferenceTask{}).Where("operator_id = ?", info.OperatorId)
	var inferenceTasks []biz.BizInferenceTask

	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	db = db.Select(`
		MIN(id) as id,
		MIN(created_at) as created_at,
		MAX(updated_at) as updated_at,
		task_hash,
		model_name,
		model_id,
		model_type,
		dataset_name,
		dataset_id,
		operator_name,
		operator_id,
		GROUP_CONCAT(algorithm_name SEPARATOR ',') as algorithm_name
	`).Group("task_hash, model_name, model_id, model_type, dataset_name, dataset_id, operator_name, operator_id")

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&inferenceTasks).Error
	return inferenceTasks, total, err
}

// GetBizInferenceRank 分页获取推理性能榜单
// Author [yourname](https://github.com/yourname)
func (inferenceTaskService *BizInferenceTaskService) GetBizInferenceRank(ctx context.Context, info bizReq.BizInferenceRank) (list []biz.BizInferenceTask, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&biz.BizInferenceTask{})
	var inferenceTasks []biz.BizInferenceTask
	db = db.Where("dataset_id = ?", info.DatasetId)
	if info.PerfType != nil {
		switch *info.PerfType {
		case 0:
			db = db.Order("average_throughput DESC")
		case 1:
			db = db.Order("average_latency ASC")
		case 2:
			db = db.Order("average_gpu_memory ASC")
		default:
			// 无效的性能类型，返回错误
			err = errors.New("invalid performance type")
			return
		}
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&inferenceTasks).Error
	return inferenceTasks, total, err
}

func (inferenceTaskService *BizInferenceTaskService) GetBizInferenceTaskPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
