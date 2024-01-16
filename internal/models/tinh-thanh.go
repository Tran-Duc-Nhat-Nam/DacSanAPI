package models

import (
	"database/sql"
	"strconv"
)

type TinhThanh struct {
	ID  int    `json:"id"`
	Ten string `json:"ten"`
}

func DocTinhThanhCSDL() ([]TinhThanh, error) {
	dsTinhThanh := []TinhThanh{}

	rows, err := db.Query("SELECT * FROM tinh_thanh ORDER BY id ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tinhThanh TinhThanh
		if err := rows.Scan(&tinhThanh.ID, &tinhThanh.Ten); err != nil {
			return nil, err
		}
		dsTinhThanh = append(dsTinhThanh, tinhThanh)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return dsTinhThanh, nil
}

func DocTinhThanhTheoIdCSDL(id int) (TinhThanh, error) {
	var tinhThanh TinhThanh

	row := db.QueryRow("SELECT * FROM tinh_thanh WHERE id = ?", strconv.Itoa(id))
	if err := row.Scan(&tinhThanh.ID, &tinhThanh.Ten); err != nil {
		if err == sql.ErrNoRows {
			return tinhThanh, err
		}
		return tinhThanh, err
	}
	return tinhThanh, nil
}

func ThemTinhThanhCSDL(tinhThanh TinhThanh) (TinhThanh, error) {
	_, err := db.Exec("INSERT INTO tinh_thanh VALUES (?, ?)", tinhThanh.ID, tinhThanh.Ten)
	return tinhThanh, err
}

func CapNhatTinhThanhCSDL(tinhThanh TinhThanh) error {
	_, err := db.Exec("UPDATE tinh_thanh SET ten = ? WHERE id = ?", tinhThanh.Ten, tinhThanh.ID)
	return err
}

func XoaTinhThanhCSDL(id int) error {
	_, err := db.Exec("DELETE FROM tinh_thanh WHERE id = ?", id)
	return err
}
