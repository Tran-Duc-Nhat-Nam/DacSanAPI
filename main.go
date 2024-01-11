package main

import (
	"database/sql"
	// "encoding/json"
	"strconv"

	"fmt"
	"log"
	"net/http"

	// "io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type TinhThanh struct {
	ID  int    `json:"id"`
	Ten string `json:"ten"`
}

type DiaChi struct {
	ID        int       `json:"id"`
	SoNha     string    `json:"so_nha"`
	TenDuong  string    `json:"ten_duong"`
	PhuongXa  string    `json:"phuong_xa"`
	QuanHuyen string    `json:"quan_huyen"`
	TinhThanh TinhThanh `json:"tinh_thanh"`
}

type VungMien struct {
	ID  int    `json:"id"`
	Ten string `json:"ten"`
}

type MuaDacSan struct {
	ID    int    `json:"id"`
	Ten   string `json:"ten"`
	Thang []int  `json:"thang"`
}

type NguyenLieu struct {
	ID  int    `json:"id"`
	Ten string `json:"ten"`
}

type HinhAnh struct {
	ID   int    `json:"id"`
	Ten  string `json:"ten"`
	MoTa string `json:"mo_ta"`
	URL  string `json:"url"`
}

type ThanhPhan struct {
	NguyenLieu NguyenLieu `json:"nguyen_lieu"`
	SoLuong    float64    `json:"so_luong"`
	DonViTinh  string     `json:"don_vi_tinh"`
}

type DacSan struct {
	ID          int         `json:"id"`
	Ten         string      `json:"ten"`
	MoTa        string      `json:"mo_ta"`
	CachCheBien string      `json:"cach_che_bien"`
	ThanhPhan   []ThanhPhan `json:"thanh_phan"`
	VungMien    []VungMien  `json:"vung_mien"`
	MuaDacSan   []MuaDacSan `json:"mua_dac_san"`
	LuotXem     int         `json:"luot_xem"`
	DiemDanhGia float64     `json:"diem_danh_gia"`
	LuotDanhGia int         `json:"luot_danh_gia"`
	HinhDaiDien HinhAnh     `json:"hinh_dai_dien"`
	HinhAnh     []HinhAnh   `json:"hinh_anh"`
}

type NoiBan struct {
	ID          int     `json:"id"`
	Ten         string  `json:"ten"`
	MoTa        string  `json:"mo_ta"`
	DiaChi      DiaChi  `json:"dia_chi"`
	LuotXem     int     `json:"luot_xem"`
	DiemDanhGia float64 `json:"diem_danh_gia"`
	LuotDanhGia int     `json:"luot_danh_gia"`
}

type NguoiDung struct {
	ID          int       `json:"id"`
	Email       string    `json:"email"`
	Ten         string    `json:"ten"`
	IsNam       bool      `json:"is_nam"`
	NgaySinh    time.Time `json:"ngay_sinh"`
	DiaChi      DiaChi    `json:"dia_chi"`
	SoDienThoai string    `json:"so_dien_thoai"`
}

type LuotXemDacSan struct {
	IdNguoiDung int       `json:"id_nguoi_dung"`
	IdDacSan    int       `json:"id_dac_san"`
	ThoiGianXem time.Time `json:"thoi_gian"`
}

type LuotXemNoiBan struct {
	IdNguoiDung int       `json:"id_nguoi_dung"`
	IdNoiBan    int       `json:"id_noi_ban"`
	ThoiGianXem time.Time `json:"thoi_gian"`
}

type LuotDanhGiaDacSan struct {
	IdNguoiDung     int       `json:"id_nguoi_dung"`
	IdDacSan        int       `json:"id_dac_san"`
	ThoiGianDanhGia time.Time `json:"thoi_gian"`
	DiemDanhGia     int       `json:"diem_danh_gia"`
}

type LuotDanhGiaNoiBan struct {
	IdNguoiDung     int       `json:"id_nguoi_dung"`
	IdNoiBan        int       `json:"id_noi_ban"`
	ThoiGianDanhGia time.Time `json:"thoi_gian"`
	DiemDanhGia     int       `json:"diem_danh_gia"`
}

func docTinhThanhCSDL() ([]TinhThanh, error) {
	var dsTinhThanh []TinhThanh

	rows, err := db.Query("SELECT * FROM tinh_thanh")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tinhThanh TinhThanh
		if err := rows.Scan(&tinhThanh.ID, &tinhThanh.Ten); err != nil {
			return nil, err
		}
		dsTinhThanh = append(dsTinhThanh, tinhThanh)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return dsTinhThanh, nil
}

func docTinhThanhTheoIdCSDL(id int) (TinhThanh, error) {
	var tinhThanh TinhThanh

	row := db.QueryRow("SELECT * FROM tinh_thanh WHERE id = ?", strconv.Itoa(id))
	if err := row.Scan(&tinhThanh.ID, &tinhThanh.Ten); err != nil {
		if err == sql.ErrNoRows {
			return tinhThanh, err
		}
		return tinhThanh, err
	}
	return tinhThanh, nil
}

func docDiaChiTheoIdCSDL(id int) (DiaChi, error) {
	var diaChi DiaChi
	var tinhThanhId int

	row := db.QueryRow("SELECT * FROM dia_chi WHERE id = ?", strconv.Itoa(id))
	if err := row.Scan(&diaChi.ID, &diaChi.SoNha, &diaChi.TenDuong, &diaChi.PhuongXa, &diaChi.QuanHuyen, &tinhThanhId); err != nil {
		if err == sql.ErrNoRows {
			return diaChi, err
		}
		return diaChi, err
	}
	tinhThanh, err := docTinhThanhTheoIdCSDL(tinhThanhId)
	if err == nil {
		diaChi.TinhThanh = tinhThanh
	}
	return diaChi, nil
}

func themDiaChiCSDL(diaChi DiaChi) error {
	var count int
	db.QueryRow("SELECT MAX(id) FROM dia_chi").Scan(&count)
	_, err := db.Exec("INSERT INTO dia_chi VALUES (?, ?, ?, ?, ?, ?)", count, diaChi.SoNha, diaChi.TenDuong, diaChi.PhuongXa, diaChi.QuanHuyen, diaChi.TinhThanh.ID)
	return err
}

func capNhatDiaChiCSDL(diaChi DiaChi) error {
	_, err := db.Exec("UPDATE dia_chi SET sp_nha = ?, ten_duong = ?, phuong_xa = ?, quan_huyen = ?, tinh_thanh = ? WHERE id = ?", diaChi.SoNha, diaChi.TenDuong, diaChi.PhuongXa, diaChi.QuanHuyen, diaChi.TinhThanh.ID, diaChi.ID)
	return err
}

func xoaDiaChiCSDL(id int) error {
	_, err := db.Exec("DELETE FROM dia_chi WHERE id = ?", id)
	return err
}

func docVungMienCSDL() ([]VungMien, error) {
	var dsVungMien []VungMien

	rows, err := db.Query("SELECT * FROM vung_mien")
	if err != nil {
		return dsVungMien, err
	}
	defer rows.Close()

	for rows.Next() {
		var vungMien VungMien
		if err := rows.Scan(&vungMien.ID, &vungMien.Ten); err != nil {
			return dsVungMien, err
		}
		dsVungMien = append(dsVungMien, vungMien)
	}

	if err := rows.Err(); err != nil {
		return dsVungMien, err
	}

	return dsVungMien, nil
}

func docVungMienTheoIdCSDL(id int) (VungMien, error) {
	var vungMien VungMien

	rows, err := db.Query("SELECT * FROM vung_mien WHERE id = " + strconv.Itoa(id))
	if err != nil {
		return vungMien, err
	}
	defer rows.Close()

	if err := rows.Scan(&vungMien.ID, &vungMien.Ten); err != nil {
		return vungMien, err
	}

	return vungMien, nil
}

func docVungMienDacSanCSDL(id int) ([]VungMien, error) {
	var dsVungMien []VungMien

	rows, err := db.Query("SELECT * FROM dac_san_thuoc_vung WHERE id_dac_san = " + strconv.Itoa(id))
	if err != nil {
		return dsVungMien, err
	}
	defer rows.Close()

	for rows.Next() {
		var idVungMien int
		if err := rows.Scan(nil, &idVungMien); err != nil {
			return nil, err
		}
		vungMien, err := docVungMienTheoIdCSDL(idVungMien)
		if err != nil {
			return dsVungMien, err
		}
		dsVungMien = append(dsVungMien, vungMien)
	}

	if err := rows.Err(); err != nil {
		return dsVungMien, err
	}

	return dsVungMien, nil
}

func themVungMienCSDL(vungMien VungMien) error {
	var count int
	db.QueryRow("SELECT MAX(id) FROM vung_mien").Scan(&count)
	_, err := db.Exec("INSERT INTO vung_mien VALUES (?, ?)", count, vungMien.Ten)
	return err
}

func capNhatVungMienCSDL(vungMien VungMien) error {
	_, err := db.Exec("UPDATE vung_mien SET ten = ? WHERE id = ?", vungMien.Ten, vungMien.ID)
	return err
}

func xoaVungMieniCSDL(id int) error {
	_, err := db.Exec("DELETE FROM vung_mien WHERE id = ?", id)
	return err
}

func docMuaCSDL() ([]MuaDacSan, error) {
	var dsMuaDacSan []MuaDacSan

	rows, err := db.Query("SELECT * FROM vung_mien")
	if err != nil {
		return dsMuaDacSan, err
	}
	defer rows.Close()

	for rows.Next() {
		var muaDacSan MuaDacSan
		if err := rows.Scan(&muaDacSan.ID, &muaDacSan.Ten); err != nil {
			return dsMuaDacSan, err
		}
		dsMuaDacSan = append(dsMuaDacSan, muaDacSan)
	}

	if err := rows.Err(); err != nil {
		return dsMuaDacSan, err
	}

	return dsMuaDacSan, nil
}

func docMuaTheoIdCSDL(id int) (MuaDacSan, error) {
	var mua MuaDacSan

	rows, err := db.Query("SELECT * FROM mua_dac_san WHERE id = " + strconv.Itoa(id))
	if err != nil {
		return mua, err
	}
	defer rows.Close()

	if err := rows.Scan(&mua.ID, &mua.Ten); err != nil {
		return mua, err
	}

	return mua, nil
}

func docMuaDacSanCSDL(id int) ([]MuaDacSan, error) {
	var dsMuaDacSan []MuaDacSan

	rows, err := db.Query("SELECT * FROM dac_san_theo_mua WHERE id_dac_san = " + strconv.Itoa(id))
	if err != nil {
		return dsMuaDacSan, err
	}
	defer rows.Close()

	for rows.Next() {
		var idMuaDacSan int
		if err := rows.Scan(nil, &idMuaDacSan); err != nil {
			return nil, err
		}
		muaDacSan, err := docMuaTheoIdCSDL(idMuaDacSan)
		if err != nil {
			return dsMuaDacSan, err
		}
		dsMuaDacSan = append(dsMuaDacSan, muaDacSan)
	}

	if err := rows.Err(); err != nil {
		return dsMuaDacSan, err
	}

	return dsMuaDacSan, nil
}

func themMuaCSDL(muaDacSan MuaDacSan) error {
	var count int
	db.QueryRow("SELECT MAX(id) FROM vung_mien").Scan(&count)
	_, err := db.Exec("INSERT INTO mua_dac_san VALUES (?, ?)", count, muaDacSan.Ten)
	return err
}

func capNhatMuaCSDL(muaDacSan MuaDacSan) error {
	_, err := db.Exec("UPDATE mua_dac_san SET ten = ? WHERE id = ?", muaDacSan.Ten, muaDacSan.ID)
	return err
}

func xoaMuaCSDL(id int) error {
	_, err := db.Exec("DELETE FROM mua_dac_san WHERE id = ?", id)
	return err
}

func docNguyenLieuCSDL() ([]NguyenLieu, error) {
	var dsNguyenLieu []NguyenLieu

	rows, err := db.Query("SELECT * FROM nguyen_lieu")
	if err != nil {
		return dsNguyenLieu, err
	}
	defer rows.Close()

	for rows.Next() {
		var nguyenLieu NguyenLieu
		if err := rows.Scan(&nguyenLieu.ID, &nguyenLieu.Ten); err != nil {
			return dsNguyenLieu, err
		}
		dsNguyenLieu = append(dsNguyenLieu, nguyenLieu)
	}

	if err := rows.Err(); err != nil {
		return dsNguyenLieu, err
	}

	return dsNguyenLieu, nil
}

func docNguyenLieuTheoIdCSDL(id int) (NguyenLieu, error) {
	var nguyenLieu NguyenLieu

	rows := db.QueryRow("SELECT * FROM nguyen_lieu WHERE id = " + strconv.Itoa(id))

	if err := rows.Scan(&nguyenLieu.ID, &nguyenLieu.Ten); err != nil {
		return nguyenLieu, err
	}

	return nguyenLieu, nil
}

func docNguyenLieuThanhPhanCSDL(id int) ([]NguyenLieu, error) {
	var dsNguyenLieu []NguyenLieu

	rows, err := db.Query("SELECT * FROM thanh_phan WHERE id_dac_san = " + strconv.Itoa(id))
	if err != nil {
		return dsNguyenLieu, err
	}
	defer rows.Close()

	for rows.Next() {
		var idNguyenLieu int
		if err := rows.Scan(nil, &idNguyenLieu); err != nil {
			return nil, err
		}
		nguyenLieu, err := docNguyenLieuTheoIdCSDL(idNguyenLieu)
		if err != nil {
			return dsNguyenLieu, err
		}
		dsNguyenLieu = append(dsNguyenLieu, nguyenLieu)
	}

	if err := rows.Err(); err != nil {
		return dsNguyenLieu, err
	}

	return dsNguyenLieu, nil
}

func themNguyenLieuCSDL(nguyenLieu NguyenLieu) error {
	var count int
	db.QueryRow("SELECT MAX(id) FROM vung_mien").Scan(&count)
	_, err := db.Exec("INSERT INTO nguyen_lieu VALUES (?, ?)", count, nguyenLieu.Ten)
	return err
}

func capNhatNguyenLieuCSDL(nguyenLieu NguyenLieu) error {
	_, err := db.Exec("UPDATE nguyen_lieu SET ten = ? WHERE id = ?", nguyenLieu.Ten, nguyenLieu.ID)
	return err
}

func xoaNguyenLieuCSDL(id int) error {
	_, err := db.Exec("DELETE FROM nguyen_lieu WHERE id = ?", id)
	return err
}

func docThanhPhanTheoIdCSDL(id int) ([]ThanhPhan, error) {
	var dsThanhPhan []ThanhPhan

	rows, err := db.Query("SELECT * FROM thanh_phan WHERE id_dac_san = " + strconv.Itoa(id))
	if err != nil {
		return dsThanhPhan, err
	}
	defer rows.Close()

	for rows.Next() {
		var thanhPhan ThanhPhan
		var temp int
		var idNguyenLieu int
		if err := rows.Scan(&temp, &idNguyenLieu, &thanhPhan.SoLuong, &thanhPhan.DonViTinh); err != nil {
			return nil, err
		}
		nguyenLieu, err := docNguyenLieuTheoIdCSDL(idNguyenLieu)
		if err == nil {
			thanhPhan.NguyenLieu = nguyenLieu
		}
		dsThanhPhan = append(dsThanhPhan, thanhPhan)
	}

	if err := rows.Err(); err != nil {
		return dsThanhPhan, err
	}

	return dsThanhPhan, nil
}

func themThanhPhanCSDL(idDacSan int, thanhPhan ThanhPhan) error {
	var count int
	db.QueryRow("SELECT MAX(id) FROM thanh_phan").Scan(&count)
	_, err := db.Exec("INSERT INTO thanh_phan VALUES (?, ?, ?, ?)", idDacSan, thanhPhan.NguyenLieu.ID, thanhPhan.SoLuong, thanhPhan.DonViTinh)
	return err
}

func capNhatThanhPhanCSDL(idDacSan int, thanhPhan ThanhPhan) error {
	_, err := db.Exec("UPDATE thanh_phan SET so_luong = ?, don_vi_tinh = ? WHERE id_dac_san = ? AND id_nguyen_lieu = ?", thanhPhan.SoLuong, thanhPhan.DonViTinh, idDacSan, thanhPhan.NguyenLieu.ID)
	return err
}

func xoaThanhPhanCSDL(idDacSan int, idNguyenLieu int) error {
	_, err := db.Exec("DELETE FROM nguyen_lieu WHERE id_dac_san = ? AND id_nguyen_lieu = ?", idDacSan, idNguyenLieu)
	return err
}

func docHinhAnhTheoIdCSDL(id int) (HinhAnh, error) {
	var hinhAhh HinhAnh

	row := db.QueryRow("SELECT * FROM hinh_anh WHERE id = " + strconv.Itoa(id))

	if err := row.Scan(&hinhAhh.ID, &hinhAhh.Ten, &hinhAhh.MoTa, &hinhAhh.URL); err != nil {
		return hinhAhh, err
	}

	return hinhAhh, nil
}

func docHinhAnhDacSanCSDL(id int) ([]HinhAnh, error) {
	var dsHinhAnh []HinhAnh

	rows, err := db.Query("SELECT * FROM hinh_anh_dac_san WHERE id_dac_san = " + strconv.Itoa(id))
	if err != nil {
		return dsHinhAnh, err
	}
	defer rows.Close()

	for rows.Next() {
		var idHinhAnh int
		if err := rows.Scan(nil, &idHinhAnh); err != nil {
			return nil, err
		}
		hinhAnh, err := docHinhAnhTheoIdCSDL(idHinhAnh)
		if err != nil {
			return dsHinhAnh, err
		}
		dsHinhAnh = append(dsHinhAnh, hinhAnh)
	}

	if err := rows.Err(); err != nil {
		return dsHinhAnh, err
	}

	return dsHinhAnh, nil
}

func themHinhAnhCSDL(hinhAnh HinhAnh) error {
	var count int
	db.QueryRow("SELECT MAX(id) FROM hinh_anh").Scan(&count)
	_, err := db.Exec("INSERT INTO hinh_anh VALUES (?, ?, ?, ?)", count, hinhAnh.Ten, hinhAnh.MoTa, hinhAnh.URL)
	return err
}

func capNhatHinhAnhCSDL(hinhAnh HinhAnh) error {
	_, err := db.Exec("UPDATE hinh_anh SET ten = ?, mota = ?, url = ? WHERE id = ?", hinhAnh.Ten, hinhAnh.MoTa, hinhAnh.URL, hinhAnh.ID)
	return err
}

func xoaHinhAnhCSDL(id int) error {
	_, err := db.Exec("DELETE FROM hinh_anh WHERE id = ?", id)
	return err
}

func docDacSanCSDL() ([]DacSan, error) {
	var dsDacSan []DacSan

	rows, err := db.Query("SELECT * FROM dac_san")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var dacSan DacSan
		var idHinhDaiDien int
		if err := rows.Scan(&dacSan.ID, &dacSan.Ten, &dacSan.MoTa, &dacSan.CachCheBien, &dacSan.LuotXem, &dacSan.DiemDanhGia, &dacSan.LuotDanhGia, &idHinhDaiDien); err != nil {
			return nil, err
		}
		thanhPhan, err := docThanhPhanTheoIdCSDL(dacSan.ID)
		if err == nil {
			dacSan.ThanhPhan = thanhPhan
		}
		vungMien, err := docVungMienDacSanCSDL(dacSan.ID)
		if err == nil {
			dacSan.VungMien = vungMien
		}
		muaDacSan, err := docMuaDacSanCSDL(dacSan.ID)
		if err == nil {
			dacSan.MuaDacSan = muaDacSan
		}
		hinhAnh, err := docHinhAnhDacSanCSDL(dacSan.ID)
		if err == nil {
			dacSan.HinhAnh = hinhAnh
		}
		hinhDaiDien, err := docHinhAnhTheoIdCSDL(idHinhDaiDien)
		if err == nil {
			dacSan.HinhDaiDien = hinhDaiDien
		}
		dsDacSan = append(dsDacSan, dacSan)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return dsDacSan, nil
}

func docDacSanTheoIdCSDL(id int) (DacSan, error) {
	var dacSan DacSan

	row := db.QueryRow("SELECT * FROM dac_san WHERE id = " + strconv.Itoa(id))

	var idHinhDaiDien int
	if err := row.Scan(&dacSan.ID, &dacSan.Ten, &dacSan.MoTa, &dacSan.CachCheBien, &dacSan.LuotXem, &dacSan.DiemDanhGia, &dacSan.LuotDanhGia, &idHinhDaiDien); err != nil {
		return dacSan, err
	}
	thanhPhan, err := docThanhPhanTheoIdCSDL(dacSan.ID)
	if err == nil {
		dacSan.ThanhPhan = thanhPhan
	}
	vungMien, err := docVungMienDacSanCSDL(dacSan.ID)
	if err == nil {
		dacSan.VungMien = vungMien
	}
	muaDacSan, err := docMuaDacSanCSDL(dacSan.ID)
	if err == nil {
		dacSan.MuaDacSan = muaDacSan
	}
	hinhAnh, err := docHinhAnhDacSanCSDL(dacSan.ID)
	if err == nil {
		dacSan.HinhAnh = hinhAnh
	}
	hinhDaiDien, err := docHinhAnhTheoIdCSDL(idHinhDaiDien)
	if err == nil {
		dacSan.HinhDaiDien = hinhDaiDien
	}

	return dacSan, nil
}

func themDacSanCSDL(dacSan DacSan) error {
	var count int
	db.QueryRow("SELECT MAX(id) FROM dac_san").Scan(&count)
	_, err := db.Exec("INSERT INTO noi_ban VALUES (?, ?, ?, ?, ?, ?)", count, dacSan.Ten, dacSan.MoTa, dacSan.CachCheBien, dacSan.LuotXem, dacSan.DiemDanhGia, dacSan.LuotDanhGia, dacSan.HinhDaiDien.ID)
	if err != nil {
		return err
	}
	for _, vungMien := range dacSan.VungMien {
		_, err = db.Exec("INSERT INTO dac_san_thuoc_vung VALUES (?, ?)", count, vungMien.ID)
		if err != nil {
			return err
		}
	}
	for _, muaDacSan := range dacSan.MuaDacSan {
		_, err = db.Exec("INSERT INTO dac_san_theo_mua VALUES (?, ?)", count, muaDacSan.ID)
		if err != nil {
			return err
		}
	}
	for _, hinhAnh := range dacSan.HinhAnh {
		_, err = db.Exec("INSERT INTO hinh_anh_dac_san VALUES (?, ?)", hinhAnh.ID, count)
		if err != nil {
			return err
		}
	}
	for _, thanhPhan := range dacSan.ThanhPhan {
		_, err = db.Exec("INSERT INTO thanh_phan VALUES (?, ?, ?, ?)", count, thanhPhan.NguyenLieu.ID, thanhPhan.SoLuong, thanhPhan.DonViTinh)
		if err != nil {
			return err
		}
	}
	return nil
}

func capNhatDacSanCSDL(dacSan DacSan) error {
	_, err := db.Exec("UPDATE dac_san SET ten = ?, mota = ?, dia_chi = ?, luot_xem = ?, diem_danh_gia = ?, luot_danh_gia = ? WHERE id = ?", dacSan.Ten, dacSan.MoTa, dacSan.CachCheBien, dacSan.LuotXem, dacSan.DiemDanhGia, dacSan.LuotDanhGia, dacSan.ID)
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE FROM dac_san_thuoc_vung WHERE id_dac_san = ?)", dacSan.ID)
	for _, vungMien := range dacSan.VungMien {
		_, err = db.Exec("INSERT INTO dac_san_thuoc_vung VALUES (?, ?)", dacSan.ID, vungMien.ID)
		if err != nil {
			return err
		}
	}
	_, err = db.Exec("DELETE FROM dac_san_theo_mua WHERE id_dac_san = ?)", dacSan.ID)
	for _, muaDacSan := range dacSan.MuaDacSan {
		_, err = db.Exec("INSERT INTO dac_san_theo_mua VALUES (?, ?)", dacSan.ID, muaDacSan.ID)
		if err != nil {
			return err
		}
	}
	_, err = db.Exec("DELETE FROM hinh_anh_dac_san WHERE id_dac_san = ?)", dacSan.ID)
	for _, hinhAnh := range dacSan.HinhAnh {
		_, err = db.Exec("INSERT INTO hinh_anh_dac_san VALUES (?, ?)", hinhAnh.ID, dacSan.ID)
		if err != nil {
			return err
		}
	}
	_, err = db.Exec("DELETE FROM thanh_phan WHERE id_dac_san = ?)", dacSan.ID)
	for _, thanhPhan := range dacSan.ThanhPhan {
		_, err = db.Exec("INSERT INTO thanh_phan VALUES (?, ?, ?, ?)", dacSan.ID, thanhPhan.NguyenLieu.ID, thanhPhan.SoLuong, thanhPhan.DonViTinh)
		if err != nil {
			return err
		}
	}
	return nil
}

func xoaDacSanCSDL(id int) error {
	_, err := db.Exec("DELETE FROM dac_san WHERE id = ?", id)
	return err
}

func docNoiBanCSDL() ([]NoiBan, error) {
	var dsNoiBan []NoiBan
	var diaChiId int

	rows, err := db.Query("SELECT * FROM noi_ban")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var noiBan NoiBan
		if err := rows.Scan(&noiBan.ID, &noiBan.Ten, &noiBan.MoTa, &diaChiId, &noiBan.LuotXem, &noiBan.DiemDanhGia, &noiBan.LuotDanhGia); err != nil {
			return nil, err
		}
		diaChi, err := docDiaChiTheoIdCSDL(diaChiId)
		if err == nil {
			noiBan.DiaChi = diaChi
		}
		dsNoiBan = append(dsNoiBan, noiBan)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return dsNoiBan, nil
}

func themNoiBanCSDL(noiBan NoiBan) error {
	var count int
	db.QueryRow("SELECT MAX(id) FROM noi_ban").Scan(&count)
	_, err := db.Exec("INSERT INTO noi_ban VALUES (?, ?, ?, ?, ?, ?)", count, noiBan.Ten, noiBan.MoTa, noiBan.DiaChi.ID, noiBan.LuotXem, noiBan.DiemDanhGia, noiBan.LuotDanhGia)
	return err
}

func capNhatNoiBanCSDL(noiBan NoiBan) error {
	_, err := db.Exec("UPDATE noi_ban SET ten = ?, mota = ?, dia_chi = ?, luot_xem = ?, diem_danh_gia = ?, luot_danh_gia = ? WHERE id = ?", noiBan.Ten, noiBan.MoTa, noiBan.DiaChi.ID, noiBan.LuotXem, noiBan.DiemDanhGia, noiBan.LuotDanhGia, noiBan.ID)
	return err
}

func xoaNoiBanCSDL(id int) error {
	_, err := db.Exec("DELETE FROM noi_ban WHERE id = ?", id)
	return err
}

func docNguoiDungCSDL(id int) (NguoiDung, error) {
	var nguoiDung NguoiDung
	var diaChiId int

	row := db.QueryRow("SELECT * FROM nguoi_dung WHERE id = ?", strconv.Itoa(id))
	if err := row.Scan(&nguoiDung.ID, &nguoiDung.Email, &nguoiDung.Ten, &nguoiDung.IsNam, &nguoiDung.NgaySinh, &diaChiId, &nguoiDung.SoDienThoai); err != nil {
		if err == sql.ErrNoRows {
			return nguoiDung, err
		}
		return nguoiDung, err
	}
	diaChi, err := docDiaChiTheoIdCSDL(diaChiId)
	if err == nil {
		nguoiDung.DiaChi = diaChi
	}
	return nguoiDung, nil
}

func docLichSuXemDacSanCSDL(idNguoiDung int) ([]LuotXemDacSan, error) {
	var lichSuXem []LuotXemDacSan

	rows, err := db.Query("SELECT * FROM lich_su_xem_dac_san WHERE id_nguoi_dung = " + strconv.Itoa(idNguoiDung))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var luotXem LuotXemDacSan
		if err := rows.Scan(&luotXem.IdNguoiDung, &luotXem.IdDacSan, &luotXem.ThoiGianXem); err != nil {
			return nil, err
		}
		lichSuXem = append(lichSuXem, luotXem)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return lichSuXem, nil
}

func themLuotXemDacSanCSDL(luotXem LuotXemDacSan) error {
	_, err := db.Exec("INSERT INTO luot_xem_dac_san VALUES (?, ?, ?, ?)", luotXem.IdNguoiDung, luotXem.IdDacSan, time.Now())
	return err
}

func docLichSuXemNoiBanCSDL(idNguoiDung int) ([]LuotXemNoiBan, error) {
	var lichSuXem []LuotXemNoiBan

	rows, err := db.Query("SELECT * FROM lich_su_xem_noi_ban WHERE id_nguoi_dung = " + strconv.Itoa(idNguoiDung))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var luotXem LuotXemNoiBan
		if err := rows.Scan(&luotXem.IdNguoiDung, &luotXem.IdNoiBan, &luotXem.ThoiGianXem); err != nil {
			return nil, err
		}
		lichSuXem = append(lichSuXem, luotXem)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return lichSuXem, nil
}

func themLuotXemNoiBanCSDL(luotXem LuotXemNoiBan) error {
	_, err := db.Exec("INSERT INTO luot_xem_noi_ban VALUES (?, ?, ?, ?)", luotXem.IdNguoiDung, luotXem.IdNoiBan, time.Now())
	return err
}

func docLichSuDanhGiaDacSanCSDL(idNguoiDung int) ([]LuotDanhGiaDacSan, error) {
	var lichSuDanhGia []LuotDanhGiaDacSan

	rows, err := db.Query("SELECT * FROM danh_gia_dac_san WHERE id_nguoi_dung = " + strconv.Itoa(idNguoiDung))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var luotDanhGia LuotDanhGiaDacSan
		if err := rows.Scan(&luotDanhGia.IdNguoiDung, &luotDanhGia.IdDacSan, &luotDanhGia.ThoiGianDanhGia, &luotDanhGia.DiemDanhGia); err != nil {
			return nil, err
		}
		lichSuDanhGia = append(lichSuDanhGia, luotDanhGia)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return lichSuDanhGia, nil
}

func docLichSuDanhGiaNoiBanCSDL(idNguoiDung int) ([]LuotDanhGiaNoiBan, error) {
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

func docDanhGiaDacSanCSDL(idDacSan int) ([]LuotDanhGiaDacSan, error) {
	var lichSuDanhGia []LuotDanhGiaDacSan

	rows, err := db.Query("SELECT * FROM danh_gia_dac_san WHERE id_dac_san = " + strconv.Itoa(idDacSan))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var luotDanhGia LuotDanhGiaDacSan
		if err := rows.Scan(&luotDanhGia.IdNguoiDung, &luotDanhGia.IdDacSan, &luotDanhGia.ThoiGianDanhGia, &luotDanhGia.DiemDanhGia); err != nil {
			return nil, err
		}
		lichSuDanhGia = append(lichSuDanhGia, luotDanhGia)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return lichSuDanhGia, nil
}

func themDanhGiaDacSanCSDL(luotDanhGia LuotDanhGiaDacSan) error {
	_, err := db.Exec("INSERT INTO luot_danh_gia_dac_san VALUES (?, ?, ?, ?)", luotDanhGia.IdNguoiDung, luotDanhGia.IdDacSan, time.Now(), luotDanhGia.DiemDanhGia)
	return err
}

func capNhatDanhGiaDacSanCSDL(luotDanhGia LuotDanhGiaDacSan) error {
	_, err := db.Exec("UPDATE luot_danh_gia_dac_san SET thoi_gian_danh_gia = ?, diem_danh_gia = ? WHERE id_nguoi_dung = ? AND id_dac_san = ?", time.Now(), luotDanhGia.DiemDanhGia, luotDanhGia.IdNguoiDung, luotDanhGia.IdDacSan)
	return err
}

func xoaDanhGiaDacSanCSDL(idNguoiDung int, idDacSan int) error {
	_, err := db.Exec("DELETE FROM luot_danh_gia_dac_san WHERE id_nguoi_dung = ? AND id_dac_san = ?", idNguoiDung, idDacSan)
	return err
}

func docDanhGiaNoiBanCSDL(idNoiBan int) ([]LuotDanhGiaNoiBan, error) {
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

func themDanhGiaNoiBanCSDL(luotDanhGia LuotDanhGiaNoiBan) error {
	_, err := db.Exec("INSERT INTO luot_danh_gia_noi_ban VALUES (?, ?, ?, ?)", luotDanhGia.IdNguoiDung, luotDanhGia.IdNoiBan, time.Now(), luotDanhGia.DiemDanhGia)
	return err
}

func capNhatDanhGiaNoiBanCSDL(luotDanhGia LuotDanhGiaNoiBan) error {
	_, err := db.Exec("UPDATE luot_danh_gia_noi_ban SET thoi_gian_danh_gia = ?, diem_danh_gia = ? WHERE id_nguoi_dung = ? AND id_dac_san = ?", time.Now(), luotDanhGia.DiemDanhGia, luotDanhGia.IdNguoiDung, luotDanhGia.IdNoiBan)
	return err
}

func xoaDanhGiaNoiBanCSDL(idNguoiDung int, idNoiBan int) error {
	_, err := db.Exec("DELETE FROM luot_danh_gia_noi_ban WHERE id_nguoi_dung = ? AND id_noi_ban = ?", idNguoiDung, idNoiBan)
	return err
}

func docTinhThanhJson(c *gin.Context) {
	tinhThanh, err := docTinhThanhCSDL()
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, tinhThanh)
}

func docDiaChiTheoIdJson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err.Error())
	}
	diaChi, err := docDiaChiTheoIdCSDL(id)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, diaChi)
}

func docVungMienJson(c *gin.Context) {
	dsVungMien, _ := docVungMienCSDL()
	c.IndentedJSON(http.StatusOK, dsVungMien)
}

func docMuaDacSanJson(c *gin.Context) {
	dsMuaDacSan, _ := docMuaCSDL()
	c.IndentedJSON(http.StatusOK, dsMuaDacSan)
}

func docNguyenLieuJson(c *gin.Context) {
	dsNguyenLieu, _ := docNguyenLieuCSDL()
	c.IndentedJSON(http.StatusOK, dsNguyenLieu)
}

func docHinhAnhTheoIdJson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err.Error())
	}
	hinhAnh, err := docHinhAnhTheoIdCSDL(id)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, hinhAnh)
}

func docDacSanJson(c *gin.Context) {
	dsDacSan, _ := docDacSanCSDL()
	c.IndentedJSON(http.StatusOK, dsDacSan)
}

func docDacSanTheoIdJson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err.Error())
	}
	nguoiDung, err := docDacSanTheoIdCSDL(id)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, nguoiDung)
}

func docNoiBanJson(c *gin.Context) {
	dsNoiBan, _ := docNoiBanCSDL()
	c.IndentedJSON(http.StatusOK, dsNoiBan)
}

func docNguoiDungTheoIdJson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err.Error())
	}
	nguoiDung, err := docNguoiDungCSDL(id)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, nguoiDung)
}

// func themTinhThanhCSDL(code int, name string) {
// 	_, err := db.Exec("INSERT INTO tinh_thanh VALUES (?, ?)", code, name)
// 	if err != nil {
// 		return
// 	}
// 	return
// }

// func themTinhThanhJson() {
// 	res, err := http.Get("https://provinces.open-api.vn/api/?depth=1")
// 	if err != nil {
// 		return
// 	}
// 	body, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		return
// 	}
// 	var items []map[string]interface{}
// 	err = json.Unmarshal(body, &items)
// 	if err != nil {
// 		return
// 	}
// 	for _, value := range items {
// 		fmt.Println(value["code"])
// 		fmt.Println(value["name"])
// 		themTinhThanhCSDL(int(value["code"].(float64)), value["name"].(string))
// 	}
// 	defer res.Body.Close()
// }

func main() {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "nhatnam2002",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "dacsandb",
		ParseTime:            true,
		AllowNativePasswords: true,
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	// themTinhThanhJson()

	router := gin.Default()
	router.GET("/dacsan", docDacSanJson)
	router.GET("/dacsan/:id", docDacSanTheoIdJson)
	router.GET("/noiban", docNoiBanJson)
	router.GET("/tinhthanh", docTinhThanhJson)
	router.GET("/vungmien", docVungMienJson)
	router.GET("/muadacsan", docMuaDacSanJson)
	router.GET("/nguyenlieu", docNguyenLieuJson)
	router.GET("/diachi/:id", docDiaChiTheoIdJson)
	router.GET("/hinhanh/:id", docHinhAnhTheoIdJson)
	router.GET("/nguoidung/:id", docNguoiDungTheoIdJson)
	router.Run("localhost:8080")
}
