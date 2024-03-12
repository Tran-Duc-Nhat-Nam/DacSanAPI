package models

import (
	"database/sql"
	"strconv"
)

type PhuongXa struct {
	ID        int       `json:"id"`
	Ten       string    `json:"ten"`
	QuanHuyen QuanHuyen `json:"quan_huyen"`
}

func DocPhuongXaTheoQuanHuyenCSDL(id int) ([]PhuongXa, error) {
	dsPhuongXa := []PhuongXa{}

	rows, err := db.Query("SELECT * FROM phuong_xa WHERE quan_huyen = ? ORDER BY ten ASC", id)
	if err != nil {
		return dsPhuongXa, err
	}
	defer rows.Close()

	for rows.Next() {
		var phuongXa PhuongXa
		var idQuanHuyen int
		if err := rows.Scan(&phuongXa.ID, &phuongXa.Ten, &idQuanHuyen); err != nil {
			return dsPhuongXa, err
		}
		quanHuyen, err := DocQuanHuyenTheoIdCSDL(idQuanHuyen)
		if err == nil {
			phuongXa.QuanHuyen = quanHuyen
		}
		dsPhuongXa = append(dsPhuongXa, phuongXa)
	}

	if err := rows.Err(); err != nil {
		return dsPhuongXa, err
	}

	return dsPhuongXa, nil
}

func DocPhuongXaTheoIdCSDL(id int) (PhuongXa, error) {
	var phuongXa PhuongXa
	var idQuanHuyen int

	row := db.QueryRow("SELECT * FROM phuong_xa WHERE id = ?", strconv.Itoa(id))
	if err := row.Scan(&phuongXa.ID, &phuongXa.Ten, &idQuanHuyen); err != nil {
		if err == sql.ErrNoRows {
			return phuongXa, err
		}
		return phuongXa, err
	}
	quanHuyen, err := DocQuanHuyenTheoIdCSDL(idQuanHuyen)
	if err == nil {
		phuongXa.QuanHuyen = quanHuyen
	}
	return phuongXa, nil
}

func ThemPhuongXaCSDL(phuongXa PhuongXa) (PhuongXa, error) {
	_, err := db.Exec("INSERT INTO phuong_xa VALUES (?, ?, ?)", phuongXa.ID, phuongXa.Ten, phuongXa.QuanHuyen.ID)
	return phuongXa, err
}

func CapNhatPhuongXaCSDL(phuongXa PhuongXa) error {
	_, err := db.Exec("UPDATE phuong_xa SET ten = ?, quan_huyen = ? WHERE id = ?", phuongXa.Ten, phuongXa.QuanHuyen.ID, phuongXa.ID)
	return err
}

func XoaPhuongXaCSDL(id int) error {
	_, err := db.Exec("DELETE FROM phuong_xa WHERE id = ?", id)
	return err
}
