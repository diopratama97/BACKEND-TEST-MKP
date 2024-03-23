package models

import (
	"backend-test-mkp/db"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ShowTime struct {
	Id          string `json:"id"`
	Bioskop_id  string `json:"bioskop_id" validate: "required"`
	Name        string `json:"name" validate:"required"`
	Show_Date   string `json:"show_date" validate:"required" `
	Show_Time   string `json:"show_time" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Status      string `json:"status" validate: "required"`
	Url_trailer string `json:"url_trailer" validate: "required"`
	Url_image   string `json:"url_image" validate:"required"`
	Created_at  string `json:"created_at"`
	Updated_at  string `json:"updated_at"`
}

func ListShow() (Response, error) {
	var data ShowTime
	var arrObj []ShowTime
	var res Response

	conn := db.CreateConn()

	sqlQuery := "SELECT * FROM film"

	rows, err := conn.Query(sqlQuery)

	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&data.Id, &data.Name, &data.Bioskop_id, &data.Price, &data.Show_Date, &data.Show_Time, &data.Status, &data.Url_image, &data.Url_trailer, &data.Created_at, &data.Updated_at)
		if err != nil {
			return res, err
		}

		arrObj = append(arrObj, data)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrObj

	return res, nil
}

func CreateShow(price int, bioskop_id, name, show_date, show_time, status, url_trailer, url_image string) (Response, error) {
	var res Response
	var findIdBioskop string

	v := validator.New()

	usr := ShowTime{
		Bioskop_id:  bioskop_id,
		Name:        name,
		Show_Date:   show_date,
		Show_Time:   show_time,
		Price:       price,
		Status:      status,
		Url_trailer: url_trailer,
		Url_image:   url_image,
	}

	err := v.Struct(usr)
	if err != nil {
		return res, err
	}

	conn := db.CreateConn()

	sqlQueryCheck := "SELECT id FROM bioskop WHERE id = $1"

	err2 := conn.QueryRow(sqlQueryCheck, bioskop_id).Scan(&findIdBioskop)

	if err2 != nil {
		res.Status = http.StatusNotFound
		res.Message = "Bioskop id not found"
		return res, nil
	}

	uuid := uuid.New()

	sqlQuery := "INSERT INTO film (id,created_at,updated_at,bioskop_id,name,show_date,show_time,price,status,url_trailer,url_image) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)"

	stmt, err := conn.Prepare(sqlQuery)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(uuid, time.Now(), time.Now(), bioskop_id, name, show_date, show_time, price, status, url_trailer, url_image)

	if err != nil {
		return res, err
	}
	fmt.Println(result)

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]string{
		"id": uuid.String(),
	}

	return res, nil
}

// func UpdateUsers(id int, name string, address string, telp string) (Response, error) {
// 	var res Response

// 	conn := db.CreateConn()

// 	sqlQuery := "UPDATE users SET name = ?, address = ?, telp = ? WHERE id = ?"

// 	stmt, err := conn.Prepare(sqlQuery)

// 	if err != nil {
// 		return res, err
// 	}

// 	result, err := stmt.Exec(name, address, telp, id)

// 	if err != nil {
// 		return res, err
// 	}

// 	rowsAffected, err := result.RowsAffected()

// 	if err != nil {
// 		return res, err
// 	}

// 	res.Status = http.StatusOK
// 	res.Message = "Success"
// 	res.Data = map[string]int64{
// 		"row_affected": rowsAffected,
// 	}

// 	return res, nil
// }

// func DeleteUsers(id int) (Response, error) {
// 	var res Response

// 	conn := db.CreateConn()

// 	sqlQuery := "DELETE FROM users WHERE id = ?"

// 	stmt, err := conn.Prepare(sqlQuery)
// 	if err != nil {
// 		return res, err
// 	}

// 	result, err := stmt.Exec(id)
// 	if err != nil {
// 		return res, err
// 	}

// 	rowAffected, err := result.RowsAffected()
// 	if err != nil {
// 		return res, err
// 	}

// 	res.Status = http.StatusOK
// 	res.Message = "Success"
// 	res.Data = rowAffected

// 	return res, nil
// }

// func DetailUsers(id int) (Response, error) {
// 	var res Response
// 	var obj Users

// 	conn := db.CreateConn()

// 	sqlQuery := "SELECT id,name,address,telp FROM users WHERE id = ?"

// 	err := conn.QueryRow(sqlQuery, id).Scan(&obj.Id, &obj.Name, &obj.Address, &obj.Telp)

// 	if err != nil {
// 		return res, err
// 	}

// 	res.Status = http.StatusOK
// 	res.Message = "Success"
// 	res.Data = obj

// 	return res, nil
// }
