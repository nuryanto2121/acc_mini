package usebarangd

import (
	"context"
	"math"
	ibarangd "nuryanto2121/acc_mini/interface/barang_detail"
	"nuryanto2121/acc_mini/models"
	querywhere "nuryanto2121/acc_mini/pkg/query"
	"reflect"
	"time"

	"github.com/mitchellh/mapstructure"
)

type useMstBarangD struct {
	repoBarangD    ibarangd.Repository
	contextTimeOut time.Duration
}

func NewMstBarangD(a ibarangd.Repository, timeout time.Duration) ibarangd.Usecase {
	return &useMstBarangD{repoBarangD: a, contextTimeOut: timeout}
}

func (u *useMstBarangD) GetDataBy(ctx context.Context, ID int) (result *models.MstBarangD, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	result, err = u.repoBarangD.GetDataBy(ID)
	if err != nil {
		return result, err
	}

	return result, nil
}
func (u *useMstBarangD) GetList(ctx context.Context, queryparam models.ParamList) (result models.ResponseModelList, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	var tBarangD = models.MstBarangD{}
	/*membuat Where like dari struct*/
	if queryparam.Search != "" {
		value := reflect.ValueOf(tBarangD)
		types := reflect.TypeOf(&tBarangD)
		queryparam.Search = querywhere.GetWhereLikeStruct(value, types, queryparam.Search, "")
	}
	result.Data, err = u.repoBarangD.GetList(queryparam)
	if err != nil {
		return result, err
	}

	result.Total, err = u.repoBarangD.Count(queryparam)
	if err != nil {
		return result, err
	}

	// d := float64(result.Total) / float64(queryparam.PerPage)
	result.LastPage = int(math.Ceil(float64(result.Total) / float64(queryparam.PerPage)))
	result.Page = queryparam.Page

	return result, nil
}
func (u *useMstBarangD) Create(ctx context.Context, data *models.MstBarangD) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	err = u.repoBarangD.Create(data)
	if err != nil {
		return err
	}
	return nil

}
func (u *useMstBarangD) Update(ctx context.Context, ID int, data interface{}) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	var form = models.MstBarangD{}
	err = mapstructure.Decode(data, &form)
	if err != nil {
		return err
		// return appE.ResponseError(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil)

	}
	err = u.repoBarangD.Update(ID, form)
	return nil
}
func (u *useMstBarangD) Delete(ctx context.Context, ID int) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	err = u.repoBarangD.Delete(ID)
	if err != nil {
		return err
	}
	return nil
}
