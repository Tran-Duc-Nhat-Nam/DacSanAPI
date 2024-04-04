package models

import "database/sql"

type YeuThichNoiBan struct {
	IdNguoiDung string `json:"id_nguoi_dung"`
	IdNoiBan    int    `json:"id_noi_ban"`
}

func DocYeuThichNoiBan(rows *sql.Rows, err error) ([]YeuThichNoiBan, error) {
	danhSachYeuThichNoiBan := []YeuThichNoiBan{}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var yeuThichNoiBan YeuThichNoiBan
		if err := rows.Scan(&yeuThichNoiBan.IdNoiBan, &yeuThichNoiBan.IdNguoiDung); err != nil {
			return nil, err
		}
		danhSachYeuThichNoiBan = append(danhSachYeuThichNoiBan, yeuThichNoiBan)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return danhSachYeuThichNoiBan, nil
}

func DocYeuThichNoiBanTheoNoiBan(idNoiBan int) ([]YeuThichNoiBan, error) {
	return DocYeuThichNoiBan(db.Query("SELECT * FROM danh_gia_noi_ban WHERE id_noi_ban = ?", idNoiBan))
}

func DocYeuThichNoiBanTheoNguoiDung(idNguoiDung string) ([]YeuThichNoiBan, error) {
	return DocYeuThichNoiBan(db.Query("SELECT * FROM yeu_thich_noi_ban WHERE id_nguoi_dung = ?", idNguoiDung))
}

func TimKiemYeuThichNoiBanTheoNoiBan(soTrang int, doDaiTrang int, idNoiBan int) ([]YeuThichNoiBan, error) {
	return DocYeuThichNoiBan(db.Query("SELECT * FROM danh_gia_noi_ban WHERE id_noi_ban = ? ORDER BY id ASC LIMIT ?, ?", idNoiBan, soTrang*doDaiTrang, doDaiTrang))
}

func TimKiemYeuThichNoiBanTheoNguoiDung(soTrang int, doDaiTrang int, idNguoiDung string) ([]YeuThichNoiBan, error) {
	return DocYeuThichNoiBan(db.Query("SELECT * FROM danh_gia_noi_ban WHERE id_nguoi_dung = ? ORDER BY id ASC LIMIT ?, ?", idNguoiDung, soTrang*doDaiTrang, doDaiTrang))
}

func DocDanhSachNoiBanYeuThich(idNguoiDung string) ([]NoiBan, error) {
	danhSachNoiBan := []NoiBan{}

	danhSachYeuThichNoiBan, err := DocYeuThichNoiBanTheoNguoiDung(idNguoiDung)

	if err == nil {
		for _, item := range danhSachYeuThichNoiBan {
			noiBan, err := DocNoiBanTheoId(item.IdNoiBan)
			if err == nil {
				danhSachNoiBan = append(danhSachNoiBan, noiBan)
			}

		}
	}
	return danhSachNoiBan, err
}

func TimKiemDanhSachNoiBanYeuThich(soTrang int, doDaiTrang int, idNguoiDung string) ([]NoiBan, error) {
	danhSachNoiBan := []NoiBan{}

	danhSachYeuThichNoiBan, err := TimKiemYeuThichNoiBanTheoNguoiDung(soTrang, doDaiTrang, idNguoiDung)

	if err == nil {
		for _, item := range danhSachYeuThichNoiBan {
			noiBan, err := DocNoiBanTheoId(item.IdNoiBan)
			if err == nil {
				danhSachNoiBan = append(danhSachNoiBan, noiBan)
			}

		}
	}
	return danhSachNoiBan, err
}

func DocYeuThichNoiBanTheoID(yeuThichNoiBan YeuThichNoiBan) error {
	rows := db.QueryRow("SELECT * FROM yeu_thich_noi_ban WHERE id_noi_ban = ? AND id_nguoi_dung = ?", yeuThichNoiBan.IdNoiBan, yeuThichNoiBan.IdNguoiDung)
	err := rows.Scan(&yeuThichNoiBan.IdNoiBan, &yeuThichNoiBan.IdNguoiDung)
	return err
}

func ThemYeuThichNoiBan(yeuThichNoiBan YeuThichNoiBan) error {
	_, err := db.Exec("INSERT INTO yeu_thich_noi_ban VALUES (?, ?)", yeuThichNoiBan.IdNoiBan, yeuThichNoiBan.IdNguoiDung)
	return err
}

func XoaYeuThichNoiBan(yeuThichNoiBan YeuThichNoiBan) error {
	_, err := db.Exec("DELETE FROM yeu_thich_noi_ban WHERE id_nguoi_dung = ? AND id_noi_ban = ?", yeuThichNoiBan.IdNguoiDung, yeuThichNoiBan.IdNoiBan)
	return err
}
