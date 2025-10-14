package event

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/lin-snow/ech0/internal/async"
	webhookModel "github.com/lin-snow/ech0/internal/model/webhook"
	queueRepository "github.com/lin-snow/ech0/internal/repository/queue"
	webhookRepository "github.com/lin-snow/ech0/internal/repository/webhook"
	"github.com/lin-snow/ech0/internal/transaction"
	logUtil "github.com/lin-snow/ech0/internal/util/log"
)

type WebhookDispatcher struct {
	bus       IEventBus                                    // 事件总线
	client    *http.Client                                 // HTTP 客户端
	repo      webhookRepository.WebhookRepositoryInterface // Webhook 仓储层
	pool      *async.WorkerPool                            // 任务池pool
	queueRepo queueRepository.QueueRepositoryInterface     // 死信任务仓储
	txManager transaction.TransactionManager               // 事务管理器
}

func NewWebhookDispatcher(
	ebp func() IEventBus,
	repo webhookRepository.WebhookRepositoryInterface,
	queueRepo queueRepository.QueueRepositoryInterface,
	txManager transaction.TransactionManager,
) *WebhookDispatcher {
	return &WebhookDispatcher{
		bus:       ebp(),     // 获取事件总线实例
		repo:      repo,      // 注入仓储层
		queueRepo: queueRepo, // 注入死信任务仓储
		client: &http.Client{ // 配置 HTTP 客户端
			Timeout: 5 * time.Second, // 请求超时时间
			Transport: &http.Transport{ // 自定义传输设置（使用连接池）
				MaxIdleConns:        10,               // 最大空闲连接数
				MaxIdleConnsPerHost: 10,               // 每个主机的最大空闲连接数
				IdleConnTimeout:     30 * time.Second, // 空闲连接超时时间
			},
		},
		pool:      async.NewWorkerPool(6, 6), // 假设最大并发数为 6，任务队列大小为 6
		txManager: txManager,                 // 注入事务管理器
	}
}

// Handle 由事件总线调用，负责调度事件到每个活跃的 webhook
func (wd *WebhookDispatcher) Handle(ctx context.Context, e *Event) error {
	// 获取所有开启的webhook
	webhooks, err := wd.repo.ListActiveWebhooks()
	if err != nil {
		return err
	}
	for _, wh := range webhooks {
		// 可以根据 event type 做过滤
		// if !wh.ShouldHandle(e.Type) {
		// 	continue
		// }
		wh := wh // 捕获变量
		// 提交任务到池中异步处理
		wd.pool.Submit(func() error {
			wd.Dispatch(ctx, &wh, e)
			return nil
		})
	}

	return nil
}

// Dispatch 负责将事件发送到指定的 webhook
func (wd *WebhookDispatcher) Dispatch(ctx context.Context, wh *webhookModel.Webhook, e *Event) {
	// 构建 HTTP 请求
	req, err := wd.buildRequest(wh, e)
	if err != nil {
		// 记录日志或处理错误
		logUtil.GetLogger().Error(err.Error())
		return
	}

	// 发送请求，带重试机制
	wd.retryWithBackoff(3, 500*time.Millisecond, func() error {
		// 发送 HTTP 请求
		resp, err := wd.client.Do(req)
		if err != nil {
			// 记录日志或处理错误
			logUtil.GetLogger().Error(err.Error())
			return err
		}
		defer resp.Body.Close()

		// 处理响应
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			// 成功处理
			return nil
		} else {
			// 记录失败日志
			logUtil.GetLogger().Error("Webhook Handle Failed: ", zap.String("name", wh.Name), zap.String("url", wh.URL))
			return err
		}
	})
}

// buildRequest 构建 HTTP 请求(POST)
func (wd *WebhookDispatcher) buildRequest(wh *webhookModel.Webhook, e *Event) (*http.Request, error) {
	// 构造 HTTP 请求头
	headers := make(http.Header)
	headers.Set("Content-Type", "application/json")  // 内容类型
	headers.Set("X-Ech0-Event", string(e.Type))      // 事件类型
	headers.Set("User-Agent", "Ech0-Webhook-Client") // 自定义 User-Agent
	headers.Set("E-Ech0-Event-ID", e.ID)             // 唯一事件 ID，便于接收方去重，保证幂等性

	// 构造 HTTP 请求体
	body, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	bodyReader := io.NopCloser(bytes.NewReader(body))

	// 构造 HTTP 请求
	req, err := http.NewRequest("POST", wh.URL, bodyReader)
	if err != nil {
		return nil, err
	}
	req.Header = headers

	// 设置 GetBody 以支持重试
	req.GetBody = func() (bodyReader io.ReadCloser, err error) {
		return io.NopCloser(bytes.NewReader(body)), nil
	}

	// 返回请求对象
	return req, nil
}

// retryWithBackoff 带指数退避的重试机制
func (wd *WebhookDispatcher) retryWithBackoff(maxRetries int, initialBackoff time.Duration, fn func() error) error {
	var err error
	delay := initialBackoff
	for range maxRetries {
		if err := fn(); err == nil {
			return nil // 成功
		}
		time.Sleep(delay)
		delay *= 2 // 指数退避
	}

	return err // 返回最后一次的错误
}

// Wait 等待所有事件处理完成
func (wd *WebhookDispatcher) Wait() {
	wd.pool.Wait()
}
