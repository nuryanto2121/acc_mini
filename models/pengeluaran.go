package models

import "time"

type TblPengeluaran struct {
	Model
	Tanggal     time.Time `json:"tanggal" gorm:"type:timestamp(0) without time zone;default:now()"`
	Descs       string    `json:"descs" gorm:"type:varchar(255);not null"`
	Pengeluaran float32   `json:"keuntungan"`
}
