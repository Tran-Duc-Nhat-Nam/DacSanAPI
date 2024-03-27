package routes

import (
	"fmt"
	"nam/dac_san_api/internal/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func DocNoiBanAPI(c *gin.Context) {
	dsNoiBan, _ := models.DocDanhSachNoiBan()
	c.IndentedJSON(http.StatusOK, dsNoiBan)
}

func DocNoiBanTheoIdAPI(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}
	dacSan, err := models.DocNoiBanTheoId(id)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, dacSan)
}

func TimKiemNoiBanAPI(c *gin.Context) {
	ten, soTrang, kichThuocTrang, err := docTuKhoaTimKiem(c)
	if err == nil {
		dacSan, err := models.TimKiemNoiBan(soTrang, kichThuocTrang, ten)
		if err != nil {
			fmt.Print(err.Error())
			c.IndentedJSON(http.StatusConflict, err.Error())
		}
		c.IndentedJSON(http.StatusOK, dacSan)
	}
}

func XemNoiBan(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err.Error())
	}
	idNguoiDung := c.Param("idnguoidung")
	dacSan, err := models.DocNoiBanTheoId(id)
	if err != nil {
		fmt.Print(err.Error())
	}
	err = models.ThemLuotXemNoiBan(models.LuotXemNoiBan{IdNguoiDung: idNguoiDung, IdNoiBan: id, ThoiGianXem: time.Now()})
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, dacSan)
}

func XemNoiBanKhach(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err.Error())
	}
	dacSan, err := models.DocNoiBanTheoId(id)
	if err != nil {
		fmt.Print(err.Error())
	}
	err = models.ThemLuotXemNoiBanKhach(id)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, dacSan)
}

func ThemNoiBanAPI(c *gin.Context) {
	var noiBan models.NoiBan

	if err := c.BindJSON(&noiBan); err != nil {
		fmt.Print(err.Error())
		return
	}

	noiBan, err := models.ThemNoiBan(noiBan)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, noiBan)
	} else {
		c.IndentedJSON(http.StatusCreated, noiBan)
	}
}

func CapNhatNoiBanAPI(c *gin.Context) {
	var noiBan models.NoiBan

	if err := c.BindJSON(&noiBan); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.CapNhatNoiBan(noiBan)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}

func XoaNoiBanAPI(c *gin.Context) {
	var Doc map[string]int

	if err := c.BindJSON(&Doc); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.XoaNoiBan(Doc["id"])
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}
