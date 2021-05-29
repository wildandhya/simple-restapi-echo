package models

import (
	"fmt"
	"go-echo/db"
	"go-echo/entity"
	"go-echo/helper"
	"net/http"
)

func GetAllUser() (helper.Response, error) {
	var obj entity.User
	var arrObj []entity.User
	var res helper.Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM user"

	rows, err := con.Query(sqlStatement)

	if err != nil {
		fmt.Println(err)
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Username, &obj.Email, &obj.Password)
		if err != nil {
			fmt.Println(err)
			return res, err
		}

		arrObj = append(arrObj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrObj

	return res, nil
}

func InsertPegawai(nama string, alamat string, telepon string) (helper.Response, error) {
	var res helper.Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO pegawai (nama, alamat, telepon) VALUES(?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama, alamat, telepon)

	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success created"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertedId,
	}

	return res, nil

}

func UpdateUser(id int, username string, email string, password string) (helper.Response, error) {
	var res helper.Response

	con := db.CreateCon()

	sqlStatement := "UPDATE user SET username= ?, email= ?, password= ? WHERE id= ?"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	hashedPassword, _ := helper.HashPassword(password)
	result, err := stmt.Exec(username, email, hashedPassword, id)

	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success update user"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil

}

func DeleteUser(id int) (helper.Response, error) {
	var res helper.Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM user WHERE id= ?"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)

	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success deleted"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil

}
