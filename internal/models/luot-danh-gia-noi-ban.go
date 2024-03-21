package models

import (
	"database/sql"
	"strconv"
	"time"
)

type LuotDanhGiaNoiBan struct {
	IdNguoiDung     string    `json:"id_nguoi_dung"`
	IdNoiBan        int       `json:"id_noi_ban"`
	ThoiGianDanhGia time.Time `json:"thoi_gian_danh_gia"`
	DiemDanhGia     int       `json:"diem_danh_gia"`
	NoiDung         string    `json:"noi_dung"`
}

func DocLichSuDanhGiaNoiBanCSDL(idNguoiDung string) ([]LuotDanhGiaNoiBan, error) {
	lichSuDanhGia := []LuotDanhGiaNoiBan{}

	rows, err := db.Query("SELECT * FROM luot_danh_gia_noi_ban WHERE id_nguoi_dung = " + idNguoiDung)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var luotDanhGia LuotDanhGiaNoiBan
		var temp sql.NullString
		if err := rows.Scan(&luotDanhGia.IdNguoiDung, &luotDanhGia.IdNoiBan, &luotDanhGia.ThoiGianDanhGia, &luotDanhGia.DiemDanhGia, &temp); err != nil {
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

func DocDanhGiaNoiBanCSDL(idNoiBan int) ([]LuotDanhGiaNoiBan, error) {
	lichSuDanhGia := []LuotDanhGiaNoiBan{}

	rows, err := db.Query("SELECT * FROM luot_danh_gia_noi_ban WHERE id_noi_ban = ?", idNoiBan)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var luotDanhGia LuotDanhGiaNoiBan
		var temp sql.NullString
		if err := rows.Scan(&luotDanhGia.IdNguoiDung, &luotDanhGia.IdNoiBan, &luotDanhGia.ThoiGianDanhGia, &luotDanhGia.DiemDanhGia, &temp); err != nil {
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

func DocNoiBanTheoNguoiDungCSDL(idNoiBan int, idNguoiDung string) (LuotDanhGiaNoiBan, error) {
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

func TinhDiemDanhGiaNoiBanCSDL(idNoiBan int) float64 {
	lichSuDanhGia := []LuotDanhGiaNoiBan{}
	tongDiem := 0.0

	rows, err := db.Query("SELECT * FROM luot_danh_gia_noi_ban WHERE id_noi_ban = " + strconv.Itoa(idNoiBan))
	if err != nil {
		return -1
	}
	defer rows.Close()

	for rows.Next() {
		var luotDanhGia LuotDanhGiaNoiBan
		var temp sql.NullString
		if err := rows.Scan(&luotDanhGia.IdNguoiDung, &luotDanhGia.IdNoiBan, &luotDanhGia.ThoiGianDanhGia, &luotDanhGia.DiemDanhGia, &temp); err != nil {
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

func ThemDanhGiaNoiBanCSDL(luotDanhGia LuotDanhGiaNoiBan) error {
	_, err := DocNoiBanTheoNguoiDungCSDL(luotDanhGia.IdNoiBan, luotDanhGia.IdNguoiDung)
	if err != nil {
		_, err = db.Exec("INSERT INTO luot_danh_gia_noi_ban VALUES (?, ?, ?, ?, ?)", luotDanhGia.IdNguoiDung, luotDanhGia.IdNoiBan, luotDanhGia.ThoiGianDanhGia, luotDanhGia.DiemDanhGia, luotDanhGia.NoiDung)
		return err
	} else {
		return CapNhatDanhGiaNoiBanCSDL(luotDanhGia)
	}
}

func CapNhatDanhGiaNoiBanCSDL(luotDanhGia LuotDanhGiaNoiBan) error {
	_, err := db.Exec("UPDATE luot_danh_gia_noi_ban SET thoi_gian_danh_gia = ?, diem_danh_gia = ?, noi_dung = ? WHERE id_nguoi_dung = ? AND id_noi_ban = ?", luotDanhGia.ThoiGianDanhGia, luotDanhGia.DiemDanhGia, luotDanhGia.NoiDung, luotDanhGia.IdNguoiDung, luotDanhGia.IdNoiBan)
	return err
}

func XoaDanhGiaNoiBanCSDL(idNguoiDung int, idNoiBan int) error {
	_, err := db.Exec("DELETE FROM luot_danh_gia_noi_ban WHERE id_nguoi_dung = ? AND id_noi_ban = ?", idNguoiDung, idNoiBan)
	return err
}
