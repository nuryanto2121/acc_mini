package routes

import (
	"nuryanto2121/acc_mini/pkg/postgresdb"
	"nuryanto2121/acc_mini/pkg/setting"

	_contUser "nuryanto2121/acc_mini/controllers/user"
	_repoUser "nuryanto2121/acc_mini/repository/sys_user"
	_useUser "nuryanto2121/acc_mini/usecase/sys_user"

	_contBarangH "nuryanto2121/acc_mini/controllers/barang_header"
	_repoBarangH "nuryanto2121/acc_mini/repository/barang_header"
	_useBarangH "nuryanto2121/acc_mini/usecase/barang_header"

	_contBarangD "nuryanto2121/acc_mini/controllers/barang_detail"
	_repoBarangD "nuryanto2121/acc_mini/repository/barang_detail"
	_useBarangD "nuryanto2121/acc_mini/usecase/barang_detail"

	_contTransaksi "nuryanto2121/acc_mini/controllers/transaksi"
	_repoTransaksi "nuryanto2121/acc_mini/repository/transaksi"
	_useTransaksi "nuryanto2121/acc_mini/usecase/transaksi"

	_contPengeluaran "nuryanto2121/acc_mini/controllers/pengeluaran"
	_repoPengeluaran "nuryanto2121/acc_mini/repository/pengeluaran"
	_usePengeluaran "nuryanto2121/acc_mini/usecase/pengeluaran"

	_saauthcont "nuryanto2121/acc_mini/controllers/auth"
	_authuse "nuryanto2121/acc_mini/usecase/auth"

	"time"

	"github.com/labstack/echo/v4"
)

//Echo :
type EchoRoutes struct {
	E *echo.Echo
}

func (e *EchoRoutes) InitialRouter() {
	timeoutContext := time.Duration(setting.FileConfigSetting.Server.ReadTimeout) * time.Second

	repoUser := _repoUser.NewRepoSysUser(postgresdb.Conn)
	useUser := _useUser.NewUserSysUser(repoUser, timeoutContext)
	_contUser.NewContUser(e.E, useUser)

	repoBarangH := _repoBarangH.NewRepoMstBarangH(postgresdb.Conn)
	useBarangH := _useBarangH.NewMstBarangH(repoBarangH, timeoutContext)
	_contBarangH.NewContBarangH(e.E, useBarangH)

	repoBarangD := _repoBarangD.NewRepoMstBarangD(postgresdb.Conn)
	useBarangD := _useBarangD.NewMstBarangD(repoBarangD, timeoutContext)
	_contBarangD.NewContBarangD(e.E, useBarangD)

	repoTransaksi := _repoTransaksi.NewRepoTblTransaksi(postgresdb.Conn)
	useTransaksi := _useTransaksi.NewUserTblTransaksi(repoTransaksi, timeoutContext)
	_contTransaksi.NewContTransaksi(e.E, useTransaksi)

	repoPengeluaran := _repoPengeluaran.NewRepoTblPengeluaran(postgresdb.Conn)
	usePengeluaran := _usePengeluaran.NewUserTblPengeluaran(repoPengeluaran, timeoutContext)
	_contPengeluaran.NewContPengeluaran(e.E, usePengeluaran)

	//_saauthcont
	useAuth := _authuse.NewUserAuth(repoUser, timeoutContext)
	_saauthcont.NewContAuth(e.E, useAuth)

}
