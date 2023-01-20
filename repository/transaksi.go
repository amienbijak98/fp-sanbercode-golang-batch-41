package repository

import (
	"database/sql"

	"github.com/wisedevguy/fp-sanbercode-golang-batch-41/structs"
)

// ===  TRANSAKSI CATAT PENGGUNAAN PRINT/CETAK ===

// --Transaksi Umum
func GetAllRecordCetak(db *sql.DB, results []structs.Transaksi) {
	sql := "SELECT * FROM transaksi"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var trx = structs.Transaksi{}

		//sql2 := "SELECT * FROM"

		err = rows.Scan(&trx.IDTransaksi, &trx.IDUser, &trx.Tanggal, &trx.Keterangan, &trx.TotalTransaksi, &trx.TotalBayar, &trx.StatusBayar)
		if err != nil {
			panic(err)
		}

		results = append(results, trx)
	}

	return
}

func GetRecordCetakByID(db *sql.DB) {

}

func InsertTransaksiCetak(db *sql.DB) {

}

func DeleteTransaksiCetak(db *sql.DB) {

}

func UpdateTransaksiCetak(db *sql.DB) {

}

// --Transaksi Detail
// func GetAllTransaksiDetail(db *sql.DB, trx structs.Transaksi) (results []structs.TransaksiDetail, err error) {

// }

// === TRANSAKSI CATAT TAMBAH SALDO/AMBIL SALDO/BAYAR HUTANG ===
func InsertTambahSaldo(db *sql.DB) {

}

func UpdateAmbilSaldo(db *sql.DB) {

}

func DeleteTransaksiSaldo(db *sql.DB) {

}
