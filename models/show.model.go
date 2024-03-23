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
	Id          string    `json:"id"`
	Bioskop_id  string    `json:"bioskop_id" validate: "required"`
	Name        string    `json:"name" validate:"required"`
	Show_Date   string    `json:"show_date" validate:"required" `
	Show_Time   string    `json:"show_time" validate:"required"`
	Price       int       `json:"price" validate:"required"`
	Status      string    `json:"status" validate: "required"`
	Url_trailer string    `json:"url_trailer" validate: "required"`
	Url_image   string    `json:"url_image" validate:"required"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}

func ListShow() (Response, error) {
	var data ShowTime
	var arrObj []ShowTime
	var res Response

	conn := db.CreateConn()

	sqlQuery := "SELECT id,bioskop_id,name,price,show_date,show_time,status,url_image,url_trailer,created_at,updated_at FROM film"

	rows, err := conn.Query(sqlQuery)

	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&data.Id, &data.Bioskop_id, &data.Name, &data.Price, &data.Show_Date, &data.Show_Time, &data.Status, &data.Url_image, &data.Url_trailer, &data.Created_at, &data.Updated_at)
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

	shw := ShowTime{
		Bioskop_id:  bioskop_id,
		Name:        name,
		Show_Date:   show_date,
		Show_Time:   show_time,
		Price:       price,
		Status:      status,
		Url_trailer: url_trailer,
		Url_image:   url_image,
	}

	err := v.Struct(shw)
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

func UpdateShow(price int, id, bioskop_id, name, show_date, show_time, status, url_trailer, url_image string) (Response, error) {
	var res Response
	var findIdBioskop string
	var findIdFilm string

	v := validator.New()

	shw := ShowTime{
		Bioskop_id:  bioskop_id,
		Name:        name,
		Show_Date:   show_date,
		Show_Time:   show_time,
		Price:       price,
		Status:      status,
		Url_trailer: url_trailer,
		Url_image:   url_image,
	}

	err := v.Struct(shw)
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

	sqlQueryCheck2 := "SELECT id FROM film WHERE id = $1"

	err3 := conn.QueryRow(sqlQueryCheck2, id).Scan(&findIdFilm)

	if err3 != nil {
		res.Status = http.StatusNotFound
		res.Message = "ID not found"
		return res, nil
	}

	sqlQuery := "UPDATE film SET name = $1, price = $2, bioskop_id = $3, show_date = $4,show_time = $5,status = $6,url_trailer = $7,url_image = $8 WHERE id = $9"

	stmt, err := conn.Prepare(sqlQuery)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, price, bioskop_id, show_date, show_time, status, url_trailer, url_image, id)

	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"row_affected": rowsAffected,
	}

	return res, nil
}

func DeleteShow(id string) (Response, error) {
	var res Response

	conn := db.CreateConn()

	sqlQuery := "DELETE FROM film WHERE id = $1"

	stmt, err := conn.Prepare(sqlQuery)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = rowAffected

	return res, nil
}

func DetailShow(id string) (Response, error) {
	var res Response
	var obj ShowTime

	conn := db.CreateConn()

	sqlQuery := "SELECT * FROM film WHERE id = $1"

	err := conn.QueryRow(sqlQuery, id).Scan(&obj.Id, &obj.Created_at, &obj.Updated_at, &obj.Bioskop_id, &obj.Name, &obj.Show_Date, &obj.Show_Time, &obj.Price, &obj.Status, &obj.Url_trailer, &obj.Url_image)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = obj

	return res, nil
}
