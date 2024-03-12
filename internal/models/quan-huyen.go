package models

import (
	"database/sql"
	"strconv"
)

type QuanHuyen struct {
	ID        int       `json:"id"`
	Ten       string    `json:"ten"`
	TinhThanh TinhThanh `json:"tinh_thanh"`
}

func DocQuanHuyenTheoTinhThanhCSDL(id int) ([]QuanHuyen, error) {
	dsQuanHuyen := []QuanHuyen{}

	rows, err := db.Query("SELECT * FROM quan_huyen WHERE tinh_thanh = ? ORDER BY ten ASC", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var quanHuyen QuanHuyen
		var idTinhThanh int
		if err := rows.Scan(&quanHuyen.ID, &quanHuyen.Ten, &idTinhThanh); err != nil {
			return nil, err
		}
		tinhThanh, err := DocTinhThanhTheoIdCSDL(idTinhThanh)
		if err == nil {
			quanHuyen.TinhThanh = tinhThanh
		}
		dsQuanHuyen = append(dsQuanHuyen, quanHuyen)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return dsQuanHuyen, nil
}

func DocQuanHuyenTheoIdCSDL(id int) (QuanHuyen, error) {
	var quanHuyen QuanHuyen
	var idTinhThanh int

	row := db.QueryRow("SELECT * FROM quan_huyen WHERE id = ?", strconv.Itoa(id))
	if err := row.Scan(&quanHuyen.ID, &quanHuyen.Ten, &idTinhThanh); err != nil {
		if err == sql.ErrNoRows {
			return quanHuyen, err
		}
		return quanHuyen, err
	}
	tinhThanh, err := DocTinhThanhTheoIdCSDL(idTinhThanh)
	if err == nil {
		quanHuyen.TinhThanh = tinhThanh
	}
	return quanHuyen, nil
}

func ThemQuanHuyenCSDL(quanHuyen QuanHuyen) (QuanHuyen, error) {
	_, err := db.Exec("INSERT INTO quan_huyen VALUES (?, ?, ?)", quanHuyen.ID, quanHuyen.Ten, quanHuyen.TinhThanh.ID)
	return quanHuyen, err
}

func CapNhatQuanHuyenCSDL(quanHuyen QuanHuyen) error {
	_, err := db.Exec("UPDATE quan_huyen SET ten = ?, tinh_thanh = ? WHERE id = ?", quanHuyen.Ten, quanHuyen.TinhThanh.ID, quanHuyen.ID)
	return err
}

func XoaQuanHuyenCSDL(id int) error {
	_, err := db.Exec("DELETE FROM quan_huyen WHERE id = ?", id)
	return err
}
