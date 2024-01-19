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

func DocNguoiDungCSDL() ([]NguoiDung, error) {
	dsNguoiDung := []NguoiDung{}

	row, err := db.Query("SELECT * FROM nguoi_dung ORDER BY id ASC")
	if err != nil {
		return dsNguoiDung, err
	}
	for row.Next() {
		var nguoiDung NguoiDung
		var idDiaChi int
		if err := row.Scan(&nguoiDung.ID, &nguoiDung.Email, &nguoiDung.Ten, &nguoiDung.IsNam, &nguoiDung.NgaySinh, &idDiaChi, &nguoiDung.SoDienThoai); err != nil {
			return dsNguoiDung, err
		}
		lichSuXemDacSan, err := DocLichSuXemDacSanCSDL(nguoiDung.ID)
		if err == nil {
			nguoiDung.LichSuXemDacSan = lichSuXemDacSan
		}
		lichSuXemNoiBan, err := DocLichSuXemNoiBanCSDL(nguoiDung.ID)
		if err == nil {
			nguoiDung.LichSuXemNoiBan = lichSuXemNoiBan
		}
		lichSuDanhGiaDacSan, err := DocLichSuDanhGiaDacSanCSDL(nguoiDung.ID)
		if err == nil {
			nguoiDung.LichSuDanhGiaDacSan = lichSuDanhGiaDacSan
		}
		lichSuDanhGiaNoiBan, err := DocLichSuDanhGiaNoiBanCSDL(nguoiDung.ID)
		if err == nil {
			nguoiDung.LichSuDanhGiaNoiBan = lichSuDanhGiaNoiBan
		}
	}

	return dsNguoiDung, nil
}

func DocNguoiDungTheoIdCSDL(id int) (NguoiDung, error) {
	var nguoiDung NguoiDung
	var idDiaChi int

	row := db.QueryRow("SELECT * FROM nguoi_dung WHERE id = ?", strconv.Itoa(id))
	if err := row.Scan(&nguoiDung.ID, &nguoiDung.Email, &nguoiDung.Ten, &nguoiDung.IsNam, &nguoiDung.NgaySinh, &idDiaChi, &nguoiDung.SoDienThoai); err != nil {
		if err == sql.ErrNoRows {
			return nguoiDung, err
		}
		return nguoiDung, err
	}
	lichSuXemDacSan, err := DocLichSuXemDacSanCSDL(id)
	if err == nil {
		nguoiDung.LichSuXemDacSan = lichSuXemDacSan
	}
	lichSuXemNoiBan, err := DocLichSuXemNoiBanCSDL(id)
	if err == nil {
		nguoiDung.LichSuXemNoiBan = lichSuXemNoiBan
	}
	lichSuDanhGiaDacSan, err := DocLichSuDanhGiaDacSanCSDL(id)
	if err == nil {
		nguoiDung.LichSuDanhGiaDacSan = lichSuDanhGiaDacSan
	}
	lichSuDanhGiaNoiBan, err := DocLichSuDanhGiaNoiBanCSDL(id)
	if err == nil {
		nguoiDung.LichSuDanhGiaNoiBan = lichSuDanhGiaNoiBan
	}
	diaChi, err := DocDiaChiTheoIdCSDL(idDiaChi)
	if err == nil {
		nguoiDung.DiaChi = diaChi
	}
	return nguoiDung, nil
}

func ThemNguoiDungCSDL(nguoiDung NguoiDung) (NguoiDung, error) {
	id := TaoIdMoi("nguoi_dung")
	nguoiDung.ID = id
	diaChi, err := TimDiaChiCSDL(nguoiDung.DiaChi)
	if err != nil {
		ThemDiaChiCSDL(nguoiDung.DiaChi)
	} else {
		nguoiDung.DiaChi = diaChi
	}
	_, err = db.Exec("INSERT INTO nguoi_dung VALUES (?, ?, ?, ?, ?, ?, ?)", id, nguoiDung.Email, nguoiDung.Ten, nguoiDung.IsNam, nguoiDung.NgaySinh, nguoiDung.DiaChi.ID, nguoiDung.SoDienThoai)
	return nguoiDung, err
}

func CapNhatNguoiDungCSDL(nguoiDung NguoiDung) error {
	diaChi, err := DocDiaChiTheoIdCSDL(nguoiDung.DiaChi.ID)
	if diaChi != nguoiDung.DiaChi {
		diaChiMoi, err := ThemDiaChiCSDL(nguoiDung.DiaChi)
		if err != nil {
			nguoiDung.DiaChi = diaChiMoi
		}
	}
	_, err = db.Exec("UPDATE nguoi_dung SET ten = ?, is_nam = ?, ngay_sinh = ?, dia_chi = ?, so_dien_thoai = ? WHERE id = ?", nguoiDung.Ten, nguoiDung.IsNam, nguoiDung.NgaySinh, nguoiDung.DiaChi.ID, nguoiDung.SoDienThoai, nguoiDung.ID)
	return err
}

func XoaNguoiDungCSDL(id int) error {
	_, err := db.Exec("DELETE FROM nguoi_dung WHERE id = ?", id)
	return err
}
