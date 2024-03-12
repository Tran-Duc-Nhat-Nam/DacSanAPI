package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nam/dac_san_api/internal/models"
	"net/http"
	"strconv"
)

func TinhDiemDanhGiaDacSanTheoIdJson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusConflict, 0)
	}
	c.IndentedJSON(http.StatusOK, models.TinhDiemDanhGiaDacSanCSDL(id))
}

func DanhGiaDacSanAPI(c *gin.Context) {
	var danhGia models.LuotDanhGiaDacSan

	if err := c.BindJSON(&danhGia); err != nil {
		fmt.Print(err.Error())
		return
	}

	if err := models.ThemDanhGiaDacSanCSDL(danhGia); err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, true)
}

func DocDanhGiaDacSanTheoNguoiDungAPI(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("idDacSan"))
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}
	idNguoiDung := c.Param("idNguoiDung")
	diem, err := models.DocDacSanTheoNguoiDungCSDL(id, idNguoiDung)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, diem)
}

func DocDanhSachDanhGiaDacSanAPI(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("idDacSan"))
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}
	diem, err := models.DocDanhGiaDacSanCSDL(id)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, diem)
}
