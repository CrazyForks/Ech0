package service

import (
	authModel "github.com/lin-snow/ech0/internal/model/auth"
	model "github.com/lin-snow/ech0/internal/model/user"
)

type UserServiceInterface interface {
	// Login 用户登录
	Login(user *authModel.LoginDto) (string, error)

	// GetUserByID 根据用户ID获取用户信息
	GetUserByID(userId int) (model.User, error)

	// Register 用户注册
	Register(registerDto *authModel.RegisterDto) error

	// UpdateUser 更新用户信息
	UpdateUser(userid uint, userdto model.UserInfoDto) error

	// UpdateUserAdmin 更新用户的管理员权限
	UpdateUserAdmin(userid uint, id uint) error

	// GetAllUsers 获取所有用户
	GetAllUsers() ([]model.User, error)

	// GetSysAdmin 获取系统管理员
	GetSysAdmin() (model.User, error)

	// DeleteUser 删除用户
	DeleteUser(userid, id uint) error

	// BindGitHub 绑定 GitHub 账号
	BindGitHub(userID uint, redirect_URI string) (string, error)

	// 获取 GitHub 登录 URL
	GetGitHubLoginURL(redirect_URI string) (string, error)

	// HandleGitHubCallback 处理 GitHub OAuth2 回调
	HandleGitHubCallback(code string, state string) (string)

	// GetOAuthInfo 获取 OAuth2 配置信息
	GetOAuthInfo(userId uint) (model.OAuthInfoDto, error)
}
