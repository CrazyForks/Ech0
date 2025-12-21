package repository

import (
	"context"

	model "github.com/lin-snow/ech0/internal/model/user"
)

type UserRepositoryInterface interface {
	// GetUserByID 根据用户ID获取用户
	GetUserByID(id int) (model.User, error)

	// GetUserByUsername 根据用户名获取用户
	GetUserByUsername(username string) (model.User, error)

	// GetAllUsers 获取所有用户
	GetAllUsers() ([]model.User, error)

	// CreateUser 创建一个新的用户
	CreateUser(ctx context.Context, newUser *model.User) error

	// GetSysAdmin 获取系统管理员
	GetSysAdmin() (model.User, error)

	// UpdateUser 更新用户
	UpdateUser(ctx context.Context, user *model.User) error

	// DeleteUser 删除用户
	DeleteUser(ctx context.Context, id uint) error

	// BindOAuth 绑定 OAuth 账号
	BindOAuth(ctx context.Context, userID uint, provider, oauthID, issuer, authType string) error

	// GetUserByOAuthID 根据 OAuth 提供商和 OAuth ID 获取用户
	GetUserByOAuthID(ctx context.Context, provider, oauthID string) (model.User, error)

	// GetOAuthInfo 获取 OAuth2 信息
	GetOAuthInfo(userId uint, provider string) (model.OAuthBinding, error)

	// GetOAuthOIDCInfo 获取 OIDC 信息
	GetOAuthOIDCInfo(userId uint, provider string, issuer string) (model.OAuthBinding, error)
}
