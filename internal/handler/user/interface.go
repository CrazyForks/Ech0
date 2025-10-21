package handler

import "github.com/gin-gonic/gin"

type UserHandlerInterface interface {
	// Login 用户登录
	Login() gin.HandlerFunc

	// Register 用户注册
	Register() gin.HandlerFunc

	// UpdateUser 更新用户信息
	UpdateUser() gin.HandlerFunc

	// UpdateUserAdmin 更新用户权限
	UpdateUserAdmin() gin.HandlerFunc

	// GetAllUsers 获取所有用户
	GetAllUsers() gin.HandlerFunc

	// DeleteUser 删除用户
	DeleteUser() gin.HandlerFunc

	// GetUserInfo 获取用户信息
	GetUserInfo() gin.HandlerFunc

	// GitHubLogin 处理 GitHub OAuth2 登录请求
	GitHubLogin() gin.HandlerFunc

	// GitHubCallback 处理 GitHub OAuth2 回调
	GitHubCallback() gin.HandlerFunc

	// BindGitHub 绑定 GitHub 账号
	BindGitHub() gin.HandlerFunc

	// GoogleLogin 处理 Google OAuth2 登录请求
	GoogleLogin() gin.HandlerFunc

	// GoogleCallback 处理 Google OAuth2 回调
	GoogleCallback() gin.HandlerFunc

	// BindGoogle 绑定 Google 账号
	BindGoogle() gin.HandlerFunc

	// CustomOAuthLogin 处理自定义 OAuth2 登录请求
	CustomOAuthLogin() gin.HandlerFunc

	// CustomOAuthCallback 处理自定义 OAuth2 回调
	CustomOAuthCallback() gin.HandlerFunc

	// BindCustomOAuth 绑定自定义 OAuth2 账号
	BindCustomOAuth() gin.HandlerFunc

	// GetOAuthInfo 获取 OAuth2 配置信息
	GetOAuthInfo() gin.HandlerFunc
}
