package models

type NoiBanDacSan struct {
	IdNoiBan  int    `json:"id_noi_ban"`
	IdDacSan  int    `json:"id_dac_san"`
	GiaBan    int    `json:"gia_ban"`
	DonViTinh string `json:"don_vi_tinh"`
}

func DocNoiBanDacSanTheoIdCSDL(id int) ([]NoiBanDacSan, error) {
	dsNoiBanDacSan := []NoiBanDacSan{}

	rows, err := db.Query("SELECT * FROM thanh_phan WHERE id_dac_san = ?", id)
	if err != nil {
		return dsNoiBanDacSan, err
	}
	defer rows.Close()

	for rows.Next() {
		var noiBanDacSan NoiBanDacSan
		if err := rows.Scan(&noiBanDacSan.IdNoiBan, &noiBanDacSan.IdDacSan, &noiBanDacSan.GiaBan, &noiBanDacSan.DonViTinh); err != nil {
			return nil, err
		}
		dsNoiBanDacSan = append(dsNoiBanDacSan, noiBanDacSan)
	}

	if err := rows.Err(); err != nil {
		return dsNoiBanDacSan, err
	}

	return dsNoiBanDacSan, nil
}

func ThemNoiBanDacSanCSDL(noiBanDacSan NoiBanDacSan) error {
	_, err := db.Exec("INSERT INTO noi_ban_dac_san VALUES (?, ?, ?, ?)", noiBanDacSan.IdNoiBan, noiBanDacSan.IdDacSan, noiBanDacSan.GiaBan, noiBanDacSan.DonViTinh)
	return err
}

func CapNhatNoiBanDacSanCSDL(noiBanDacSan NoiBanDacSan) error {
	_, err := db.Exec("UPDATE noi_ban_dac_san SET so_luong = ?, don_vi_tinh = ? WHERE id_noi_ban = ? AND id_dac_san = ?", noiBanDacSan.GiaBan, noiBanDacSan.DonViTinh, noiBanDacSan.IdNoiBan, noiBanDacSan.IdDacSan)
	return err
}

func XoaNoiBanDacSanCSDL(idNoiBan int, idDacSan int) error {
	_, err := db.Exec("DELETE FROM noi_ban_dac_san WHERE id_noi_ban = ? AND id_dac_san = ?", idDacSan, idDacSan)
	return err
}
