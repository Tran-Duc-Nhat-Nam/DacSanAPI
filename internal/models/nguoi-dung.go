package models

import (
	"database/sql"
	"time"
)

type NguoiDung struct {
	ID                  string              `json:"id"`
	Email               string              `json:"email"`
	Ten                 string              `json:"ten"`
	IsNam               bool                `json:"is_nam"`
	NgaySinh            time.Time           `json:"ngay_sinh"`
	DiaChi              DiaChi              `json:"dia_chi"`
	SoDienThoai         string              `json:"so_dien_thoai"`
	LichSuXemDacSan     []LuotXemDacSan     `json:"lich_su_xem_dac_san"`
	LichSuXemNoiBan     []LuotXemNoiBan     `json:"lich_su_xem_noi_ban"`
	LichSuDanhGiaDacSan []LuotDanhGiaDacSan `json:"lich_su_danh_gia_dac_san"`
	LichSuDanhGiaNoiBan []LuotDanhGiaNoiBan `json:"lich_su_danh_gia_noi_ban"`
}

func DocNguoiDung(rows *sql.Rows, err error) ([]NguoiDung, error) {
	dsNguoiDung := []NguoiDung{}

	if err != nil {
		return dsNguoiDung, err
	}

	defer rows.Close()

	for rows.Next() {
		var nguoiDung NguoiDung
		var idDiaChi int
		if err := rows.Scan(&nguoiDung.ID, &nguoiDung.Email, &nguoiDung.Ten, &nguoiDung.IsNam, &nguoiDung.NgaySinh, &idDiaChi, &nguoiDung.SoDienThoai); err != nil {
			return dsNguoiDung, err
		}
		nguoiDung.LichSuXemDacSan, err = DocLichSuXemDacSan(nguoiDung.ID)
		nguoiDung.LichSuXemNoiBan, err = DocLichSuXemNoiBan(nguoiDung.ID)
		nguoiDung.LichSuDanhGiaDacSan, err = DocLichSuDanhGiaDacSan(nguoiDung.ID)
		nguoiDung.LichSuDanhGiaNoiBan, err = DocLichSuDanhGiaNoiBan(nguoiDung.ID)
		diaChi, err := DocDiaChiTheoId(idDiaChi)
		if err == nil {
			nguoiDung.DiaChi = diaChi
		}
		dsNguoiDung = append(dsNguoiDung, nguoiDung)
	}
	return dsNguoiDung, err
}

func DocDanhSachNguoiDung() ([]NguoiDung, error) {
	return DocNguoiDung(db.Query("SELECT * FROM nguoi_dung ORDER BY id ASC"))
}

func TimKiemNguoiDung(soTrang int, doDaiTrang int, ten string) ([]NguoiDung, error) {
	return DocNguoiDung(db.Query("SELECT * FROM nguoi_dung WHERE ten LIKE ? ORDER BY id ASC LIMIT ?, ?", "%"+ten+"%", soTrang*doDaiTrang, doDaiTrang))
}

func DocNguoiDungTheoId(id string) (NguoiDung, error) {
	var nguoiDung NguoiDung
	var idDiaChi int

	row := db.QueryRow("SELECT * FROM nguoi_dung WHERE id = ?", id)
	if err := row.Scan(&nguoiDung.ID, &nguoiDung.Email, &nguoiDung.Ten, &nguoiDung.IsNam, &nguoiDung.NgaySinh, &idDiaChi, &nguoiDung.SoDienThoai); err != nil {
		if err == sql.ErrNoRows {
			return nguoiDung, err
		}
		return nguoiDung, err
	}
	lichSuXemDacSan, err := DocLichSuXemDacSan(id)
	if err == nil {
		nguoiDung.LichSuXemDacSan = lichSuXemDacSan
	}
	lichSuXemNoiBan, err := DocLichSuXemNoiBan(id)
	if err == nil {
		nguoiDung.LichSuXemNoiBan = lichSuXemNoiBan
	}
	lichSuDanhGiaDacSan, err := DocLichSuDanhGiaDacSan(id)
	if err == nil {
		nguoiDung.LichSuDanhGiaDacSan = lichSuDanhGiaDacSan
	}
	lichSuDanhGiaNoiBan, err := DocLichSuDanhGiaNoiBan(id)
	if err == nil {
		nguoiDung.LichSuDanhGiaNoiBan = lichSuDanhGiaNoiBan
	}
	diaChi, err := DocDiaChiTheoId(idDiaChi)
	if err == nil {
		nguoiDung.DiaChi = diaChi
	}
	return nguoiDung, nil
}

func ThemNguoiDung(nguoiDung NguoiDung) (NguoiDung, error) {
	diaChi, err := TimDiaChi(nguoiDung.DiaChi)
	if err != nil {
		nguoiDung.DiaChi, err = ThemDiaChi(nguoiDung.DiaChi)
	} else {
		nguoiDung.DiaChi = diaChi
	}
	_, err = db.Exec("INSERT INTO nguoi_dung VALUES (?, ?, ?, ?, ?, ?, ?)", nguoiDung.ID, nguoiDung.Email, nguoiDung.Ten, nguoiDung.IsNam, nguoiDung.NgaySinh, nguoiDung.DiaChi.ID, nguoiDung.SoDienThoai)
	return nguoiDung, err
}

func CapNhatNguoiDung(nguoiDung NguoiDung) error {
	CapNhatDiaChi(nguoiDung.DiaChi)
	_, err := db.Exec("UPDATE nguoi_dung SET ten = ?, is_nam = ?, ngay_sinh = ?, dia_chi = ?, so_dien_thoai = ? WHERE id = ?", nguoiDung.Ten, nguoiDung.IsNam, nguoiDung.NgaySinh, nguoiDung.DiaChi.ID, nguoiDung.SoDienThoai, nguoiDung.ID)
	return err
}

func XoaNguoiDung(id int) error {
	_, err := db.Exec("DELETE FROM nguoi_dung WHERE id = ?", id)
	return err
}
