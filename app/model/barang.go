package model

type Barang struct {
	ID        string `json:"id"`
	Nama      string `json:"nama"`
	Deskripsi string `json:"deskripsi"`
	Harga     int    `json:"harga"`
	Stok      int    `json:"stok"`
}
