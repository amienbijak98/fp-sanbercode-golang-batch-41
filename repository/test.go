package repository

import (
	"database/sql"
	"time"

	"github.com/wisedevguy/fp-sanbercode-golang-batch-41/structs"
)

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

func main() {
	var transaksi structs.Transaksi
	db, err := sql.Open("postgres", "postgres://username:password@host:port/database")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Create
	insertTransaksi, err := db.Prepare("INSERT INTO transaksi (id_user, tanggal, keterangan, total_transaksi, total_bayar, status_bayar) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id_transaksi")
	if err != nil {
		panic(err)
	}
	defer insertTransaksi.Close()

	var idTransaksi int64
	err = insertTransaksi.QueryRow(transaksi.Users, transaksi.Tanggal, transaksi.Keterangan, transaksi.TotalTransaksi, transaksi.TotalBayar, transaksi.StatusBayar).Scan(&idTransaksi)
	if err != nil {
		panic(err)
	}

	for _, detail := range transaksi.TransaksiDetail {
		_, err := db.Exec("INSERT INTO transaksi_detail (id_transaksi, id_item, harga, qty, total) VALUES ($1, $2, $3, $4, $5)", idTransaksi, detail.IDItem, detail.Harga, detail.Qty, detail.Total)
		if err != nil {
			panic(err)
		}
	}

	// Read
	var t Transaksi
	err = db.QueryRow("SELECT id_transaksi, id_user, tanggal, keterangan, total_transaksi, total_bayar, status_bayar FROM transaksi WHERE id_transaksi = $1", idTransaksi).Scan(&t.IDTransaksi, &t.IDUser, &t.Tanggal, &t.Keterangan, &t.TotalTransaksi, &t.TotalBayar, &t.StatusBayar)
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("SELECT id_transaksi_detail, id_transaksi, id_item, harga, qty, total FROM transaksi_detail WHERE id_transaksi = $1", idTransaksi)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var details []TransaksiDetail
	for rows.Next() {
		var detail TransaksiDetail
		err := rows.Scan(&detail.IDTransaksiDetail, &detail.IDTransaksi, &detail.IDItem, &detail.Harga, &detail.Qty, &detail.Total)
		if err != nil {
			panic(err)
		}
		details = append(details, detail)
	}

	t.TransaksiDetail = details

	// Update
	updateTransaksi, err := db.Prepare("UPDATE transaksi SET id_user = $1, tanggal = $2, keterangan = $3, total_transaksi = $4, total_bayar = $5, status_bayar = $6 WHERE id_transaksi = $7")
	if err != nil {
		panic(err)
	}
	defer updateTransaksi.Close()

	_, err = updateTransaksi.Exec(t.IDUser, t.Tanggal, t.Keterangan, t.TotalTransaksi, t.TotalBayar, t.StatusBayar, t.IDTransaksi)
	if err != nil {
		panic(err)
	}

	for _, detail := range t.TransaksiDetail {
		updateTransaksiDetail, err := db.Prepare("UPDATE transaksi_detail SET id_item = $1, harga = $2, qty = $3, total = $4 WHERE id_transaksi_detail = $5")
		if err != nil {
			panic(err)
		}
		defer updateTransaksiDetail.Close()

		_, err = updateTransaksiDetail.Exec(detail.IDItem, detail.Harga, detail.Qty, detail.Total, detail.IDTransaksiDetail)
		if err != nil {
			panic(err)
		}
	}

	// Delete
	_, err = db.Exec("DELETE FROM transaksi_detail WHERE id_transaksi = $1", idTransaksi)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("DELETE FROM transaksi WHERE id_transaksi = $1", idTransaksi)
	if err != nil {
		panic(err)
	}
}
