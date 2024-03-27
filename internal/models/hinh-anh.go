package models

type HinhAnh struct {
	ID   int    `json:"id"`
	Ten  string `json:"ten"`
	MoTa string `json:"mo_ta"`
	URL  string `json:"url"`
}

func DocHinhAnhTheoId(id int) (HinhAnh, error) {
	var hinhAhh HinhAnh

	row := db.QueryRow("SELECT * FROM hinh_anh WHERE id = ?", id)

	if err := row.Scan(&hinhAhh.ID, &hinhAhh.Ten, &hinhAhh.MoTa, &hinhAhh.URL); err != nil {
		return hinhAhh, err
	}

	return hinhAhh, nil
}

func DocHinhAnhDacSan(id int) ([]HinhAnh, error) {
	dsHinhAnh := []HinhAnh{}

	rows, err := db.Query("SELECT * FROM hinh_anh_dac_san WHERE id_dac_san = ?", id)
	if err != nil {
		return dsHinhAnh, err
	}
	defer rows.Close()

	for rows.Next() {
		var idHinhAnh int
		var idDacSan int
		if err := rows.Scan(&idDacSan, &idHinhAnh); err != nil {
			return nil, err
		}
		hinhAnh, err := DocHinhAnhTheoId(idHinhAnh)
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

func ThemHinhAnh(hinhAnh HinhAnh) error {
	hinhAnh.ID = TaoIdMoi("hinh_anh")
	_, err := db.Exec("INSERT INTO hinh_anh VALUES (?, ?, ?, ?)", hinhAnh.ID, hinhAnh.Ten, hinhAnh.MoTa, hinhAnh.URL)
	return err
}

func CapNhatHinhAnh(hinhAnh HinhAnh) error {
	_, err := db.Exec("UPDATE hinh_anh SET ten = ?, mo_ta = ?, url = ? WHERE id = ?", hinhAnh.Ten, hinhAnh.MoTa, hinhAnh.URL, hinhAnh.ID)
	return err
}

func XoaHinhAnh(id int) error {
	_, err := db.Exec("DELETE FROM hinh_anh WHERE id = ?", id)
	return err
}
