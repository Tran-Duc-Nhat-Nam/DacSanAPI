package models

import (
	"database/sql"
	"strconv"
	"time"
)

type NguoiDung struct {
	ID                  int                 `json:"id"`
	Email               string              `json:"email"`
	Ten                 string              `json:"ten"`
	IsNam               bool                `json:"is_nam"`
	NgaySinh            time.Time           `json:"ngay_sinh"`
	DiaChi              DiaChi              `json:"dia_chi"`
	SoDienThoai         string              `json:"so_dien_thoai"`
	LichSuXemDacSan     []LuotXemDacSan     `json:"lich_su_xem_dac_san"`
	LichSuXemNoiBan     []LuotXemNoiBan     `json:"lich_su_xem_noi_ban"`
	LichSuDanhGiaDacSan []LuotDanhGiaDacSan `json:"lich_su_danh_gia_dac_san"`
	LichSuDanhGiaNoiBan []LuotDanhGiaNoiBan `json:"lich_su_danh_gia_san"`
}

func DocNguoiDungCSDL(id int) (NguoiDung, error) {
	var nguoiDung NguoiDung
	var diaChiId int

	row := db.QueryRow("SELECT * FROM nguoi_dung WHERE id = ?", strconv.Itoa(id))
	if err := row.Scan(&nguoiDung.ID, &nguoiDung.Email, &nguoiDung.Ten, &nguoiDung.IsNam, &nguoiDung.NgaySinh, &diaChiId, &nguoiDung.SoDienThoai); err != nil {
		if err == sql.ErrNoRows {
			return nguoiDung, err
		}
		return nguoiDung, err
	}
	lichSuXemDacSan, err := DocLichSuXemDacSanCSDL(diaChiId)
	if err == nil {
		nguoiDung.LichSuXemDacSan = lichSuXemDacSan
	}
	lichSuXemNoiBan, err := DocLichSuXemNoiBanCSDL(diaChiId)
	if err == nil {
		nguoiDung.LichSuXemNoiBan = lichSuXemNoiBan
	}
	lichSuDanhGiaDacSan, err := DocLichSuDanhGiaDacSanCSDL(diaChiId)
	if err == nil {
		nguoiDung.LichSuDanhGiaDacSan = lichSuDanhGiaDacSan
	}
	lichSuDanhGiaNoiBan, err := DocLichSuDanhGiaNoiBanCSDL(diaChiId)
	if err == nil {
		nguoiDung.LichSuDanhGiaNoiBan = lichSuDanhGiaNoiBan
	}
	return nguoiDung, nil
}

func ThemNguoiDungCSDL(nguoiDung NguoiDung) error {
	var count int
	db.QueryRow("SELECT MAX(id) FROM nguoi_dung").Scan(&count)
	_, err := db.Exec("INSERT INTO nguoi_dung VALUES (?, ?, ?, ?, ?, ?, ?)", count, nguoiDung.Email, nguoiDung.Ten, nguoiDung.IsNam, nguoiDung.NgaySinh, nguoiDung.DiaChi.ID, nguoiDung.SoDienThoai)
	return err
}
