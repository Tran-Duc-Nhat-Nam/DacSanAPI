package models

import "strconv"

type VungMien struct {
	ID  int    `json:"id"`
	Ten string `json:"ten"`
}

func DocVungMienCSDL() ([]VungMien, error) {
	var dsVungMien []VungMien

	rows, err := db.Query("SELECT * FROM vung_mien ORDER BY id ASC")
	if err != nil {
		return dsVungMien, err
	}
	defer rows.Close()

	for rows.Next() {
		var vungMien VungMien
		if err := rows.Scan(&vungMien.ID, &vungMien.Ten); err != nil {
			return dsVungMien, err
		}
		dsVungMien = append(dsVungMien, vungMien)
	}

	if err := rows.Err(); err != nil {
		return dsVungMien, err
	}

	return dsVungMien, nil
}

func DocVungMienTheoIdCSDL(id int) (VungMien, error) {
	var vungMien VungMien

	rows, err := db.Query("SELECT * FROM vung_mien WHERE id = " + strconv.Itoa(id))
	if err != nil {
		return vungMien, err
	}
	defer rows.Close()

	if err := rows.Scan(&vungMien.ID, &vungMien.Ten); err != nil {
		return vungMien, err
	}

	return vungMien, nil
}

func DocVungMienDacSanCSDL(id int) ([]VungMien, error) {
	var dsVungMien []VungMien

	rows, err := db.Query("SELECT * FROM dac_san_thuoc_vung WHERE id_dac_san = " + strconv.Itoa(id) + " ORDER BY id ASC")
	if err != nil {
		return dsVungMien, err
	}
	defer rows.Close()

	for rows.Next() {
		var idVungMien int
		if err := rows.Scan(nil, &idVungMien); err != nil {
			return nil, err
		}
		vungMien, err := DocVungMienTheoIdCSDL(idVungMien)
		if err != nil {
			return dsVungMien, err
		}
		dsVungMien = append(dsVungMien, vungMien)
	}

	if err := rows.Err(); err != nil {
		return dsVungMien, err
	}

	return dsVungMien, nil
}

func ThemVungMienCSDL(vungMien VungMien) (VungMien, error) {
	var count int
	db.QueryRow("SELECT MAX(id) FROM vung_mien").Scan(&count)
	count++
	_, err := db.Exec("INSERT INTO vung_mien VALUES (?, ?)", count, vungMien.Ten)
	vungMien.ID = count
	return vungMien, err
}

func CapNhatVungMienCSDL(vungMien VungMien) error {
	_, err := db.Exec("UPDATE vung_mien SET ten = ? WHERE id = ?", vungMien.Ten, vungMien.ID)
	return err
}

func XoaVungMienCSDL(id int) error {
	_, err := db.Exec("DELETE FROM vung_mien WHERE id = ?", id)
	return err
}
