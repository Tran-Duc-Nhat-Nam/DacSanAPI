package models

import (
	"database/sql"
	"strconv"
)

type DiaChi struct {
	ID       int      `json:"id"`
	SoNha    string   `json:"so_nha"`
	TenDuong string   `json:"ten_duong"`
	PhuongXa PhuongXa `json:"phuong_xa"`
}

func DocDiaChiTheoIdCSDL(id int) (DiaChi, error) {
	var diaChi DiaChi
	var idPhuongXa int
	row := db.QueryRow("SELECT * FROM dia_chi WHERE id = ?", strconv.Itoa(id))
	if err := row.Scan(&diaChi.ID, &diaChi.SoNha, &diaChi.TenDuong, &idPhuongXa); err != nil {
		if err == sql.ErrNoRows {
			return diaChi, err
		}
		return diaChi, err
	}
	phuongXa, err := DocPhuongXaTheoIdCSDL(idPhuongXa)
	if err == nil {
		diaChi.PhuongXa = phuongXa
	}
	return diaChi, nil
}

func TimDiaChiCSDL(diaChi DiaChi) (DiaChi, error) {
	row := db.QueryRow("SELECT id FROM dia_chi WHERE so_nha = ? AND ten_duong = ? AND phuong_xa = ?", diaChi.SoNha, diaChi.TenDuong, diaChi.PhuongXa.ID)
	if err := row.Scan(&diaChi.ID); err != nil {
		if err == sql.ErrNoRows {
			return diaChi, err
		}
		return diaChi, err
	}
	return diaChi, nil
}

func ThemDiaChiCSDL(diaChi DiaChi) (DiaChi, error) {
	id := TaoIdMoi("dia_chi")
	diaChi.ID = id
	_, err := db.Exec("INSERT INTO dia_chi VALUES (?, ?, ?, ?)", id, diaChi.SoNha, diaChi.TenDuong, diaChi.PhuongXa.ID)
	return diaChi, err
}

func CapNhatDiaChiCSDL(diaChi DiaChi) error {
	_, err := db.Exec("UPDATE dia_chi SET so_nha = ?, ten_duong = ?, phuong_xa = ? WHERE id = ?", diaChi.SoNha, diaChi.TenDuong, diaChi.PhuongXa.ID)
	return err
}

func XoaDiaChiCSDL(id int) error {
	_, err := db.Exec("DELETE FROM dia_chi WHERE id = ?", id)
	return err
}
