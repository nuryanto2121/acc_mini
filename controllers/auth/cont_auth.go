package contauth

import (
	"context"
	"fmt"
	"net/http"
	iauth "nuryanto2121/acc_mini/interface/auth"
	"nuryanto2121/acc_mini/models"
	app "nuryanto2121/acc_mini/pkg"
	tool "nuryanto2121/acc_mini/pkg/tools"

	"github.com/labstack/echo/v4"
)

type ContAuth struct {
	useAuth iauth.Usecase
}

func NewContAuth(e *echo.Echo, useAuth iauth.Usecase) {
	cont := &ContAuth{
		useAuth: useAuth,
		// useSaClient:     useSaClient,
		// useSaUser:       useSaUser,
		// useSaFileUpload: useSaFileUpload,
	}

	// e.POST("/api/auth/register", cont.Register)
	e.POST("/api/auth/login", cont.Login)
	// e.POST("/api/auth/forgot", cont.ForgotPassword)
	e.POST("/api/auth/reset", cont.ResetPasswd)
	// e.POST("/api/auth/verify", cont.Verify)
}

// Register :
// @Summary Login
// @Tags Auth
// @Produce json
// @Param req body models.LoginForm true "req param #changes are possible to adjust the form of the registration form from frontend"
// @Success 200 {object} tool.ResponseModel
// @Router /api/auth/login [post]
func (u *ContAuth) Login(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		// logger = logging.Logger{} // wajib
		appE = tool.Res{R: e} // wajib
		// client sa_models.SaClient

		form = models.LoginForm{}
		// dataFiles = sa_models.SaFileOutput{}
	)

	// validasi and bind to struct
	httpCode, errMsg := app.BindAndValid(e, &form)
	// logger.Info(util.Stringify(form))
	if httpCode != 200 {
		return appE.ResponseError(http.StatusBadRequest, errMsg, nil)
	}

	out, err := u.useAuth.Login(ctx, &form)
	if err != nil {
		// return appE.Response(out)
		// return appE.ResponseError(util.GetStatusCode(err), fmt.Sprintf("%v", err), nil)
		return appE.ResponseError(http.StatusUnauthorized, fmt.Sprintf("%v", err), nil)
	}

	return appE.Response(http.StatusOK, "Ok", out)
}

// Register :
// @Summary Reset Password
// @Tags Auth
// @Produce json
// @Param req body models.ResetPasswd true "req param #changes are possible to adjust the form of the registration form from frontend"
// @Success 200 {object} tool.ResponseModel
// @Router /api/auth/reset [post]
func (u *ContAuth) ResetPasswd(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		// logger = logging.Logger{} // wajib
		appE = tool.Res{R: e} // wajib
		// client sa_models.SaClient

		form = models.ResetPasswd{}
	)
	httpCode, errMsg := app.BindAndValid(e, &form)
	// logger.Info(util.Stringify(form))
	if httpCode != 200 {
		return appE.ResponseError(http.StatusBadRequest, errMsg, nil)
	}
	err := u.useAuth.ResetPassword(ctx, &form)
	if err != nil {
		return appE.ResponseError(http.StatusUnauthorized, fmt.Sprintf("%v", err), nil)
	}

	return appE.Response(http.StatusOK, "Ok", "Please Login")
}
