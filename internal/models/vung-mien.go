package models

import "strconv"

type VungMien struct {
	ID  int    `json:"id"`
	Ten string `json:"ten"`
}

func DocVungMien() ([]VungMien, error) {
	dsVungMien := []VungMien{}

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

func DocVungMienTheoId(id int) (VungMien, error) {
	var vungMien VungMien

	row := db.QueryRow("SELECT * FROM vung_mien WHERE id = ?", id)

	if err := row.Scan(&vungMien.ID, &vungMien.Ten); err != nil {
		return vungMien, err
	}

	return vungMien, nil
}

func DocVungMienDacSan(id int) ([]VungMien, error) {
	dsVungMien := []VungMien{}

	rows, err := db.Query("SELECT * FROM dac_san_thuoc_vung WHERE id_dac_san = " + strconv.Itoa(id) + " ORDER BY id_dac_san ASC")
	if err != nil {
		return dsVungMien, err
	}
	defer rows.Close()

	for rows.Next() {
		var idVungMien int
		if err := rows.Scan(&id, &idVungMien); err != nil {
			return dsVungMien, err
		}
		vungMien, err := DocVungMienTheoId(idVungMien)
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

func ThemVungMien(vungMien VungMien) (VungMien, error) {
	var count int
	db.QueryRow("SELECT MAX(id) FROM vung_mien").Scan(&count)
	count++
	_, err := db.Exec("INSERT INTO vung_mien VALUES (?, ?)", count, vungMien.Ten)
	vungMien.ID = count
	return vungMien, err
}

func CapNhatVungMien(vungMien VungMien) error {
	_, err := db.Exec("UPDATE vung_mien SET ten = ? WHERE id = ?", vungMien.Ten, vungMien.ID)
	return err
}

func XoaVungMien(id int) error {
	_, err := db.Exec("DELETE FROM vung_mien WHERE id = ?", id)
	return err
}
