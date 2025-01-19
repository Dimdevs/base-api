package model

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Barang struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key;" json:"id"`
	Nama      string    `json:"nama"`
	Deskripsi string    `json:"deskripsi"`
	Harga     int       `json:"harga"`
	Stok      int       `json:"stok"`
}

func (barang *Barang) BeforeCreate(tx *gorm.DB) (err error) {
	if barang.ID == uuid.Nil {
		barang.ID = uuid.New()
	}
	return
}
