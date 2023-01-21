package repository

import (
	"database/sql"

	"github.com/wisedevguy/fp-sanbercode-golang-batch-41/structs"
)

// ===  TRANSAKSI CATAT PENGGUNAAN PRINT/CETAK ===

// --Transaksi Umum
func GetAllTransaksi(db *sql.DB) (results []structs.Transaksi, err error) {

	sql := "SELECT * FROM transaksi"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var trx = structs.Transaksi{}

		err = rows.Scan(&trx.IDTransaksi, &trx.Users, &trx.Tanggal, &trx.Keterangan, &trx.TotalTransaksi, &trx.TotalBayar, &trx.StatusBayar)

		results = append(results, trx)
	}

	return
}

func GetRecordCetakByID(db *sql.DB) {

}

func InsertTransaksiCetak(db *sql.DB, transaksi structs.Transaksi) (err error) {

	insertTransaksi, err := db.Prepare("INSERT INTO transaksi (users, keterangan, total_bayar) VALUES ($1, $2, $3) RETURNING id_transaksi")
	if err != nil {
		panic(err)
	}

	transaksi.Keterangan = "PRINT"
	var idTransaksi int64
	err = insertTransaksi.QueryRow(transaksi.Users, transaksi.Keterangan, transaksi.TotalBayar).Scan(&idTransaksi)
	if err != nil {
		panic(err)
	}

	var totalTrx int = 0

	for _, detail := range transaksi.TransaksiDetail {
		totalTrx += detail.Total
		_, err := db.Exec("INSERT INTO transaksi_detail (id_transaksi, id_item, harga, qty, total) VALUES ($1, $2, $3, $4, $5)", idTransaksi, detail.IDItem, detail.Harga, detail.Qty, detail.Total)
		if err != nil {
			panic(err)
		}
	}

	statusBayar := "SIMPAN SALDO"

	if transaksi.TotalBayar == totalTrx {
		statusBayar = "LUNAS"
	} else if transaksi.TotalBayar < totalTrx {
		statusBayar = "HUTANG/KURANG BAYAR"
	}

	updateTransaksi := "UPDATE transaksi SET total_transaksi=$1, status_bayar=$2 WHERE id_transaksi=$3"

	errs := db.QueryRow(updateTransaksi, totalTrx, statusBayar, idTransaksi)

	return errs.Err()
}

func DeleteTransaksiCetak(db *sql.DB) {

}

func UpdateTransaksiCetak(db *sql.DB) {

}

// === TRANSAKSI CATAT TAMBAH SALDO/AMBIL SALDO/BAYAR HUTANG ===
func InsertTambahSaldo(db *sql.DB) {

}

func UpdateAmbilSaldo(db *sql.DB) {

}

func DeleteTransaksiSaldo(db *sql.DB) {

}
