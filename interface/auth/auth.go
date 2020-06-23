package iauth

import (
	"context"
	"nuryanto2121/acc_mini/models"
)

type Usecase interface {
	Login(ctx context.Context, dataLogin *models.LoginForm) (output interface{}, err error)
	ForgotPassword(ctx context.Context, dataForgot *models.ForgotForm) (err error)
	ResetPassword(ctx context.Context, dataReset *models.ResetPasswd) (err error)
}
