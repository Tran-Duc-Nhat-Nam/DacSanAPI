package routes

import (
	"fmt"
	"nam/dac_san_api/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func TinhDiemDanhGiaNoiBanTheoIdAPI(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusConflict, 0)
	}
	c.IndentedJSON(http.StatusOK, models.TinhDiemDanhGiaNoiBan(id))
}

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
