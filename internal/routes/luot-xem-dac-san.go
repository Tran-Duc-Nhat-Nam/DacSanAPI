package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nam/dac_san_api/internal/models"
	"net/http"
	"strconv"
)

func DocDanhSachDacSanDaXemAPI(c *gin.Context) {
	idNguoiDung := c.Param("idNguoiDung")

	ds, err := models.DocDanhSachDacSanDaXem(idNguoiDung)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, ds)
}

func TimKiemDanhSachDacSanDaXemAPI(c *gin.Context) {
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

	ds, err := models.TimKiemDanhSachDacSanDaXem(soTrang, kichThuocTrang, idNguoiDung)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, ds)
}
