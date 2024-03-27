package models

import "database/sql"

type NoiBan struct {
	ID          int     `json:"id"`
	Ten         string  `json:"ten"`
	MoTa        string  `json:"mo_ta"`
	DiaChi      DiaChi  `json:"dia_chi"`
	DsDacSan    []int   `json:"ds_dac_san"`
	LuotXem     int     `json:"luot_xem"`
	DiemDanhGia float64 `json:"diem_danh_gia"`
	LuotDanhGia int     `json:"luot_danh_gia"`
}

func DocNoiBan(rows *sql.Rows, err error) ([]NoiBan, error) {
	dsNoiBan := []NoiBan{}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var noiBan NoiBan
		var diaChiId int
		if err := rows.Scan(&noiBan.ID, &noiBan.Ten, &noiBan.MoTa, &diaChiId, &noiBan.LuotXem, &noiBan.DiemDanhGia, &noiBan.LuotDanhGia); err != nil {
			return nil, err
		}
		diaChi, err := DocDiaChiTheoId(diaChiId)
		if err == nil {
			noiBan.DiaChi = diaChi
		}
		dsID, err := DocNoiBanDacSanTheoNoiBan(noiBan.ID)
		if err == nil {
			noiBan.DsDacSan = dsID
		}
		dsNoiBan = append(dsNoiBan, noiBan)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return dsNoiBan, nil
}

func DocDanhSachNoiBan() ([]NoiBan, error) {
	return DocNoiBan(db.Query("SELECT * FROM noi_ban ORDER BY id ASC"))
}

func TimKiemNoiBan(soTrang int, doDaiTrang int, ten string) ([]NoiBan, error) {
	return DocNoiBan(db.Query("SELECT * FROM noi_ban WHERE ten LIKE ? ORDER BY id ASC LIMIT ?, ?", "%"+ten+"%", soTrang*doDaiTrang, doDaiTrang))
}

func DocNoiBanTheoId(id int) (NoiBan, error) {
	var noiBan NoiBan
	var idDiaChi int

	row := db.QueryRow("SELECT * FROM noi_ban WHERE id = ?", id)
	if err := row.Scan(&noiBan.ID, &noiBan.Ten, &noiBan.MoTa, &idDiaChi, &noiBan.LuotXem, &noiBan.DiemDanhGia, &noiBan.LuotDanhGia); err != nil {
		if err == sql.ErrNoRows {
			return noiBan, err
		}
		return noiBan, err
	}
	diaChi, err := DocDiaChiTheoId(idDiaChi)
	if err == nil {
		noiBan.DiaChi = diaChi
	}
	dsID, err := DocNoiBanDacSanTheoNoiBan(noiBan.ID)
	if err == nil {
		noiBan.DsDacSan = dsID
	}
	return noiBan, nil
}

func ThemNoiBan(noiBan NoiBan) (NoiBan, error) {
	diaChi, err := TimDiaChi(noiBan.DiaChi)
	if err != nil {
		diaChi, err = ThemDiaChi(noiBan.DiaChi)
	} else {
		noiBan.DiaChi = diaChi
	}
	noiBan.ID = TaoIdMoi("noi_ban")
	_, err = db.Exec("INSERT INTO noi_ban VALUES (?, ?, ?, ?, ?, ?, ?)", noiBan.ID, noiBan.Ten, noiBan.MoTa, diaChi.ID, noiBan.LuotXem, noiBan.DiemDanhGia, noiBan.LuotDanhGia)
	if err == nil {
		for _, idDacSan := range noiBan.DsDacSan {
			_, err = db.Exec("INSERT INTO noi_ban_dac_san VALUES (?, ?)", noiBan.ID, idDacSan)
			if err != nil {
				return noiBan, err
			}
		}
	}
	return noiBan, err
}

func CapNhatNoiBan(noiBan NoiBan) error {
	err := CapNhatDiaChi(noiBan.DiaChi)
	if err != nil {
		return err
	}
	noiBan.ID = TaoIdMoi("noi_ban")
	_, err = db.Exec("UPDATE noi_ban SET ten = ?, mo_ta = ?, dia_chi = ?, luot_xem = ?, diem_danh_gia = ?, luot_danh_gia = ? WHERE id = ?", noiBan.Ten, noiBan.MoTa, noiBan.DiaChi.ID, noiBan.LuotXem, noiBan.DiemDanhGia, noiBan.LuotDanhGia, noiBan.ID)
	if err != nil {
		_, err = db.Exec("DELETE FROM noi_ban_dac_san WHERE id_noi_ban = ?)", noiBan.ID)
		for _, idDacSan := range noiBan.DsDacSan {
			_, err = db.Exec("INSERT INTO noi_ban_dac_san VALUES (?, ?)", noiBan.ID, idDacSan)
			if err != nil {
				return err
			}
		}
	}
	return err
}

func XoaNoiBan(id int) error {
	_, err := db.Exec("DELETE FROM noi_ban WHERE id = ?", id)
	return err
}
