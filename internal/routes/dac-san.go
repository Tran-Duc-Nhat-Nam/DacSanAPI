package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nam/dac_san_api/internal/models"
	"net/http"
	"strconv"
	"time"
)

func DocDacSanJson(c *gin.Context) {
	dsDacSan, _ := models.DocDanhSachDacSan()
	c.IndentedJSON(http.StatusOK, dsDacSan)
}

func DocTrangDacSanJson(c *gin.Context) {
	soTrang, err := strconv.Atoi(c.Param("index"))
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
	}
	kichThuocTrang, err := strconv.Atoi(c.Param("size"))
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
	}
	dacSan, err := models.DocDanhSachDacSanTheoTrang(soTrang, kichThuocTrang)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
	}
	c.IndentedJSON(http.StatusOK, dacSan)
}

func TimKiemDacSanAPI(c *gin.Context) {
	ten := c.Param("ten")
	soTrang, err := strconv.Atoi(c.Param("index"))
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}
	kichThuocTrang, err := strconv.Atoi(c.Param("size"))
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}
	dacSan, err := models.TimKiemDacSan(soTrang, kichThuocTrang, ten)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, dacSan)
}

func TimKiemDacSanTheoVungMienAPI(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
	}
	soTrang, err := strconv.Atoi(c.Param("index"))
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
	}
	kichThuocTrang, err := strconv.Atoi(c.Param("size"))
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
	}
	dacSan, err := models.TimKiemDacSanTheoVungMien(soTrang, kichThuocTrang, id)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
	}
	c.IndentedJSON(http.StatusOK, dacSan)
}

func TimKiemDacSanTheoMuaAPI(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
	}
	soTrang, err := strconv.Atoi(c.Param("index"))
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
	}
	kichThuocTrang, err := strconv.Atoi(c.Param("size"))
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
	}
	dacSan, err := models.TimKiemDacSanTheoMua(soTrang, kichThuocTrang, id)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
	}
	c.IndentedJSON(http.StatusOK, dacSan)
}

func TimKiemDacSanTheoNguyenLieuAPI(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
	}
	soTrang, err := strconv.Atoi(c.Param("index"))
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
	}
	kichThuocTrang, err := strconv.Atoi(c.Param("size"))
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
	}
	dacSan, err := models.TimKiemDacSanTheoNguyenLieu(soTrang, kichThuocTrang, id)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
	}
	c.IndentedJSON(http.StatusOK, dacSan)
}

func DocDacSanTheoIdJson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err.Error())
	}
	dacSan, err := models.DocDacSanTheoId(id)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, dacSan)
}

func XemDacSan(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err.Error())
	}
	idNguoiDung := c.Param("idnguoidung")
	dacSan, err := models.DocDacSanTheoId(id)
	if err != nil {
		fmt.Print(err.Error())
	}
	err = models.ThemLuotXemDacSan(models.LuotXemDacSan{IdNguoiDung: idNguoiDung, IdDacSan: id, ThoiGianXem: time.Now()})
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, dacSan)
}

func XemDacSanKhach(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err.Error())
	}
	dacSan, err := models.DocDacSanTheoId(id)
	if err != nil {
		fmt.Print(err.Error())
	}
	err = models.ThemLuotXemDacSanKhach(id)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, dacSan)
}

func DocDacSanTheoNoiBanAPI(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err.Error())
	}
	dacSan, err := models.DocDacSanTheoNoiBanCSDL(id)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, dacSan)
}

func DocDacSanTheoTenJson(c *gin.Context) {
	ten := c.Param("ten")
	dacSan, err := models.DocDacSanTheoTenCSDL(ten)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, dacSan)
}

func DocDacSanTheoMoTaJson(c *gin.Context) {
	moTa := c.Param("moTa")
	dacSan, err := models.DocDacSanTheoMoTaCSDL(moTa)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, dacSan)
}

func DocDacSanTheoCachCheBienJson(c *gin.Context) {
	cachCheBien := c.Param("cachCheBien")
	dacSan, err := models.DocDacSanTheoCachCheBienCSDL(cachCheBien)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, dacSan)
}

func TimKiemDacSanTheoDanhSachVungMienAPI(c *gin.Context) {
	var Doc map[string][]int

	if err := c.BindJSON(&Doc); err != nil {
		fmt.Print(err.Error())
		return
	}

	var dsID []int

	for _, id := range Doc["dsVungMien"] {
		dsID = append(dsID, id)
	}

	ten, soTrang, kichThuocTrang, err := docTuKhoaTimKiem(c)
	if err == nil {
		dacSan, err := models.TimKiemDacSanTheoDanhSachVungMien(soTrang, kichThuocTrang, ten, dsID)
		if err != nil {
			fmt.Print(err.Error())
		}
		c.IndentedJSON(http.StatusOK, dacSan)
	}
}

func TimKiemDacSanTheoDanhSachMuaAPI(c *gin.Context) {
	var Doc map[string][]int

	if err := c.BindJSON(&Doc); err != nil {
		fmt.Print(err.Error())
		return
	}

	var dsID []int

	for _, id := range Doc["dsMuaDacSan"] {
		dsID = append(dsID, id)
	}

	ten, soTrang, kichThuocTrang, err := docTuKhoaTimKiem(c)
	if err == nil {
		dacSan, err := models.TimKiemDacSanTheoDanhSachMua(soTrang, kichThuocTrang, ten, dsID)
		if err != nil {
			fmt.Print(err.Error())
		}
		c.IndentedJSON(http.StatusOK, dacSan)
	}
}

func TimKiemDacSanTheoMuaVungMienAPI(c *gin.Context) {
	var Doc map[string][]int

	if err := c.BindJSON(&Doc); err != nil {
		fmt.Print(err.Error())
		return
	}

	var dsVungMien []int
	var dsMuaDacSan []int

	for _, id := range Doc["dsVungMien"] {
		dsVungMien = append(dsVungMien, id)
	}

	for _, id := range Doc["dsMuaDacSan"] {
		dsMuaDacSan = append(dsMuaDacSan, id)
	}

	ten := c.Param("ten")
	soTrang, err := strconv.Atoi(c.Param("index"))
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
	}
	kichThuocTrang, err := strconv.Atoi(c.Param("size"))
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
	}

	dacSan, err := models.TimKiemDacSanTheoMuaVungMien(soTrang, kichThuocTrang, dsVungMien, dsMuaDacSan, ten)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, dacSan)
}

func TimKiemDacSanTheoDanhSachNguyenLieuAPI(c *gin.Context) {
	var Doc map[string][]int

	if err := c.BindJSON(&Doc); err != nil {
		fmt.Print(err.Error())
		return
	}

	var dsID []int

	for _, id := range Doc["dsNguyenLieu"] {
		dsID = append(dsID, id)
	}

	ten, soTrang, kichThuocTrang, err := docTuKhoaTimKiem(c)
	if err == nil {
		dacSan, err := models.TimKiemDacSanTheoDanhSachNguyenLieu(soTrang, kichThuocTrang, ten, dsID)
		if err != nil {
			fmt.Print(err.Error())
		}
		c.IndentedJSON(http.StatusOK, dacSan)
	}
}

func TimKiemDacSanTheoNguyenLieuVungMienAPI(c *gin.Context) {
	var Doc map[string][]int

	if err := c.BindJSON(&Doc); err != nil {
		fmt.Print(err.Error())
		return
	}

	var dsNguyenLieu []int
	var dsVungMien []int

	for _, id := range Doc["dsNguyenLieu"] {
		dsNguyenLieu = append(dsNguyenLieu, id)
	}

	for _, id := range Doc["dsVungMien"] {
		dsVungMien = append(dsVungMien, id)
	}

	ten, soTrang, kichThuocTrang, err := docTuKhoaTimKiem(c)
	if err == nil {
		dacSan, err := models.TimKiemDacSanTheoNguyenLieuVungMien(soTrang, kichThuocTrang, dsNguyenLieu, dsVungMien, ten)
		if err != nil {
			fmt.Print(err.Error())
		}
		c.IndentedJSON(http.StatusOK, dacSan)
	}
}

func TimKiemDacSanTheoNguyenLieuMuaAPI(c *gin.Context) {
	var Doc map[string][]int

	if err := c.BindJSON(&Doc); err != nil {
		fmt.Print(err.Error())
		return
	}

	var dsNguyenLieu []int
	var dsMuaDacSan []int

	for _, id := range Doc["dsNguyenLieu"] {
		dsNguyenLieu = append(dsNguyenLieu, id)
	}

	for _, id := range Doc["dsMuaDacSan"] {
		dsMuaDacSan = append(dsMuaDacSan, id)
	}

	ten, soTrang, kichThuocTrang, err := docTuKhoaTimKiem(c)
	if err == nil {
		dacSan, err := models.TimKiemDacSanTheoNguyenLieuMua(soTrang, kichThuocTrang, dsNguyenLieu, dsMuaDacSan, ten)
		if err != nil {
			fmt.Print(err.Error())
		}
		c.IndentedJSON(http.StatusOK, dacSan)
	}
}

func TimKiemDacSanTheoDieuKien(c *gin.Context) {
	var Doc map[string][]int

	if err := c.BindJSON(&Doc); err != nil {
		fmt.Print(err.Error())
		return
	}

	var dsNguyenLieu []int
	var dsVungMien []int
	var dsMuaDacSan []int

	for _, id := range Doc["dsNguyenLieu"] {
		dsNguyenLieu = append(dsNguyenLieu, id)
	}

	for _, id := range Doc["dsVungMien"] {
		dsVungMien = append(dsVungMien, id)
	}

	for _, id := range Doc["dsMuaDacSan"] {
		dsMuaDacSan = append(dsMuaDacSan, id)
	}

	ten, soTrang, kichThuocTrang, err := docTuKhoaTimKiem(c)
	if err == nil {
		dacSan, err := models.TimKiemDacSanTheoDieuKien(soTrang, kichThuocTrang, dsNguyenLieu, dsVungMien, dsMuaDacSan, ten)
		if err != nil {
			fmt.Print(err.Error())
		}
		c.IndentedJSON(http.StatusOK, dacSan)
	}
}

func ThemDacSanJson(c *gin.Context) {
	var dacSan models.DacSan

	if err := c.BindJSON(&dacSan); err != nil {
		fmt.Print(err.Error())
		return
	}

	dacSan, err := models.ThemDacSanCSDL(dacSan)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, dacSan)
	} else {
		c.IndentedJSON(http.StatusCreated, dacSan)
	}
}

func CapNhatDacSanJson(c *gin.Context) {
	var dacSan models.DacSan

	if err := c.BindJSON(&dacSan); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.CapNhatDacSan(dacSan)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}

func XoaDacSanJson(c *gin.Context) {
	var Doc map[string]int

	if err := c.BindJSON(&Doc); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.XoaDacSanCSDL(Doc["id"])
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}
