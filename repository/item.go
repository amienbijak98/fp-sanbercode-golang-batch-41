package repository

import (
	"database/sql"

	"github.com/wisedevguy/fp-sanbercode-golang-batch-41/structs"
)

func InsertItem(db *sql.DB, item structs.Item) (err error) {
	sql := "INSERT INTO item (nama_item, harga_item) VALUES ($1, $2)"

	errs := db.QueryRow(sql, item.NamaItem, item.HargaItem)

	return errs.Err()
}

func GetAllItem(db *sql.DB) (results []structs.Item, err error) {
	sql := "SELECT * FROM item"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var item = structs.Item{}

		err = rows.Scan(&item.IDItem, &item.NamaItem, &item.HargaItem)
		if err != nil {
			panic(err)
		}

		results = append(results, item)

	}
	return
}

func UpdateItem(db *sql.DB, item structs.Item) (err error) {
	sql := "UPDATE item SET nama_item=$1, harga_item=$2 WHERE id_item=$3"

	errs := db.QueryRow(sql, item.NamaItem, item.HargaItem, item.IDItem)

	return errs.Err()
}

func DeleteItem(db *sql.DB, item structs.Item) (err error) {
	sql := "DELETE FROM item WHERE id_item=$1"

	_, err = db.Exec(sql, item.IDItem)

	return
}
