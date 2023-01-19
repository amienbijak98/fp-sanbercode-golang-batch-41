package structs

import "time"

type User struct {
	Username string `json:"username"`
	Password string `json:"pasword"`
	FullName string `json:"full_name"`
	Saldo    int    `json:"saldo"`
}

type Transaksi struct {
	IDTransaksi    int       `json:"id_transaksi"`
	IDUser         string    `json:"id_user"`
	Tanggal        time.Time `json:"tanggal"`
	Keterangan     string    `json:"keterangan"`
	TotalTransaksi int       `json:"total_transaksi"`
	TotalBayar     int       `json:"total_bayar"`
	StatusBayar    string    `json:"status_bayar"`
}

type TransaksiDetail struct {
	IDTransaksiDetail int `json:"id_transaksi_detail"`
	IDTransaksi       int `json:"id_transaksi"`
	IDPaketCetak      int `json:"id_paket_cetak"`
	Harga             int `json:"harga"`
	Qty               int `json:"qty"`
	Total             int `json:"total"`
}

type ItemCetak struct {
	IDItemCetak int    `json:"id_item_cetak"`
	NamaItem    string `json:"nama_item"`
	HargaItem   int    `json:"harga_item"`
}
