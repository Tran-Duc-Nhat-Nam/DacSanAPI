package models

import (
	"database/sql"
	"fmt"
	"strconv"
)

type DacSan struct {
	ID          int         `json:"id"`
	Ten         string      `json:"ten"`
	MoTa        string      `json:"mo_ta"`
	CachCheBien string      `json:"cach_che_bien"`
	ThanhPhan   []ThanhPhan `json:"thanh_phan"`
	VungMien    []VungMien  `json:"vung_mien"`
	MuaDacSan   []MuaDacSan `json:"mua_dac_san"`
	DsNoiBan    []int       `json:"ds_noi_ban"`
	LuotXem     int         `json:"luot_xem"`
	DiemDanhGia float64     `json:"diem_danh_gia"`
	LuotDanhGia int         `json:"luot_danh_gia"`
	HinhDaiDien HinhAnh     `json:"hinh_dai_dien"`
	HinhAnh     []HinhAnh   `json:"hinh_anh"`
}

func DocDacSan(rows *sql.Rows, err error) ([]DacSan, error) {
	dsDacSan := []DacSan{}

	if err != nil {
		return dsDacSan, err
	}

	defer rows.Close()

	for rows.Next() {
		var dacSan DacSan
		var idHinhDaiDien int
		if err := rows.Scan(&dacSan.ID, &dacSan.Ten, &dacSan.MoTa, &dacSan.CachCheBien, &dacSan.LuotXem, &dacSan.DiemDanhGia, &dacSan.LuotDanhGia, &idHinhDaiDien); err != nil {
			return dsDacSan, err
		}
		thanhPhan, err := DocThanhPhanTheoId(dacSan.ID)
		if err == nil {
			dacSan.ThanhPhan = thanhPhan
		} else {
			fmt.Println(err)
		}
		vungMien, err := DocVungMienDacSan(dacSan.ID)
		if err == nil {
			dacSan.VungMien = vungMien
		} else {
			fmt.Println(err)
		}
		muaDacSan, err := DocMuaDacSan(dacSan.ID)
		if err == nil {
			dacSan.MuaDacSan = muaDacSan
		} else {
			fmt.Println(err)
		}
		hinhAnh, err := DocHinhAnhDacSan(dacSan.ID)
		if err == nil {
			dacSan.HinhAnh = hinhAnh
		} else {
			fmt.Println(err)
		}
		hinhDaiDien, err := DocHinhAnhTheoId(idHinhDaiDien)
		if err == nil {
			dacSan.HinhDaiDien = hinhDaiDien
		} else {
			fmt.Println(err)
		}
		dsDacSan = append(dsDacSan, dacSan)
	}

	if err := rows.Err(); err != nil {
		return dsDacSan, err
	}

	return dsDacSan, nil
}

func DocDanhSachDacSan() ([]DacSan, error) {
	return DocDacSan(db.Query("SELECT * FROM dac_san ORDER BY id ASC"))
}

func DocDanhSachDacSanTheoTrang(soTrang int, doDaiTrang int) ([]DacSan, error) {
	return DocDacSan(db.Query("SELECT * FROM dac_san ORDER BY id ASC LIMIT ?, ?", soTrang*doDaiTrang, doDaiTrang))
}

func TimKiemDacSan(soTrang int, doDaiTrang int, ten string) ([]DacSan, error) {
	return DocDacSan(db.Query("SELECT * FROM dac_san WHERE ten LIKE ? ORDER BY id ASC LIMIT ?, ?", "%"+ten+"%", soTrang*doDaiTrang, doDaiTrang))
}

func TimKiemDacSanTheoVungMien(soTrang int, doDaiTrang int, id int) ([]DacSan, error) {
	return DocDacSan(db.Query("SELECT id, ten, mo_ta, cach_che_bien, luot_xem, diem_danh_gia, luot_danh_gia, hinh_dai_dien FROM dac_san, dac_san_thuoc_vung WHERE id = id_dac_san AND id_vung_mien = ? ORDER BY id ASC LIMIT ?, ?", id, soTrang*doDaiTrang, doDaiTrang))
}

func TimKiemDacSanTheoNguyenLieu(soTrang int, doDaiTrang int, id int) ([]DacSan, error) {
	return DocDacSan(db.Query("SELECT id, ten, mo_ta, cach_che_bien, luot_xem, diem_danh_gia, luot_danh_gia, hinh_dai_dien FROM dac_san, thanh_phan WHERE id = id_dac_san AND id_nguyen_lieu = ? ORDER BY id ASC LIMIT ?, ?", id, soTrang*doDaiTrang, doDaiTrang))
}

func TimKiemDacSanTheoMua(soTrang int, doDaiTrang int, id int) ([]DacSan, error) {
	return DocDacSan(db.Query("SELECT id, ten, mo_ta, cach_che_bien, luot_xem, diem_danh_gia, luot_danh_gia, hinh_dai_dien FROM dac_san, dac_san_theo_mua WHERE id = id_dac_san AND id_mua_dac_san = ? ORDER BY id ASC LIMIT ?, ?", id, soTrang*doDaiTrang, doDaiTrang))
}

func TimKiemDacSanTheoDanhSachVungMien(soTrang int, doDaiTrang int, ten string, dsID []int) ([]DacSan, error) {
	var argString string
	for i, id := range dsID {
		if i == 0 {
			argString += strconv.Itoa(id)
		} else {
			argString += ", " + strconv.Itoa(id)
		}
	}
	return DocDacSan(db.Query("SELECT id, ten, mo_ta, cach_che_bien, luot_xem, diem_danh_gia, luot_danh_gia, hinh_dai_dien FROM dac_san, dac_san_thuoc_vung WHERE id = id_dac_san AND id_vung_mien in ("+argString+") AND ten LIKE ? ORDER BY id ASC LIMIT ?, ?", "%"+ten+"%", soTrang*doDaiTrang, doDaiTrang))
}

func TimKiemDacSanTheoDanhSachMua(soTrang int, doDaiTrang int, ten string, dsID []int) ([]DacSan, error) {
	var argString string
	for i, id := range dsID {
		if i == 0 {
			argString += strconv.Itoa(id)
		} else {
			argString += ", " + strconv.Itoa(id)
		}
	}
	return DocDacSan(db.Query("SELECT id, ten, mo_ta, cach_che_bien, luot_xem, diem_danh_gia, luot_danh_gia, hinh_dai_dien FROM dac_san, dac_san_theo_mua WHERE id = id_dac_san AND id_mua_dac_san in ("+argString+") AND ten LIKE ? ORDER BY id ASC LIMIT ?, ?", "%"+ten+"%", soTrang*doDaiTrang, doDaiTrang))
}

func TimKiemDacSanTheoMuaVungMien(soTrang int, doDaiTrang int, dsVungMien []int, dsMuaDacSan []int, ten string) ([]DacSan, error) {
	var argStringVM string
	for i, id := range dsVungMien {
		if i == 0 {
			argStringVM += strconv.Itoa(id)
		} else {
			argStringVM += ", " + strconv.Itoa(id)
		}
	}
	var argStringMDS string
	for i, id := range dsMuaDacSan {
		if i == 0 {
			argStringMDS += strconv.Itoa(id)
		} else {
			argStringMDS += ", " + strconv.Itoa(id)
		}
	}
	return DocDacSan(db.Query("SELECT id, ten, mo_ta, cach_che_bien, luot_xem, diem_danh_gia, luot_danh_gia, hinh_dai_dien FROM dac_san, dac_san_thuoc_vung, dac_san_theo_mua WHERE id = id_dac_san AND id_vung_mien in ("+argStringVM+") AND id_mua_dac_san in ("+argStringMDS+") AND ten LIKE ? ORDER BY id ASC LIMIT ?, ?", "%"+ten+"%", soTrang*doDaiTrang, doDaiTrang))
}

func TimKiemDacSanTheoDanhSachNguyenLieu(soTrang int, doDaiTrang int, ten string, dsID []int) ([]DacSan, error) {
	var argString string
	for i, id := range dsID {
		if i == 0 {
			argString += strconv.Itoa(id)
		} else {
			argString += ", " + strconv.Itoa(id)
		}
	}
	return DocDacSan(db.Query("SELECT id, ten, mo_ta, cach_che_bien, luot_xem, diem_danh_gia, luot_danh_gia, hinh_dai_dien FROM dac_san, thanh_phan WHERE id = id_dac_san AND id_nguyen_lieu in ("+argString+") AND ten LIKE ? ORDER BY id ASC LIMIT ?, ?", "%"+ten+"%", soTrang*doDaiTrang, doDaiTrang))
}

func TimKiemDacSanTheoNguyenLieuVungMien(soTrang int, doDaiTrang int, dsNguyenLieu []int, dsMuaDacSan []int, ten string) ([]DacSan, error) {
	var argStringNL string
	for i, id := range dsNguyenLieu {
		if i == 0 {
			argStringNL += strconv.Itoa(id)
		} else {
			argStringNL += ", " + strconv.Itoa(id)
		}
	}
	var argStringVM string
	for i, id := range dsMuaDacSan {
		if i == 0 {
			argStringVM += strconv.Itoa(id)
		} else {
			argStringVM += ", " + strconv.Itoa(id)
		}
	}
	return DocDacSan(db.Query("SELECT id, ten, mo_ta, cach_che_bien, luot_xem, diem_danh_gia, luot_danh_gia, hinh_dai_dien FROM dac_san, thanh_phan, dac_san_thuoc_vung WHERE id = id_dac_san AND id_nguyen_lieu in ("+argStringNL+") AND id_vung_mien in ("+argStringVM+") AND ten LIKE ? ORDER BY id ASC LIMIT ?, ?", "%"+ten+"%", soTrang*doDaiTrang, doDaiTrang))
}

func TimKiemDacSanTheoNguyenLieuMua(soTrang int, doDaiTrang int, dsNguyenLieu []int, dsVungMien []int, ten string) ([]DacSan, error) {
	var argStringNL string
	for i, id := range dsNguyenLieu {
		if i == 0 {
			argStringNL += strconv.Itoa(id)
		} else {
			argStringNL += ", " + strconv.Itoa(id)
		}
	}
	var argStringMDS string
	for i, id := range dsVungMien {
		if i == 0 {
			argStringMDS += strconv.Itoa(id)
		} else {
			argStringMDS += ", " + strconv.Itoa(id)
		}
	}
	return DocDacSan(db.Query("SELECT id, ten, mo_ta, cach_che_bien, luot_xem, diem_danh_gia, luot_danh_gia, hinh_dai_dien FROM dac_san, thanh_phan, dac_san_theo_mua WHERE id = id_dac_san AND id_nguyen_lieu in ("+argStringNL+") AND id_mua_dac_san in ("+argStringMDS+") AND ten LIKE ? ORDER BY id ASC LIMIT ?, ?", "%"+ten+"%", soTrang*doDaiTrang, doDaiTrang))
}

func TimKiemDacSanTheoDieuKien(soTrang int, doDaiTrang int, dsNguyenLieu []int, dsVungMien []int, dsMuaDacSan []int, ten string) ([]DacSan, error) {
	var argStringNL string
	for i, id := range dsNguyenLieu {
		if i == 0 {
			argStringNL += strconv.Itoa(id)
		} else {
			argStringNL += ", " + strconv.Itoa(id)
		}
	}
	var argStringMDS string
	for i, id := range dsVungMien {
		if i == 0 {
			argStringMDS += strconv.Itoa(id)
		} else {
			argStringMDS += ", " + strconv.Itoa(id)
		}
	}
	var argStringVM string
	for i, id := range dsMuaDacSan {
		if i == 0 {
			argStringVM += strconv.Itoa(id)
		} else {
			argStringVM += ", " + strconv.Itoa(id)
		}
	}
	return DocDacSan(db.Query("SELECT id, ten, mo_ta, cach_che_bien, luot_xem, diem_danh_gia, luot_danh_gia, hinh_dai_dien FROM dac_san, thanh_phan, dac_san_thuoc_vung, dac_san_theo_mua WHERE id = id_dac_san AND id_nguyen_lieu in ("+argStringNL+") AND id_vung_mien in ("+argStringVM+") AND id_mua_dac_san in ("+argStringMDS+") AND ten LIKE ? ORDER BY id ASC LIMIT ?, ?", "%"+ten+"%", soTrang*doDaiTrang, doDaiTrang))
}

func DocDacSanTheoId(id int) (DacSan, error) {
	var dacSan DacSan

	row := db.QueryRow("SELECT * FROM dac_san WHERE id = ?", id)

	var idHinhDaiDien int
	if err := row.Scan(&dacSan.ID, &dacSan.Ten, &dacSan.MoTa, &dacSan.CachCheBien, &dacSan.LuotXem, &dacSan.DiemDanhGia, &dacSan.LuotDanhGia, &idHinhDaiDien); err != nil {
		return dacSan, err
	}
	thanhPhan, err := DocThanhPhanTheoId(dacSan.ID)
	if err == nil {
		dacSan.ThanhPhan = thanhPhan
	} else {
		fmt.Println(err)
	}
	vungMien, err := DocVungMienDacSan(dacSan.ID)
	if err == nil {
		dacSan.VungMien = vungMien
	} else {
		fmt.Println(err)
	}
	muaDacSan, err := DocMuaDacSan(dacSan.ID)
	if err == nil {
		dacSan.MuaDacSan = muaDacSan
	} else {
		fmt.Println(err)
	}
	hinhAnh, err := DocHinhAnhDacSan(dacSan.ID)
	if err == nil {
		dacSan.HinhAnh = hinhAnh
	} else {
		fmt.Println(err)
	}
	hinhDaiDien, err := DocHinhAnhTheoId(idHinhDaiDien)
	if err == nil {
		dacSan.HinhDaiDien = hinhDaiDien
	} else {
		fmt.Println(err)
	}

	return dacSan, nil
}

func DocDacSanTheoNoiBan(id int) ([]DacSan, error) {
	dsDacSan := []DacSan{}

	rows, err := db.Query("SELECT * FROM noi_ban_dac_san where id_noi_ban = ? ORDER BY id_dac_san ASC", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var idDacSan int
		if err := rows.Scan(&id, &idDacSan); err != nil {
			return nil, err
		}
		dacSan, err := DocDacSanTheoId(idDacSan)
		if err == nil {
			dsDacSan = append(dsDacSan, dacSan)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return dsDacSan, nil
}

func DocDacSanTheoTen(text string) ([]DacSan, error) {
	return DocLike(text, "ten", "dac_san", DocDacSanTheoId)
}

func DocDacSanTheoMoTa(text string) ([]DacSan, error) {
	return DocLike(text, "mo_ta", "dac_san", DocDacSanTheoId)
}

func DocDacSanTheoCachCheBien(text string) ([]DacSan, error) {
	return DocLike(text, "cach_che_bien", "dac_san", DocDacSanTheoId)
}

func DocLike(text string, cot string, bang string, timTheoID func(int) (DacSan, error)) ([]DacSan, error) {
	return DocDacSan(db.Query("SELECT id FROM "+bang+" WHERE "+cot+" LIKE ? ORDER BY id ASC", "%"+text+"%"))
}

func ThemDacSan(dacSan DacSan) (DacSan, error) {
	for _, hinhAnh := range dacSan.HinhAnh {
		var count int
		_ = db.QueryRow("SELECT COUNT(*) FROM hinh_anh WHERE ten = ?", hinhAnh.Ten).Scan(&count)
		if count == 0 {
			_, err := db.Exec("INSERT INTO hinh_anh VALUES (?, ?, ?, ?)", TaoIdMoi("hinh_anh"), hinhAnh.Ten, hinhAnh.MoTa, hinhAnh.URL)
			if err != nil {
				return dacSan, err
			}
		}
	}
	var count int
	_ = db.QueryRow("SELECT COUNT(*) FROM hinh_anh WHERE ten = ?", dacSan.HinhDaiDien.Ten).Scan(&count)
	if count == 0 {
		dacSan.HinhDaiDien.ID = TaoIdMoi("hinh_anh")
		_, err := db.Exec("INSERT INTO hinh_anh VALUES (?, ?, ?, ?)", dacSan.HinhDaiDien.ID, dacSan.HinhDaiDien.Ten, dacSan.HinhDaiDien.MoTa, dacSan.HinhDaiDien.URL)
		if err != nil {
			return dacSan, err
		}
	} else {
		_ = db.QueryRow("SELECT id FROM hinh_anh WHERE ten = ?", dacSan.HinhDaiDien.Ten).Scan(&dacSan.HinhDaiDien.ID)
	}
	id := TaoIdMoi("dac_san")
	dacSan.ID = id
	_, err := db.Exec("INSERT INTO dac_san VALUES (?, ?, ?, ?, ?, ?, ?, ?)", id, dacSan.Ten, dacSan.MoTa, dacSan.CachCheBien, dacSan.LuotXem, dacSan.DiemDanhGia, dacSan.LuotDanhGia, dacSan.HinhDaiDien.ID)
	if err != nil {
		return dacSan, err
	}
	for _, vungMien := range dacSan.VungMien {
		_, err = db.Exec("INSERT INTO dac_san_thuoc_vung VALUES (?, ?)", id, vungMien.ID)
		if err != nil {
			return dacSan, err
		}
	}
	for _, muaDacSan := range dacSan.MuaDacSan {
		_, err = db.Exec("INSERT INTO dac_san_theo_mua VALUES (?, ?)", id, muaDacSan.ID)
		if err != nil {
			return dacSan, err
		}
	}
	for _, hinhAnh := range dacSan.HinhAnh {
		_, err = db.Exec("INSERT INTO hinh_anh_dac_san VALUES (?, ?)", hinhAnh.ID, id)
		if err != nil {
			return dacSan, err
		}
	}
	for _, thanhPhan := range dacSan.ThanhPhan {
		_, err = db.Exec("INSERT INTO thanh_phan VALUES (?, ?, ?, ?)", id, thanhPhan.NguyenLieu.ID, thanhPhan.SoLuong, thanhPhan.DonViTinh)
		if err != nil {
			return dacSan, err
		}
	}
	for _, idNoiBan := range dacSan.DsNoiBan {
		_, err = db.Exec("INSERT INTO noi_ban_dac_san VALUES (?, ?)", idNoiBan, dacSan.ID)
		if err != nil {
			return dacSan, err
		}
	}
	return dacSan, nil
}

func CapNhatDacSan(dacSan DacSan) error {
	_, err := db.Exec("UPDATE dac_san SET ten = ?, mo_ta = ?, cach_che_bien = ?, luot_xem = ?, diem_danh_gia = ?, luot_danh_gia = ? WHERE id = ?", dacSan.Ten, dacSan.MoTa, dacSan.CachCheBien, dacSan.LuotXem, dacSan.DiemDanhGia, dacSan.LuotDanhGia, dacSan.ID)
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
	_, err = db.Exec("DELETE FROM noi_ban_dac_san WHERE id_dac_san = ?)", dacSan.ID)
	for _, idNoiBan := range dacSan.DsNoiBan {
		_, err = db.Exec("INSERT INTO noi_ban_dac_san VALUES (?, ?)", idNoiBan, dacSan.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

func XoaDacSan(id int) error {
	_, err := db.Exec("DELETE FROM dac_san_thuoc_vung WHERE id_dac_san = ?", id)
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE FROM dac_san_theo_mua WHERE id_dac_san = ?", id)
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE FROM hinh_anh_dac_san WHERE id_dac_san = ?", id)
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE FROM thanh_phan WHERE id_dac_san = ?", id)
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE FROM noi_ban_dac_san WHERE id_dac_san = ?", id)
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE FROM dac_san WHERE id = ?", id)
	return err
}
