package models

import (
	"database/sql"
	"time"
)

type LuotDanhGiaNoiBan struct {
	IdNguoiDung     string    `json:"id_nguoi_dung"`
	IdNoiBan        int       `json:"id_noi_ban"`
	ThoiGianDanhGia time.Time `json:"thoi_gian_danh_gia"`
	DiemDanhGia     int       `json:"diem_danh_gia"`
	NoiDung         string    `json:"noi_dung"`
}

func DocLuotDanhGiaNoiBan(rows *sql.Rows, err error) ([]LuotDanhGiaNoiBan, error) {
	lichSuDanhGia := []LuotDanhGiaNoiBan{}
	if err != nil {
		return lichSuDanhGia, err
	}
	defer rows.Close()

	for rows.Next() {
		var luotDanhGia LuotDanhGiaNoiBan
		var temp sql.NullString
		if err := rows.Scan(&luotDanhGia.IdNguoiDung, &luotDanhGia.IdNoiBan, &luotDanhGia.ThoiGianDanhGia, &luotDanhGia.DiemDanhGia, &temp); err != nil {
			return lichSuDanhGia, err
		}
		if temp.Valid {
			luotDanhGia.NoiDung = temp.String
		}
		lichSuDanhGia = append(lichSuDanhGia, luotDanhGia)
	}

	if err := rows.Err(); err != nil {
		return lichSuDanhGia, err
	}
	return lichSuDanhGia, nil
}

func DocLichSuDanhGiaNoiBan(idNguoiDung string) ([]LuotDanhGiaNoiBan, error) {
	return DocLuotDanhGiaNoiBan(db.Query("SELECT * FROM luot_danh_gia_noi_ban WHERE id_nguoi_dung = ?", idNguoiDung))
}

func DocDanhGiaNoiBan(idNoiBan int) ([]LuotDanhGiaNoiBan, error) {
	return DocLuotDanhGiaNoiBan(db.Query("SELECT * FROM luot_danh_gia_noi_ban WHERE id_noi_ban = ?", idNoiBan))
}

func DocNoiBanTheoNguoiDung(idNoiBan int, idNguoiDung string) (LuotDanhGiaNoiBan, error) {
	var luotDanhGia LuotDanhGiaNoiBan

	rows := db.QueryRow("SELECT * FROM luot_danh_gia_noi_ban WHERE id_noi_ban = ? AND id_nguoi_dung = ?", idNoiBan, idNguoiDung)

	var temp sql.NullString
	if err := rows.Scan(&luotDanhGia.IdNguoiDung, &luotDanhGia.IdNoiBan, &luotDanhGia.ThoiGianDanhGia, &luotDanhGia.DiemDanhGia, &temp); err != nil {
		return luotDanhGia, err
	}
	if temp.Valid {
		luotDanhGia.NoiDung = temp.String
	}

	return luotDanhGia, nil
}

func ThemDanhGiaNoiBan(luotDanhGia LuotDanhGiaNoiBan) error {
	_, err := DocNoiBanTheoNguoiDung(luotDanhGia.IdNoiBan, luotDanhGia.IdNguoiDung)
	if err != nil {
		_, err = db.Exec("INSERT INTO luot_danh_gia_noi_ban VALUES (?, ?, ?, ?, ?)", luotDanhGia.IdNguoiDung, luotDanhGia.IdNoiBan, luotDanhGia.ThoiGianDanhGia, luotDanhGia.DiemDanhGia, luotDanhGia.NoiDung)
		return err
	} else {
		return CapNhatDanhGiaNoiBan(luotDanhGia)
	}
}

func CapNhatDanhGiaNoiBan(luotDanhGia LuotDanhGiaNoiBan) error {
	_, err := db.Exec("UPDATE luot_danh_gia_noi_ban SET thoi_gian_danh_gia = ?, diem_danh_gia = ?, noi_dung = ? WHERE id_nguoi_dung = ? AND id_noi_ban = ?", luotDanhGia.ThoiGianDanhGia, luotDanhGia.DiemDanhGia, luotDanhGia.NoiDung, luotDanhGia.IdNguoiDung, luotDanhGia.IdNoiBan)
	return err
}

func XoaDanhGiaNoiBan(idNguoiDung string, idDacSan int) error {
	_, err := db.Exec("DELETE FROM luot_danh_gia_noi_ban WHERE id_nguoi_dung = ? AND id_noi_ban = ?", idNguoiDung, idDacSan)
	return err
}
