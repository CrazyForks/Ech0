package repository

import (
	"context"

	"gorm.io/gorm"

	model "github.com/lin-snow/ech0/internal/model/webhook"
	"github.com/lin-snow/ech0/internal/transaction"
)

type WebhookRepository struct {
	db func() *gorm.DB
}

func NewWebhookRepository(dbProvider func() *gorm.DB) WebhookRepositoryInterface {
	return &WebhookRepository{
		db: dbProvider,
	}
}

func (webhookRepository *WebhookRepository) getDB(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(transaction.TxKey).(*gorm.DB); ok {
		return tx
	}
	return webhookRepository.db()
}

// CreateWebhook 创建一个webhook
func (webhookRepository *WebhookRepository) CreateWebhook(ctx context.Context, webhook *model.Webhook) error {
	if err := webhookRepository.getDB(ctx).Create(webhook).Error; err != nil {
		return err
	}

	return nil
}

// GetAllWebhooks 获取所有webhooks
func (webhookRepository *WebhookRepository) GetAllWebhooks() ([]model.Webhook, error) {
	var webhooks []model.Webhook
	if err := webhookRepository.db().Find(&webhooks).Error; err != nil {
		return nil, err
	}

	return webhooks, nil
}

// DeleteWebhookByID 根据ID删除webhook
func (webhookRepository *WebhookRepository) DeleteWebhookByID(ctx context.Context, id uint) error {
	if err := webhookRepository.getDB(ctx).Delete(&model.Webhook{}, id).Error; err != nil {
		return err
	}

	return nil
}
