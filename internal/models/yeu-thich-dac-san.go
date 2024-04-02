package models

type YeuThichDacSan struct {
	IdNguoiDung string `json:"id_nguoi_dung"`
	IdDacSan    int    `json:"id_dac_san"`
}

func DocYeuThichDacSanTheoDacSan(idDacSan int) ([]YeuThichDacSan, error) {
	danhSachYeuThichDacSan := []YeuThichDacSan{}

	rows, err := db.Query("SELECT * FROM danh_gia_dac_san WHERE id_dac_san = ?", idDacSan)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var yeuThichDacSan YeuThichDacSan
		if err := rows.Scan(&yeuThichDacSan.IdDacSan, &yeuThichDacSan.IdNguoiDung); err != nil {
			return nil, err
		}
		danhSachYeuThichDacSan = append(danhSachYeuThichDacSan, yeuThichDacSan)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return danhSachYeuThichDacSan, nil
}

func DocYeuThichDacSanTheoNguoiDung(idNguoiDung string) ([]YeuThichDacSan, error) {
	danhSachYeuThichDacSan := []YeuThichDacSan{}

	rows, err := db.Query("SELECT * FROM yeu_thich_dac_san WHERE id_nguoi_dung = ?", idNguoiDung)
	if err != nil {
		return danhSachYeuThichDacSan, err
	}
	defer rows.Close()

	for rows.Next() {
		var yeuThichDacSan YeuThichDacSan
		if err := rows.Scan(&yeuThichDacSan.IdDacSan, &yeuThichDacSan.IdNguoiDung); err != nil {
			return danhSachYeuThichDacSan, err
		}
		danhSachYeuThichDacSan = append(danhSachYeuThichDacSan, yeuThichDacSan)
	}

	if err := rows.Err(); err != nil {
		return danhSachYeuThichDacSan, err
	}

	return danhSachYeuThichDacSan, err
}

func DocDanhSachDacSanYeuThich(idNguoiDung string) ([]DacSan, error) {
	danhSachDacSan := []DacSan{}

	danhSachYeuThichDacSan, err := DocYeuThichDacSanTheoNguoiDung(idNguoiDung)

	if err == nil {
		for _, item := range danhSachYeuThichDacSan {
			dacSan, err := DocDacSanTheoId(item.IdDacSan)
			if err == nil {
				danhSachDacSan = append(danhSachDacSan, dacSan)
			}

		}
	}
	return danhSachDacSan, err
}

func DocYeuThichDacSan(yeuThichDacSan YeuThichDacSan) error {
	rows := db.QueryRow("SELECT * FROM yeu_thich_dac_san WHERE id_dac_san = ? AND id_nguoi_dung = ?", yeuThichDacSan.IdDacSan, yeuThichDacSan.IdNguoiDung)
	err := rows.Scan(&yeuThichDacSan.IdDacSan, &yeuThichDacSan.IdNguoiDung)
	return err
}

func ThemYeuThichDacSan(yeuThichDacSan YeuThichDacSan) error {
	_, err := db.Exec("INSERT INTO yeu_thich_dac_san VALUES (?, ?)", yeuThichDacSan.IdDacSan, yeuThichDacSan.IdNguoiDung)
	return err
}

func XoaYeuThichDacSan(yeuThichDacSan YeuThichDacSan) error {
	_, err := db.Exec("DELETE FROM yeu_thich_dac_san WHERE id_nguoi_dung = ? AND id_dac_san = ?", yeuThichDacSan.IdNguoiDung, yeuThichDacSan.IdDacSan)
	return err
}
