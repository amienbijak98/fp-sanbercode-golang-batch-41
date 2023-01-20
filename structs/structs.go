package structs

import "time"

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Saldo    int    `json:"saldo"`
	Role     string `json:"role"`
}

type Transaksi struct {
	IDTransaksi     int64             `json:"id_transaksi"`
	IDUser          string            `json:"id_user"`
	Tanggal         time.Time         `json:"tanggal"`
	Keterangan      string            `json:"keterangan"`
	TotalTransaksi  int               `json:"total_transaksi"`
	TotalBayar      int               `json:"total_bayar"`
	StatusBayar     string            `json:"status_bayar"`
	TransaksiDetail []TransaksiDetail `json:"transaksi-detail"`
}

type TransaksiDetail struct {
	IDTransaksiDetail int64 `json:"id_transaksi_detail"`
	IDTransaksi       int64 `json:"id_transaksi"`
	IDItem            int64 `json:"id_item"`
	Harga             int   `json:"harga"`
	Qty               int   `json:"qty"`
	Total             int   `json:"total"`
}

type Item struct {
	IDItem    int64  `json:"id_item"`
	NamaItem  string `json:"nama_item"`
	HargaItem int    `json:"harga_item"`
}
