package routes

import (
	"fmt"
	"nam/dac_san_api/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DocNguoiDungAPI(c *gin.Context) {
	nguoiDung, err := models.DocDanhSachNguoiDung()
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, nguoiDung)
}

func DocNguoiDungTheoIdAPI(c *gin.Context) {
	id := c.Param("id")
	nguoiDung, err := models.DocNguoiDungTheoId(id)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, nguoiDung)
}

func TimKiemNguoiDungAPI(c *gin.Context) {
	ten, soTrang, kichThuocTrang, err := docTuKhoaTimKiem(c)
	if err == nil {
		dacSan, err := models.TimKiemNguoiDung(soTrang, kichThuocTrang, ten)
		if err != nil {
			fmt.Print(err.Error())
			c.IndentedJSON(http.StatusConflict, err.Error())
		}
		c.IndentedJSON(http.StatusOK, dacSan)
	}
}

func ThemNguoiDungAPI(c *gin.Context) {
	var nguoiDung models.NguoiDung

	if err := c.BindJSON(&nguoiDung); err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	nguoiDung, err := models.ThemNguoiDung(nguoiDung)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, nguoiDung)
	} else {
		c.IndentedJSON(http.StatusCreated, nguoiDung)
	}
}

func CapNhatNguoiDungAPI(c *gin.Context) {
	var nguoiDung models.NguoiDung

	if err := c.BindJSON(&nguoiDung); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.CapNhatNguoiDung(nguoiDung)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}

func XoaNguoiDungAPI(c *gin.Context) {
	var Doc map[string]int

	if err := c.BindJSON(&Doc); err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
		return
	}

	err := models.XoaNguoiDung(Doc["id"])
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}
