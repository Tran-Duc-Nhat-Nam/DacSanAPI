package models

import (
	"strconv"
	"time"
)

type LuotXemNoiBan struct {
	IdNguoiDung int       `json:"id_nguoi_dung"`
	IdNoiBan    int       `json:"id_noi_ban"`
	ThoiGianXem time.Time `json:"thoi_gian"`
}

func DocLichSuXemNoiBanCSDL(idNguoiDung int) ([]LuotXemNoiBan, error) {
	var lichSuXem []LuotXemNoiBan

	rows, err := db.Query("SELECT * FROM lich_su_xem_noi_ban WHERE id_nguoi_dung = " + strconv.Itoa(idNguoiDung))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var luotXem LuotXemNoiBan
		if err := rows.Scan(&luotXem.IdNguoiDung, &luotXem.IdNoiBan, &luotXem.ThoiGianXem); err != nil {
			return nil, err
		}
		lichSuXem = append(lichSuXem, luotXem)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return lichSuXem, nil
}

func ThemLuotXemNoiBanCSDL(luotXem LuotXemNoiBan) error {
	_, err := db.Exec("INSERT INTO luot_xem_noi_ban VALUES (?, ?, ?, ?)", luotXem.IdNguoiDung, luotXem.IdNoiBan, time.Now())
	return err
}
