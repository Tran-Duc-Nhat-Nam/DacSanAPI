package routes

import (
	"fmt"
	"nam/dac_san_api/internal/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func DocDacSanJson(c *gin.Context) {
	dsDacSan, _ := models.DocDacSanCSDL()
	c.IndentedJSON(http.StatusOK, dsDacSan)
}

func DocDacSanTheoTrangJson(c *gin.Context) {
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
	dacSan, err := models.DocDacSanTheoTrangCSDL(soTrang, kichThuocTrang)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
	}
	c.IndentedJSON(http.StatusOK, dacSan)
}

func DocDacSanTheoIdJson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err.Error())
	}
	dacSan, err := models.DocDacSanTheoIdCSDL(id)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, dacSan)
}

func DocDacSanTheoNoiBanJson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err.Error())
	}
	dacSan, err := models.DocDacSanTheoNoiBanCSDL(id)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, dacSan)
}

func DocDacSanTheoTenJson(c *gin.Context) {
	ten := c.Param("ten")
	nguoiDung, err := models.DocDacSanTheoTenCSDL(ten)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, nguoiDung)
}

func DocDacSanTheoMoTaJson(c *gin.Context) {
	moTa := c.Param("moTa")
	nguoiDung, err := models.DocDacSanTheoMoTaCSDL(moTa)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, nguoiDung)
}

func DocDacSanTheoCachCheBienJson(c *gin.Context) {
	cachCheBien := c.Param("cachCheBien")
	nguoiDung, err := models.DocDacSanTheoCachCheBienCSDL(cachCheBien)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, nguoiDung)
}

func XemDacSan(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err.Error())
	}
	idNguoiDung := c.Param("idnguoidung")
	dacSan, err := models.DocDacSanTheoIdCSDL(id)
	if err != nil {
		fmt.Print(err.Error())
	}
	err = models.ThemLuotXemDacSanCSDL(models.LuotXemDacSan{IdNguoiDung: idNguoiDung, IdDacSan: dacSan.ID, ThoiGianXem: time.Now()})
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, dacSan)
}

func ThemDacSanJson(c *gin.Context) {
	var dacSan models.DacSan

	if err := c.BindJSON(&dacSan); err != nil {
		fmt.Print(err.Error())
		return
	}

	dacSan, err := models.ThemDacSanCSDL(dacSan)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, dacSan)
	} else {
		c.IndentedJSON(http.StatusCreated, dacSan)
	}
}

func CapNhatDacSanJson(c *gin.Context) {
	var dacSan models.DacSan

	if err := c.BindJSON(&dacSan); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.CapNhatDacSanCSDL(dacSan)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}

func XoaDacSanJson(c *gin.Context) {
	var Doc map[string]int

	if err := c.BindJSON(&Doc); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.XoaDacSanCSDL(Doc["id"])
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}
