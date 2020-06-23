package models

type MstBarangD struct {
	Model
	BarangHID  int     `json:"barang_h_id" gorm:"type:integer;not null"`
	BarangCD   string  `json:"barang_cd" gorm:"type:varchar(20);not null"`
	Descs      string  `json:"descs" gorm:"type:varchar(100)"`
	HargaModal float32 `json:"harga_modal"`
	HargaJual  float32 `json:"harga_jual"`
	Keuntungan float32 `json:"keuntungan"`
}
