package models

import "strconv"

type NguyenLieu struct {
	ID  int    `json:"id"`
	Ten string `json:"ten"`
}

func DocNguyenLieuCSDL() ([]NguyenLieu, error) {
	dsNguyenLieu := []NguyenLieu{}

	rows, err := db.Query("SELECT * FROM nguyen_lieu ORDER BY id ASC")
	if err != nil {
		return dsNguyenLieu, err
	}
	defer rows.Close()

	for rows.Next() {
		var nguyenLieu NguyenLieu
		if err := rows.Scan(&nguyenLieu.ID, &nguyenLieu.Ten); err != nil {
			return dsNguyenLieu, err
		}
		dsNguyenLieu = append(dsNguyenLieu, nguyenLieu)
	}

	if err := rows.Err(); err != nil {
		return dsNguyenLieu, err
	}

	return dsNguyenLieu, nil
}

func DocNguyenLieuTheoIdCSDL(id int) (NguyenLieu, error) {
	var nguyenLieu NguyenLieu

	rows := db.QueryRow("SELECT * FROM nguyen_lieu WHERE id = " + strconv.Itoa(id))

	if err := rows.Scan(&nguyenLieu.ID, &nguyenLieu.Ten); err != nil {
		return nguyenLieu, err
	}

	return nguyenLieu, nil
}

func DocNguyenLieuThanhPhanCSDL(id int) ([]NguyenLieu, error) {
	dsNguyenLieu := []NguyenLieu{}

	rows, err := db.Query("SELECT * FROM thanh_phan WHERE id_dac_san = ?", id)
	if err != nil {
		return dsNguyenLieu, err
	}
	defer rows.Close()

	for rows.Next() {
		var idNguyenLieu int
		if err := rows.Scan(&id, &idNguyenLieu); err != nil {
			return dsNguyenLieu, err
		}
		nguyenLieu, err := DocNguyenLieuTheoIdCSDL(idNguyenLieu)
		if err != nil {
			return dsNguyenLieu, err
		}
		dsNguyenLieu = append(dsNguyenLieu, nguyenLieu)
	}

	if err := rows.Err(); err != nil {
		return dsNguyenLieu, err
	}

	return dsNguyenLieu, nil
}

func ThemNguyenLieuCSDL(nguyenLieu NguyenLieu) (NguyenLieu, error) {
	var count int
	db.QueryRow("SELECT MAX(id) FROM nguyen_lieu").Scan(&count)
	count++
	_, err := db.Exec("INSERT INTO nguyen_lieu VALUES (?, ?)", count, nguyenLieu.Ten)
	nguyenLieu.ID = count
	return nguyenLieu, err
}

func CapNhatNguyenLieuCSDL(nguyenLieu NguyenLieu) error {
	_, err := db.Exec("UPDATE nguyen_lieu SET ten = ? WHERE id = ?", nguyenLieu.Ten, nguyenLieu.ID)
	return err
}

func XoaNguyenLieuCSDL(id int) error {
	_, err := db.Exec("DELETE FROM nguyen_lieu WHERE id = ?", id)
	return err
}
