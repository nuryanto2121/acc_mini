package models

type TblTransaksi struct {
	Model
	Pembeli    string  `json:"pembeli" gorm:"type:varchar(60);not null"`
	BarangCD   string  `json:"barang_cd" gorm:"type:varchar(20);not null"`
	HargaModal float32 `json:"harga_modal"`
	HargaJual  float32 `json:"harga_jual"`
	Keuntungan float32 `json:"keuntungan"`
}
