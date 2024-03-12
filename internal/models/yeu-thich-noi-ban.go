package models

type YeuThichNoiBan struct {
	IdNguoiDung string `json:"id_nguoi_dung"`
	IdNoiBan    int    `json:"id_noi_ban"`
}

func DocYeuThichNoiBanTheoNoiBanCSDL(idNoiBan int) ([]YeuThichNoiBan, error) {
	danhSachYeuThichNoiBan := []YeuThichNoiBan{}

	rows, err := db.Query("SELECT * FROM danh_gia_noi_ban WHERE id_noi_ban = ?", idNoiBan)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var yeuThichNoiBan YeuThichNoiBan
		if err := rows.Scan(&yeuThichNoiBan.IdNguoiDung, &yeuThichNoiBan.IdNoiBan); err != nil {
			return nil, err
		}
		danhSachYeuThichNoiBan = append(danhSachYeuThichNoiBan, yeuThichNoiBan)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return danhSachYeuThichNoiBan, nil
}

func DocYeuThichNoiBanTheoNguoiDungCSDL(idNguoiDung string) ([]YeuThichNoiBan, error) {
	danhSachYeuThichNoiBan := []YeuThichNoiBan{}

	rows, err := db.Query("SELECT * FROM danh_gia_noi_ban WHERE id_noi_ban = ?", idNguoiDung)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var yeuThichNoiBan YeuThichNoiBan
		if err := rows.Scan(&yeuThichNoiBan.IdNguoiDung, &yeuThichNoiBan.IdNoiBan); err != nil {
			return nil, err
		}
		danhSachYeuThichNoiBan = append(danhSachYeuThichNoiBan, yeuThichNoiBan)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return danhSachYeuThichNoiBan, nil
}

func DocYeuThichNoiBanCSDL(yeuThichNoiBan YeuThichNoiBan) error {
	rows := db.QueryRow("SELECT * FROM yeu_thich_noi_ban WHERE id_noi_ban = ? AND id_nguoi_dung = ?", yeuThichNoiBan.IdNoiBan, yeuThichNoiBan.IdNguoiDung)
	err := rows.Scan(&yeuThichNoiBan.IdNoiBan, &yeuThichNoiBan.IdNguoiDung)
	return err
}

func ThemYeuThichNoiBanCSDL(yeuThichNoiBan YeuThichNoiBan) error {
	_, err := db.Exec("INSERT INTO yeu_thich_noi_ban VALUES (?, ?)", yeuThichNoiBan.IdNoiBan, yeuThichNoiBan.IdNguoiDung)
	return err
}

func XoaYeuThichNoiBanCSDL(yeuThichNoiBan YeuThichNoiBan) error {
	_, err := db.Exec("DELETE FROM yeu_thich_noi_ban WHERE id_nguoi_dung = ? AND id_noi_ban = ?", yeuThichNoiBan.IdNguoiDung, yeuThichNoiBan.IdNoiBan)
	return err
}
