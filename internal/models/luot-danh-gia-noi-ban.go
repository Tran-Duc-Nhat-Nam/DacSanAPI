package models

import (
	"strconv"
	"time"
)

type LuotDanhGiaNoiBan struct {
	IdNguoiDung     int       `json:"id_nguoi_dung"`
	IdNoiBan        int       `json:"id_noi_ban"`
	ThoiGianDanhGia time.Time `json:"thoi_gian"`
	DiemDanhGia     int       `json:"diem_danh_gia"`
}

func DocLichSuDanhGiaNoiBanCSDL(idNguoiDung int) ([]LuotDanhGiaNoiBan, error) {
	var lichSuDanhGia []LuotDanhGiaNoiBan

	rows, err := db.Query("SELECT * FROM danh_gia_noi_ban WHERE id_nguoi_dung = " + strconv.Itoa(idNguoiDung))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var luotDanhGia LuotDanhGiaNoiBan
		if err := rows.Scan(&luotDanhGia.IdNguoiDung, &luotDanhGia.IdNoiBan, &luotDanhGia.ThoiGianDanhGia, &luotDanhGia.DiemDanhGia); err != nil {
			return nil, err
		}
		lichSuDanhGia = append(lichSuDanhGia, luotDanhGia)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return lichSuDanhGia, nil
}

func DocDanhGiaNoiBanCSDL(idNoiBan int) ([]LuotDanhGiaNoiBan, error) {
	var lichSuDanhGia []LuotDanhGiaNoiBan

	rows, err := db.Query("SELECT * FROM danh_gia_noi_ban WHERE id_noi_ban = " + strconv.Itoa(idNoiBan))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var luotDanhGia LuotDanhGiaNoiBan
		if err := rows.Scan(&luotDanhGia.IdNguoiDung, &luotDanhGia.IdNoiBan, &luotDanhGia.ThoiGianDanhGia, &luotDanhGia.DiemDanhGia); err != nil {
			return nil, err
		}
		lichSuDanhGia = append(lichSuDanhGia, luotDanhGia)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return lichSuDanhGia, nil
}

func ThemDanhGiaNoiBanCSDL(luotDanhGia LuotDanhGiaNoiBan) error {
	_, err := db.Exec("INSERT INTO luot_danh_gia_noi_ban VALUES (?, ?, ?, ?)", luotDanhGia.IdNguoiDung, luotDanhGia.IdNoiBan, time.Now(), luotDanhGia.DiemDanhGia)
	return err
}

func CapNhatDanhGiaNoiBanCSDL(luotDanhGia LuotDanhGiaNoiBan) error {
	_, err := db.Exec("UPDATE luot_danh_gia_noi_ban SET thoi_gian_danh_gia = ?, diem_danh_gia = ? WHERE id_nguoi_dung = ? AND id_dac_san = ?", time.Now(), luotDanhGia.DiemDanhGia, luotDanhGia.IdNguoiDung, luotDanhGia.IdNoiBan)
	return err
}

func XoaDanhGiaNoiBanCSDL(idNguoiDung int, idNoiBan int) error {
	_, err := db.Exec("DELETE FROM luot_danh_gia_noi_ban WHERE id_nguoi_dung = ? AND id_noi_ban = ?", idNguoiDung, idNoiBan)
	return err
}
