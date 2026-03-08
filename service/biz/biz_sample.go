package biz

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz"
	bizReq "github.com/flipped-aurora/gin-vue-admin/server/model/biz/request"
)

type BizSampleService struct{}

// CreateBizSample 创建数据集样本管理记录
// Author [yourname](https://github.com/yourname)
func (sampleService *BizSampleService) CreateBizSample(ctx context.Context, sample *biz.BizSample) (err error) {
	err = global.GVA_DB.Create(sample).Error
	return err
}

// DeleteBizSample 删除数据集样本管理记录
// Author [yourname](https://github.com/yourname)
func (sampleService *BizSampleService) DeleteBizSample(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&biz.BizSample{}, "id = ?", ID).Error
	return err
}

// DeleteBizSampleByIds 批量删除数据集样本管理记录
// Author [yourname](https://github.com/yourname)
func (sampleService *BizSampleService) DeleteBizSampleByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]biz.BizSample{}, "id in ?", IDs).Error
	return err
}

// UpdateBizSample 更新数据集样本管理记录
// Author [yourname](https://github.com/yourname)
func (sampleService *BizSampleService) UpdateBizSample(ctx context.Context, sample biz.BizSample) (err error) {
	err = global.GVA_DB.Model(&biz.BizSample{}).Where("id = ?", sample.ID).Updates(&sample).Error
	return err
}

// GetBizSample 根据ID获取数据集样本管理记录
// Author [yourname](https://github.com/yourname)
func (sampleService *BizSampleService) GetBizSample(ctx context.Context, ID string) (sample biz.BizSample, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&sample).Error
	return
}

// GetBizSampleInfoList 分页获取数据集样本管理记录
// Author [yourname](https://github.com/yourname)
func (sampleService *BizSampleService) GetBizSampleInfoList(ctx context.Context, info bizReq.BizSampleSearch) (list []biz.BizSample, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&biz.BizSample{})
	var samples []biz.BizSample
	// 如果有条件搜索 下方会自动创建搜索语句
	db = db.Where("dataset_id = ?", info.DatasetId)

	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&samples).Error
	return samples, total, err
}
func (sampleService *BizSampleService) GetBizSamplePublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
