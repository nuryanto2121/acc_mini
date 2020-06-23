package contuser

import (
	"context"
	"fmt"
	"net/http"
	iusers "nuryanto2121/acc_mini/interface/user"
	midd "nuryanto2121/acc_mini/middleware"
	"nuryanto2121/acc_mini/models"
	app "nuryanto2121/acc_mini/pkg"
	tool "nuryanto2121/acc_mini/pkg/tools"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"

	_ "nuryanto2121/acc_mini/docs"

	echoSwagger "github.com/swaggo/echo-swagger"
)

type ContUser struct {
	useUser iusers.Usecase
}

func NewContUser(e *echo.Echo, a iusers.Usecase) {
	controller := &ContUser{
		useUser: a,
	}
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/health_check", controller.HealthCheck)

	r := e.Group("/api/user")
	r.Use(midd.JWT)
	r.GET("/:id", controller.GetDataBy)
	r.GET("", controller.GetList)
	r.POST("", controller.Create)
	r.PUT("/:id", controller.Update)
	r.DELETE("", controller.Delete)
}

// GetDataByID :
// @Summary GetById
// @Security ApiKeyAuth
// @Tags User
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} tool.ResponseModel
// @Router /api/user/{id} [get]
func (u *ContUser) GetDataBy(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		// logger = logging.Logger{}
		appE = tool.Res{R: e} // wajib
		id   = e.Param("id")  //kalo bukan int => 0
		// valid  validation.Validation                 // wajib
	)
	ID, err := strconv.Atoi(id)
	if err != nil {
		return appE.Response(http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
	}

	data, err := u.useUser.GetDataBy(ctx, ID)
	if err != nil {
		return appE.Response(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil)
	}

	return appE.Response(http.StatusOK, "Ok", data)
}

// GetList :
// @Summary GetList User
// @Security ApiKeyAuth
// @Tags User
// @Produce  json
// @Param page query int true "Page"
// @Param perpage query int true "PerPage"
// @Param search query string false "Search"
// @Param initsearch query string false "InitSearch"
// @Param sortfield query string false "SortField"
// @Success 200 {object} models.ResponseModelList
// @Router /api/user [get]
func (u *ContUser) GetList(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		// logger = logging.Logger{}
		appE = tool.Res{R: e} // wajib
		//valid      validation.Validation // wajib
		paramquery   = models.ParamList{} // ini untuk list
		responseList = models.ResponseModelList{}
		err          error
	)

	httpCode, errMsg := app.BindAndValid(e, &paramquery)
	// logger.Info(util.Stringify(paramquery))
	if httpCode != 200 {
		return appE.ResponseErrorList(http.StatusBadRequest, errMsg, responseList)
	}
	claims, err := app.GetClaims(e)
	if err != nil {
		return appE.ResponseErrorList(http.StatusBadRequest, fmt.Sprintf("%v", err), responseList)
	}
	if !claims.IsAdmin {
		paramquery.InitSearch = " id_created = " + strconv.Itoa(claims.UserID)
	}

	responseList, err = u.useUser.GetList(ctx, paramquery)
	if err != nil {
		// return e.JSON(http.StatusBadRequest, err.Error())
		return appE.ResponseErrorList(tool.GetStatusCode(err), fmt.Sprintf("%v", err), responseList)
	}

	// return e.JSON(http.StatusOK, ListDataUser)
	return appE.Response(http.StatusOK, "", responseList)
}

// CreateSaUser :
// @Summary Add User
// @Security ApiKeyAuth
// @Tags User
// @Produce json
// @Param req body models.AddUser true "req param #changes are possible to adjust the form of the registration form from frontend"
// @Success 200 {object} tool.ResponseModel
// @Router /api/user [post]
func (u *ContUser) Create(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		// logger     = logging.Logger{} // wajib
		appE    = tool.Res{R: e} // wajib
		sysUser models.SysUser
		form    models.AddUser
	)

	// user := e.Get("user").(*jwt.Token)
	// claims := user.Claims.(*util.Claims)
	// validasi and bind to struct
	httpCode, errMsg := app.BindAndValid(e, &form)
	// logger.Info(util.Stringify(form))
	if httpCode != 200 {
		return appE.ResponseError(http.StatusBadRequest, errMsg, nil)
	}
	claims, err := app.GetClaims(e)
	if err != nil {
		return appE.ResponseError(http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
	}
	// mapping to struct model saRole
	err = mapstructure.Decode(form, &sysUser)
	if err != nil {
		return appE.ResponseError(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil)

	}
	sysUser.IDCreated = claims.UserID
	err = u.useUser.Create(ctx, &sysUser)
	if err != nil {
		return appE.ResponseError(tool.GetStatusCode(err), fmt.Sprintf("%v", err), nil)
	}

	return appE.Response(http.StatusCreated, "Ok", sysUser)
}

// UpdateSaUser :
// @Summary Update User
// @Security ApiKeyAuth
// @Tags User
// @Produce json
// @Param id path string true "ID"
// @Param req body models.AddUser true "req param #changes are possible to adjust the form of the registration form from frontend"
// @Success 200 {object} tool.ResponseModel
// @Router /api/user/{id} [put]
func (u *ContUser) Update(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		// logger = logging.Logger{} // wajib
		appE = tool.Res{R: e} // wajib
		err  error
		// valid  validation.Validation                 // wajib
		id   = e.Param("id") //kalo bukan int => 0
		form = models.AddUser{}
	)
	// user := e.Get("user").(*jwt.Token)
	// claims := user.Claims.(*util.Claims)

	MenuID, _ := strconv.Atoi(id)
	// logger.Info(id)
	if err != nil {
		return appE.ResponseError(http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
	}

	// validasi and bind to struct
	httpCode, errMsg := app.BindAndValid(e, &form)
	// logger.Info(util.Stringify(form))
	if httpCode != 200 {
		return appE.ResponseError(http.StatusBadRequest, errMsg, nil)
	}

	// form.UpdatedBy = claims.UserName
	err = u.useUser.Update(ctx, MenuID, &form)
	if err != nil {
		return appE.ResponseError(tool.GetStatusCode(err), fmt.Sprintf("%v", err), nil)
	}
	return appE.Response(http.StatusCreated, "Ok", nil)
}

// DeleteSaUser :
// @Summary Delete User
// @Security ApiKeyAuth
// @Tags User
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} tool.ResponseModel
// @Router /api/user [delete]
func (u *ContUser) Delete(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		// logger = logging.Logger{}
		appE = tool.Res{R: e} // wajib
		id   = e.Param("id")  //kalo bukan int => 0
		// valid  validation.Validation                 // wajib
	)
	ID, err := strconv.Atoi(id)
	if err != nil {
		return appE.Response(http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
	}

	err = u.useUser.Delete(ctx, ID)
	if err != nil {
		return appE.Response(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil)
	}

	return appE.Response(http.StatusOK, "Ok", nil)
}

func (u *ContUser) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "success")
}
