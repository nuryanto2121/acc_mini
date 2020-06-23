package usetransaksi

import (
	"context"
	"math"
	itransaksi "nuryanto2121/acc_mini/interface/transaksi"
	"nuryanto2121/acc_mini/models"
	querywhere "nuryanto2121/acc_mini/pkg/query"
	"reflect"
	"time"
)

type useTblTransaksi struct {
	repoUser       itransaksi.Repository
	contextTimeOut time.Duration
}

func NewUserTblTransaksi(a itransaksi.Repository, timeout time.Duration) itransaksi.Usecase {
	return &useTblTransaksi{repoUser: a, contextTimeOut: timeout}
}

func (u *useTblTransaksi) GetDataBy(ctx context.Context, ID int) (result *models.TblTransaksi, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	result, err = u.repoUser.GetDataBy(ID)
	if err != nil {
		return result, err
	}

	return result, nil
}
func (u *useTblTransaksi) GetList(ctx context.Context, queryparam models.ParamList) (result models.ResponseModelList, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	var tUser = models.TblTransaksi{}
	/*membuat Where like dari struct*/
	if queryparam.Search != "" {
		value := reflect.ValueOf(tUser)
		types := reflect.TypeOf(&tUser)
		queryparam.Search = querywhere.GetWhereLikeStruct(value, types, queryparam.Search, "")
	}
	result.Data, err = u.repoUser.GetList(queryparam)
	if err != nil {
		return result, err
	}

	result.Total, err = u.repoUser.Count(queryparam)
	if err != nil {
		return result, err
	}

	// d := float64(result.Total) / float64(queryparam.PerPage)
	result.LastPage = int(math.Ceil(float64(result.Total) / float64(queryparam.PerPage)))
	result.Page = queryparam.Page

	return result, nil
}
func (u *useTblTransaksi) Create(ctx context.Context, data *models.TblTransaksi) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	err = u.repoUser.Create(data)
	if err != nil {
		return err
	}
	return nil

}
func (u *useTblTransaksi) Update(ctx context.Context, ID int, data interface{}) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	// var form = models.AddUser{}
	// err = mapstructure.Decode(data, &form)
	// if err != nil {
	// 	return err
	// 	// return appE.ResponseError(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil)

	// }
	err = u.repoUser.Update(ID, data)
	return nil
}
func (u *useTblTransaksi) Delete(ctx context.Context, ID int) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	err = u.repoUser.Delete(ID)
	if err != nil {
		return err
	}
	return nil
}
