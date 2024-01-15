package models

import "strconv"

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

func DocDacSanCSDL() ([]DacSan, error) {
	var dsDacSan []DacSan

	rows, err := db.Query("SELECT * FROM dac_san ORDER BY id ASC")
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
		thanhPhan, err := DocThanhPhanTheoIdCSDL(dacSan.ID)
		if err == nil {
			dacSan.ThanhPhan = thanhPhan
		}
		vungMien, err := DocVungMienDacSanCSDL(dacSan.ID)
		if err == nil {
			dacSan.VungMien = vungMien
		}
		muaDacSan, err := DocMuaDacSanCSDL(dacSan.ID)
		if err == nil {
			dacSan.MuaDacSan = muaDacSan
		}
		hinhAnh, err := DocHinhAnhDacSanCSDL(dacSan.ID)
		if err == nil {
			dacSan.HinhAnh = hinhAnh
		}
		hinhDaiDien, err := DocHinhAnhTheoIdCSDL(idHinhDaiDien)
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

func DocDacSanTheoIdCSDL(id int) (DacSan, error) {
	var dacSan DacSan

	row := db.QueryRow("SELECT * FROM dac_san WHERE id = " + strconv.Itoa(id))

	var idHinhDaiDien int
	if err := row.Scan(&dacSan.ID, &dacSan.Ten, &dacSan.MoTa, &dacSan.CachCheBien, &dacSan.LuotXem, &dacSan.DiemDanhGia, &dacSan.LuotDanhGia, &idHinhDaiDien); err != nil {
		return dacSan, err
	}
	thanhPhan, err := DocThanhPhanTheoIdCSDL(dacSan.ID)
	if err == nil {
		dacSan.ThanhPhan = thanhPhan
	}
	vungMien, err := DocVungMienDacSanCSDL(dacSan.ID)
	if err == nil {
		dacSan.VungMien = vungMien
	}
	muaDacSan, err := DocMuaDacSanCSDL(dacSan.ID)
	if err == nil {
		dacSan.MuaDacSan = muaDacSan
	}
	hinhAnh, err := DocHinhAnhDacSanCSDL(dacSan.ID)
	if err == nil {
		dacSan.HinhAnh = hinhAnh
	}
	hinhDaiDien, err := DocHinhAnhTheoIdCSDL(idHinhDaiDien)
	if err == nil {
		dacSan.HinhDaiDien = hinhDaiDien
	}

	return dacSan, nil
}

func ThemDacSanCSDL(dacSan DacSan) error {
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

func CapNhatDacSanCSDL(dacSan DacSan) error {
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

func XoaDacSanCSDL(id int) error {
	_, err := db.Exec("DELETE FROM dac_san WHERE id = ?", id)
	return err
}
