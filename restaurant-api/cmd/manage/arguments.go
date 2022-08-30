package main

import (
	"time"

	"lightsaid.com/restaurant-app/restaurant-api/internal/model"
)

// 新增管理员入参
type createAdminRequest struct {
	Name     string `json:"name" label:"用户名 " binding:"required,min=2,max=6"`
	Phone    string `json:"phone" label:"手机号 " binding:"required,len=11"`
	Password string `json:"password" label:"密码 " binding:"required,min=6,max=16"`
	AuthCode string `json:"auth_code" label:"验证码 "  binding:"required,min=4,max=6"`
}

// 登录入参
type loginRequest struct {
	Phone    string `json:"phone" label:"手机号 " binding:"required,len=11"`
	Password string `json:"password" label:"密码 " binding:"required,min=6,max=16"`
}

// 登录出参
type loginResponse struct {
	Admin                 *model.Admin `json:"admin"`
	Token                 string       `json:"token"`
	SessionID             string       `json:"session_id"`
	ExpiredAt             time.Time    `json:"expired_at"`
	RefreshToken          string       `json:"refresh_token"`
	RefreshTokenExpiredAt time.Time    `json:"refresh_token_expired_at"`
}

// 创建商铺入参
type createShopRequest struct {
	Name string `json:"name" label:"商铺名" binding:"required,min=2,max=10"`
	Logo string `json:"logo" binding:"required"`
}

// 从新获取token入参
type renewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" label:"refresh_token" binding:"required"`
}

// 从新获取token出参
type renewAccessTokenResponse struct {
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}
