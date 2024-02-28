package models

import (
	"database/sql"
	"fmt"
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

func DocDacSanCSDL(rows *sql.Rows, err error) ([]DacSan, error) {
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
		thanhPhan, err := DocThanhPhanTheoIdCSDL(dacSan.ID)
		if err == nil {
			dacSan.ThanhPhan = thanhPhan
		} else {
			fmt.Println(err)
		}
		vungMien, err := DocVungMienDacSanCSDL(dacSan.ID)
		if err == nil {
			dacSan.VungMien = vungMien
		} else {
			fmt.Println(err)
		}
		muaDacSan, err := DocMuaDacSanCSDL(dacSan.ID)
		if err == nil {
			dacSan.MuaDacSan = muaDacSan
		} else {
			fmt.Println(err)
		}
		hinhAnh, err := DocHinhAnhDacSanCSDL(dacSan.ID)
		if err == nil {
			dacSan.HinhAnh = hinhAnh
		} else {
			fmt.Println(err)
		}
		hinhDaiDien, err := DocHinhAnhTheoIdCSDL(idHinhDaiDien)
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

func DocDanhSachDacSanCSDL() ([]DacSan, error) {
	return DocDacSanCSDL(db.Query("SELECT * FROM dac_san ORDER BY id ASC"))
}

func DocTrangDacSanCSDL(soTrang int, doDaiTrang int) ([]DacSan, error) {
	return DocDacSanCSDL(db.Query("SELECT * FROM dac_san ORDER BY id ASC LIMIT ?, ?", soTrang*doDaiTrang, doDaiTrang))
}

func DocTrangDacSanTheoTenCSDL(soTrang int, doDaiTrang int, ten string) ([]DacSan, error) {
	return DocDacSanCSDL(db.Query("SELECT * FROM dac_san WHERE ten LIKE ? ORDER BY id ASC LIMIT ?, ?", "%"+ten+"%", soTrang*doDaiTrang, doDaiTrang))
}

func DocTrangDacSanTheoVungMienCSDL(soTrang int, doDaiTrang int, id int) ([]DacSan, error) {
	return DocDacSanCSDL(db.Query("SELECT * FROM dac_san, dac_san_thuoc_vung WHERE id = id_dac_san AND id_vung_mien = ? ORDER BY id ASC LIMIT ?, ?", id, soTrang*doDaiTrang, doDaiTrang))
}

func DocTrangDacSanTheoNguyenLieuCSDL(soTrang int, doDaiTrang int, id int) ([]DacSan, error) {
	return DocDacSanCSDL(db.Query("SELECT * FROM dac_san, thanh_phan WHERE id = id_dac_san AND id_nguyen_lieu = ? ORDER BY id ASC LIMIT ?, ?", id, soTrang*doDaiTrang, doDaiTrang))
}

func DocTrangDacSanTheoMuaDacSanCSDL(soTrang int, doDaiTrang int, id int) ([]DacSan, error) {
	return DocDacSanCSDL(db.Query("SELECT * FROM dac_san, dac_san_theo_mua WHERE id = id_dac_san AND id_mua_dac_san = ? ORDER BY id ASC LIMIT ?, ?", id, soTrang*doDaiTrang, doDaiTrang))
}

func DocDacSanTheoIdCSDL(id int) (DacSan, error) {
	var dacSan DacSan

	row := db.QueryRow("SELECT * FROM dac_san WHERE id = ?", id)

	var idHinhDaiDien int
	if err := row.Scan(&dacSan.ID, &dacSan.Ten, &dacSan.MoTa, &dacSan.CachCheBien, &dacSan.LuotXem, &dacSan.DiemDanhGia, &dacSan.LuotDanhGia, &idHinhDaiDien); err != nil {
		return dacSan, err
	}
	thanhPhan, err := DocThanhPhanTheoIdCSDL(dacSan.ID)
	if err == nil {
		dacSan.ThanhPhan = thanhPhan
	} else {
		fmt.Println(err)
	}
	vungMien, err := DocVungMienDacSanCSDL(dacSan.ID)
	if err == nil {
		dacSan.VungMien = vungMien
	} else {
		fmt.Println(err)
	}
	muaDacSan, err := DocMuaDacSanCSDL(dacSan.ID)
	if err == nil {
		dacSan.MuaDacSan = muaDacSan
	} else {
		fmt.Println(err)
	}
	hinhAnh, err := DocHinhAnhDacSanCSDL(dacSan.ID)
	if err == nil {
		dacSan.HinhAnh = hinhAnh
	} else {
		fmt.Println(err)
	}
	hinhDaiDien, err := DocHinhAnhTheoIdCSDL(idHinhDaiDien)
	if err == nil {
		dacSan.HinhDaiDien = hinhDaiDien
	} else {
		fmt.Println(err)
	}
	dacSan.DiemDanhGia = TinhDiemDanhGiaDacSanCSDL(dacSan.ID)

	return dacSan, nil
}

func DocDacSanTheoNoiBanCSDL(id int) ([]DacSan, error) {
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
		dacSan, err := DocDacSanTheoIdCSDL(idDacSan)
		if err == nil {
			dsDacSan = append(dsDacSan, dacSan)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return dsDacSan, nil
}

func DocDacSanTheoTenCSDL(text string) ([]DacSan, error) {
	return DocLike(text, "ten", "dac_san", DocDacSanTheoIdCSDL)
}

func DocDacSanTheoMoTaCSDL(text string) ([]DacSan, error) {
	return DocLike(text, "mo_ta", "dac_san", DocDacSanTheoIdCSDL)
}

func DocDacSanTheoCachCheBienCSDL(text string) ([]DacSan, error) {
	return DocLike(text, "cach_che_bien", "dac_san", DocDacSanTheoIdCSDL)
}

func DocLike(text string, cot string, bang string, timTheoID func(int) (DacSan, error)) ([]DacSan, error) {
	dsDacSan := []DacSan{}

	rows, err := db.Query("SELECT id FROM "+bang+" WHERE "+cot+" LIKE ? ORDER BY id ASC", "%"+text+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var idDacSan int
		if err := rows.Scan(&idDacSan); err != nil {
			return nil, err
		}
		dacSan, err := timTheoID(idDacSan)
		if err == nil {
			dsDacSan = append(dsDacSan, dacSan)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return dsDacSan, nil
}

func ThemDacSanCSDL(dacSan DacSan) (DacSan, error) {
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

func CapNhatDacSanCSDL(dacSan DacSan) error {
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

func XoaDacSanCSDL(id int) error {
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
