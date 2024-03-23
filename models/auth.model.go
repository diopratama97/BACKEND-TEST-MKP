package models

import (
	"backend-test-mkp/db"
	"backend-test-mkp/helpers"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type RegisterValidator struct {
	Id       string `json:"id"`
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Address  string `json:"address" validate:"required"`
	Telp     string `json:"telp" validate:"required"`
	Role     string `json:"role" validate:"required"`
	Password string `json:"password" validate: "required"`
}

type LoginValidator struct {
	Id       string `json:"id"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate: "required"`
}

func Login(email, password string) (bool, error) {
	var obj LoginValidator
	var pwd string

	v := validator.New()

	usr := LoginValidator{
		Email:    email,
		Password: password,
	}

	err := v.Struct(usr)
	if err != nil {
		return false, err
	}

	conn := db.CreateConn()

	sqlQuery := "SELECT id,email,password FROM users WHERE email = $1"

	err2 := conn.QueryRow(sqlQuery, email).Scan(&obj.Id, &obj.Email, &pwd)

	if err2 == sql.ErrNoRows {
		return false, err
	}

	if err2 != nil {
		return false, err
	}

	match, err := helpers.ComparePassword(password, pwd)
	if !match {
		return false, err
	}

	return true, nil
}

func Register(email, password, address, telp, role, name string) (Response, error) {
	var res Response
	var findEmail string

	v := validator.New()

	usr := RegisterValidator{
		Name:     name,
		Address:  address,
		Telp:     telp,
		Email:    email,
		Role:     role,
		Password: password,
	}

	err := v.Struct(usr)
	if err != nil {
		return res, err
	}

	uuid := uuid.New()

	conn := db.CreateConn()

	sqlQueryCheck := "SELECT email FROM users WHERE email = $1"

	err2 := conn.QueryRow(sqlQueryCheck, email).Scan(&findEmail)

	if err2 == nil {
		res.Status = http.StatusConflict
		res.Message = "Duplicated email"
		return res, nil
	}

	sqlQuery := "INSERT INTO users (id,name,address,telp,email,role,password,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)"

	stmt, err := conn.Prepare(sqlQuery)

	if err != nil {
		fmt.Println(err)
		return res, err
	}

	hash, _ := helpers.HashPassword(password)

	result, err := stmt.Exec(uuid, name, address, telp, email, role, hash, time.Now(), time.Now())

	fmt.Println(result)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]string{
		"id": uuid.String(),
	}

	return res, nil
}
