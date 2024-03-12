package routes

import (
	"fmt"
	"nam/dac_san_api/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DocNguyenLieuJson(c *gin.Context) {
	dsNguyenLieu, _ := models.DocDanhSachNguyenLieuCSDL()
	c.IndentedJSON(http.StatusOK, dsNguyenLieu)
}

func DocNguyenLieuTheoTenJson(c *gin.Context) {
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
	dacSan, err := models.DocNguyenLieuTheoTenCSDL(soTrang, kichThuocTrang, ten)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
	}
	c.IndentedJSON(http.StatusOK, dacSan)
}

func ThemNguyenLieuJson(c *gin.Context) {
	var nguyenLieu models.NguyenLieu

	if err := c.BindJSON(&nguyenLieu); err != nil {
		fmt.Print(err.Error())
		return
	}

	nguyenLieu, err := models.ThemNguyenLieuCSDL(nguyenLieu)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, nguyenLieu)
	} else {
		c.IndentedJSON(http.StatusCreated, nguyenLieu)
	}
}

func CapNhatNguyenLieuJson(c *gin.Context) {
	var nguyenLieu models.NguyenLieu

	if err := c.BindJSON(&nguyenLieu); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.CapNhatNguyenLieuCSDL(nguyenLieu)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}

func XoaNguyenLieuJson(c *gin.Context) {
	var Doc map[string]int

	if err := c.BindJSON(&Doc); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.XoaNguyenLieuCSDL(Doc["id"])
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}
