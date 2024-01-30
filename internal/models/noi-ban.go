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

func DocNoiBanCSDL() ([]NoiBan, error) {
	dsNoiBan := []NoiBan{}
	var diaChiId int

	rows, err := db.Query("SELECT * FROM noi_ban ORDER BY id ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var noiBan NoiBan
		if err := rows.Scan(&noiBan.ID, &noiBan.Ten, &noiBan.MoTa, &diaChiId, &noiBan.LuotXem, &noiBan.DiemDanhGia, &noiBan.LuotDanhGia); err != nil {
			return nil, err
		}
		diaChi, err := DocDiaChiTheoIdCSDL(diaChiId)
		if err == nil {
			noiBan.DiaChi = diaChi
		}
		dsID, err := DocNoiBanDacSanTheoNoiBanCSDL(noiBan.ID)
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

func DocNoiBanTheoIdCSDL(id int) (NoiBan, error) {
	var noiBan NoiBan
	var idDiaChi int

	row := db.QueryRow("SELECT * FROM noi_ban WHERE id = ?", id)
	if err := row.Scan(&noiBan.ID, &noiBan.Ten, &noiBan.MoTa, &idDiaChi, &noiBan.LuotXem, &noiBan.DiemDanhGia, &noiBan.LuotDanhGia); err != nil {
		if err == sql.ErrNoRows {
			return noiBan, err
		}
		return noiBan, err
	}
	diaChi, err := DocDiaChiTheoIdCSDL(idDiaChi)
	if err == nil {
		noiBan.DiaChi = diaChi
	}
	dsID, err := DocNoiBanDacSanTheoNoiBanCSDL(noiBan.ID)
	if err == nil {
		noiBan.DsDacSan = dsID
	}
	return noiBan, nil
}

func ThemNoiBanCSDL(noiBan NoiBan) (NoiBan, error) {
	diaChi, err := TimDiaChiCSDL(noiBan.DiaChi)
	if err != nil {
		diaChi, err = ThemDiaChiCSDL(noiBan.DiaChi)
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

func CapNhatNoiBanCSDL(noiBan NoiBan) error {
	_, err := TimDiaChiCSDL(noiBan.DiaChi)
	if err != nil {
		ThemDiaChiCSDL(noiBan.DiaChi)
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

func XoaNoiBanCSDL(id int) error {
	_, err := db.Exec("DELETE FROM noi_ban WHERE id = ?", id)
	return err
}
