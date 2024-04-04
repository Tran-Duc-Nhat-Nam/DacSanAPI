package routes

import (
	"fmt"
	"nam/dac_san_api/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func YeuThichDacSanAPI(c *gin.Context) {
	idDacSan, err := strconv.Atoi(c.Param("idDacSan"))
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}
	idNguoiDung := c.Param("idNguoiDung")

	if err := models.ThemYeuThichDacSan(models.YeuThichDacSan{IdNguoiDung: idNguoiDung, IdDacSan: idDacSan}); err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}

func BoYeuThichDacSanAPI(c *gin.Context) {
	idDacSan, err := strconv.Atoi(c.Param("idDacSan"))
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}
	idNguoiDung := c.Param("idNguoiDung")

	if err := models.XoaYeuThichDacSan(models.YeuThichDacSan{IdNguoiDung: idNguoiDung, IdDacSan: idDacSan}); err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}

func DocYeuThichDacSanAPI(c *gin.Context) {
	idDacSan, err := strconv.Atoi(c.Param("idDacSan"))
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}
	idNguoiDung := c.Param("idNguoiDung")

	if err := models.DocYeuThichDacSanTheoID(models.YeuThichDacSan{IdNguoiDung: idNguoiDung, IdDacSan: idDacSan}); err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}

func TimKiemDanhSachDacSanYeuThichAPI(c *gin.Context) {
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
	idNguoiDung := c.Param("idNguoiDung")

	ds, err := models.TimKiemDanhSachDacSanYeuThich(soTrang, kichThuocTrang, idNguoiDung)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, ds)
}

func DocDanhSachDacSanYeuThichAPI(c *gin.Context) {
	idNguoiDung := c.Param("idNguoiDung")

	ds, err := models.DocDanhSachDacSanYeuThich(idNguoiDung)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, ds)
}
