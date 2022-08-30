package security

import (
	"crypto/md5"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const minSecretKeySize = 32

// 错误类型
var (
	ErrExpiredToken = errors.New("token 过期")
	ErrInvalidToken = errors.New("token 无效")
)

// Token 负载的数据结构
type TokenPayload struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// JWTMaker 管理 jwtToken
type JWTMaker struct {
	secret_key string
}

// NewTokenPayload  创建一个 TokenPayload 负载数据
func NewTokenPayload(userID string, duration time.Duration) *TokenPayload {
	return &TokenPayload{
		ID:        primitive.NewObjectID().Hex(),
		UserID:    userID,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
}

// Valid 验证 token 是否失效, 是为了实现 jwt.Claims 接口，使用 jwt.NewWithClaims 创建Token
func (payload *TokenPayload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

// NewJWTMaker 传递一个密钥创建一个Token管理者，密钥长度不小于32
func NewJWTMaker(secretKey string) (*JWTMaker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("jwt 密钥长度必须大于%d", minSecretKeySize-1)
	}

	bytes := md5.Sum([]byte(secretKey))
	key := fmt.Sprintf("%x", bytes)

	return &JWTMaker{secret_key: key}, nil
}

// CreateToken 创建 token
func (maker *JWTMaker) CreateToken(userID string, duration time.Duration) (string, *TokenPayload, error) {
	payload := NewTokenPayload(userID, duration)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(maker.secret_key))
	return token, payload, err
}

// VerifyToken 检查token是否有效
func (maker *JWTMaker) VerifyToken(token string) (*TokenPayload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secret_key), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &TokenPayload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*TokenPayload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
