package repository

import (
	"database/sql"

	"github.com/wisedevguy/fp-sanbercode-golang-batch-41/structs"
)

func GetAllUser(db *sql.DB) (results []structs.User, err error) {
	sql := "SELECT * FROM user_app"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var user = structs.User{}

		err = rows.Scan(&user.Username, &user.Password, &user.FullName, &user.Saldo, &user.Role)
		if err != nil {
			panic(err)
		}

		results = append(results, user)

	}
	return
}

func InsertUser(db *sql.DB, user structs.User) (err error) {
	sql := "INSERT INTO user_app (username, password, full_name, saldo, role) VALUES ($1, $2, $3, $4, $5)"

	errs := db.QueryRow(sql, user.Username, user.Password, user.FullName, user.Saldo, user.Role)

	return errs.Err()
}

func UpdateUser(db *sql.DB, user structs.User) (err error) {
	sql := "UPDATE user_app SET password=$1, full_name=$2, saldo=$3, role=$4 WHERE username=$5"

	errs := db.QueryRow(sql, user.Password, user.FullName, user.Saldo, user.Role, user.Username)

	return errs.Err()
}

func DeleteUser(db *sql.DB, user structs.User) (err error) {
	sql := "DELETE FROM user_app WHERE username = $1;"
	_, err = db.Exec(sql, user.Username)
	return
}

func GetUserByUsername(db *sql.DB, input structs.User) (result structs.User, err error) {
	var person = structs.User{}
	sql := "SELECT * FROM user_app WHERE username =$1;"

	errs := db.QueryRow(sql, input.Username).Scan(&person.Username, &person.Password, &person.FullName, &person.Saldo, &person.Role)
	if errs != nil {
		panic(errs)
	}
	result = person

	return
}
