package biz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz"
	bizReq "github.com/flipped-aurora/gin-vue-admin/server/model/biz/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BizModelApi struct{}

// CreateBizModel 创建模型管理
// @Tags BizModel
// @Summary 创建模型管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body bizReq.BizModelCreate true "创建模型管理"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /model/createBizModel [post]
func (modelApi *BizModelApi) CreateBizModel(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var modelReq bizReq.BizModelCreate
	err := c.ShouldBindJSON(&modelReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	creatorName := utils.GetUserName(c)
	creatorId := int64(utils.GetUserID(c))
	model := biz.BizModel{
		ModelName:     modelReq.ModelName,
		AlgorithmName: modelReq.AlgorithmName,
		AlgorithmId:   modelReq.AlgorithmId,
		Type:          modelReq.Type,
		Enable:        modelReq.Enable,
		CreatorName:   &creatorName,
		CreatorId:     &creatorId,
	}

	err = modelService.CreateBizModel(ctx, &model)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteBizModel 删除模型管理
// @Tags BizModel
// @Summary 删除模型管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body biz.BizModel true "删除模型管理"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /model/deleteBizModel [delete]
func (modelApi *BizModelApi) DeleteBizModel(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := modelService.DeleteBizModel(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteBizModelByIds 批量删除模型管理
// @Tags BizModel
// @Summary 批量删除模型管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /model/deleteBizModelByIds [delete]
func (modelApi *BizModelApi) DeleteBizModelByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := modelService.DeleteBizModelByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateBizModel 更新模型管理
// @Tags BizModel
// @Summary 更新模型管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body biz.BizModel true "更新模型管理"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /model/updateBizModel [put]
func (modelApi *BizModelApi) UpdateBizModel(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var model biz.BizModel
	err := c.ShouldBindJSON(&model)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = modelService.UpdateBizModel(ctx, model)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindBizModel 用id查询模型管理
// @Tags BizModel
// @Summary 用id查询模型管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询模型管理"
// @Success 200 {object} response.Response{data=biz.BizModel,msg=string} "查询成功"
// @Router /model/findBizModel [get]
func (modelApi *BizModelApi) FindBizModel(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	remodel, err := modelService.GetBizModel(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(remodel, c)
}

// GetBizModelList 分页获取模型管理列表
// @Tags BizModel
// @Summary 分页获取模型管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query bizReq.BizModelSearch true "分页获取模型管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /model/getBizModelList [get]
func (modelApi *BizModelApi) GetBizModelList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo bizReq.BizModelSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := modelService.GetBizModelInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetBizModelPublic 不需要鉴权的模型管理接口
// @Tags BizModel
// @Summary 不需要鉴权的模型管理接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /model/getBizModelPublic [get]
func (modelApi *BizModelApi) GetBizModelPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	modelService.GetBizModelPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的模型管理接口信息",
	}, "获取成功", c)
}
