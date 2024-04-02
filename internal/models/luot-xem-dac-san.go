package models

import (
	"time"
)

type LuotXemDacSan struct {
	IdNguoiDung string    `json:"id_nguoi_dung"`
	IdDacSan    int       `json:"id_dac_san"`
	ThoiGianXem time.Time `json:"thoi_gian"`
}

func DocLichSuXemDacSan(idNguoiDung string) ([]LuotXemDacSan, error) {
	lichSuXem := []LuotXemDacSan{}

	rows, err := db.Query("SELECT * FROM luot_xem_dac_san WHERE id_nguoi_dung = ?", idNguoiDung)
	if err != nil {
		return lichSuXem, err
	}
	defer rows.Close()

	for rows.Next() {
		var luotXem LuotXemDacSan
		if err := rows.Scan(&luotXem.IdNguoiDung, &luotXem.IdDacSan, &luotXem.ThoiGianXem); err != nil {
			return lichSuXem, err
		}
		lichSuXem = append(lichSuXem, luotXem)
	}

	if err := rows.Err(); err != nil {
		return lichSuXem, err
	}

	return lichSuXem, nil
}

func DocDanhSachDacSanDaXem(idNguoiDung string) ([]DacSan, error) {
	danhSachDacSan := []DacSan{}

	danhSachYeuThichDacSan, err := DocLichSuXemDacSan(idNguoiDung)

	if err == nil {
		for _, item := range danhSachYeuThichDacSan {
			dacSan, err := DocDacSanTheoId(item.IdDacSan)
			if err == nil {
				danhSachDacSan = append(danhSachDacSan, dacSan)
			}

		}
	}
	return danhSachDacSan, err
}

func ThemLuotXemDacSan(luotXem LuotXemDacSan) error {
	_, err := db.Exec("INSERT INTO luot_xem_dac_san VALUES (?, ?, ?)", luotXem.IdNguoiDung, luotXem.IdDacSan, time.Now())
	return err
}

func ThemLuotXemDacSanKhach(id int) error {
	_, err := db.Exec("CALL watch(?)", id)
	return err
}
