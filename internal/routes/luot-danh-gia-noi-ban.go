package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nam/dac_san_api/internal/models"
	"net/http"
	"strconv"
)

func TinhDiemDanhGiaNoiBanTheoIdJson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusConflict, 0)
	}
	c.IndentedJSON(http.StatusOK, models.TinhDiemDanhGiaNoiBanCSDL(id))
}

func DanhGiaNoiBanAPI(c *gin.Context) {
	var danhGia models.LuotDanhGiaNoiBan

	if err := c.BindJSON(&danhGia); err != nil {
		fmt.Print(err.Error())
		return
	}

	if err := models.ThemDanhGiaNoiBanCSDL(danhGia); err != nil {
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
	diem, err := models.DocNoiBanTheoNguoiDungCSDL(id, idNguoiDung)
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
	diem, err := models.DocDanhGiaNoiBanCSDL(id)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, diem)
}
