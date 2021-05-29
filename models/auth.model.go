package models

import (
	"database/sql"
	"go-echo/db"
	"go-echo/entity"
	"go-echo/helper"

	"github.com/go-playground/validator/v10"
)

func Register(username, email, password string) (helper.Response, error) {
	var res helper.Response

	// validation

	v := validator.New()

	user := entity.User{
		Username: username,
		Email:    email,
		Password: password,
	}

	err := v.Struct(user)
	if err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := "INSERT INTO user SET username=?, email=?, password=?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	hashPassword, _ := helper.HashPassword(password)
	result, err := stmt.Exec(username, email, hashPassword)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = 201
	res.Message = "Register Success"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertedId,
	}

	return res, nil
}

func Login(email, password string) (bool, error) {
	var obj entity.User
	var pwd string
	con := db.CreateCon()

	sqlStatement := "SELECT * FROM user WHERE email=?"

	err := con.QueryRow(sqlStatement, email).Scan(&obj.Id, &obj.Username, &obj.Email, &pwd)

	if err == sql.ErrNoRows {
		return false, err
	}

	if err != nil {
		return false, err
	}

	match, err := helper.CheckPasswordHash(password, pwd)
	if !match {
		return false, err
	}
	return true, nil
}
