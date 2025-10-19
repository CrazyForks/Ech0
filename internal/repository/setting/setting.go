package repository

import (
	"context"

	"gorm.io/gorm"

	model "github.com/lin-snow/ech0/internal/model/setting"
	"github.com/lin-snow/ech0/internal/transaction"
)

type SettingRepository struct {
	db func() *gorm.DB
}

func NewSettingRepository(dbProvider func() *gorm.DB) SettingRepositoryInterface {
	return &SettingRepository{
		db: dbProvider,
	}
}

func (settingRepository *SettingRepository) getDB(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(transaction.TxKey).(*gorm.DB); ok {
		return tx
	}
	return settingRepository.db()
}

// ListAccessTokens 列出访问令牌
func (settingRepository *SettingRepository) ListAccessTokens(userID uint) ([]model.AccessTokenSetting, error) {
	var tokens []model.AccessTokenSetting
	// 查询所有访问令牌
	if err := settingRepository.db().Where("user_id = ?", userID).Find(&tokens).Error; err != nil {
		return nil, err
	}
	return tokens, nil
}

// CreateAccessToken 创建访问令牌
func (settingRepository *SettingRepository) CreateAccessToken(
	ctx context.Context,
	token *model.AccessTokenSetting,
) error {
	db := settingRepository.getDB(ctx)
	return db.Create(token).Error
}

// DeleteAccessTokenByID 删除访问令牌
func (settingRepository *SettingRepository) DeleteAccessTokenByID(ctx context.Context, id uint) error {
	db := settingRepository.getDB(ctx)
	return db.Delete(&model.AccessTokenSetting{}, id).Error
}
