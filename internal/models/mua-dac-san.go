package models

type MuaDacSan struct {
	ID    int    `json:"id"`
	Ten   string `json:"ten"`
	Thang []int  `json:"thang"`
}

func DocMua() ([]MuaDacSan, error) {
	dsMuaDacSan := []MuaDacSan{}

	rows, err := db.Query("SELECT * FROM mua_dac_san ORDER BY id ASC")
	if err != nil {
		return dsMuaDacSan, err
	}
	defer rows.Close()

	for rows.Next() {
		var muaDacSan MuaDacSan
		if err := rows.Scan(&muaDacSan.ID, &muaDacSan.Ten); err != nil {
			return dsMuaDacSan, err
		}
		dsMuaDacSan = append(dsMuaDacSan, muaDacSan)
	}

	if err := rows.Err(); err != nil {
		return dsMuaDacSan, err
	}

	return dsMuaDacSan, nil
}

func DocMuaTheoId(id int) (MuaDacSan, error) {
	var mua MuaDacSan

	row := db.QueryRow("SELECT * FROM mua_dac_san WHERE id = ?", id)

	if err := row.Scan(&mua.ID, &mua.Ten); err != nil {
		return mua, err
	}

	return mua, nil
}

func DocMuaDacSan(id int) ([]MuaDacSan, error) {
	dsMuaDacSan := []MuaDacSan{}

	rows, err := db.Query("SELECT * FROM dac_san_theo_mua WHERE id_dac_san = ?", id)
	if err != nil {
		return dsMuaDacSan, err
	}
	defer rows.Close()

	for rows.Next() {
		var idMuaDacSan int
		if err := rows.Scan(&id, &idMuaDacSan); err != nil {
			return dsMuaDacSan, err
		}
		muaDacSan, err := DocMuaTheoId(idMuaDacSan)
		if err != nil {
			return dsMuaDacSan, err
		}
		dsMuaDacSan = append(dsMuaDacSan, muaDacSan)
	}

	if err := rows.Err(); err != nil {
		return dsMuaDacSan, err
	}

	return dsMuaDacSan, nil
}

func ThemMua(muaDacSan MuaDacSan) (MuaDacSan, error) {
	var count int
	db.QueryRow("SELECT MAX(id) FROM mua_dac_san").Scan(&count)
	count++
	_, err := db.Exec("INSERT INTO mua_dac_san VALUES (?, ?)", count, muaDacSan.Ten)
	muaDacSan.ID = count
	return muaDacSan, err
}

func CapNhatMua(muaDacSan MuaDacSan) error {
	_, err := db.Exec("UPDATE mua_dac_san SET ten = ? WHERE id = ?", muaDacSan.Ten, muaDacSan.ID)
	return err
}

func XoaMua(id int) error {
	_, err := db.Exec("DELETE FROM mua_dac_san WHERE id = ?", id)
	return err
}
