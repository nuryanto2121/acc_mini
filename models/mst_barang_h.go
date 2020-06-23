package models

type MstBarangH struct {
	Model
	GroupBarang string `json:"group_barang" gorm:"type:varchar(100);not null"`
}
