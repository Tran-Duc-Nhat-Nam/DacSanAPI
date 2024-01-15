package models

import (
	"database/sql"
	"strconv"
)

type DiaChi struct {
	ID        int       `json:"id"`
	SoNha     string    `json:"so_nha"`
	TenDuong  string    `json:"ten_duong"`
	PhuongXa  string    `json:"phuong_xa"`
	QuanHuyen string    `json:"quan_huyen"`
	TinhThanh TinhThanh `json:"tinh_thanh"`
}

func DocDiaChiTheoIdCSDL(id int) (DiaChi, error) {
	var diaChi DiaChi
	var tinhThanhId int

	row := db.QueryRow("SELECT * FROM dia_chi WHERE id = ?", strconv.Itoa(id))
	if err := row.Scan(&diaChi.ID, &diaChi.SoNha, &diaChi.TenDuong, &diaChi.PhuongXa, &diaChi.QuanHuyen, &tinhThanhId); err != nil {
		if err == sql.ErrNoRows {
			return diaChi, err
		}
		return diaChi, err
	}
	tinhThanh, err := DocTinhThanhTheoIdCSDL(tinhThanhId)
	if err == nil {
		diaChi.TinhThanh = tinhThanh
	}
	return diaChi, nil
}

func ThemDiaChiCSDL(diaChi DiaChi) error {
	var count int
	db.QueryRow("SELECT MAX(id) FROM dia_chi").Scan(&count)
	_, err := db.Exec("INSERT INTO dia_chi VALUES (?, ?, ?, ?, ?, ?)", count, diaChi.SoNha, diaChi.TenDuong, diaChi.PhuongXa, diaChi.QuanHuyen, diaChi.TinhThanh.ID)
	return err
}

func CapNhatDiaChiCSDL(diaChi DiaChi) error {
	_, err := db.Exec("UPDATE dia_chi SET sp_nha = ?, ten_duong = ?, phuong_xa = ?, quan_huyen = ?, tinh_thanh = ? WHERE id = ?", diaChi.SoNha, diaChi.TenDuong, diaChi.PhuongXa, diaChi.QuanHuyen, diaChi.TinhThanh.ID, diaChi.ID)
	return err
}

func XoaDiaChiCSDL(id int) error {
	_, err := db.Exec("DELETE FROM dia_chi WHERE id = ?", id)
	return err
}
