package models

type ThanhPhan struct {
	NguyenLieu NguyenLieu `json:"nguyen_lieu"`
	SoLuong    float64    `json:"so_luong"`
	DonViTinh  string     `json:"don_vi_tinh"`
}

func DocThanhPhanTheoId(id int) ([]ThanhPhan, error) {
	dsThanhPhan := []ThanhPhan{}

	rows, err := db.Query("SELECT * FROM thanh_phan WHERE id_dac_san = ?", id)
	if err != nil {
		return dsThanhPhan, err
	}
	defer rows.Close()

	for rows.Next() {
		var thanhPhan ThanhPhan
		var temp int
		var idNguyenLieu int
		if err := rows.Scan(&temp, &idNguyenLieu, &thanhPhan.SoLuong, &thanhPhan.DonViTinh); err != nil {
			return dsThanhPhan, err
		}
		nguyenLieu, err := DocNguyenLieuTheoId(idNguyenLieu)
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

func ThemThanhPhan(idDacSan int, thanhPhan ThanhPhan) error {
	_, err := db.Exec("INSERT INTO thanh_phan VALUES (?, ?, ?, ?)", idDacSan, thanhPhan.NguyenLieu.ID, thanhPhan.SoLuong, thanhPhan.DonViTinh)
	return err
}

func CapNhatThanhPhan(idDacSan int, thanhPhan ThanhPhan) error {
	_, err := db.Exec("UPDATE thanh_phan SET so_luong = ?, don_vi_tinh = ? WHERE id_dac_san = ? AND id_nguyen_lieu = ?", thanhPhan.SoLuong, thanhPhan.DonViTinh, idDacSan, thanhPhan.NguyenLieu.ID)
	return err
}

func XoaThanhPhan(idDacSan int, idNguyenLieu int) error {
	_, err := db.Exec("DELETE FROM thanh_phan WHERE id_dac_san = ? AND id_nguyen_lieu = ?", idDacSan, idNguyenLieu)
	return err
}
