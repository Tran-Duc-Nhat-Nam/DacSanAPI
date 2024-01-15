package models

import "strconv"

type HinhAnh struct {
	ID   int    `json:"id"`
	Ten  string `json:"ten"`
	MoTa string `json:"mo_ta"`
	URL  string `json:"url"`
}

func DocHinhAnhTheoIdCSDL(id int) (HinhAnh, error) {
	var hinhAhh HinhAnh

	row := db.QueryRow("SELECT * FROM hinh_anh WHERE id = " + strconv.Itoa(id))

	if err := row.Scan(&hinhAhh.ID, &hinhAhh.Ten, &hinhAhh.MoTa, &hinhAhh.URL); err != nil {
		return hinhAhh, err
	}

	return hinhAhh, nil
}

func DocHinhAnhDacSanCSDL(id int) ([]HinhAnh, error) {
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
		hinhAnh, err := DocHinhAnhTheoIdCSDL(idHinhAnh)
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

func ThemHinhAnhCSDL(hinhAnh HinhAnh) error {
	var count int
	db.QueryRow("SELECT MAX(id) FROM hinh_anh").Scan(&count)
	_, err := db.Exec("INSERT INTO hinh_anh VALUES (?, ?, ?, ?)", count, hinhAnh.Ten, hinhAnh.MoTa, hinhAnh.URL)
	return err
}

func CapNhatHinhAnhCSDL(hinhAnh HinhAnh) error {
	_, err := db.Exec("UPDATE hinh_anh SET ten = ?, mota = ?, url = ? WHERE id = ?", hinhAnh.Ten, hinhAnh.MoTa, hinhAnh.URL, hinhAnh.ID)
	return err
}

func XoaHinhAnhCSDL(id int) error {
	_, err := db.Exec("DELETE FROM hinh_anh WHERE id = ?", id)
	return err
}
