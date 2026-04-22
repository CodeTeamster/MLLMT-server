package biz

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"path/filepath"
	"strings"
	"sync"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz"
	bizReq "github.com/flipped-aurora/gin-vue-admin/server/model/biz/request"
	bizRes "github.com/flipped-aurora/gin-vue-admin/server/model/biz/response"
	"gorm.io/datatypes"
)

type BizInferenceTaskService struct{}

func sendRequest(
	baseURL string,
	isStream bool,
	modelName string,
	text string,
	imageURL string,
) (string, float64, float64, float64) {
	requestBody := map[string]interface{}{
		"model": modelName,
		"messages": []interface{}{
			map[string]interface{}{
				"role": "user",
				"content": []interface{}{
					map[string]interface{}{
						"type": "text",
						"text": text,
					},
					map[string]interface{}{
						"type": "image_url",
						"image_url": map[string]interface{}{
							"url": imageURL,
						},
					},
				},
			},
		},
		"stream":     isStream,
		"max_tokens": 256,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Printf("Error marshalling JSON for %s: %v\n", baseURL, err)
		return "", 0, 0, 0
	}

	req, err := http.NewRequest("POST", baseURL+"/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error creating request for %s: %v\n", baseURL, err)
		return "", 0, 0, 0
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer <random_string>")
	if isStream {
		req.Header.Set("Accept", "text/event-stream")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request to %s: %v\n", baseURL, err)
		return "", 0, 0, 0
	}
	defer resp.Body.Close()

	if isStream {
		type ResponseContent struct {
			Seq                    int     `json:"seq"`
			Content                string  `json:"content"`
			PerTokenGenTime        float64 `json:"per_token_gen_time"`
			PerTokenGPUMem         float64 `json:"per_token_gpu_mem"`
			AverageThroughputToNow float64 `json:"avg_throughput_to_now"`
			AverageLatencyToNow    float64 `json:"avg_latency_to_now"`
		}

		var contentArray []ResponseContent
		var finalMem, finalThroughput, finalLatency float64
		scanner := bufio.NewScanner(resp.Body)

		for scanner.Scan() {
			line := scanner.Text()
			// 过滤空行和非 data: 开头的行
			if line == "" || !strings.HasPrefix(line, "data: ") {
				continue
			}

			dataStr := strings.TrimPrefix(line, "data: ")
			// 判断流是否结束
			if dataStr == "[DONE]" {
				break
			}

			// 解析单一 chunk 的 JSON 结构
			var chunk struct {
				Choices []struct {
					Delta struct {
						Content string `json:"content"`
					} `json:"delta"`
				} `json:"choices"`
				Metrics struct {
					GeneratedTokens        int       `json:"generated_tokens"`
					PerTokenGenTime        []float64 `json:"per_token_gen_time"`
					PerTokenGPUMem         []float64 `json:"per_token_gpu_mem"`
					AverageThroughputToNow []float64 `json:"avg_throughput_to_now"`
					AverageLatencyToNow    []float64 `json:"avg_latency_to_now"`
				} `json:"metrics"`
			}

			if err := json.Unmarshal([]byte(dataStr), &chunk); err == nil {
				if len(chunk.Choices) > 0 && len(chunk.Choices[0].Delta.Content) > 0 {
					var perTokenGenTime, perTokenGPUMem, avgThroughput, avgLatency float64

					if len(chunk.Metrics.PerTokenGenTime) > 0 {
						perTokenGenTime = math.Round(chunk.Metrics.PerTokenGenTime[0]*1000*10000) / 10000
					}
					if len(chunk.Metrics.PerTokenGPUMem) > 0 {
						perTokenGPUMem = math.Round(chunk.Metrics.PerTokenGPUMem[0]*10000) / 10000
						finalMem = perTokenGPUMem
					}
					if len(chunk.Metrics.AverageThroughputToNow) > 0 {
						avgThroughput = math.Round(chunk.Metrics.AverageThroughputToNow[0]*10000) / 10000
						finalThroughput = avgThroughput
					}
					if len(chunk.Metrics.AverageLatencyToNow) > 0 {
						avgLatency = math.Round(chunk.Metrics.AverageLatencyToNow[0]*1000*10000) / 10000
						finalLatency = avgLatency
					}

					content := ResponseContent{
						Seq:                    chunk.Metrics.GeneratedTokens,
						Content:                chunk.Choices[0].Delta.Content,
						PerTokenGenTime:        perTokenGenTime,
						PerTokenGPUMem:         perTokenGPUMem,
						AverageThroughputToNow: avgThroughput,
						AverageLatencyToNow:    avgLatency,
					}
					contentArray = append(contentArray, content)
				}
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading stream from %s: %v\n", baseURL, err)
		}

		finalResBytes, _ := json.Marshal(contentArray)
		return string(finalResBytes), finalMem, finalThroughput, finalLatency
	}

	bodyByte, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response from %s: %v\n", baseURL, err)
		return "", 0, 0, 0
	}

	return string(bodyByte), 0, 0, 0
}

// RunBizInferenceTask 运行推理任务
// Author [yourname](https://github.com/yourname)
func (inferenceTaskService *BizInferenceTaskService) RunBizInferenceTask(ctx context.Context, inferenceTask *biz.BizRunInferenceTask) (logs []biz.BizInferenceLog, err error) {
	// 并发发送请求
	targetServer := map[string]string{
		"http://localhost:24321/v1": "/nfs6/yrc/model/Qwen/Qwen2.5-VL-7B-Instruct",
		"http://localhost:24322/v1": "/nfs6/yrc/model/Qwen/Qwen2.5-VL-7B-Instruct-AWQ",
	}
	isStream := true

	sample := biz.BizSample{}
	err = global.GVA_DB.Model(&biz.BizSample{}).
		Where("dataset_id = ?", inferenceTask.DatasetId).
		Where("id = ?", inferenceTask.SampleId).First(&sample).Error
	if err != nil {
		return
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	type Result struct {
		Content         json.RawMessage
		FinalMem        float64
		FinalThroughput float64
		FinalLatency    float64
	}
	results := make(map[string]Result)

	for url, modelName := range targetServer {
		wg.Add(1)
		go func(targetUrl, modelName string) {
			defer wg.Done()
			resContent, finalMem, finalThroughput, finalLatency := sendRequest(
				targetUrl,
				isStream,
				modelName,
				*sample.Prompt,
				"http://localhost:28080/api/"+*sample.Img,
			)

			mu.Lock()
			results[filepath.Base(modelName)] = Result{
				Content:         json.RawMessage(resContent),
				FinalMem:        finalMem,
				FinalThroughput: finalThroughput,
				FinalLatency:    finalLatency,
			}
			mu.Unlock()
		}(url, modelName)
	}
	wg.Wait()

	for modelName, res := range results {
		var algorithmName string
		var algorithmId int64
		fakeModelId := int64(4)
		fakeModelName := "Qwen2.5-VL-7B-Instruct"
		if modelName == "Qwen2.5-VL-7B-Instruct" {
			algorithmName = "unquantized"
			algorithmId = int64(1)
		} else {
			algorithmName = "AWQ"
			algorithmId = int64(7)
		}

		inferenceLog := biz.BizInferenceLog{
			TaskHash:      inferenceTask.TaskHash,
			DatasetId:     inferenceTask.DatasetId,
			SampleId:      inferenceTask.SampleId,
			ModelName:     &fakeModelName,
			ModelId:       &fakeModelId,
			AlgorithmName: &algorithmName,
			AlgorithmId:   &algorithmId,
			Throughput:    &res.FinalThroughput,
			Latency:       &res.FinalLatency,
			GpuMemory:     &res.FinalMem,
			TaskInnerSeq:  inferenceTask.TaskInnerSeq,
			JsonLog:       datatypes.JSON(res.Content),
		}
		err = global.GVA_DB.Create(&inferenceLog).Error
		if err != nil {
			fmt.Printf("Error saving inference log for model %s: %v\n", modelName, err)
		}
		logs = append(logs, inferenceLog)
	}

	return
}

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

// GetBizInferenceCompleteRecord 获取推理任务详细记录
// Author [yourname](https://github.com/yourname)
func (inferenceTaskService *BizInferenceTaskService) GetBizInferenceCompleteRecord(
	ctx context.Context,
	info bizReq.BizInferenceCompleteRecord,
) (inferenceCompleteRecord bizRes.BizInferenceCompleteRecord, err error) {
	err = global.GVA_DB.Model(&biz.BizInferenceTask{}).
		Where("task_hash = ?", info.TaskHash).
		Find(&inferenceCompleteRecord.InferenceTasks).Error
	if err != nil {
		return
	}

	db := global.GVA_DB.Model(&biz.BizInferenceLog{}).
		Where("task_hash = ?", info.TaskHash)

	if info.TaskInnerSeq != nil {
		db = db.Where("task_inner_seq = ?", info.TaskInnerSeq)
	}

	var flatLogs []biz.BizInferenceLog
	err = db.Order("task_inner_seq ASC").
		Find(&flatLogs).Error
	if err != nil {
		return
	}

	if len(flatLogs) > 0 {
		var currentGroup []biz.BizInferenceLog
		currentSeq := flatLogs[0].TaskInnerSeq

		for _, log := range flatLogs {
			if (log.TaskInnerSeq == nil && currentSeq == nil) ||
				(log.TaskInnerSeq != nil && currentSeq != nil && *log.TaskInnerSeq == *currentSeq) {
				currentGroup = append(currentGroup, log)
			} else {
				inferenceCompleteRecord.InferenceLogs = append(
					inferenceCompleteRecord.InferenceLogs,
					currentGroup,
				)
				currentGroup = []biz.BizInferenceLog{log}
				currentSeq = log.TaskInnerSeq
			}
		}
		if len(currentGroup) > 0 {
			inferenceCompleteRecord.InferenceLogs = append(
				inferenceCompleteRecord.InferenceLogs,
				currentGroup,
			)
		}
	}

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
