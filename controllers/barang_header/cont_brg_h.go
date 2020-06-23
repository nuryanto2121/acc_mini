package contbarangh

import (
	"context"
	"fmt"
	"net/http"
	ibarangh "nuryanto2121/acc_mini/interface/barang_header"
	midd "nuryanto2121/acc_mini/middleware"
	"nuryanto2121/acc_mini/models"
	app "nuryanto2121/acc_mini/pkg"
	tool "nuryanto2121/acc_mini/pkg/tools"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ContBarangH struct {
	useBarangH ibarangh.Usecase
}

func NewContBarangH(e *echo.Echo, a ibarangh.Usecase) {
	controller := &ContBarangH{
		useBarangH: a,
	}

	r := e.Group("/api/barangh")
	r.Use(midd.JWT)
	r.GET("/:id", controller.GetDataBy)
	r.GET("", controller.GetList)
	r.POST("", controller.Create)
	r.PUT("/:id", controller.Update)
	r.DELETE("", controller.Delete)
}

// GetDataByID :
// @Summary GetById
// @Tags BarangH
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} tool.ResponseModel
// @Router /api/barangh/{id} [get]
func (u *ContBarangH) GetDataBy(e echo.Context) error {
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

	data, err := u.useBarangH.GetDataBy(ctx, ID)
	if err != nil {
		return appE.Response(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil)
	}

	return appE.Response(http.StatusOK, "Ok", data)
}

// GetList :
// @Summary GetList BarangH
// @Security ApiKeyAuth
// @Tags BarangH
// @Produce  json
// @Param page query int true "Page"
// @Param perpage query int true "PerPage"
// @Param search query string false "Search"
// @Param initsearch query string false "InitSearch"
// @Param sortfield query string false "SortField"
// @Success 200 {object} models.ResponseModelList
// @Router /api/barangh [get]
func (u *ContBarangH) GetList(e echo.Context) error {
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
		return appE.ResponseError(http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
	}
	if !claims.IsAdmin {
		paramquery.InitSearch = " id_created = " + strconv.Itoa(claims.UserID)
	}
	responseList, err = u.useBarangH.GetList(ctx, paramquery)
	if err != nil {
		// return e.JSON(http.StatusBadRequest, err.Error())
		return appE.ResponseErrorList(tool.GetStatusCode(err), fmt.Sprintf("%v", err), responseList)
	}

	// return e.JSON(http.StatusOK, ListDataBarangH)
	return appE.Response(http.StatusOK, "", responseList)
}

// CreateSaBarangH :
// @Summary Add BarangH
// @Security ApiKeyAuth
// @Tags BarangH
// @Produce json
// @Param req body models.MstBarangH true "req param #changes are possible to adjust the form of the registration form from frontend"
// @Success 200 {object} tool.ResponseModel
// @Router /api/barangh [post]
func (u *ContBarangH) Create(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		// logger     = logging.Logger{} // wajib
		appE = tool.Res{R: e} // wajib
		// sysBarangH models.MstBarangH
		form models.MstBarangH
	)

	// barangh := e.Get("barangh").(*jwt.Token)
	// claims := barangh.Claims.(*util.Claims)
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

	form.IDCreated = claims.UserID
	err = u.useBarangH.Create(ctx, &form)
	if err != nil {
		return appE.ResponseError(tool.GetStatusCode(err), fmt.Sprintf("%v", err), nil)
	}

	return appE.Response(http.StatusCreated, "Ok", form)
}

// UpdateSaBarangH :
// @Summary Update BarangH
// @Security ApiKeyAuth
// @Tags BarangH
// @Produce json
// @Param id path string true "ID"
// @Param req body models.MstBarangH true "req param #changes are possible to adjust the form of the registration form from frontend"
// @Success 200 {object} tool.ResponseModel
// @Router /api/barangh/{id} [put]
func (u *ContBarangH) Update(e echo.Context) error {
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
		form = models.MstBarangH{}
	)
	// barangh := e.Get("barangh").(*jwt.Token)
	// claims := barangh.Claims.(*util.Claims)

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

	// form.UpdatedBy = claims.BarangHName
	err = u.useBarangH.Update(ctx, MenuID, &form)
	if err != nil {
		return appE.ResponseError(tool.GetStatusCode(err), fmt.Sprintf("%v", err), nil)
	}
	return appE.Response(http.StatusCreated, "Ok", nil)
}

// DeleteSaBarangH :
// @Summary Delete BarangH
// @Security ApiKeyAuth
// @Tags BarangH
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} tool.ResponseModel
// @Router /api/barangh [delete]
func (u *ContBarangH) Delete(e echo.Context) error {
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

	err = u.useBarangH.Delete(ctx, ID)
	if err != nil {
		return appE.Response(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil)
	}

	return appE.Response(http.StatusOK, "Ok", nil)
}

func (u *ContBarangH) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "success")
}
