package biz

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz"
	bizReq "github.com/flipped-aurora/gin-vue-admin/server/model/biz/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BizInferenceTaskApi struct{}

// CreateBizInferenceTask 创建推理任务记录
// @Tags BizInferenceTask
// @Summary 创建推理任务记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body bizReq.BizInferenceTaskCreate true "创建推理任务记录"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /inferenceTask/createBizInferenceTask [post]
func (inferenceTaskApi *BizInferenceTaskApi) CreateBizInferenceTask(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var inferenceTaskReq bizReq.BizInferenceTaskCreate
	err := c.ShouldBindJSON(&inferenceTaskReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	operatorName := utils.GetUserName(c)
	operatorId := int64(utils.GetUserID(c))
	taskhash := utils.MD5Encode(
		*inferenceTaskReq.ModelName,
		*inferenceTaskReq.DatasetName,
		operatorName,
		operatorId,
		time.Now().UnixNano(),
	)

	inferenceTask := biz.BizInferenceTask{
		ModelName:         inferenceTaskReq.ModelName,
		ModelId:           inferenceTaskReq.ModelId,
		AlgorithmName:     inferenceTaskReq.AlgorithmName,
		AlgorithmId:       inferenceTaskReq.AlgorithmId,
		ModelType:         inferenceTaskReq.ModelType,
		DatasetName:       inferenceTaskReq.DatasetName,
		DatasetId:         inferenceTaskReq.DatasetId,
		TaskHash:          &taskhash,
		AverageThroughput: inferenceTaskReq.AverageThroughput,
		AverageLatency:    inferenceTaskReq.AverageLatency,
		AverageGpuMemory:  inferenceTaskReq.AverageGpuMemory,
		OperatorName:      &operatorName,
		OperatorId:        &operatorId,
	}

	err = inferenceTaskService.CreateBizInferenceTask(ctx, &inferenceTask)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteBizInferenceTask 删除推理任务记录
// @Tags BizInferenceTask
// @Summary 删除推理任务记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body biz.BizInferenceTask true "删除推理任务记录"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /inferenceTask/deleteBizInferenceTask [delete]
func (inferenceTaskApi *BizInferenceTaskApi) DeleteBizInferenceTask(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := inferenceTaskService.DeleteBizInferenceTask(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteBizInferenceTaskByIds 批量删除推理任务记录
// @Tags BizInferenceTask
// @Summary 批量删除推理任务记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /inferenceTask/deleteBizInferenceTaskByIds [delete]
func (inferenceTaskApi *BizInferenceTaskApi) DeleteBizInferenceTaskByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := inferenceTaskService.DeleteBizInferenceTaskByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateBizInferenceTask 更新推理任务记录
// @Tags BizInferenceTask
// @Summary 更新推理任务记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body biz.BizInferenceTask true "更新推理任务记录"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /inferenceTask/updateBizInferenceTask [put]
func (inferenceTaskApi *BizInferenceTaskApi) UpdateBizInferenceTask(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var inferenceTask biz.BizInferenceTask
	err := c.ShouldBindJSON(&inferenceTask)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = inferenceTaskService.UpdateBizInferenceTask(ctx, inferenceTask)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindBizInferenceTask 用id查询推理任务记录
// @Tags BizInferenceTask
// @Summary 用id查询推理任务记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询推理任务记录"
// @Success 200 {object} response.Response{data=biz.BizInferenceTask,msg=string} "查询成功"
// @Router /inferenceTask/findBizInferenceTask [get]
func (inferenceTaskApi *BizInferenceTaskApi) FindBizInferenceTask(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reinferenceTask, err := inferenceTaskService.GetBizInferenceTask(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reinferenceTask, c)
}

// GetBizInferenceTaskList 分页获取推理任务记录列表
// @Tags BizInferenceTask
// @Summary 分页获取推理任务记录列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query bizReq.BizInferenceTaskSearch true "分页获取推理任务记录列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /inferenceTask/getBizInferenceTaskList [get]
func (inferenceTaskApi *BizInferenceTaskApi) GetBizInferenceTaskList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo bizReq.BizInferenceTaskSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	operatorId := int64(utils.GetUserID(c))
	pageInfo.OperatorId = &operatorId

	list, total, err := inferenceTaskService.GetBizInferenceTaskInfoList(ctx, pageInfo)
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

// GetBizInferenceRank 分页获取推理性能榜单
// @Tags BizInferenceTask
// @Summary 分页获取推理性能榜单
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query bizReq.GetBizInferenceRank true "分页获取推理性能榜单"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /inferenceTask/getBizInferenceRank [get]
func (inferenceTaskApi *BizInferenceTaskApi) GetBizInferenceRank(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var inferenceRankReq bizReq.BizInferenceRank
	err := c.ShouldBindQuery(&inferenceRankReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := inferenceTaskService.GetBizInferenceRank(ctx, inferenceRankReq)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     inferenceRankReq.Page,
		PageSize: inferenceRankReq.PageSize,
	}, "获取成功", c)
}

// GetBizInferenceTaskPublic 不需要鉴权的推理任务记录接口
// @Tags BizInferenceTask
// @Summary 不需要鉴权的推理任务记录接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /inferenceTask/getBizInferenceTaskPublic [get]
func (inferenceTaskApi *BizInferenceTaskApi) GetBizInferenceTaskPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	inferenceTaskService.GetBizInferenceTaskPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的推理任务记录接口信息",
	}, "获取成功", c)
}
