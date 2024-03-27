package routes

import (
	"fmt"
	"nam/dac_san_api/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DocNguyenLieuAPI(c *gin.Context) {
	dsNguyenLieu, _ := models.DocDanhSachNguyenLieu()
	c.IndentedJSON(http.StatusOK, dsNguyenLieu)
}

func DocNguyenLieuTheoTenAPI(c *gin.Context) {
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
	dacSan, err := models.DocNguyenLieuTheoTen(soTrang, kichThuocTrang, ten)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
	}
	c.IndentedJSON(http.StatusOK, dacSan)
}

func ThemNguyenLieuAPI(c *gin.Context) {
	var nguyenLieu models.NguyenLieu

	if err := c.BindJSON(&nguyenLieu); err != nil {
		fmt.Print(err.Error())
		return
	}

	nguyenLieu, err := models.ThemNguyenLieu(nguyenLieu)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, nguyenLieu)
	} else {
		c.IndentedJSON(http.StatusCreated, nguyenLieu)
	}
}

func CapNhatNguyenLieuAPI(c *gin.Context) {
	var nguyenLieu models.NguyenLieu

	if err := c.BindJSON(&nguyenLieu); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.CapNhatNguyenLieu(nguyenLieu)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}

func XoaNguyenLieuAPI(c *gin.Context) {
	var Doc map[string]int

	if err := c.BindJSON(&Doc); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.XoaNguyenLieu(Doc["id"])
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}
