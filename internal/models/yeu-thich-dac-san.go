package models

import "database/sql"

type YeuThichDacSan struct {
	IdNguoiDung string `json:"id_nguoi_dung"`
	IdDacSan    int    `json:"id_dac_san"`
}

func DocYeuThichDacSan(rows *sql.Rows, err error) ([]YeuThichDacSan, error) {
	danhSachYeuThichDacSan := []YeuThichDacSan{}
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

func DocYeuThichDacSanTheoDacSan(idDacSan int) ([]YeuThichDacSan, error) {
	return DocYeuThichDacSan(db.Query("SELECT * FROM danh_gia_dac_san WHERE id_dac_san = ?", idDacSan))
}

func DocYeuThichDacSanTheoNguoiDung(idNguoiDung string) ([]YeuThichDacSan, error) {
	return DocYeuThichDacSan(db.Query("SELECT * FROM yeu_thich_dac_san WHERE id_nguoi_dung = ?", idNguoiDung))
}

func TimKiemYeuThichDacSanTheoDacSan(soTrang int, doDaiTrang int, idDacSan int) ([]YeuThichDacSan, error) {
	return DocYeuThichDacSan(db.Query("SELECT * FROM danh_gia_dac_san WHERE id_dac_san = ? ORDER BY id ASC LIMIT ?, ?", idDacSan, soTrang*doDaiTrang, doDaiTrang))
}

func TimKiemYeuThichDacSanTheoNguoiDung(soTrang int, doDaiTrang int, idNguoiDung string) ([]YeuThichDacSan, error) {
	return DocYeuThichDacSan(db.Query("SELECT * FROM danh_gia_dac_san WHERE id_nguoi_dung = ? ORDER BY id ASC LIMIT ?, ?", idNguoiDung, soTrang*doDaiTrang, doDaiTrang))
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

func TimKiemDanhSachDacSanYeuThich(soTrang int, doDaiTrang int, idNguoiDung string) ([]DacSan, error) {
	danhSachDacSan := []DacSan{}

	danhSachYeuThichDacSan, err := TimKiemYeuThichDacSanTheoNguoiDung(soTrang, doDaiTrang, idNguoiDung)

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

func DocYeuThichDacSanTheoID(yeuThichDacSan YeuThichDacSan) error {
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
