package model

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// MyClaims 是自定义的 JWT 声明结构体
type MyClaims struct {
	Userid   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

const (
	// MAX_USER_COUNT 定义最大用户数量
	MAX_USER_COUNT = 5
	// NO_USER_LOGINED 定义未登录用户的 ID
	NO_USER_LOGINED = uint(0)
)

type OAuth2Action string
type AuthType string

const (
	// OAuth2ActionLogin 表示登录操作
	OAuth2ActionLogin OAuth2Action = "login"
	// OAuth2ActionRegister 表示注册操作
	OAuth2ActionRegister OAuth2Action = "register"
	// OAuth2ActionBind 表示绑定操作
	OAuth2ActionBind OAuth2Action = "bind"

	AuthTypeOAuth2 AuthType = "oauth2"
	AuthTypeOIDC   AuthType = "oidc"
)

type OAuthState struct {
	Action   string `json:"action"`
	UserID   uint   `json:"user_id,omitempty"`
	Nonce    string `json:"nonce"`
	Redirect string `json:"redirect,omitempty"`
	Exp      int64  `json:"exp"`
	Provider string `json:"provider,omitempty"`
}

// GitHubTokenResponse GitHub token 响应结构
type GitHubTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

// GitHubUser GitHub 用户信息
type GitHubUser struct {
	ID        int64  `json:"id"`
	Login     string `json:"login"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
}

// GoogleTokenResponse Google token 响应结构
type GoogleTokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	IDToken      string `json:"id_token"`
}

// GoogleUser Google 用户信息
type GoogleUser struct {
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
}

// QQTokenResponse QQ token 响应结构
type QQTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenID       string `json:"openid,omitempty"`
}

// QQOpenIDResponse QQ OpenID 响应结构
type QQOpenIDResponse struct {
	ClientID string `json:"client_id"`
	OpenID   string `json:"openid"`
}

// QQUser QQ 用户信息
type QQUser struct {
	Nickname     string `json:"nickname"`
	FigureURL    string `json:"figureurl"`
	FigureURL1   string `json:"figureurl_1"`
	FigureURL2   string `json:"figureurl_2"`
	FigureURLQQ1 string `json:"figureurl_qq_1"`
	FigureURLQQ2 string `json:"figureurl_qq_2"`
	Gender       string `json:"gender"`
}

// Passkey/WebAuthn 定义 Passkey/WebAuthn 实体，用于存储 Passkey/WebAuthn 凭证信息和绑定已有用户
type Passkey struct {
	ID           uint   `gorm:"primaryKey"`
	UserID       uint   `gorm:"not null;index"`
	CredentialID string `gorm:"size:255;not null;uniqueIndex:uid_cred"`
	PublicKey    string `gorm:"type:text;not null"` // 或 []byte
	SignCount    uint32 `gorm:"not null;default:0"`
	LastUsedAt   time.Time
	DeviceName   string `gorm:"size:128"`
	AAGUID       string `gorm:"size:36"`
}
