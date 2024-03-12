package models

import (
	"database/sql"
	"strconv"
	"time"
)

type LuotDanhGiaDacSan struct {
	IdNguoiDung     string    `json:"id_nguoi_dung"`
	IdDacSan        int       `json:"id_dac_san"`
	ThoiGianDanhGia time.Time `json:"thoi_gian_danh_gia"`
	DiemDanhGia     int       `json:"diem_danh_gia"`
	NoiDung         string    `json:"noi_dung"`
}

func DocLichSuDanhGiaDacSanCSDL(idNguoiDung string) ([]LuotDanhGiaDacSan, error) {
	lichSuDanhGia := []LuotDanhGiaDacSan{}

	rows, err := db.Query("SELECT * FROM danh_gia_dac_san WHERE id_nguoi_dung = " + idNguoiDung)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var luotDanhGia LuotDanhGiaDacSan
		var temp sql.NullString
		if err := rows.Scan(&luotDanhGia.IdNguoiDung, &luotDanhGia.IdDacSan, &luotDanhGia.ThoiGianDanhGia, &luotDanhGia.DiemDanhGia, &temp); err != nil {
			return nil, err
		}
		if temp.Valid {
			luotDanhGia.NoiDung = temp.String
		}
		lichSuDanhGia = append(lichSuDanhGia, luotDanhGia)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return lichSuDanhGia, nil
}

func DocDanhGiaDacSanCSDL(idDacSan int) ([]LuotDanhGiaDacSan, error) {
	lichSuDanhGia := []LuotDanhGiaDacSan{}

	rows, err := db.Query("SELECT * FROM danh_gia_dac_san WHERE id_dac_san = ?", idDacSan)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var luotDanhGia LuotDanhGiaDacSan
		var temp sql.NullString
		if err := rows.Scan(&luotDanhGia.IdNguoiDung, &luotDanhGia.IdDacSan, &luotDanhGia.ThoiGianDanhGia, &luotDanhGia.DiemDanhGia, &temp); err != nil {
			return nil, err
		}
		if temp.Valid {
			luotDanhGia.NoiDung = temp.String
		}
		lichSuDanhGia = append(lichSuDanhGia, luotDanhGia)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return lichSuDanhGia, nil
}

func DocDacSanTheoNguoiDungCSDL(idDacSan int, idNguoiDung string) (LuotDanhGiaDacSan, error) {
	var luotDanhGia LuotDanhGiaDacSan

	rows := db.QueryRow("SELECT * FROM luot_danh_gia_dac_san WHERE id_dac_san = ? AND id_nguoi_dung = ?", idDacSan, idNguoiDung)

	var temp sql.NullString
	if err := rows.Scan(&luotDanhGia.IdNguoiDung, &luotDanhGia.IdDacSan, &luotDanhGia.ThoiGianDanhGia, &luotDanhGia.DiemDanhGia, &temp); err != nil {
		return luotDanhGia, err
	}
	if temp.Valid {
		luotDanhGia.NoiDung = temp.String
	}

	return luotDanhGia, nil
}

func TinhDiemDanhGiaDacSanCSDL(idDacSan int) float64 {
	lichSuDanhGia := []LuotDanhGiaDacSan{}
	tongDiem := 0.0

	rows, err := db.Query("SELECT * FROM danh_gia_dac_san WHERE id_dac_san = " + strconv.Itoa(idDacSan))
	if err != nil {
		return -1
	}
	defer rows.Close()

	for rows.Next() {
		var luotDanhGia LuotDanhGiaDacSan
		var temp sql.NullString
		if err := rows.Scan(&luotDanhGia.IdNguoiDung, &luotDanhGia.IdDacSan, &luotDanhGia.ThoiGianDanhGia, &luotDanhGia.DiemDanhGia, &temp); err != nil {
			return -1
		}
		if temp.Valid {
			luotDanhGia.NoiDung = temp.String
		}
		lichSuDanhGia = append(lichSuDanhGia, luotDanhGia)
		tongDiem += float64(luotDanhGia.DiemDanhGia)
	}

	if err := rows.Err(); err != nil {
		return -1
	}

	return tongDiem / float64(len(lichSuDanhGia))
}

func ThemDanhGiaDacSanCSDL(luotDanhGia LuotDanhGiaDacSan) error {
	_, err := DocDacSanTheoNguoiDungCSDL(luotDanhGia.IdDacSan, luotDanhGia.IdNguoiDung)
	if err != nil {
		_, err = db.Exec("INSERT INTO luot_danh_gia_dac_san VALUES (?, ?, ?, ?, ?)", luotDanhGia.IdNguoiDung, luotDanhGia.IdDacSan, luotDanhGia.ThoiGianDanhGia, luotDanhGia.DiemDanhGia, luotDanhGia.NoiDung)
		return err
	} else {
		return CapNhatDanhGiaDacSanCSDL(luotDanhGia)
	}
}

func CapNhatDanhGiaDacSanCSDL(luotDanhGia LuotDanhGiaDacSan) error {
	_, err := db.Exec("UPDATE luot_danh_gia_dac_san SET thoi_gian_danh_gia = ?, diem_danh_gia = ?, noi_dung = ? WHERE id_nguoi_dung = ? AND id_dac_san = ?", luotDanhGia.ThoiGianDanhGia, luotDanhGia.DiemDanhGia, luotDanhGia.NoiDung, luotDanhGia.IdNguoiDung, luotDanhGia.IdDacSan)
	return err
}

func XoaDanhGiaDacSanCSDL(idNguoiDung int, idDacSan int) error {
	_, err := db.Exec("DELETE FROM luot_danh_gia_dac_san WHERE id_nguoi_dung = ? AND id_dac_san = ?", idNguoiDung, idDacSan)
	return err
}
