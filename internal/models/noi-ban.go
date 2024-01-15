package models

type NoiBan struct {
	ID          int     `json:"id"`
	Ten         string  `json:"ten"`
	MoTa        string  `json:"mo_ta"`
	DiaChi      DiaChi  `json:"dia_chi"`
	LuotXem     int     `json:"luot_xem"`
	DiemDanhGia float64 `json:"diem_danh_gia"`
	LuotDanhGia int     `json:"luot_danh_gia"`
}

func DocNoiBanCSDL() ([]NoiBan, error) {
	var dsNoiBan []NoiBan
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
		dsNoiBan = append(dsNoiBan, noiBan)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return dsNoiBan, nil
}

func ThemNoiBanCSDL(noiBan NoiBan) (NoiBan, error) {
	var count int
	db.QueryRow("SELECT MAX(id) FROM dia_chi").Scan(&count)
	count++
	result, err := db.Exec("INSERT INTO dia_chi VALUES (?, ?, ?, ?, ?, ?)", count, noiBan.DiaChi.SoNha, noiBan.DiaChi.TenDuong, noiBan.DiaChi.PhuongXa, noiBan.DiaChi.QuanHuyen, noiBan.DiaChi.TinhThanh.ID)
	idDiaChi, err := result.LastInsertId()
	if err != nil {
		return noiBan, err
	}
	db.QueryRow("SELECT MAX(id) FROM noi_ban").Scan(&count)
	count++
	_, err = db.Exec("INSERT INTO noi_ban VALUES (?, ?, ?, ?, ?, ?)", count, noiBan.Ten, noiBan.MoTa, idDiaChi, noiBan.LuotXem, noiBan.DiemDanhGia, noiBan.LuotDanhGia)
	noiBan.ID = count
	return noiBan, err
}

func CapNhatNoiBanCSDL(noiBan NoiBan) error {
	_, err := db.Exec("UPDATE noi_ban SET ten = ?, mo_ta = ?, dia_chi = ?, luot_xem = ?, diem_danh_gia = ?, luot_danh_gia = ? WHERE id = ?", noiBan.Ten, noiBan.MoTa, noiBan.DiaChi.ID, noiBan.LuotXem, noiBan.DiemDanhGia, noiBan.LuotDanhGia, noiBan.ID)
	return err
}

func XoaNoiBanCSDL(id int) error {
	_, err := db.Exec("DELETE FROM noi_ban WHERE id = ?", id)
	return err
}
