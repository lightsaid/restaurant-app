package main

// 新增管理员入参
type createAdminRequest struct {
	Name     string `json:"name" label:"用户名 " binding:"required,min=2,max=6"`
	Phone    string `json:"phone" label:"手机号 " binding:"required,len=11"`
	Password string `json:"password" label:"密码 " binding:"required,min=6,max=16"`
	AuthCode string `json:"auth_code" label:"验证码 "  binding:"required,min=4,max=6"`
}
