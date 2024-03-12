package models

import (
	"database/sql"
	"strconv"
)

type NguyenLieu struct {
	ID  int    `json:"id"`
	Ten string `json:"ten"`
}

func DocNguyenLieuCSDL(rows *sql.Rows, err error) ([]NguyenLieu, error) {
	dsNguyenLieu := []NguyenLieu{}

	if err != nil {
		return dsNguyenLieu, err
	}

	defer rows.Close()

	for rows.Next() {
		var nguyenLieu NguyenLieu
		if err := rows.Scan(&nguyenLieu.ID, &nguyenLieu.Ten); err != nil {
			return dsNguyenLieu, err
		}
		dsNguyenLieu = append(dsNguyenLieu, nguyenLieu)
	}

	if err := rows.Err(); err != nil {
		return dsNguyenLieu, err
	}

	return dsNguyenLieu, nil
}

func DocNguyenLieuTheoIdCSDL(id int) (NguyenLieu, error) {
	var nguyenLieu NguyenLieu

	rows := db.QueryRow("SELECT * FROM nguyen_lieu WHERE id = " + strconv.Itoa(id))

	if err := rows.Scan(&nguyenLieu.ID, &nguyenLieu.Ten); err != nil {
		return nguyenLieu, err
	}

	return nguyenLieu, nil
}

func DocDanhSachNguyenLieuCSDL() ([]NguyenLieu, error) {
	return DocNguyenLieuCSDL(db.Query("SELECT * FROM nguyen_lieu ORDER BY id ASC"))
}

func DocNguyenLieuThanhPhanCSDL(id int) ([]NguyenLieu, error) {
	return DocNguyenLieuCSDL(db.Query("SELECT * FROM thanh_phan WHERE id_dac_san = ? ORDER BY id ASC", id))
}

func DocNguyenLieuTheoTenCSDL(soTrang int, doDaiTrang int, ten string) ([]NguyenLieu, error) {
	return DocNguyenLieuCSDL(db.Query("SELECT * FROM nguyen_lieu WHERE ten LIKE ? ORDER BY id ASC LIMIT ?, ?", "%"+ten+"%", soTrang*doDaiTrang, doDaiTrang))
}

func ThemNguyenLieuCSDL(nguyenLieu NguyenLieu) (NguyenLieu, error) {
	nguyenLieu.ID = TaoIdMoi("nguyen_lieu")
	_, err := db.Exec("INSERT INTO nguyen_lieu VALUES (?, ?)", nguyenLieu.ID, nguyenLieu.Ten)
	return nguyenLieu, err
}

func CapNhatNguyenLieuCSDL(nguyenLieu NguyenLieu) error {
	_, err := db.Exec("UPDATE nguyen_lieu SET ten = ? WHERE id = ?", nguyenLieu.Ten, nguyenLieu.ID)
	return err
}

func XoaNguyenLieuCSDL(id int) error {
	_, err := db.Exec("DELETE FROM nguyen_lieu WHERE id = ?", id)
	return err
}
