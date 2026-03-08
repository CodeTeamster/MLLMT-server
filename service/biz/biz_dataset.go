package biz

import (
	"context"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz"
	bizReq "github.com/flipped-aurora/gin-vue-admin/server/model/biz/request"
)

type BizDatasetService struct{}

// CreateBizDataset 创建数据集管理记录
// Author [yourname](https://github.com/yourname)
func (datasetService *BizDatasetService) CreateBizDataset(ctx context.Context, dataset *biz.BizDataset) (err error) {
	err = global.GVA_DB.Create(dataset).Error
	return err
}

// DeleteBizDataset 删除数据集管理记录
// Author [yourname](https://github.com/yourname)
func (datasetService *BizDatasetService) DeleteBizDataset(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&biz.BizDataset{}, "id = ?", ID).Error
	return err
}

// DeleteBizDatasetByIds 批量删除数据集管理记录
// Author [yourname](https://github.com/yourname)
func (datasetService *BizDatasetService) DeleteBizDatasetByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]biz.BizDataset{}, "id in ?", IDs).Error
	return err
}

// UpdateBizDataset 更新数据集管理记录
// Author [yourname](https://github.com/yourname)
func (datasetService *BizDatasetService) UpdateBizDataset(ctx context.Context, dataset biz.BizDataset) (err error) {
	err = global.GVA_DB.Model(&biz.BizDataset{}).Where("id = ?", dataset.ID).Updates(&dataset).Error
	return err
}

// GetBizDataset 根据ID获取数据集管理记录
// Author [yourname](https://github.com/yourname)
func (datasetService *BizDatasetService) GetBizDataset(ctx context.Context, ID string) (dataset biz.BizDataset, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&dataset).Error
	return
}

// GetBizDatasetInfoList 分页获取数据集管理记录
// Author [yourname](https://github.com/yourname)
func (datasetService *BizDatasetService) GetBizDatasetInfoList(ctx context.Context, info bizReq.BizDatasetSearch, userId uint) (list []biz.BizDataset, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&biz.BizDataset{})
	fmt.Printf("ctx: %v\n", ctx)
	var datasets []biz.BizDataset

	db = db.Where("creator_id = ?", userId).Or("scope = ?", 1)

	// 如果有条件搜索 下方会自动创建搜索语句
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

	err = db.Find(&datasets).Error
	return datasets, total, err
}
func (datasetService *BizDatasetService) GetBizDatasetPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
