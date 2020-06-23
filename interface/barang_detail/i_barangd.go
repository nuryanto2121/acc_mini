package ibarangd

import (
	"context"
	"nuryanto2121/acc_mini/models"
)

type Repository interface {
	GetDataBy(ID int) (result *models.MstBarangD, err error)
	GetList(queryparam models.ParamList) (result []*models.MstBarangD, err error)
	Create(data *models.MstBarangD) (err error)
	Update(ID int, data interface{}) (err error)
	Delete(ID int) (err error)
	Count(queryparam models.ParamList) (result int, err error)
}

type Usecase interface {
	GetDataBy(ctx context.Context, ID int) (result *models.MstBarangD, err error)
	GetList(ctx context.Context, queryparam models.ParamList) (result models.ResponseModelList, err error)
	Create(ctx context.Context, data *models.MstBarangD) (err error)
	Update(ctx context.Context, ID int, data interface{}) (err error)
	Delete(ctx context.Context, ID int) (err error)
}
