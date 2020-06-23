package contpengeluaran

import (
	"context"
	"fmt"
	"net/http"
	ipengeluaran "nuryanto2121/acc_mini/interface/pengeluaran"
	midd "nuryanto2121/acc_mini/middleware"
	"nuryanto2121/acc_mini/models"
	app "nuryanto2121/acc_mini/pkg"
	tool "nuryanto2121/acc_mini/pkg/tools"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ContPengeluaran struct {
	usePengeluaran ipengeluaran.Usecase
}

func NewContPengeluaran(e *echo.Echo, a ipengeluaran.Usecase) {
	controller := &ContPengeluaran{
		usePengeluaran: a,
	}

	r := e.Group("/api/pengeluaran")
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
// @Tags Pengeluaran
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} tool.ResponseModel
// @Router /api/pengeluaran/{id} [get]
func (u *ContPengeluaran) GetDataBy(e echo.Context) error {
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

	data, err := u.usePengeluaran.GetDataBy(ctx, ID)
	if err != nil {
		return appE.Response(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil)
	}

	return appE.Response(http.StatusOK, "Ok", data)
}

// GetList :
// @Summary GetList Pengeluaran
// @Security ApiKeyAuth
// @Tags Pengeluaran
// @Produce  json
// @Param page query int true "Page"
// @Param perpage query int true "PerPage"
// @Param search query string false "Search"
// @Param initsearch query string false "InitSearch"
// @Param sortfield query string false "SortField"
// @Success 200 {object} models.ResponseModelList
// @Router /api/pengeluaran [get]
func (u *ContPengeluaran) GetList(e echo.Context) error {
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
	responseList, err = u.usePengeluaran.GetList(ctx, paramquery)
	if err != nil {
		// return e.JSON(http.StatusBadRequest, err.Error())
		return appE.ResponseErrorList(tool.GetStatusCode(err), fmt.Sprintf("%v", err), responseList)
	}

	// return e.JSON(http.StatusOK, ListDataPengeluaran)
	return appE.Response(http.StatusOK, "", responseList)
}

// CreateSaPengeluaran :
// @Summary Add Pengeluaran
// @Security ApiKeyAuth
// @Tags Pengeluaran
// @Produce json
// @Param req body models.TblPengeluaran true "req param #changes are possible to adjust the form of the registration form from frontend"
// @Success 200 {object} tool.ResponseModel
// @Router /api/pengeluaran [post]
func (u *ContPengeluaran) Create(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		// logger     = logging.Logger{} // wajib
		appE = tool.Res{R: e} // wajib
		// sysPengeluaran models.TblPengeluaran
		form models.TblPengeluaran
	)

	// pengeluaran := e.Get("pengeluaran").(*jwt.Token)
	// claims := pengeluaran.Claims.(*util.Claims)
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
	// err := mapstructure.Decode(form, &sysPengeluaran)
	// if err != nil {
	// 	return appE.ResponseError(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil)

	// }
	form.IDCreated = claims.UserID
	err = u.usePengeluaran.Create(ctx, &form)
	if err != nil {
		return appE.ResponseError(tool.GetStatusCode(err), fmt.Sprintf("%v", err), nil)
	}

	return appE.Response(http.StatusCreated, "Ok", form)
}

// UpdateSaPengeluaran :
// @Summary Update Pengeluaran
// @Security ApiKeyAuth
// @Tags Pengeluaran
// @Produce json
// @Param id path string true "ID"
// @Param req body models.TblPengeluaran true "req param #changes are possible to adjust the form of the registration form from frontend"
// @Success 200 {object} tool.ResponseModel
// @Router /api/pengeluaran/{id} [put]
func (u *ContPengeluaran) Update(e echo.Context) error {
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
		form = models.TblPengeluaran{}
	)
	// pengeluaran := e.Get("pengeluaran").(*jwt.Token)
	// claims := pengeluaran.Claims.(*util.Claims)

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

	// form.UpdatedBy = claims.PengeluaranName
	err = u.usePengeluaran.Update(ctx, MenuID, &form)
	if err != nil {
		return appE.ResponseError(tool.GetStatusCode(err), fmt.Sprintf("%v", err), nil)
	}
	return appE.Response(http.StatusCreated, "Ok", nil)
}

// DeleteSaPengeluaran :
// @Summary Delete Pengeluaran
// @Security ApiKeyAuth
// @Tags Pengeluaran
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} tool.ResponseModel
// @Router /api/pengeluaran [delete]
func (u *ContPengeluaran) Delete(e echo.Context) error {
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

	err = u.usePengeluaran.Delete(ctx, ID)
	if err != nil {
		return appE.Response(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil)
	}

	return appE.Response(http.StatusOK, "Ok", nil)
}

func (u *ContPengeluaran) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "success")
}
