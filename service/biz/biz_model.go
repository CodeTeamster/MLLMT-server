package biz

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz"
	bizReq "github.com/flipped-aurora/gin-vue-admin/server/model/biz/request"
)

type BizModelService struct{}

// CreateBizModel 创建模型管理记录
// Author [yourname](https://github.com/yourname)
func (modelService *BizModelService) CreateBizModel(ctx context.Context, model *biz.BizModel) (err error) {
	err = global.GVA_DB.Create(model).Error
	return err
}

// DeleteBizModel 删除模型管理记录
// Author [yourname](https://github.com/yourname)
func (modelService *BizModelService) DeleteBizModel(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&biz.BizModel{}, "id = ?", ID).Error
	return err
}

// DeleteBizModelByIds 批量删除模型管理记录
// Author [yourname](https://github.com/yourname)
func (modelService *BizModelService) DeleteBizModelByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]biz.BizModel{}, "id in ?", IDs).Error
	return err
}

// UpdateBizModel 更新模型管理记录
// Author [yourname](https://github.com/yourname)
func (modelService *BizModelService) UpdateBizModel(ctx context.Context, model biz.BizModel) (err error) {
	err = global.GVA_DB.Model(&biz.BizModel{}).Where("id = ?", model.ID).Updates(&model).Error
	return err
}

// GetBizModel 根据ID获取模型管理记录
// Author [yourname](https://github.com/yourname)
func (modelService *BizModelService) GetBizModel(ctx context.Context, ID string) (model biz.BizModel, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&model).Error
	return
}

// GetBizModelInfoList 分页获取模型管理记录
// Author [yourname](https://github.com/yourname)
func (modelService *BizModelService) GetBizModelInfoList(ctx context.Context, info bizReq.BizModelSearch) (list []biz.BizModel, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&biz.BizModel{})
	var models []biz.BizModel
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

	err = db.Find(&models).Error
	return models, total, err
}
func (modelService *BizModelService) GetBizModelPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
