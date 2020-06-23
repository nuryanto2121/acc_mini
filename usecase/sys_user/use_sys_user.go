package usesysuser

import (
	"context"
	"math"
	iusers "nuryanto2121/acc_mini/interface/user"
	"nuryanto2121/acc_mini/models"
	querywhere "nuryanto2121/acc_mini/pkg/query"
	"reflect"
	"time"

	"github.com/mitchellh/mapstructure"
)

type useSysUser struct {
	repoUser       iusers.Repository
	contextTimeOut time.Duration
}

func NewUserSysUser(a iusers.Repository, timeout time.Duration) iusers.Usecase {
	return &useSysUser{repoUser: a, contextTimeOut: timeout}
}

func (u *useSysUser) GetByEmailSaUser(ctx context.Context, email string) (result models.SysUser, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	a := models.SysUser{}
	result, err = u.repoUser.GetByEmailSaUser(email)
	if err != nil {
		return a, err
	}
	return result, nil
}

func (u *useSysUser) GetDataBy(ctx context.Context, ID int) (result *models.SysUser, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	result, err = u.repoUser.GetDataBy(ID)
	if err != nil {
		return result, err
	}

	return result, nil
}
func (u *useSysUser) GetList(ctx context.Context, queryparam models.ParamList) (result models.ResponseModelList, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	var tUser = models.SysUser{}
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
func (u *useSysUser) Create(ctx context.Context, data *models.SysUser) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	err = u.repoUser.Create(data)
	if err != nil {
		return err
	}
	return nil

}
func (u *useSysUser) Update(ctx context.Context, ID int, data interface{}) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	var form = models.AddUser{}
	err = mapstructure.Decode(data, &form)
	if err != nil {
		return err
		// return appE.ResponseError(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil)

	}
	err = u.repoUser.Update(ID, form)
	return nil
}
func (u *useSysUser) Delete(ctx context.Context, ID int) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	err = u.repoUser.Delete(ID)
	if err != nil {
		return err
	}
	return nil
}
