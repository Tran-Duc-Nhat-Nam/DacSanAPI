package models

func DocNoiBanDacSanTheoDacSanCSDL(id int) ([]int, error) {
	dsID := []int{}

	rows, err := db.Query("SELECT * FROM noi_ban_dac_san WHERE id_dac_san = ?", id)
	if err != nil {
		return dsID, err
	}
	defer rows.Close()

	for rows.Next() {
		var kq int
		if err := rows.Scan(&kq, &id); err != nil {
			return nil, err
		}
		dsID = append(dsID, kq)
	}

	if err := rows.Err(); err != nil {
		return dsID, err
	}

	return dsID, nil
}

func DocNoiBanDacSanTheoNoiBanCSDL(id int) ([]int, error) {
	dsID := []int{}

	rows, err := db.Query("SELECT * FROM noi_ban_dac_san WHERE id_noi_ban = ?", id)
	if err != nil {
		return dsID, err
	}
	defer rows.Close()

	for rows.Next() {
		var kq int
		if err := rows.Scan(&id, &kq); err != nil {
			return nil, err
		}
		dsID = append(dsID, kq)
	}

	if err := rows.Err(); err != nil {
		return dsID, err
	}

	return dsID, nil
}

func ThemNoiBanDacSanCSDL(idNoiBan int, idDacSan int) error {
	_, err := db.Exec("INSERT INTO noi_ban_dac_san VALUES (?, ?)", idNoiBan, idDacSan)
	return err
}

func XoaNoiBanDacSanCSDL(idNoiBan int, idDacSan int) error {
	_, err := db.Exec("DELETE FROM noi_ban_dac_san WHERE id_noi_ban = ? AND id_dac_san = ?", idDacSan, idDacSan)
	return err
}
