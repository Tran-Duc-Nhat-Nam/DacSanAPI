package models

import (
	"database/sql"
)

type DiaChi struct {
	ID       int      `json:"id"`
	SoNha    string   `json:"so_nha"`
	TenDuong string   `json:"ten_duong"`
	PhuongXa PhuongXa `json:"phuong_xa"`
}

func DocDiaChiTheoId(id int) (DiaChi, error) {
	var diaChi DiaChi
	var idPhuongXa int
	row := db.QueryRow("SELECT * FROM dia_chi WHERE id = ?", id)
	if err := row.Scan(&diaChi.ID, &diaChi.SoNha, &diaChi.TenDuong, &idPhuongXa); err != nil {
		if err == sql.ErrNoRows {
			return diaChi, err
		}
		return diaChi, err
	}
	phuongXa, err := DocPhuongXaTheoId(idPhuongXa)
	if err == nil {
		diaChi.PhuongXa = phuongXa
	}
	return diaChi, nil
}

func TimDiaChi(diaChi DiaChi) (DiaChi, error) {
	row := db.QueryRow("SELECT id FROM dia_chi WHERE so_nha = ? AND ten_duong = ? AND phuong_xa = ?", diaChi.SoNha, diaChi.TenDuong, diaChi.PhuongXa.ID)
	if err := row.Scan(&diaChi.ID); err != nil {
		if err == sql.ErrNoRows {
			return diaChi, err
		}
		return diaChi, err
	}
	return diaChi, nil
}

func ThemDiaChi(diaChi DiaChi) (DiaChi, error) {
	diaChi.ID = TaoIdMoi("dia_chi")
	_, err := db.Exec("INSERT INTO dia_chi VALUES (?, ?, ?, ?)", diaChi.ID, diaChi.SoNha, diaChi.TenDuong, diaChi.PhuongXa.ID)
	return diaChi, err
}

func CapNhatDiaChi(diaChi DiaChi) error {
	_, err := db.Exec("UPDATE dia_chi SET so_nha = ?, ten_duong = ?, phuong_xa = ? WHERE id = ?", diaChi.SoNha, diaChi.TenDuong, diaChi.PhuongXa.ID)
	return err
}

func XoaDiaChi(id int) error {
	_, err := db.Exec("DELETE FROM dia_chi WHERE id = ?", id)
	return err
}
