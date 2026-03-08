
package biz

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz"
    bizReq "github.com/flipped-aurora/gin-vue-admin/server/model/biz/request"
)

type BizAlgorithmService struct {}
// CreateBizAlgorithm 创建推理加速算法记录
// Author [yourname](https://github.com/yourname)
func (algorithmService *BizAlgorithmService) CreateBizAlgorithm(ctx context.Context, algorithm *biz.BizAlgorithm) (err error) {
	err = global.GVA_DB.Create(algorithm).Error
	return err
}

// DeleteBizAlgorithm 删除推理加速算法记录
// Author [yourname](https://github.com/yourname)
func (algorithmService *BizAlgorithmService)DeleteBizAlgorithm(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&biz.BizAlgorithm{},"id = ?",ID).Error
	return err
}

// DeleteBizAlgorithmByIds 批量删除推理加速算法记录
// Author [yourname](https://github.com/yourname)
func (algorithmService *BizAlgorithmService)DeleteBizAlgorithmByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]biz.BizAlgorithm{},"id in ?",IDs).Error
	return err
}

// UpdateBizAlgorithm 更新推理加速算法记录
// Author [yourname](https://github.com/yourname)
func (algorithmService *BizAlgorithmService)UpdateBizAlgorithm(ctx context.Context, algorithm biz.BizAlgorithm) (err error) {
	err = global.GVA_DB.Model(&biz.BizAlgorithm{}).Where("id = ?",algorithm.ID).Updates(&algorithm).Error
	return err
}

// GetBizAlgorithm 根据ID获取推理加速算法记录
// Author [yourname](https://github.com/yourname)
func (algorithmService *BizAlgorithmService)GetBizAlgorithm(ctx context.Context, ID string) (algorithm biz.BizAlgorithm, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&algorithm).Error
	return
}
// GetBizAlgorithmInfoList 分页获取推理加速算法记录
// Author [yourname](https://github.com/yourname)
func (algorithmService *BizAlgorithmService)GetBizAlgorithmInfoList(ctx context.Context, info bizReq.BizAlgorithmSearch) (list []biz.BizAlgorithm, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&biz.BizAlgorithm{})
    var algorithms []biz.BizAlgorithm
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&algorithms).Error
	return  algorithms, total, err
}
func (algorithmService *BizAlgorithmService)GetBizAlgorithmPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
