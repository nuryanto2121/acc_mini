package useauth //iauth

import (
	"context"
	"errors"
	iauth "nuryanto2121/acc_mini/interface/auth"
	iusers "nuryanto2121/acc_mini/interface/user"
	"nuryanto2121/acc_mini/models"
	util "nuryanto2121/acc_mini/pkg/utils"
	"nuryanto2121/acc_mini/redisdb"
	"time"
)

type useAuht struct {
	repoUser       iusers.Repository
	contextTimeOut time.Duration
}

func NewUserAuth(a iusers.Repository, timeout time.Duration) iauth.Usecase {
	return &useAuht{
		repoUser:       a,
		contextTimeOut: timeout,
	}
}

func (u *useAuht) Login(ctx context.Context, dataLogin *models.LoginForm) (output interface{}, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	DataUser, err := u.repoUser.GetByEmailSaUser(dataLogin.UserName)
	if err != nil {
		// return util.GoutputErrCode(http.StatusUnauthorized, "Your User/Email not valid.") //appE.ResponseError(util.GetStatusCode(err), fmt.Sprintf("%v", err), nil)
		return nil, errors.New("Your Email not valid.")
	}

	if !util.ComparePassword(DataUser.Password, util.GetPassword(dataLogin.Password)) {
		return nil, errors.New("Your Password not valid.")
	}

	token, err := util.GenerateToken(DataUser.ID, DataUser.Email, DataUser.IsAdmin)
	if err != nil {
		return nil, err
	}

	redisdb.AddSession(token, DataUser.ID)

	restUser := map[string]interface{}{
		"id":        DataUser.ID,
		"email":     DataUser.Email,
		"user_name": DataUser.FullName,
		"is_admin":  DataUser.IsAdmin,
	}
	response := map[string]interface{}{
		"token":     token,
		"data_user": restUser,
	}

	return response, nil
}
func (u *useAuht) ForgotPassword(ctx context.Context, dataForgot *models.ForgotForm) (err error) {
	return nil
}

func (u *useAuht) ResetPassword(ctx context.Context, dataReset *models.ResetPasswd) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	if dataReset.Passwd != dataReset.ConfirmPasswd {
		return errors.New("Password and Confirm Password not same.")
	}

	email, err := util.ParseEmailToken(dataReset.TokenEmail)
	if err != nil {
		email = dataReset.TokenEmail
	}

	dataUser, err := u.repoUser.GetByEmailSaUser(email)
	if err != nil {
		return err
	}

	dataUser.Password, _ = util.Hash(dataReset.Passwd)

	err = u.repoUser.Update(dataUser.ID, &dataUser)
	if err != nil {
		return err
	}

	return nil
}
