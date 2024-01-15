package models

import "strconv"

type MuaDacSan struct {
	ID    int    `json:"id"`
	Ten   string `json:"ten"`
	Thang []int  `json:"thang"`
}

func DocMuaCSDL() ([]MuaDacSan, error) {
	var dsMuaDacSan []MuaDacSan

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

func DocMuaTheoIdCSDL(id int) (MuaDacSan, error) {
	var mua MuaDacSan

	rows, err := db.Query("SELECT * FROM mua_dac_san WHERE id = " + strconv.Itoa(id))
	if err != nil {
		return mua, err
	}
	defer rows.Close()

	if err := rows.Scan(&mua.ID, &mua.Ten); err != nil {
		return mua, err
	}

	return mua, nil
}

func DocMuaDacSanCSDL(id int) ([]MuaDacSan, error) {
	var dsMuaDacSan []MuaDacSan

	rows, err := db.Query("SELECT * FROM dac_san_theo_mua WHERE id_dac_san = " + strconv.Itoa(id))
	if err != nil {
		return dsMuaDacSan, err
	}
	defer rows.Close()

	for rows.Next() {
		var idMuaDacSan int
		if err := rows.Scan(nil, &idMuaDacSan); err != nil {
			return nil, err
		}
		muaDacSan, err := DocMuaTheoIdCSDL(idMuaDacSan)
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

func ThemMuaCSDL(muaDacSan MuaDacSan) (MuaDacSan, error) {
	var count int
	db.QueryRow("SELECT MAX(id) FROM mua_dac_san").Scan(&count)
	count++
	_, err := db.Exec("INSERT INTO mua_dac_san VALUES (?, ?)", count, muaDacSan.Ten)
	muaDacSan.ID = count
	return muaDacSan, err
}

func CapNhatMuaCSDL(muaDacSan MuaDacSan) error {
	_, err := db.Exec("UPDATE mua_dac_san SET ten = ? WHERE id = ?", muaDacSan.Ten, muaDacSan.ID)
	return err
}

func XoaMuaCSDL(id int) error {
	_, err := db.Exec("DELETE FROM mua_dac_san WHERE id = ?", id)
	return err
}
