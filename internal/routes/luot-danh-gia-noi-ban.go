package routes

import (
	"fmt"
	"nam/dac_san_api/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DanhGiaNoiBanAPI(c *gin.Context) {
	var danhGia models.LuotDanhGiaNoiBan

	if err := c.BindJSON(&danhGia); err != nil {
		fmt.Print(err.Error())
		return
	}

	if err := models.ThemDanhGiaNoiBan(danhGia); err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, true)
}

func CapNhatDanhGiaNoiBanAPI(c *gin.Context) {
	var danhGia models.LuotDanhGiaDacSan

	if err := c.BindJSON(&danhGia); err != nil {
		fmt.Print(err.Error())
		return
	}

	if err := models.CapNhatDanhGiaDacSan(danhGia); err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, true)
}

func HuyDanhGiaNoiBanAPI(c *gin.Context) {
	idNoiBan, err := strconv.Atoi(c.Param("idNoiBan"))
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}
	idNguoiDung := c.Param("idNguoiDung")

	if err := models.XoaDanhGiaNoiBan(idNguoiDung, idNoiBan); err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, true)
}

func DocDanhGiaNoiBanTheoNguoiDungAPI(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("idNoiBan"))
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}
	idNguoiDung := c.Param("idNguoiDung")
	diem, err := models.DocNoiBanTheoNguoiDung(id, idNguoiDung)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, diem)
}

func DocDanhSachDanhGiaNoiBanAPI(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("idNoiBan"))
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}
	diem, err := models.DocDanhGiaNoiBan(id)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, diem)
}
