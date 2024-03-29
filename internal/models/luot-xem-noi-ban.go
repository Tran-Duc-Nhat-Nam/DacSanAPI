package models

import (
	"time"
)

type LuotXemNoiBan struct {
	IdNguoiDung string    `json:"id_nguoi_dung"`
	IdNoiBan    int       `json:"id_noi_ban"`
	ThoiGianXem time.Time `json:"thoi_gian"`
}

func DocLichSuXemNoiBan(idNguoiDung string) ([]LuotXemNoiBan, error) {
	lichSuXem := []LuotXemNoiBan{}

	rows, err := db.Query("SELECT * FROM luot_xem_noi_ban WHERE id_nguoi_dung = ?", idNguoiDung)
	if err != nil {
		return lichSuXem, err
	}
	defer rows.Close()

	for rows.Next() {
		var luotXem LuotXemNoiBan
		if err := rows.Scan(&luotXem.IdNguoiDung, &luotXem.IdNoiBan, &luotXem.ThoiGianXem); err != nil {
			return lichSuXem, err
		}
		lichSuXem = append(lichSuXem, luotXem)
	}

	if err := rows.Err(); err != nil {
		return lichSuXem, err
	}

	return lichSuXem, nil
}

func DemLuotXemNoiBan(idNguoiDung string) (int, error) {
	soLuotXem := -1
	row := db.QueryRow("SELECT COUNT(*) FROM luot_xem_noi_ban WHERE id_nguoi_dung = ?", idNguoiDung)
	err := row.Scan(&soLuotXem)
	return soLuotXem, err
}

func ThemLuotXemNoiBan(luotXem LuotXemNoiBan) error {
	_, err := db.Exec("INSERT INTO luot_xem_noi_ban VALUES (?, ?, ?)", luotXem.IdNguoiDung, luotXem.IdNoiBan, time.Now())
	return err
}

func ThemLuotXemNoiBanKhach(id int) error {
	_, err := db.Exec("CALL watch_nb(?)", id)
	return err
}
