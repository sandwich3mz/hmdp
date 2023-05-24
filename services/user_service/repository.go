package user_service

import (
	"context"
	"hmdp/dto"
)

type Repository interface {
	SendCode(ctx context.Context, phone string)
	Login(ctx context.Context, loginForm dto.LoginFormDTO) string
}
