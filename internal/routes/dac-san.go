package routes

import (
	"fmt"
	"nam/dac_san_api/internal/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func DocDacSanJson(c *gin.Context) {
	dsDacSan, _ := models.DocDanhSachDacSanCSDL()
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
	dacSan, err := models.DocTrangDacSanCSDL(soTrang, kichThuocTrang)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
	}
	c.IndentedJSON(http.StatusOK, dacSan)
}

func DocTrangDacSanTheoTenJson(c *gin.Context) {
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
	dacSan, err := models.DocTrangDacSanTheoTenCSDL(soTrang, kichThuocTrang, ten)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
	}
	c.IndentedJSON(http.StatusOK, dacSan)
}

func DocTrangDacSanTheoVungMienJson(c *gin.Context) {
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
	dacSan, err := models.DocTrangDacSanTheoVungMienCSDL(soTrang, kichThuocTrang, id)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
	}
	c.IndentedJSON(http.StatusOK, dacSan)
}

func DocTrangDacSanTheoMuaDacSanJson(c *gin.Context) {
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
	dacSan, err := models.DocTrangDacSanTheoMuaDacSanCSDL(soTrang, kichThuocTrang, id)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
	}
	c.IndentedJSON(http.StatusOK, dacSan)
}

func DocTrangDacSanTheoNguyenLieuJson(c *gin.Context) {
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
	dacSan, err := models.DocTrangDacSanTheoNguyenLieuCSDL(soTrang, kichThuocTrang, id)
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
	dacSan, err := models.DocDacSanTheoIdCSDL(id)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, dacSan)
}

func DocDacSanTheoNoiBanJson(c *gin.Context) {
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

func DocDacSanTheoVungMien(c *gin.Context) {
	var Doc map[string][]int

	if err := c.BindJSON(&Doc); err != nil {
		fmt.Print(err.Error())
		return
	}

	var dsID []int

	for _, id := range Doc["dsVungMien"] {
		dsID = append(dsID, id)
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

	dacSan, err := models.DocTrangDacSanTheoTenVaDanhSachVungMien(soTrang, kichThuocTrang, ten, dsID)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, dacSan)
}

func DocDacSanTheoMua(c *gin.Context) {
	var Doc map[string]map[string]int

	if err := c.BindJSON(&Doc); err != nil {
		fmt.Print(err.Error())
		return
	}

	var dsID []int

	for _, id := range Doc["dsMuaDacSan"] {
		dsID = append(dsID, id)
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

	dacSan, err := models.DocTrangDacSanTheoTenVaDanhSachMua(soTrang, kichThuocTrang, ten, dsID)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, dacSan)
}

func DocDacSanTheoThoiKhong(c *gin.Context) {
	var Doc map[string]map[string]int

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

	dacSan, err := models.DocTrangDacSanTheoTenVaThoiKhong(soTrang, kichThuocTrang, dsVungMien, dsMuaDacSan, ten)
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
	dacSan, err := models.DocDacSanTheoIdCSDL(id)
	if err != nil {
		fmt.Print(err.Error())
	}
	err = models.ThemLuotXemDacSanCSDL(models.LuotXemDacSan{IdNguoiDung: idNguoiDung, IdDacSan: id, ThoiGianXem: time.Now()})
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, dacSan)
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

	err := models.CapNhatDacSanCSDL(dacSan)
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
