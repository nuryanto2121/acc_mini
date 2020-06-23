package usebarangh

import (
	"context"
	"math"
	ibarangh "nuryanto2121/acc_mini/interface/barang_header"
	"nuryanto2121/acc_mini/models"
	querywhere "nuryanto2121/acc_mini/pkg/query"
	"reflect"
	"time"

	"github.com/mitchellh/mapstructure"
)

type useMstBarangH struct {
	repoBarangH    ibarangh.Repository
	contextTimeOut time.Duration
}

func NewMstBarangH(a ibarangh.Repository, timeout time.Duration) ibarangh.Usecase {
	return &useMstBarangH{repoBarangH: a, contextTimeOut: timeout}
}

func (u *useMstBarangH) GetDataBy(ctx context.Context, ID int) (result *models.MstBarangH, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	result, err = u.repoBarangH.GetDataBy(ID)
	if err != nil {
		return result, err
	}

	return result, nil
}
func (u *useMstBarangH) GetList(ctx context.Context, queryparam models.ParamList) (result models.ResponseModelList, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	var tBarangH = models.MstBarangH{}
	/*membuat Where like dari struct*/
	if queryparam.Search != "" {
		value := reflect.ValueOf(tBarangH)
		types := reflect.TypeOf(&tBarangH)
		queryparam.Search = querywhere.GetWhereLikeStruct(value, types, queryparam.Search, "")
	}
	result.Data, err = u.repoBarangH.GetList(queryparam)
	if err != nil {
		return result, err
	}

	result.Total, err = u.repoBarangH.Count(queryparam)
	if err != nil {
		return result, err
	}

	// d := float64(result.Total) / float64(queryparam.PerPage)
	result.LastPage = int(math.Ceil(float64(result.Total) / float64(queryparam.PerPage)))
	result.Page = queryparam.Page

	return result, nil
}
func (u *useMstBarangH) Create(ctx context.Context, data *models.MstBarangH) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	err = u.repoBarangH.Create(data)
	if err != nil {
		return err
	}
	return nil

}
func (u *useMstBarangH) Update(ctx context.Context, ID int, data interface{}) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	var form = models.MstBarangH{}
	err = mapstructure.Decode(data, &form)
	if err != nil {
		return err
		// return appE.ResponseError(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil)

	}
	err = u.repoBarangH.Update(ID, form)
	return nil
}
func (u *useMstBarangH) Delete(ctx context.Context, ID int) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	err = u.repoBarangH.Delete(ID)
	if err != nil {
		return err
	}
	return nil
}
